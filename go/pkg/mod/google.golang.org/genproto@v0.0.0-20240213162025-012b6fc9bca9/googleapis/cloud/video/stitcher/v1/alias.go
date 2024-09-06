// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by aliasgen. DO NOT EDIT.

// Package stitcher aliases all exported identifiers in package
// "cloud.google.com/go/video/stitcher/apiv1/stitcherpb".
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb.
// Please read https://github.com/googleapis/google-cloud-go/blob/main/migration.md
// for more details.
package stitcher

import (
	src "cloud.google.com/go/video/stitcher/apiv1/stitcherpb"
	grpc "google.golang.org/grpc"
)

// Deprecated: Please use consts in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
const (
	CompanionAds_ALL                             = src.CompanionAds_ALL
	CompanionAds_ANY                             = src.CompanionAds_ANY
	CompanionAds_DISPLAY_REQUIREMENT_UNSPECIFIED = src.CompanionAds_DISPLAY_REQUIREMENT_UNSPECIFIED
	CompanionAds_NONE                            = src.CompanionAds_NONE
	Event_ACCEPT_INVITATION                      = src.Event_ACCEPT_INVITATION
	Event_BREAK_END                              = src.Event_BREAK_END
	Event_BREAK_START                            = src.Event_BREAK_START
	Event_CLICK                                  = src.Event_CLICK
	Event_CLICK_THROUGH                          = src.Event_CLICK_THROUGH
	Event_CLOSE                                  = src.Event_CLOSE
	Event_CLOSE_LINEAR                           = src.Event_CLOSE_LINEAR
	Event_COLLAPSE                               = src.Event_COLLAPSE
	Event_COMPLETE                               = src.Event_COMPLETE
	Event_CREATIVE_VIEW                          = src.Event_CREATIVE_VIEW
	Event_ERROR                                  = src.Event_ERROR
	Event_EVENT_TYPE_UNSPECIFIED                 = src.Event_EVENT_TYPE_UNSPECIFIED
	Event_EXPAND                                 = src.Event_EXPAND
	Event_FIRST_QUARTILE                         = src.Event_FIRST_QUARTILE
	Event_IMPRESSION                             = src.Event_IMPRESSION
	Event_MIDPOINT                               = src.Event_MIDPOINT
	Event_MUTE                                   = src.Event_MUTE
	Event_PAUSE                                  = src.Event_PAUSE
	Event_PROGRESS                               = src.Event_PROGRESS
	Event_RESUME                                 = src.Event_RESUME
	Event_REWIND                                 = src.Event_REWIND
	Event_SKIP                                   = src.Event_SKIP
	Event_START                                  = src.Event_START
	Event_THIRD_QUARTILE                         = src.Event_THIRD_QUARTILE
	Event_UNMUTE                                 = src.Event_UNMUTE
	ManifestOptions_ASCENDING                    = src.ManifestOptions_ASCENDING
	ManifestOptions_DESCENDING                   = src.ManifestOptions_DESCENDING
	ManifestOptions_ORDER_POLICY_UNSPECIFIED     = src.ManifestOptions_ORDER_POLICY_UNSPECIFIED
)

// Deprecated: Please use vars in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
var (
	CompanionAds_DisplayRequirement_name                             = src.CompanionAds_DisplayRequirement_name
	CompanionAds_DisplayRequirement_value                            = src.CompanionAds_DisplayRequirement_value
	Event_EventType_name                                             = src.Event_EventType_name
	Event_EventType_value                                            = src.Event_EventType_value
	File_google_cloud_video_stitcher_v1_ad_tag_details_proto         = src.File_google_cloud_video_stitcher_v1_ad_tag_details_proto
	File_google_cloud_video_stitcher_v1_cdn_keys_proto               = src.File_google_cloud_video_stitcher_v1_cdn_keys_proto
	File_google_cloud_video_stitcher_v1_companions_proto             = src.File_google_cloud_video_stitcher_v1_companions_proto
	File_google_cloud_video_stitcher_v1_events_proto                 = src.File_google_cloud_video_stitcher_v1_events_proto
	File_google_cloud_video_stitcher_v1_sessions_proto               = src.File_google_cloud_video_stitcher_v1_sessions_proto
	File_google_cloud_video_stitcher_v1_slates_proto                 = src.File_google_cloud_video_stitcher_v1_slates_proto
	File_google_cloud_video_stitcher_v1_stitch_details_proto         = src.File_google_cloud_video_stitcher_v1_stitch_details_proto
	File_google_cloud_video_stitcher_v1_video_stitcher_service_proto = src.File_google_cloud_video_stitcher_v1_video_stitcher_service_proto
	ManifestOptions_OrderPolicy_name                                 = src.ManifestOptions_OrderPolicy_name
	ManifestOptions_OrderPolicy_value                                = src.ManifestOptions_OrderPolicy_value
)

// Details of an ad request to an ad server.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type AdRequest = src.AdRequest

// Metadata for a stitched ad.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type AdStitchDetail = src.AdStitchDetail

// Configuration for an Akamai CDN key.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type AkamaiCdnKey = src.AkamaiCdnKey

// Configuration for a CDN key. Used by the Video Stitcher to sign URIs for
// fetching video manifests and signing media segments for playback.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type CdnKey = src.CdnKey
type CdnKey_AkamaiCdnKey = src.CdnKey_AkamaiCdnKey
type CdnKey_GoogleCdnKey = src.CdnKey_GoogleCdnKey

// Metadata for a companion.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type Companion = src.Companion

// Metadata for companion ads.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type CompanionAds = src.CompanionAds

// Indicates how many of the companions should be displayed with the ad.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type CompanionAds_DisplayRequirement = src.CompanionAds_DisplayRequirement
type Companion_HtmlAdResource = src.Companion_HtmlAdResource
type Companion_IframeAdResource = src.Companion_IframeAdResource
type Companion_StaticAdResource = src.Companion_StaticAdResource

// Request message for VideoStitcherService.createCdnKey.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type CreateCdnKeyRequest = src.CreateCdnKeyRequest

// Request message for VideoStitcherService.createLiveSession.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type CreateLiveSessionRequest = src.CreateLiveSessionRequest

// Request message for VideoStitcherService.createSlate.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type CreateSlateRequest = src.CreateSlateRequest

// Request message for VideoStitcherService.createVodSession
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type CreateVodSessionRequest = src.CreateVodSessionRequest

// Request message for VideoStitcherService.deleteCdnKey.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type DeleteCdnKeyRequest = src.DeleteCdnKeyRequest

// Request message for VideoStitcherService.deleteSlate.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type DeleteSlateRequest = src.DeleteSlateRequest

// Describes an event and a trigger URI.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type Event = src.Event

// Describes the event that occurred.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type Event_EventType = src.Event_EventType

// Request message for VideoStitcherService.getCdnKey.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type GetCdnKeyRequest = src.GetCdnKeyRequest

// Request message for VideoStitcherService.getLiveAdTagDetail
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type GetLiveAdTagDetailRequest = src.GetLiveAdTagDetailRequest

// Request message for VideoStitcherService.getSession.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type GetLiveSessionRequest = src.GetLiveSessionRequest

// Request message for VideoStitcherService.getSlate.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type GetSlateRequest = src.GetSlateRequest

// Request message for VideoStitcherService.getVodAdTagDetail
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type GetVodAdTagDetailRequest = src.GetVodAdTagDetailRequest

// Request message for VideoStitcherService.getVodSession
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type GetVodSessionRequest = src.GetVodSessionRequest

// Request message for VideoStitcherService.getVodStitchDetail.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type GetVodStitchDetailRequest = src.GetVodStitchDetailRequest

// Configuration for a Google Cloud CDN key.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type GoogleCdnKey = src.GoogleCdnKey

// Metadata for an HTML ad resource.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type HtmlAdResource = src.HtmlAdResource

// Metadata for an IFrame ad resource.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type IframeAdResource = src.IframeAdResource

// Describes what was stitched into a VOD session's manifest.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type Interstitials = src.Interstitials

// Request message for VideoStitcherService.listCdnKeys.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type ListCdnKeysRequest = src.ListCdnKeysRequest

// Response message for VideoStitcher.ListCdnKeys.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type ListCdnKeysResponse = src.ListCdnKeysResponse

// Request message for VideoStitcherService.listLiveAdTagDetails.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type ListLiveAdTagDetailsRequest = src.ListLiveAdTagDetailsRequest

// Response message for VideoStitcherService.listLiveAdTagDetails.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type ListLiveAdTagDetailsResponse = src.ListLiveAdTagDetailsResponse

// Request message for VideoStitcherService.listSlates.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type ListSlatesRequest = src.ListSlatesRequest

// Response message for VideoStitcherService.listSlates.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type ListSlatesResponse = src.ListSlatesResponse

// Request message for VideoStitcherService.listVodAdTagDetails.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type ListVodAdTagDetailsRequest = src.ListVodAdTagDetailsRequest

// Response message for VideoStitcherService.listVodAdTagDetails.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type ListVodAdTagDetailsResponse = src.ListVodAdTagDetailsResponse

// Request message for VideoStitcherService.listVodStitchDetails.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type ListVodStitchDetailsRequest = src.ListVodStitchDetailsRequest

// Response message for VideoStitcherService.listVodStitchDetails.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type ListVodStitchDetailsResponse = src.ListVodStitchDetailsResponse

// Container for a live session's ad tag detail.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type LiveAdTagDetail = src.LiveAdTagDetail

// Metadata for a live session.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type LiveSession = src.LiveSession

// Options for manifest generation.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type ManifestOptions = src.ManifestOptions

// Defines the ordering policy during manifest generation.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type ManifestOptions_OrderPolicy = src.ManifestOptions_OrderPolicy

// Indicates a time in which a list of events should be triggered during media
// playback.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type ProgressEvent = src.ProgressEvent

// Filters for a video or muxed redition.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type RenditionFilter = src.RenditionFilter

// Metadata for an ad request.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type RequestMetadata = src.RequestMetadata

// Metadata for the response of an ad request.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type ResponseMetadata = src.ResponseMetadata

// Slate object
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type Slate = src.Slate

// Metadata for a static ad resource.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type StaticAdResource = src.StaticAdResource

// UnimplementedVideoStitcherServiceServer can be embedded to have forward
// compatible implementations.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type UnimplementedVideoStitcherServiceServer = src.UnimplementedVideoStitcherServiceServer

// Request message for VideoStitcherService.updateCdnKey.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type UpdateCdnKeyRequest = src.UpdateCdnKeyRequest

// Request message for VideoStitcherService.updateSlate.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type UpdateSlateRequest = src.UpdateSlateRequest

// VideoStitcherServiceClient is the client API for VideoStitcherService
// service. For semantics around ctx use and closing/ending streaming RPCs,
// please refer to
// https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type VideoStitcherServiceClient = src.VideoStitcherServiceClient

// VideoStitcherServiceServer is the server API for VideoStitcherService
// service.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type VideoStitcherServiceServer = src.VideoStitcherServiceServer

// Information related to the details for one ad tag.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type VodAdTagDetail = src.VodAdTagDetail

// Metadata for a VOD session.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type VodSession = src.VodSession

// Metadata for an inserted ad in a VOD session.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type VodSessionAd = src.VodSessionAd

// Metadata for an inserted ad break.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type VodSessionAdBreak = src.VodSessionAdBreak

// Metadata for the entire stitched content in a VOD session.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type VodSessionContent = src.VodSessionContent

// Detailed information related to the interstitial of a VOD session.
//
// Deprecated: Please use types in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
type VodStitchDetail = src.VodStitchDetail

// Deprecated: Please use funcs in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
func NewVideoStitcherServiceClient(cc grpc.ClientConnInterface) VideoStitcherServiceClient {
	return src.NewVideoStitcherServiceClient(cc)
}

// Deprecated: Please use funcs in: cloud.google.com/go/video/stitcher/apiv1/stitcherpb
func RegisterVideoStitcherServiceServer(s *grpc.Server, srv VideoStitcherServiceServer) {
	src.RegisterVideoStitcherServiceServer(s, srv)
}
