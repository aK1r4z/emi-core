package api

import (
	milky "github.com/aK1r4z/emi-core/types"
)

type (
	// 发送好友戳一戳
	SendFriendNudgeRequest struct {
		UserID int64 `json:"user_id"` // 好友 QQ 号
		IsSelf bool  `json:"is_self"` // 是否戳自己，默认值：false
	}

	SendFriendNudgeResponse struct{}

	// 发送名片点赞
	SendProfileLikeRequest struct {
		UserID int64 `json:"user_id"` // 好友 QQ 号
		Count  int32 `json:"count"`   // 点赞数量，默认值：1
	}

	SendProfileLikeResponse struct{}

	// 删除好友
	DeleteFriendRequest struct {
		UserID int64 `json:"user_id"` // 好友 QQ 号
	}

	DeleteFriendResponse struct{}

	// 获取好友请求列表
	GetFriendRequestsRequest struct {
		Limit      int32 `json:"limit"`       // 获取的最大请求数量，默认值：20
		IsFiltered bool  `json:"is_filtered"` // true 表示只获取被过滤（由风险账号发起）的通知，false 表示只获取未被过滤的通知，默认值：false
	}

	GetFriendRequestsResponse struct {
		Requests []milky.FriendRequest `json:"requests"` // 好友请求列表
	}

	// 同意好友请求
	AcceptFriendRequestRequest struct {
		InitiatorUID int64 `json:"initiator_uid"` // 请求发起者 UID
		IsFiltered   bool  `json:"is_filtered"`   // 是否是被过滤的请求，默认值：false
	}

	AcceptFriendRequestResponse struct{}

	// 拒绝好友请求
	RejectFriendRequestRequest struct {
		InitiatorUID int64   `json:"initiator_uid"` // 请求发起者 UID
		IsFiltered   bool    `json:"is_filtered"`   // 是否是被过滤的请求，默认值：false
		Reason       *string `json:"reason"`        // 拒绝理由
	}

	RejectFriendRequestResponse struct{}
)
