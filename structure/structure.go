package structure

import (
	"database/sql"
	"encoding/json"
)


//RFTracking :
type RFTracking struct {
	IDRF        int            `json:"id" db:"id"`
	CustAccID   string         `json:"customer_acc" db:"customer_acc"`
	CustUserID  string         `json:"user_id_customer" db:"user_id_customer"`
	RefAccID    sql.NullString `json:"referral_acc" db:"referral_acc"`
	RefUserID   sql.NullString `json:"user_id_referral" db:"user_id_referral"`
	RefAmnt     int64          `json:"referral_fee_amount" db:"referral_fee_amount"`
	CBAmnt      int64          `json:"cashback_amount" db:"cashback_amount"`
	DescRF      string         `json:"description" db:"order_id"`
	BatchID     string         `json:"batch_id" db:"batch_id"`
	ListingName string         `json:"listing_name" db:"listing_name"`
}

//MarshalJSON : convert overload RFTracking
func (a RFTracking) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		IDRF        int    `json:"id"`
		CustAccID   string `json:"customer_acc"`
		RefAccID    string `json:"referral_acc"`
		CustUserID  string `json:"user_id_customer"`
		RefUserID   string `json:"user_id_referral"`
		RefAmnt     int64  `json:"referral_fee_amount"`
		CBAmnt      int64  `json:"cashback_amount"`
		DescRF      string `json:"description"`
		BatchID     string `json:"batch_id"`
		ListingName string `json:"listing_name"`
	}{
		IDRF:        a.IDRF,
		CustAccID:   a.CustAccID,
		RefAccID:    a.RefAccID.String,
		CustUserID:  a.CustUserID,
		RefUserID:   a.RefUserID.String,
		RefAmnt:     a.RefAmnt,
		CBAmnt:      a.CBAmnt,
		DescRF:      a.DescRF,
		BatchID:     a.BatchID,
		ListingName: a.ListingName,
	})
}

//OrderTracking :
type OrderTracking struct {
	OrderID     string `json:"id" db:"id"`
	UserID      string `json:"user_id_customer" db:"user_id_customer"`
	ListingName string `json:"listing_name" db:"listing_name"`
	BatchID     string `json:"batch_id" db:"batch_id"`
	AccID       string `json:"account_id_x" db:"account_id_x"`
	Amount      int64  `json:"amount" db:"amount"`
	ListingID   string `json:"listing_id" db:"listing_id"`
	TotalLot    int64  `json:"total_lot" db:"total_lot"`
}

//MarshalJSON : convert overload OrderTracking
func (a OrderTracking) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OrderID     string `json:"id"`
		UserID      string `json:"user_id_customer"`
		ListingName string `json:"listing_name"`
		BatchID     string `json:"batch_id"`
		AccID       string `json:"account_id_x"`
		Amount      int64  `json:"amount"`
		ListingID   string `json:"listing_id"`
		TotalLot    int64  `json:"total_lot"`
	}{
		OrderID:     a.OrderID,
		UserID:      a.UserID,
		ListingName: a.ListingName,
		BatchID:     a.BatchID,
		AccID:       a.AccID,
		Amount:      a.Amount,
		ListingID:   a.ListingID,
		TotalLot:    a.TotalLot,
	})
}

//Order14Tracking :
type Order14Tracking struct {
	OrderID     string `json:"order_id" db:"order_id"`
	TS          string `json:"ts" db:"ts"`
	ListingID   string `json:"listing_id" db:"listing_id"`
	ListingName string `json:"listing_name" db:"listing_name"`
	BatchID     string `json:"batch_id" db:"batch_id"`
	Amount      int64  `json:"amount" db:"amount"`
	Status      int    `json:"status" db:"status"`
	NotiFlag    int    `json:"notif_flag" db:"notif_flag"`
	UserID      string `json:"user_id" db:"user_id"`
	ID          int    `json:"id" db:"id"`
}

//MarshalJSON : convert overload Order14Tracking
func (a Order14Tracking) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OrderID     string `json:"order_id"`
		TS          string `json:"ts"`
		ListingID   string `json:"listing_id"`
		ListingName string `json:"listing_name"`
		BatchID     string `json:"batch_id"`
		Amount      int64  `json:"amount"`
		Status      int    `json:"status"`
		NotiFlag    int    `json:"notif_flag"`
		UserID      string `json:"user_id"`
		ID          int    `json:"id"`
	}{
		OrderID:     a.OrderID,
		TS:          a.TS,
		ListingID:   a.ListingID,
		ListingName: a.ListingName,
		BatchID:     a.BatchID,
		Amount:      a.Amount,
		Status:      a.Status,
		NotiFlag:    a.NotiFlag,
		UserID:      a.UserID,
		ID:          a.ID,
	})
}

//OPTracking :
type OPTracking struct {
	OrderID string `json:"order_id" db:"order_id"`
}

//MarshalJSON : convert overload OPTracking
func (a OPTracking) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OrderID string `json:"order_id"`
	}{
		OrderID: a.OrderID,
	})
}

//MailSender OPbirthDaySender :
type OPbirthDaySender struct {
	UserID string `json:"user_id" db:"user_id"`
	FullName string `json:"full_name" db:"full_name"`
	Email string `json:"email" db:"email"`
}

//MarshalJSON : convert OPbirthDaySender
func (a OPbirthDaySender) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		UserID string `json:"user_id"`
		FullName string `json:"full_name"`
		Email string `json:"email"`
	}{
		UserID: a.UserID,
		FullName: a.FullName,
		Email: a.Email,
	})
}
