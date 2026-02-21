package api

import milky_types "github.com/aK1r4z/emi-core/types"

// 协议端使用的 QQ 协议平台，可能值：windows linux macos android_pad android_phone ipad iphone harmony watch
type QQProtocolType string

const (
	WINDOWS       QQProtocolType = "windows"
	LINUX         QQProtocolType = "linux"
	MACOS         QQProtocolType = "macos"
	ANDROID_PAD   QQProtocolType = "android_pad"
	ANDROID_PHONE QQProtocolType = "android_phone"
	I_PAD         QQProtocolType = "ipad"
	I_PHONE       QQProtocolType = "iphone"
	HARMONY       QQProtocolType = "harmony"
	WATCH         QQProtocolType = "watch"
)

type (
	// 获取登录信息
	GetLoginInfoRequest struct{}

	GetLoginInfoResponse struct {
		UIN      int64  `json:"uin"`      // 登录 QQ 号
		Nickname string `json:"nickname"` // 登录昵称
	}

	// 获取协议端信息
	GetImplInfoRequest struct{}

	GetImplInfoResponse struct {
		ImplName          string         `json:"impl_name"`           // 协议端名称
		ImplVersion       string         `json:"impl_version"`        // 协议端版本
		QQProtocolVersion string         `json:"qq_protocol_version"` // 协议端使用的 QQ 协议版本
		QQProtocolType    QQProtocolType `json:"qq_protocol_type"`    // 协议端使用的 QQ 协议平台，可能值：windows linux macos android_pad android_phone ipad iphone harmony watch
		MilkyVersion      string         `json:"milky_version"`       // 协议端实现的 Milky 协议版本，目前为 "1.1"
	}

	// 获取用户个人信息
	GetUserProfileRequest struct {
		UserID int64 `json:"user_id"` // 用户 QQ 号
	}

	GetUserProfileResponse struct {
		Nickname string       `json:"nickname"` // 昵称
		QID      string       `json:"qid"`      // QID
		Age      int32        `json:"age"`      // 年龄
		Gender   milky_types.Gender `json:"sex"`      // 性别，可能值：male female unknown
		Remark   string       `json:"remark"`   // 备注
		Bio      string       `json:"bio"`      // 个性签名
		Level    int32        `json:"level"`    // QQ 等级
		Country  string       `json:"country"`  // 国家或地区
		City     string       `json:"city"`     // 城市
		School   string       `json:"school"`   // 学校
	}

	// 获取好友列表
	GetFriendListRequest struct {
		NoCache bool `json:"no_cache"` // 是否强制不使用缓存，默认值：false
	}

	GetFriendListResponse struct {
		Friends []milky_types.FriendEntity `json:"friends"` // 好友列表
	}

	// 获取好友信息
	GetFriendInfoRequest struct {
		UserID  int64 `json:"user_id"`  // 好友 QQ 号
		NoCache bool  `json:"no_cache"` // 是否强制不使用缓存，默认值：false
	}

	GetFriendInfoResponse struct {
		Friend milky_types.FriendEntity `json:"friend"` // 好友信息
	}

	// 获取群列表
	GetGroupListRequest struct {
		NoCache bool `json:"no_cache"` // 是否强制不使用缓存，默认值：false
	}

	GetGroupListResponse struct {
		Groups []milky_types.GroupEntity `json:"groups"` // 群列表
	}

	// 获取群信息
	GetGroupInfoRequest struct {
		GroupID int64 `json:"group_id"` // 群号
		NoCache bool  `json:"no_cache"` // 是否强制不使用缓存，默认值：false
	}

	GetGroupInfoResponse struct {
		Group milky_types.GroupEntity `json:"group"` // 群信息
	}

	// 获取群成员列表
	GetGroupMemberListRequest struct {
		GroupID int64 `json:"group_id"` // 群号
		NoCache bool  `json:"no_cache"` // 是否强制不使用缓存，默认值：false
	}

	GetGroupMemberListResponse struct {
		Members []milky_types.GroupMemberEntity `json:"members"` // 群成员列表
	}

	// 获取群成员信息
	GetGroupMemberInfoRequest struct {
		GroupID int64 `json:"group_id"` // 群号
		UserID  int64 `json:"user_id"`  // 群成员 QQ 号
		NoCache bool  `json:"no_cache"` // 是否强制不使用缓存，默认值：false
	}

	GetGroupMemberInfoResponse struct {
		Member milky_types.GroupMemberEntity `json:"member"` // 群成员信息
	}

	// 设置 QQ 账号头像
	SetAvatarRequest struct {
		URI string `json:"uri"` // 头像文件 URI，支持 file:// http(s):// base64:// 三种格式
	}

	SetAvatarResponse struct{}

	// 设置 QQ 账号昵称
	SetNicknameRequest struct {
		NewNickname string `json:"new_nickname"` // 新昵称
	}

	SetNicknameResponse struct{}

	// 设置 QQ 账号个性签名
	SetBioRequest struct {
		NewBio string `json:"new_bio"` // 新个性签名
	}

	SetBioResponse struct{}

	// 获取自定义表情 URL 列表
	GetCustomFaceURLListRequest struct{}

	GetCustomFaceURLListResponse struct {
		URLs []string `json:"urls"` // 自定义表情 URL 列表
	}

	// 获取 Cookies
	GetCookiesRequest struct {
		Domain string `json:"domain"` // 需要获取 Cookies 的域名
	}

	GetCookiesResponse struct {
		Cookies string `json:"cookies"` // 域名对应的 Cookies 字符串
	}

	// 获取 CSRF Token
	GetCSRFTokenRequest struct{}

	GetCSRFTokenResponse struct {
		CSRFToken string `json:"csrf_token"` // CSRF Token
	}
)
