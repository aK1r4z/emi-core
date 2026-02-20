package api

import milky "github.com/aK1r4z/emi-core/types"

// 协议端使用的 QQ 协议平台，可能值：windows linux macos android_pad android_phone ipad iphone harmony watch
type QQProtocolType string

const (
	WINDOWS       QQProtocolType = "windows"
	LINUX         QQProtocolType = "linux"
	MACOS         QQProtocolType = "macos"
	ANDROID_PAD   QQProtocolType = "android_pad"
	ANDROID_PHONE QQProtocolType = "android_phone"
	I_PAD          QQProtocolType = "ipad"
	I_PHONE        QQProtocolType = "iphone"
	HARMONY       QQProtocolType = "harmony"
	WATCH         QQProtocolType = "watch"
)

type (
	GetLoginInfoRequest struct{}

	GetLoginInfoResponse struct {
		UIN      int64  `json:"uin"`
		Nickname string `json:"nickname"`
	}

	GetImplInfoRequest struct{}

	GetImplInfoResponse struct {
		ImplName          string         `json:"impl_name"`
		ImplVersion       string         `json:"impl_version"`
		QQProtocolVersion string         `json:"qq_protocol_version"`
		QQProtocolType    QQProtocolType `json:"qq_protocol_type"`
		MilkyVersion      string         `json:"milky_version"`
	}

	GetUserProfileRequest struct {
		UserID int64 `json:"user_id"`
	}

	GetUserProfileResponse struct {
		Nickname string        `json:"nickname"`
		QID      string        `json:"qid"`
		Age      int32         `json:"age"`
		Gender   milky.Gender `json:"sex"`
		Remark   string        `json:"remark"`
		Bio      string        `json:"bio"`
		Level    int32         `json:"level"`
		Country  string        `json:"country"`
		City     string        `json:"city"`
		School   string        `json:"school"`
	}

	GetFriendListRequest struct {
		NoCache bool `json:"no_cache"`
	}

	GetFriendListResponse struct {
		Friends []milky.FriendEntity `json:"friends"`
	}

	GetFriendInfoRequest struct {
		UserID  int64 `json:"user_id"`
		NoCache bool  `json:"no_cache"`
	}

	GetFriendInfoResponse struct {
		Friend milky.FriendEntity `json:"friend"`
	}

	GetGroupListRequest struct {
		NoCache bool `json:"no_cache"`
	}

	GetGroupListResponse struct {
		Groups []milky.GroupEntity `json:"groups"`
	}

	GetGroupInfoRequest struct {
		GroupID int64 `json:"group_id"`
		NoCache bool  `json:"no_cache"`
	}

	GetGroupInfoResponse struct {
		Group milky.GroupEntity `json:"group"`
	}

	GetGroupMemberListRequest struct {
		GroupID int64 `json:"group_id"`
		NoCache bool  `json:"no_cache"`
	}

	GetGroupMemberListResponse struct {
		Members []milky.GroupMemberEntity `json:"members"`
	}

	GetGroupMemberInfoRequest struct {
		GroupID int64 `json:"group_id"`
		UserID  int64 `json:"user_id"`
		NoCache bool  `json:"no_cache"`
	}

	GetGroupMemberInfoResponse struct {
		Member milky.GroupMemberEntity `json:"member"`
	}

	SetAvatarRequest struct {
		// 头像文件 URI，支持 file:// http(s):// base64:// 三种格式
		URI string `json:"uri"`
	}

	SetAvatarResponse struct{}

	SetNicknameRequest struct {
		NewNickname string `json:"new_nickname"`
	}

	SetNicknameResponse struct{}

	SetBioRequest struct {
		NewBio string `json:"new_bio"`
	}

	SetBioResponse struct{}

	GetCustomFaceURLListRequest struct{}

	GetCustomFaceURLListResponse struct {
		URLs []string `json:"urls"`
	}

	GetCookiesRequest struct {
		Domain string `json:"domain"`
	}

	GetCookiesResponse struct {
		Cookies string `json:"cookies"`
	}

	GetCSRFTokenRequest struct{}

	GetCSRFTokenResponse struct {
		Token string `json:"token"`
	}
)
