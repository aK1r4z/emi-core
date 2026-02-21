package emi_core

type Endpoint string

const (

	// SystemAPI

	GetLoginInfo         Endpoint = "get_login_info"           // 获取登录信息
	GetImplInfo          Endpoint = "get_impl_info"            // 获取协议端信息
	GetUserProfile       Endpoint = "get_user_profile"         // 获取用户个人信息
	GetFriendList        Endpoint = "get_friend_list"          // 获取好友列表
	GetFriendInfo        Endpoint = "get_friend_info"          // 获取好友信息
	GetGroupList         Endpoint = "get_group_list"           // 获取群列表
	GetGroupInfo         Endpoint = "get_group_info"           // 获取群信息
	GetGroupMemberList   Endpoint = "get_group_member_list"    // 获取群成员列表
	GetGroupMemberInfo   Endpoint = "get_group_member_info"    // 获取群成员信息
	SetAvatar            Endpoint = "set_avatar"               // 设置 QQ 账号头像
	SetNickname          Endpoint = "set_nickname"             // 设置 QQ 账号昵称
	SetBio               Endpoint = "set_bio"                  // 设置 QQ 账号个性签名
	GetCustomFaceURLList Endpoint = "get_custom_face_url_list" // 获取自定义表情 URL 列表
	GetCookies           Endpoint = "get_cookies"              // 获取 Cookies
	GetCSRFToken         Endpoint = "get_csrf_token"           // 获取 CSRF Token

	// MessageAPI

	SendPrivateMessage   Endpoint = "send_private_message"   // 发送私聊消息
	SendGroupMessage     Endpoint = "send_group_message"     // 发送群聊消息
	RecallPrivateMessage Endpoint = "recall_private_message" // 撤回私聊消息
	RecallGroupMessage   Endpoint = "recall_group_message"   // 撤回群聊消息
	GetMessage           Endpoint = "get_message"            // 获取消息
	GetHistoryMessages   Endpoint = "get_history_messages"   // 获取历史消息列表
	GetResourceTempURL   Endpoint = "get_resource_temp_url"  // 获取临时资源链接
	GetForwardedMessages Endpoint = "get_forwarded_messages" // 获取合并转发消息内容
	MarkMessageAsRead    Endpoint = "mark_message_as_read"   // 标记消息为已读

	// FriendAPI

	SendFriendNudge     Endpoint = "send_friend_nudge"     // 发送好友戳一戳
	SendProfileLike     Endpoint = "send_profile_like"     // 发送名片点赞
	DeleteFriend        Endpoint = "delete_friend"         // 删除好友
	GetFriendRequests   Endpoint = "get_friend_requests"   // 获取好友请求列表
	AcceptFriendRequest Endpoint = "accept_friend_request" // 同意好友请求
	RejectFriendRequest Endpoint = "reject_friend_request" // 拒绝好友请求

	// GroupAPI

	SetGroupName               Endpoint = "set_group_name"                 // 设置群名称
	SetGroupAvatar             Endpoint = "set_group_avatar"               // 设置群头像
	SetGroupMemberCard         Endpoint = "set_group_member_card"          // 设置群名片
	SetGroupMemberSpecialTitle Endpoint = "set_group_member_special_title" // 设置群成员专属头衔
	SetGroupMemberAdmin        Endpoint = "set_group_member_admin"         // 设置群管理员
	SetGroupMemberMute         Endpoint = "set_group_member_mute"          // 设置群成员禁言
	SetGroupMemberWholeMute    Endpoint = "set_group_whole_mute"           // 设置群全员禁言
	KickGroupMember            Endpoint = "kick_group_member"              // 踢出群成员
	GetGroupAnnouncements      Endpoint = "get_group_announcements"        // 获取群公告列表
	SendGroupAnnouncement      Endpoint = "send_group_announcement"        // 发送群公告
	DeleteGroupAnnouncement    Endpoint = "delete_group_announcement"      // 删除群公告
	GetGroupEssenceMessages    Endpoint = "get_group_essence_messages"     // 获取群精华消息列表
	SetGroupEssenceMessage     Endpoint = "set_group_essence_message"      // 设置群精华消息
	QuitGroup                  Endpoint = "quit_group"                     // 退出群
	SendGroupMessageReaction   Endpoint = "send_group_message_reaction"    // 发送群消息表情回应
	SendGroupNudge             Endpoint = "send_group_nudge"               // 发送群戳一戳
	GetGroupNotifications      Endpoint = "get_group_notifications"        // 获取群通知列表
	AcceptGroupRequest         Endpoint = "accept_group_request"           // 同意入群/邀请他人入群请求
	RejectGroupRequest         Endpoint = "reject_group_request"           // 拒绝入群/邀请他人入群请求
	AcceptGroupInvitation      Endpoint = "accept_group_invitation"        // 同意他人邀请自身入群
	RejectGroupInvitation      Endpoint = "reject_group_invitation"        // 拒绝他人邀请自身入群

	// FileAPI

	UploadPrivateFile         Endpoint = "upload_private_file"           // 上传私聊文件
	UploadGroupFile           Endpoint = "upload_group_file"             // 上传群文件
	GetPrivateFileDownloadURL Endpoint = "get_private_file_download_url" // 获取私聊文件下载链接
	GetGroupFileDownloadURL   Endpoint = "get_group_file_download_url"   // 获取群文件下载链接
	GetGroupFiles             Endpoint = "get_group_files"               // 获取群文件列表
	MoveGroupFile             Endpoint = "move_group_file"               // 移动群文件
	RenameGroupFile           Endpoint = "rename_group_file"             // 重命名群文件
	DeleteGroupFile           Endpoint = "delete_group_file"             // 删除群文件
	CreateGroupFolder         Endpoint = "create_group_folder"           // 创建群文件夹
	RenameGroupFolder         Endpoint = "rename_group_folder"           // 重命名群文件夹
	DeleteGroupFolder         Endpoint = "delete_group_folder"           // 删除群文件夹
)
