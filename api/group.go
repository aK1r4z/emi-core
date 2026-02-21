package api

import milky "github.com/aK1r4z/emi-core/types"

// 请求对应的通知类型，可能值：join_request invited_join_request
type GroupRequestNotificationType string

const (
	JoinRequestNotify        GroupRequestNotificationType = "join_request"
	InvitedJoinRequestNotify GroupRequestNotificationType = "invited_join_request"
)

type (
	// 设置群名称
	SetGroupNameRequest struct {
		GroupID      int64  `json:"group_id"`       // 群号
		NewGroupName string `json:"new_group_name"` // 新群名称
	}

	SetGroupNameResponse struct{}

	// 设置群头像
	SetGroupAvatarRequest struct {
		GroupID  int64  `json:"group_id"`  // 群号
		ImageURI string `json:"image_uri"` // 头像文件 URI，支持 file:// http(s):// base64:// 三种格式
	}

	SetGroupAvatarResponse struct{}

	// 设置群名片
	SetGroupMemberCardRequest struct {
		GroupID int64  `json:"group_id"` // 群号
		UserID  int64  `json:"user_id"`  // 被设置的群成员 QQ 号
		Card    string `json:"card"`     // 新群名片
	}

	SetGroupMemberCardResponse struct{}

	// 设置群成员专属头衔
	SetGroupMemberSpecialTitleRequest struct {
		GroupID      int64  `json:"group_id"`      // 群号
		UserID       int64  `json:"user_id"`       // 被设置的群成员 QQ 号
		SpecialTitle string `json:"special_title"` // 新专属头衔
	}

	SetGroupMemberSpecialTitleResponse struct{}

	// 设置群管理员
	SetGroupMemberAdminRequest struct {
		GroupID int64 `json:"group_id"` // 群号
		UserID  int64 `json:"user_id"`  // 被设置的 QQ 号
		IsSet   bool  `json:"is_set"`   // 是否设置为管理员，false 表示取消管理员，默认值：true
	}

	SetGroupMemberAdminResponse struct{}

	// 设置群成员禁言
	SetGroupMemberMuteRequest struct {
		GroupID  int64 `json:"group_id"` // 群号
		UserID   int64 `json:"user_id"`  // 被设置的 QQ 号
		Duration int32 `json:"duration"` // 禁言持续时间（秒），设为 0 为取消禁言，默认值：0
	}

	SetGroupMemberMuteResponse struct{}

	// 设置群全员禁言
	SetGroupMemberWholeMuteRequest struct {
		GroupID int64 `json:"group_id"` // 群号
		IsMute  bool  `json:"is_mute"`  // 是否开启全员禁言，false 表示取消全员禁言，默认值：true
	}

	SetGroupMemberWholeMuteResponse struct{}

	// 踢出群成员
	KickGroupMemberRequest struct {
		GroupID          int64 `json:"group_id"`           // 群号
		UserID           int64 `json:"user_id"`            // 被踢的 QQ 号
		RejectAddRequest bool  `json:"reject_add_request"` // 是否拒绝加群申请，false 表示不拒绝，默认值：false
	}

	KickGroupMemberResponse struct{}

	GetGroupAnnouncementsRequest struct {
		GroupID int64 `json:"group_id"` // 群号
	}

	GetGroupAnnouncementsResponse struct {
		Announcements []milky.GroupAnnouncementEntity `json:"announcements"` // 群公告列表
	}

	// 发送群公告
	SendGroupAnnouncementRequest struct {
		GroupID  int64   `json:"group_id"`  // 群号
		Content  string  `json:"content"`   // 公告内容
		ImageURI *string `json:"image_uri"` // 公告附带图像文件 URI，支持 file:// http(s):// base64:// 三种格式
	}

	SendGroupAnnouncementResponse struct{}

	// 删除群公告
	DeleteGroupAnnouncementRequest struct {
		GroupID        int64  `json:"group_id"`        // 群号
		AnnouncementID string `json:"announcement_id"` // 公告 ID
	}

	DeleteGroupAnnouncementResponse struct{}

	// 获取群精华消息列表
	GetGroupEssenceMessagesRequest struct {
		GroupID   int64 `json:"group_id"`   // 群号
		PageIndex int32 `json:"page_index"` // 页码索引，从 0 开始
		PageSize  int32 `json:"page_size"`  // 每页包含的精华消息数量
	}

	GetGroupEssenceMessagesResponse struct {
		Messages []milky.GroupEssenceMessage `json:"messages"` // 精华消息列表
		IsEnd    bool                        `json:"is_end"`   // 是否已到最后一页
	}

	// 设置群精华消息
	SetGroupEssenceMessageRequest struct {
		GroupID    int64 `json:"group_id"`    // 群号
		MessageSeq int64 `json:"message_seq"` // 消息序列号
		IsSet      bool  `json:"is_set"`      // 是否设置为精华消息，false 表示取消精华，默认值：true
	}

	SetGroupEssenceMessageResponse struct{}

	// 退出群
	QuitGroupRequest struct {
		GroupID int64 `json:"group_id"` // 群号
	}

	QuitGroupResponse struct{}

	// 发送群消息表情回应
	SendGroupMessageReactionRequest struct {
		GroupID    int64  `json:"group_id"`    // 群号
		MessageSeq int64  `json:"message_seq"` // 要回应的消息序列号
		Reaction   string `json:"reaction"`    // 表情 ID
		IsAdd      bool   `json:"is_add"`      // 是否添加表情，false 表示取消，默认值：true
	}

	SendGroupMessageReactionResponse struct{}

	// 发送群戳一戳
	SendGroupNudgeRequest struct {
		GroupID int64 `json:"group_id"` // 群号
		UserID  int64 `json:"user_id"`  // 被戳的群成员 QQ 号
	}

	SendGroupNudgeResponse struct{}

	// 获取群通知列表
	GetGroupNotificationsRequest struct {
		StartNotificationSeq *int64 `json:"start_notification_seq"` // 起始通知序列号
		IsFiltered           bool   `json:"is_filtered"`            // true 表示只获取被过滤（由风险账号发起）的通知，false 表示只获取未被过滤的通知，默认值：false
		Limit                int32  `json:"limit"`                  // 获取的最大通知数量，默认值：20
	}

	GetGroupNotificationsResponse struct {
		Notifications       []milky.GroupNotification // 获取到的群通知（notification_seq 降序排列），序列号不一定连续
		NextNotificationSeq *int64                    // 下一页起始通知序列号
	}

	// 同意入群/邀请他人入群请求
	AcceptGroupRequestRequest struct {
		NotificationSeq  int64                        `json:"notification_seq"`  // 请求对应的通知序列号
		NotificationType GroupRequestNotificationType `json:"notification_type"` // 请求对应的通知类型，可能值：join_request invited_join_request
		GroupID          int64                        `json:"group_id"`          // 请求所在的群号
		IsFiltered       bool                         `json:"is_filtered"`       // 是否是被过滤的请求，默认值：false
	}

	AcceptGroupRequestResponse struct{}

	// 拒绝入群/邀请他人入群请求
	RejectGroupRequestRequest struct {
		NotificationSeq  int64                        `json:"notification_seq"`  // 请求对应的通知序列号
		NotificationType GroupRequestNotificationType `json:"notification_type"` // 请求对应的通知类型，可能值：join_request invited_join_request
		GroupID          int64                        `json:"group_id"`          // 请求所在的群号
		IsFiltered       bool                         `json:"is_filtered"`       // 是否是被过滤的请求，默认值：false
		Reason           *string                      `json:"reason"`            // 拒绝理由
	}

	RejectGroupRequestResponse struct{}

	// 同意他人邀请自身入群
	AcceptGroupInvitationRequest struct {
		GroupID       int64 `json:"group_id"`       // 群号
		InvitationSeq int64 `json:"invitation_seq"` // 邀请序列号
	}

	AcceptGroupInvitationResponse struct{}

	// 拒绝他人邀请自身入群
	RejectGroupInvitationRequest struct {
		GroupID       int64 `json:"group_id"`       // 群号
		InvitationSeq int64 `json:"invitation_seq"` // 邀请序列号
	}

	RejectGroupInvitationResponse struct{}
)
