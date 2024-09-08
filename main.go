package main

import (
	"bytes"
	"fmt"
	"io"
	"net/mail"
	"os"
	"text/template"
	"time"

	"runtime"

	"github.com/jmoiron/sqlx"
	"github.com/madebagus/ledgermate2/conf"
	database "github.com/madebagus/ledgermate2/db"
	"github.com/madebagus/ledgermate2/helper"
	"github.com/madebagus/ledgermate2/model"
	"github.com/madebagus/ledgermate2/structure"
	"github.com/mattn/go-colorable"
	"github.com/robfig/cron/v3"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	//DBSQL
	DBSQL  []*sqlx.DB
	Logger zerolog.Logger
)

const (
	tstamp = "02/01 15:04:05.999"
)

// main : Service initialization
func main() {

	runtime.GOMAXPROCS(4 * runtime.NumCPU())

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic occurred: %#v\n", err)
		}
	}()

	zerolog.TimeFieldFormat = time.RFC3339Nano

	var writers []io.Writer
	if terminal.IsTerminal(int(os.Stdout.Fd())) {
		if runtime.GOOS == "windows" {
			writers = append(writers, zerolog.ConsoleWriter{
				Out:        colorable.NewColorableStdout(),
				TimeFormat: tstamp,
			})
		} else {
			writers = append(writers, zerolog.ConsoleWriter{
				Out:        os.Stdout,
				TimeFormat: tstamp,
				FormatTimestamp: func(i interface{}) string {
					parse, _ := time.Parse(time.RFC3339Nano, i.(string))
					x, _ := helper.TimeInWIB(parse)
					return "\033[1;36m" + x.Format(tstamp) + "\033[0m"
				},
			})
		}
	}

	config := LogConfig{
		ConsoleLoggingEnabled: true,
		EncodeLogsAsJSON:      false,
		FileLoggingEnabled:    true,
		Directory:             "./log",
		Filename:              "services.log",
		MaxSize:               100,
		MaxBackups:            7,
		MaxAge:                30,
	}

	writers = append(writers, newRollingFile(config))
	mw := io.MultiWriter(writers...)

	wr := diode.NewWriter(mw, 500, 50*time.Millisecond, func(missed int) {
		fmt.Printf("Logger Dropped %d messages", missed)
	})
	Logger = zerolog.New(wr).With().Timestamp().Logger()
	Logger.Info().
		Bool("fileLogging", config.FileLoggingEnabled).
		Bool("jsonLogOutput", config.EncodeLogsAsJSON).
		Str("logDirectory", config.Directory).
		Str("fileName", config.Filename).
		Int("maxSizeMB", config.MaxSize).
		Int("maxBackups", config.MaxBackups).
		Int("maxAgeInDays", config.MaxAge).
		Msg("Logging system configured")

	helper.Logger = &Logger

	DBSQL = [](*sqlx.DB){
		database.PSQLDBInit(),
		database.MYSQLInit(),
	}

	Logger.Info().Msg("Cron starting")
	c := cron.New()

	c.AddFunc("@every 5m", clearanceSVC)
	c.AddFunc("@every 6m", update5MSVC)
	c.AddFunc("@every 40m", lockedlotReleaser)
	c.AddFunc("@every 60m", clearingNotification)
	c.AddFunc("@every 30m", cleanLeftoverOrder)
	c.AddFunc("@every 60m", transferProfitShare)
	c.AddFunc("@every 200m", birthDaySender)
	c.AddFunc("@every 400m", bdEmailreleaser)

	c.Start()
	Logger.Info().Msg("Cron Started")
	for {
		time.Sleep(1 * time.Second)
	}
}

const (
	// IDTypeLedgerBatch 0001
	IDTypeLedgerBatch = 0x02
	// IDTypeDeposit 0011
	IDTypeDeposit = 0x03
	// IDTypeLedgerUser 0100
	IDTypeLedgerUser = 0x08
	// IDTypeLedgerUnit 0101
	IDTypeLedgerUnit = 0x0A
	// IDTypeBlock 0111
	IDTypeBlock = 0x0E
	// IDTypeOrder 1010
	IDTypeOrder = 0x14
	// IDTypeTrading 1011
	IDTypeTrading = 0x16
)

func update5MSVC() {
	fmt.Println("##SQL 5M update##")
	DBcon := model.DBInit(DBSQL)
	_, err := DBcon.DB[conf.DBPSQL1].Exec(`REFRESH MATERIALIZED VIEW mv_public_summary`)
	if err != nil {
		Logger.Error().Msgf("update5MSVC-1: ", err)
	}
	_, err = DBcon.DB[conf.DBPSQL1].Exec(`REFRESH MATERIALIZED VIEW mv_public_listing`)
	if err != nil {
		Logger.Error().Msgf("update5MSVC-2: ", err)
	}
}

func clearanceSVC() {
	fmt.Println("##SQL garbage collector##")
	DBcon := model.DBInit(DBSQL)
	_, err := DBcon.DB[conf.DBMYSQL].Exec(`DELETE FROM lost_recovery WHERE TIMESTAMPDIFF(Second,created_at,NOW()) > 10800`) // 3 hours
	if err != nil {
		Logger.Error().Msgf("clearanceSVC-1: ", err)
	}
	_, err = DBcon.DB[conf.DBMYSQL].Exec(`DELETE FROM browser WHERE TIMESTAMPDIFF(Second,ts_introduced,NOW()) > 10800 AND is_allowed=0`)
	if err != nil {
		Logger.Error().Msgf("clearanceSVC-2: ", err)
	}
}

func lockedlotReleaser() {
	fmt.Println("##Scanning Locked lot##")
	DBcon := model.DBInit(DBSQL)
	_, err := DBcon.DB[conf.DBPSQL1].Exec(`DELETE FROM d_locked_lot where lock_end < now()`)
	if err != nil {
		Logger.Error().Msgf("lockedlotReleaser: ", err)
	}
}

func clearingNotification() {
	fmt.Println("##Clearing Notification##")
	DBcon := model.DBInit(DBSQL)
	_, err := DBcon.DB[conf.DBMYSQL].Exec(`DELETE FROM user_notification WHERE TIMESTAMPDIFF(Second,ts,NOW()) > 1036800`) // 12 days
	if err != nil {
		Logger.Error().Msgf("clearingNotification-1: ", err)
	}
	_, err = DBcon.DB[conf.DBMYSQL].Exec(`DELETE FROM user_notification_recipient WHERE TIMESTAMPDIFF(Second,ts,NOW()) > 1036800`) // 12 days
	if err != nil {
		Logger.Error().Msgf("clearingNotification-2: ", err)
	}
	_, err = DBcon.DB[conf.DBMYSQL].Exec(`DELETE FROM user_notification WHERE id not in (select id from user_notification_recipient)`) // delete leftover
	if err != nil {
		Logger.Error().Msgf("clearingNotification-3: ", err)
	}
	_, err = DBcon.DB[conf.DBMYSQL].Exec(`DELETE FROM user_notification_recipient WHERE id not in (select id from user_notification)`) // delete leftover
	if err != nil {
		Logger.Error().Msgf("clearingNotification-4: ", err)
	}
}

func cleanLeftoverOrder() {
	fmt.Println("##Scanning Leftover Orders##")
	DBcon := model.DBInit(DBSQL)

	br := []structure.OrderTracking{}

	err := DBcon.DB[conf.DBPSQL1].Select(&br, `select a.id,b.user_id as user_id_customer, a.batch_id, a.listing_id, c.listing_name, a.account_id_x, a.amount, a.total_lot 
	from d_property_order a, user_key b, property_listing c
	where (select (DATE_PART('day', now()::timestamp - a.ts::timestamp) * 24 + 
       DATE_PART('hour', now()::timestamp - a.ts::timestamp)) * 60 +
       DATE_PART('minute', now()::timestamp - a.ts::timestamp)) > 30 and a.status = 0 and a.account_id_x = b.account_id 
	and a.listing_id = c.listing_id`)
	if err != nil {
		Logger.Error().Msgf("cleanLeftoverOrder_6: ", err)
	}

	for _, x := range br {

		tx, errtx := DBcon.DB[conf.DBPSQL1].Beginx()
		if errtx != nil {
			Logger.Error().Msgf("tx begin : ", errtx)
			continue
		}

		tx2, errtx2 := DBcon.DB[conf.DBMYSQL].Beginx()
		if errtx2 != nil {
			Logger.Error().Msgf("tx2 begin : ", errtx2)
			continue
		}

		amnt := helper.NumFormat(x.Amount)
		_, err7 := tx.Exec(`update d_property_order set status = 3 where id=$1`, x.OrderID)
		if err7 != nil {
			tx.Rollback()
			Logger.Error().Msgf("cleanLeftoverOrder_7: ", err7)
			continue
		}

		_, err8 := tx.Exec(`delete from d_tracking_order where id=$1`, x.OrderID)
		if err8 != nil {
			tx.Rollback()
			Logger.Error().Msgf("cleanLeftoverOrder_8: ", err8)
			continue
		}

		_, err9 := tx.Exec(`update d_deposit_transaction set status = 28 where hash = $1`, x.OrderID)
		if err8 != nil {
			tx.Rollback()
			Logger.Error().Msgf("cleanLeftoverOrder_9: ", err9)
			continue
		}

		msg := "Your Buy Order on Primary Market FAIL! on Batch ID : " + x.BatchID + " - " + x.ListingName + ", please try again later. Your Deposit IDR. " + amnt + " has been return to your Deposit Balance"
		id := helper.NewUUID()
		_, err10 := tx2.Exec(`insert into user_notification (id,
			ts,
			notif_type,
			sender_user_id,
			recipient_group,
			message_body,
			status) values (
			UUID_TO_BIN(?),
			now(),
			43,
			UUID_TO_BIN('00000000-0000-0000-0000-000000000000'),
			1,
			?,
			1)`, id, msg)
		if err10 != nil {
			tx.Rollback()
			tx2.Rollback()
			Logger.Error().Msgf("cleanLeftoverOrder_10: ", err10)
			continue
		}

		_, err11 := tx2.Exec(`INSERT INTO user_notification_recipient (id,
			ts,
			recipient_user_id,
			read_status) VALUES (UUID_TO_BIN(?),now(),UUID_TO_BIN(?),0)`, id, x.UserID)
		if err11 != nil {
			tx.Rollback()
			tx2.Rollback()
			Logger.Error().Msgf("cleanLeftoverOrder_11: ", err11)
			continue
		}

		_, err14 := tx.Exec(`update d_tracking_deposit set spendable=(spendable+$2) where user_id = $1`, x.UserID, x.Amount)
		if err14 != nil {
			tx.Rollback()
			tx2.Rollback()
			Logger.Error().Msgf("cleanLeftoverOrder_14: ", err14)
			continue
		}

		_, err16 := tx.Exec(`UPDATE d_tracking_batch SET lot_remaining=(lot_remaining + $2) WHERE batch_id=$1`, x.BatchID, x.TotalLot)
		if err16 != nil {
			tx.Rollback()
			tx2.Rollback()
			Logger.Error().Msgf("cleanLeftoverOrder_16: ", err16)
			continue
		}
		tx.Commit()
		tx2.Commit()
	}

}

//TODO : clear not used OTP
//TODO : table auto partitioning

func transferProfitShare() {
	fmt.Println("## Scanning and transfer profit share 2 ##")
	DBcon := model.DBInit(DBSQL)

	br := []structure.PSTracking{}

	err := DBcon.DB[conf.DBPSQL1].Select(&br, `select ps.id, uk.user_id, b.account_id,b.listing_id ,b.batch_id ,to_char(b.profit_share_date,'MM-YYYY') as share_period, b.lots, 
	floor(round((b.lots*a.profit_lot_weight)*b.weight))::integer as total_shares 
	from va_profit_lot_weight a, va_profit_user_lots_weight b, user_key uk, profit_share ps 
	where a.listing_id = b.listing_id and a.profit_share_date = b.profit_share_date and b.account_id = uk.account_id and a.listing_id = ps.listing_id and a.profit_share_date = ps.profit_share_date 
	and a.profit_share_date::date <= now()::date`)
	if err != nil {
		Logger.Error().Msgf("transferProfit_6: ", err)
	}

	for _, x := range br {

		tx, errtx := DBcon.DB[conf.DBPSQL1].Beginx()
		if errtx != nil {
			Logger.Error().Msgf("tx begin: ", errtx)
			continue
		}

		tx2, errtx2 := DBcon.DB[conf.DBMYSQL].Beginx()
		if errtx2 != nil {
			Logger.Error().Msgf("tx2 begin: ", errtx2)
			continue
		}

		amnt := helper.NumFormat(x.TotalShare)
		id := helper.GenRandomHashID(IDTypeDeposit)
		desc := "Profit Share at Listing : " + x.ListingID + " Batch : " + x.BatchID + " Share Period : " + string(x.SharePeriod)

		_, err9 := tx.Exec(`insert into d_deposit_transaction(
			hash,
			account_id,
			acc_name,
			trx_sub_type,
			trx_value,
			status,
			trx_type,
			listing_id,
			description, 
			user_id) values (
			$1,	
			$2,
			$3,
			2,
			$4,
			1,
			1,
			$5,
			$6,
			$7)
		`, id, x.AccID, x.BatchID, x.TotalShare, x.ListingID, desc, x.UserID)
		if err9 != nil {
			tx.Rollback()
			Logger.Error().Msgf("transferProfit_7: ", err9)
			continue
		}

		_, err10 := tx.Exec(`update profit_share set trf_status = 1,trf_date=now() where id = $1 and trf_status <> 1`, x.IDPS)
		if err10 != nil {
			tx.Rollback()
			Logger.Error().Msgf("transferProfit_8: ", err10)
			continue
		}

		_, err13 := tx.Exec(`update d_tracking_deposit set spendable=(spendable+$2) where user_id = $1`, x.UserID, x.TotalShare)
		if err13 != nil {
			tx.Rollback()
			Logger.Error().Msgf("transferProfit_11: ", err13)
			continue
		}

		msg := "Congratulation!, PROFIT SHARE of <b>IDR. " + amnt + "</b> based on your ownership at : <b>" + x.ListingID + " Batch : <b>" + x.BatchID + "</b> for Share Period : <b>" + string(x.SharePeriod) + "</b>, Already Transfered to your Balance. Please check your Account Statement for more detail"
		id2 := helper.NewUUID()
		_, err11 := tx2.Exec(`insert into user_notification (id,
			ts,
			notif_type,
			sender_user_id,
			recipient_group,
			message_body,
			status) values (
			UUID_TO_BIN(?),
			now(),
			52,
			UUID_TO_BIN('00000000-0000-0000-0000-000000000000'),
			1,
			?,
			1)`, id2, msg)
		if err11 != nil {
			tx.Rollback()
			tx2.Rollback()
			Logger.Error().Msgf("transferProfit_9: ", err11)
			continue
		}

		_, err12 := tx2.Exec(`INSERT INTO user_notification_recipient (id,
			ts,
			recipient_user_id,
			read_status) VALUES (UUID_TO_BIN(?),now(),UUID_TO_BIN(?),0)`, id2, x.UserID)
		if err12 != nil {
			tx.Rollback()
			tx2.Rollback()
			Logger.Error().Msgf("transferProfit_10: ", err12)
			continue
		}
		tx.Commit()
		tx2.Commit()
	}

}

func birthDaySender() {
	fmt.Println("## Scanning profile birthday and sending email##")
	DBcon := model.DBInit(DBSQL)

	br := []structure.OPbirthDaySender{}

	// send email verification
	data := map[string]interface{}{
		"name":  "FullName",
		"Title": "Birthday",
	}
	t := template.Must(template.New("email").Parse(helper.Emailultah))
	buf := &bytes.Buffer{}
	if err := t.Execute(buf, data); err != nil {
		Logger.Error().Msgf("Send email err : %s", err.Error())
		return
	}
	body := buf.String()

	err := DBcon.DB[conf.DBPSQL1].Select(&br, `select a.user_id, full_name, email from user_profile_decrypt a  WHERE
    DATE_PART('day', a.birth_date) = date_part('day', CURRENT_DATE) AND
    DATE_PART('month', a.birth_date) = date_part('month', CURRENT_DATE) and user_id not in (select user_id from d_email_bd_send)`)

	if err != nil {
		Logger.Error().Msgf("mail sender: ", err)
	}

	for _, x := range br {

		_, errx := DBcon.DB[conf.DBPSQL1].Exec(`insert into d_email_bd_send (user_id, create_at) values ($1,now())`, x.UserID)
		if errx != nil {
			Logger.Error().Msgf("insert into email sender : ", errx)
		}

		tox := mail.Address{Name: x.FullName, Address: x.Email}
		data["name"] = x.FullName
		buf := &bytes.Buffer{}
		if err := t.Execute(buf, data); err != nil {
			Logger.Error().Msgf("Send email err : %s", err.Error())
			return
		}
		body = buf.String()

		helper.SendMailGun(&tox, "Ini Hari Spesialmu!", body, "Birthday")

	}

}

func bdEmailreleaser() {
	fmt.Println("##Clearing birthday email records##")
	DBcon := model.DBInit(DBSQL)
	_, err := DBcon.DB[conf.DBPSQL1].Exec(`delete from d_email_bd_send where create_at::date < now()::date;`)
	if err != nil {
		Logger.Error().Msgf("bdEmailreleaser: ", err)
	}
}
