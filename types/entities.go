package types

import (
	"encoding/json"
)

type (
	Gender                string
	GroupRole             string
	JoinRequestState      string
	GroupNotificationType string
	MessageScene          string
)

const (
	GenderMale    Gender = "male"
	GenderFemale  Gender = "female"
	GenderUnknown Gender = "unknown"

	GroupOwner  GroupRole = "owner"
	GroupAdmin  GroupRole = "admin"
	GroupMember GroupRole = "member"

	JoinRequestPending  JoinRequestState = "pending"
	JoinRequestAccepted JoinRequestState = "accepted"
	JoinRequestRejected JoinRequestState = "rejected"
	JoinRequestIgnored  JoinRequestState = "ignored"

	GroupJoinRequestNotify        GroupNotificationType = "join_request"
	GroupAdminChangeNotify        GroupNotificationType = "admin_change"
	GroupKickNotify               GroupNotificationType = "kick"
	GroupQuitNotify               GroupNotificationType = "quit"
	GroupInvitedJoinRequestNotify GroupNotificationType = "invited_join_request"

	MessageSceneFriend MessageScene = "friend"
	MessageSceneGroup  MessageScene = "group"
	MessageSceneTemp   MessageScene = "temp"
)

type (
	FriendEntity struct {
		UserID   int64                `json:"user_id"`
		Nickname string               `json:"nickname"`
		Gender   Gender               `json:"sex"`
		QID      string               `json:"qid"`
		Remark   string               `json:"remark"`
		Category FriendCategoryEntity `json:"category"`
	}

	FriendCategoryEntity struct {
		CategoryID   int64  `json:"category_id"`
		CategoryName string `json:"category_name"`
	}

	GroupEntity struct {
		GroupID        int64  `json:"group_id"`
		GroupName      string `json:"group_name"`
		MemberCount    int32  `json:"member_count"`
		MaxMemberCount int32  `json:"max_member_count"`
	}

	GroupMemberEntity struct {
		UserID        int64     `json:"user_id"`
		Nickname      string    `json:"nickname"`
		Gender        Gender    `json:"sex"`
		GroupID       int64     `json:"group_id"`
		Card          string    `json:"card"`
		Title         string    `json:"title"`
		Level         int32     `json:"level"`
		Role          GroupRole `json:"role"`
		JoinTime      int64     `json:"join_time"`
		LastSentTime  int64     `json:"last_sent_time"`
		ShutUpEndTime *int64    `json:"shut_up_end_time"`
	}

	GroupAnnouncementEntity struct {
		GroupID        int64   `json:"group_id"`
		AnnouncementID string  `json:"announcement_id"`
		UserID         int64   `json:"user_id"`
		Time           int64   `json:"time"`
		Content        string  `json:"content"`
		ImageURL       *string `json:"image_url"`
	}

	GroupFileEntity struct {
		GroupID         int64  `json:"group_id"`
		FileID          string `json:"file_id"`
		FileName        string `json:"file_name"`
		ParentFolderID  string `json:"parent_folder_id"`
		FileSize        int64  `json:"file_size"`
		UploadedTime    int64  `json:"uploaded_time"`
		ExpireTime      *int64 `json:"expire_time"`
		UploaderID      int64  `json:"uploader_id"`
		DownloadedTimes int32  `json:"downloaded_times"`
	}

	GroupFolderEntity struct {
		GroupID          int64  `json:"group_id"`
		FolderID         string `json:"folder_id"`
		ParentFolderID   string `json:"parent_folder_id"`
		FolderName       string `json:"folder_name"`
		CreatedTime      int64  `json:"created_time"`
		LastModifiedTime int64  `json:"last_modified_time"`
		CreatorID        int64  `json:"creator_id"`
		FileCount        int32  `json:"file_count"`
	}

	FriendRequest struct {
		Time          int64            `json:"time"`
		InitiatorID   int64            `json:"initiator_id"`
		InitiatorUID  string           `json:"initiator_uid"`
		TargetUserID  int64            `json:"target_user_id"`
		TargetUserUID string           `json:"target_user_uid"`
		State         JoinRequestState `json:"state"`
		Comment       string           `json:"comment"`
		Via           string           `json:"via"`
		IsFiltered    bool             `json:"is_filtered"`
	}

	GroupNotification struct {
		Type            GroupNotificationType `json:"type"`
		GroupID         int64                 `json:"group_id"`
		NotificationSeq int64                 `json:"notification_seq"`

		IsFiltered   *bool             `json:"is_filtered"`    // JoinRequest,
		Comment      *string           `json:"comment"`        // JoinRequest,
		State        *JoinRequestState `json:"state"`          // JoinRequest, InvitedJoinRequest,
		InitiatorID  *int64            `json:"initiator_id"`   // JoinRequest, InvitedJoinRequest,
		OperatorID   *int64            `json:"operator_id"`    // JoinRequest, AdminChange, Kick, InvitedJoinRequest,
		TargetUserID *int64            `json:"target_user_id"` // AdminChange, Kick, Quit, InvitedJoinRequest,
		IsSet        *bool             `json:"is_set"`         // AdminChange,
	}

	IncomingMessage struct {
		MessageScene MessageScene      `json:"message_scene"`
		MessageSeq   int64             `json:"message_seq"`
		SenderID     int64             `json:"sender_id"`
		PeerID       int64             `json:"peer_id"`
		Time         int64             `json:"time"`
		Segments     []IncomingSegment `json:"segments"`

		Friend      *FriendEntity      `json:"friend"`
		Group       *GroupEntity       `json:"group"`
		GroupMember *GroupMemberEntity `json:"group_member"`
	}

	IncomingForwardedMessage struct {
		SenderName string            `json:"sender_name"`
		AvatarURL  string            `json:"avatar_url"`
		Time       int64             `json:"time"`
		Segments   []IncomingSegment `json:"segments"`
	}

	GroupEssenceMessage struct {
		GroupID       int64             `json:"group_id"`
		MessageSeq    int64             `json:"message_seq"`
		MessageTime   int64             `json:"message_time"`
		SenderID      int64             `json:"sender_id"`
		SenderName    string            `json:"sender_name"`
		OperatorID    int64             `json:"operator_id"`
		OperatorName  string            `json:"operator_name"`
		OperationTime int64             `json:"operation_time"`
		Segments      []IncomingSegment `json:"segments"`
	}

	IncomingSegment struct {
		Type SegmentType     `json:"type"`
		Data json.RawMessage `json:"data"`
	}

	OutgoingForwardedMessage struct {
		UserID     int64             `json:"user_id"`
		SenderName string            `json:"sender_name"`
		Segments   []OutgoingSegment `json:"segments"`
	}

	OutgoingSegment struct {
		Type SegmentType     `json:"type"`
		Data json.RawMessage `json:"data"`
	}
)
