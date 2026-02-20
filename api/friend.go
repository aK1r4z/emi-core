package api

import (
	milky "github.com/aK1r4z/emi-core/types"
)

type (
	SendFriendNudgeRequest struct {
		UserID int64 `json:"user_id"`
		IsSelf bool  `json:"is_self"`
	}

	SendFriendNudgeResponse struct{}

	SendProfileLikeRequest struct {
		UserID int64 `json:"user_id"`
		Count  int32 `json:"count"`
	}

	SendProfileLikeResponse struct{}

	DeleteFriendRequest struct {
		UserID int64 `json:"user_id"`
	}

	DeleteFriendResponse struct{}

	GetFriendRequestsRequest struct {
		Limit      int32 `json:"limit"`
		IsFiltered bool  `json:"is_filtered"`
	}

	GetFriendRequestsResponse struct {
		Requests []milky.FriendRequest `json:"requests"`
	}

	AcceptFriendRequestRequest struct {
		InitiatorUID int64 `json:"initiator_uid"`
		IsFiltered   bool  `json:"is_filtered"`
	}

	AcceptFriendRequestResponse struct{}

	RejectFriendRequestRequest struct {
		InitiatorUID int64   `json:"initiator_uid"`
		IsFiltered   bool    `json:"is_filtered"`
		Reason       *string `json:"reason"`
	}

	RejectFriendRequestResponse struct{}
)
