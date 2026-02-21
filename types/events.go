package types

import (
	"encoding/json"
	"errors"
)

var ErrInvalidEventType = errors.New("invalid event type")

type EventType string

const (
	EventBotOffline                EventType = "bot_offline"
	EventMessageReceive            EventType = "message_receive"
	EventMessageRecall             EventType = "message_recall"
	EventFriendRequest             EventType = "friend_request"
	EventGroupJoinRequest          EventType = "group_join_request"
	EventGroupInvitedJoinRequest   EventType = "group_invited_join_request"
	EventGroupInvitation           EventType = "group_invitation"
	EventFriendNudge               EventType = "friend_nudge"
	EventFriendFileUpload          EventType = "friend_file_upload"
	EventGroupAdminChange          EventType = "group_admin_change"
	EventGroupEssenceMessageChange EventType = "group_essence_message_change"
	EventGroupMemberIncrease       EventType = "group_member_increase"
	EventGroupMemberDecrease       EventType = "group_member_decrease"
	EventGroupNameChange           EventType = "group_name_change"
	EventGroupMessageReaction      EventType = "group_message_reaction"
	EventGroupMute                 EventType = "group_mute"
	EventGroupWholeMute            EventType = "group_whole_mute"
	EventGroupNudge                EventType = "group_nudge"
	EventGroupFileUpload           EventType = "group_file_upload"
)

func (e *EventType) IsValid() bool {
	switch *e {
	case
		EventFriendRequest,
		EventGroupJoinRequest,
		EventGroupInvitedJoinRequest,
		EventGroupInvitation,
		EventFriendNudge,
		EventFriendFileUpload,
		EventGroupAdminChange,
		EventGroupEssenceMessageChange,
		EventGroupMemberIncrease,
		EventGroupMemberDecrease,
		EventGroupNameChange,
		EventGroupMessageReaction,
		EventGroupMute,
		EventGroupWholeMute,
		EventGroupNudge,
		EventGroupFileUpload:
		return true
	}
	return false
}

func FromString(s string) (EventType, error) {
	e := EventType(s)
	if !e.IsValid() {
		return "", ErrInvalidEventType
	}
	return e, nil
}

type EventDescriber interface {
	Type() EventType
}

type Event interface {
	EventDescriber
	New() Event
}

type RawEvent struct {
	Type   EventType       `json:"event_type"`
	Time   int64           `json:"time"`
	SelfID int64           `json:"self_id"`
	Data   json.RawMessage `json:"data"`
}

type BotOfflineEvent struct {
	Reason string `json:"reason"`
}

func (e *BotOfflineEvent) Type() EventType { return EventBotOffline }
func (e *BotOfflineEvent) New() Event      { return &BotOfflineEvent{} }

type MessageReceiveEvent struct {
	IncomingMessage
}

func (e *MessageReceiveEvent) Type() EventType { return EventMessageReceive }
func (e *MessageReceiveEvent) New() Event      { return &MessageReceiveEvent{} }

type MessageRecallEvent struct {
	MessageScene  MessageScene `json:"message_scene"`
	PeerID        int64        `json:"peer_id"`
	MessageSeq    int64        `json:"message_seq"`
	SenderID      int64        `json:"sender_id"`
	OperatorID    int64        `json:"operator_id"`
	DisplaySuffix string       `json:"display_suffix"`
}

func (e *MessageRecallEvent) Type() EventType { return EventMessageRecall }
func (e *MessageRecallEvent) New() Event      { return &MessageRecallEvent{} }

type FriendRequestEvent struct {
	InitiatorID  int64  `json:"initiator_id"`
	InitiatorUID string `json:"initiator_uid"`
	Comment      string `json:"comment"`
	Via          string `json:"via"`
}

func (e *FriendRequestEvent) Type() EventType { return EventFriendRequest }
func (e *FriendRequestEvent) New() Event      { return &FriendRequestEvent{} }

type GroupJoinRequestEvent struct {
	GroupID         int64  `json:"group_id"`
	NotificationSeq int64  `json:"notification_seq"`
	IsFiltered      bool   `json:"is_filtered"`
	InitiatorID     int64  `json:"initiator_id"`
	Comment         string `json:"comment"`
}

func (e *GroupJoinRequestEvent) Type() EventType { return EventGroupJoinRequest }
func (e *GroupJoinRequestEvent) New() Event      { return &GroupJoinRequestEvent{} }

type GroupInvitedJoinRequestEvent struct {
	GroupID         int64 `json:"group_id"`
	NotificationSeq int64 `json:"notification_seq"`
	InitiatorID     int64 `json:"initiator_id"`
	TargetUserID    int64 `json:"target_user_id"`
}

func (e *GroupInvitedJoinRequestEvent) Type() EventType { return EventGroupInvitedJoinRequest }
func (e *GroupInvitedJoinRequestEvent) New() Event      { return &GroupInvitedJoinRequestEvent{} }

type GroupInvitationEvent struct {
	GroupID       int64 `json:"group_id"`
	InvitationSeq int64 `json:"invitation_seq"`
	InitiatorID   int64 `json:"initiator_id"`
}

func (e *GroupInvitationEvent) Type() EventType { return EventGroupInvitation }
func (e *GroupInvitationEvent) New() Event      { return &GroupInvitationEvent{} }

type FriendNudgeEvent struct {
	UserID                int64  `json:"user_id"`
	IsSelfSend            bool   `json:"is_self_send"`
	IsSelfReceive         bool   `json:"is_self_receive"`
	DisplayAction         string `json:"display_action"`
	DisplaySuffix         string `json:"display_suffix"`
	DisplayActionImageURL string `json:"display_action_img_url"`
}

func (e *FriendNudgeEvent) Type() EventType { return EventFriendNudge }
func (e *FriendNudgeEvent) New() Event      { return &FriendNudgeEvent{} }

type FriendFileUploadEvent struct {
	UserID   int64  `json:"user_id"`
	FileID   string `json:"file_id"`
	FileName string `json:"file_name"`
	FileSize int64  `json:"file_size"`
	FileHash string `json:"file_hash"` // 文件的 TriSHA1 哈希值
	IsSelf   bool   `json:"is_self"`
}

func (e *FriendFileUploadEvent) Type() EventType { return EventFriendFileUpload }
func (e *FriendFileUploadEvent) New() Event      { return &FriendFileUploadEvent{} }

type GroupAdminChangeEvent struct {
	GroupID    int64 `json:"group_id"`
	UserID     int64 `json:"user_id"`
	OperatorID int64 `json:"operator_id"`
	IsSet      bool  `json:"is_set"`
}

func (e *GroupAdminChangeEvent) Type() EventType { return EventGroupAdminChange }
func (e *GroupAdminChangeEvent) New() Event      { return &GroupAdminChangeEvent{} }

type GroupEssenceMessageChangeEvent struct {
	GroupID    int64 `json:"group_id"`
	MessageSeq int64 `json:"message_seq"`
	OperatorID int64 `json:"operator_id"`
	IsSet      bool  `json:"is_set"`
}

func (e *GroupEssenceMessageChangeEvent) Type() EventType { return EventGroupEssenceMessageChange }
func (e *GroupEssenceMessageChangeEvent) New() Event      { return &GroupEssenceMessageChangeEvent{} }

type GroupMemberIncreaseEvent struct {
	GroupID    int64  `json:"group_id"`
	UserID     int64  `json:"user_id"`
	OperatorID *int64 `json:"operator_id"` // 管理员 QQ 号，如果是管理员同意入群
	InvitorID  *int64 `json:"invitor_id"`  // 邀请者 QQ 号，如果是邀请入群
}

func (e *GroupMemberIncreaseEvent) Type() EventType { return EventGroupMemberIncrease }
func (e *GroupMemberIncreaseEvent) New() Event      { return &GroupMemberIncreaseEvent{} }

type GroupMemberDecreaseEvent struct {
	GroupID    int64  `json:"group_id"`
	UserID     int64  `json:"user_id"`
	OperatorID *int64 `json:"operator_id"`
}

func (e *GroupMemberDecreaseEvent) Type() EventType { return EventGroupMemberDecrease }
func (e *GroupMemberDecreaseEvent) New() Event      { return &GroupMemberDecreaseEvent{} }

type GroupNameChangeEvent struct {
	GroupId      int64  `json:"group_id"`
	NewGroupName string `json:"new_group_name"`
	OperatorID   int64  `json:"operator_id"`
}

func (e *GroupNameChangeEvent) Type() EventType { return EventGroupNameChange }
func (e *GroupNameChangeEvent) New() Event      { return &GroupNameChangeEvent{} }

type GroupMessageReactionEvent struct {
	GroupID    int64  `json:"group_id"`
	UserID     int64  `json:"user_id"`
	MessageSeq int64  `json:"message_seq"`
	FaceID     string `json:"face_id"`
	IsAdd      bool   `json:"is_add"`
}

func (e *GroupMessageReactionEvent) Type() EventType { return EventGroupMessageReaction }
func (e *GroupMessageReactionEvent) New() Event      { return &GroupMessageReactionEvent{} }

type GroupMuteEvent struct {
	GroupID    int64 `json:"group_id"`
	UserID     int64 `json:"user_id"`
	OperatorID int64 `json:"operator_id"`
	Duration   int32 `json:"duration"` // 禁言时长（秒），为 0 表示取消禁言
}

func (e *GroupMuteEvent) Type() EventType { return EventGroupMute }
func (e *GroupMuteEvent) New() Event      { return &GroupMuteEvent{} }

type GroupWholeMuteEvent struct {
	GroupId    int64 `json:"group_id"`
	OperatorID int64 `json:"operator_id"`
	IsMute     bool  `json:"is_mute"`
}

func (e *GroupWholeMuteEvent) Type() EventType { return EventGroupWholeMute }
func (e *GroupWholeMuteEvent) New() Event      { return &GroupWholeMuteEvent{} }

type GroupNudgeEvent struct {
	GroupID               int64  `json:"group_id"`
	SenderID              int64  `json:"sender_id"`
	ReceiverID            int64  `json:"receiver_id"`
	DisplayAction         string `json:"display_action"`
	DisplaySuffix         string `json:"display_suffix"`
	DisplayActionImageURL string `json:"display_action_img_url"`
}

func (e *GroupNudgeEvent) Type() EventType { return EventGroupNudge }
func (e *GroupNudgeEvent) New() Event      { return &GroupNudgeEvent{} }

type GroupFileUploadEvent struct {
	GroupID  int64  `json:"group_id"`
	UserID   int64  `json:"user_id"`
	FileID   string `json:"file_id"`
	FileName string `json:"file_name"`
	FileSize int64  `json:"file_size"`
}

func (e *GroupFileUploadEvent) Type() EventType { return EventGroupFileUpload }
func (e *GroupFileUploadEvent) New() Event      { return &GroupFileUploadEvent{} }
