package emi_core

type Endpoint string

const (

	// System

	SystemLoginInfoGet Endpoint = "get_login_info" // 获取登录信息
	SystemImplInfoGet  Endpoint = "get_impl_info"  // 获取协议端信息
	SystemCookiesGet   Endpoint = "get_cookies"    // 获取 Cookies
	SystemCSRFTokenGet Endpoint = "get_csrf_token" // 获取 CSRF Token

	// Self

	SelfProfileGet  Endpoint = "get_user_profile" // 获取用户个人信息
	SelfAvatarSet   Endpoint = "set_avatar"       // 设置 QQ 账号头像
	SelfNicknameSet Endpoint = "set_nickname"     // 设置 QQ 账号昵称
	SelfBioSet      Endpoint = "set_bio"          // 设置 QQ 账号个性签名

	// Friend

	FriendListGet Endpoint = "get_friend_list" // 获取好友列表
	FriendInfoGet Endpoint = "get_friend_info" // 获取好友信息

	// Group

	GroupAnnouncementSend    Endpoint = "send_group_announcement"     // 发送群公告
	GroupMessageReactionSend Endpoint = "send_group_message_reaction" // 发送群消息表情回应
	GroupNudgeSend           Endpoint = "send_group_nudge"            // 发送群戳一戳
	GroupListGet             Endpoint = "get_group_list"              // 获取群列表
	GroupInfoGet             Endpoint = "get_group_info"              // 获取群信息
	GroupAnnouncementsGet    Endpoint = "get_group_announcements"     // 获取群公告列表
	GroupEssenceMessagesGet  Endpoint = "get_group_essence_messages"  // 获取群精华消息列表
	GroupNotificationsGet    Endpoint = "get_group_notifications"     // 获取群通知列表
	GroupNameSet             Endpoint = "set_group_name"              // 设置群名称
	GroupAvatarSet           Endpoint = "set_group_avatar"            // 设置群头像
	GroupWholeMuteSet        Endpoint = "set_group_whole_mute"        // 设置群全员禁言
	GroupEssenceMessageSet   Endpoint = "set_group_essence_message"   // 设置群精华消息
	GroupAnnouncementDelete  Endpoint = "delete_group_announcement"   // 删除群公告
	GroupRequestAccept       Endpoint = "accept_group_request"        // 同意入群/邀请他人入群请求
	GroupInvitationAccept    Endpoint = "accept_group_invitation"     // 同意他人邀请自身入群
	GroupRequestReject       Endpoint = "reject_group_request"        // 拒绝入群/邀请他人入群请求
	GroupInvitationReject    Endpoint = "reject_group_invitation"     // 拒绝他人邀请自身入群
	GroupQuit                Endpoint = "quit_group"                  // 退出群

	// GroupMember

	GroupMemberListGet         Endpoint = "get_group_member_list"          // 获取群成员列表
	GroupMemberInfoGet         Endpoint = "get_group_member_info"          // 获取群成员信息
	GroupMemberCardSet         Endpoint = "set_group_member_card"          // 设置群名片
	GroupMemberSpecialTitleSet Endpoint = "set_group_member_special_title" // 设置群成员专属头衔
	GroupMemberAdminSet        Endpoint = "set_group_member_admin"         // 设置群管理员
	GroupMemberMuteSet         Endpoint = "set_group_member_mute"          // 设置群成员禁言
	GroupMemberKick            Endpoint = "kick_group_member"              // 踢出群成员

	// Resource
	ResourceCustomFaceURLListGet Endpoint = "get_custom_face_url_list" // 获取自定义表情 URL 列表
	ResourceTempURLGet           Endpoint = "get_resource_temp_url"    // 获取临时资源链接

	// Files & Folder

	FileUploadPrivate         Endpoint = "upload_private_file"           // 上传私聊文件
	FileUploadGroup           Endpoint = "upload_group_file"             // 上传群文件
	FileDownloadURLGetPrivate Endpoint = "get_private_file_download_url" // 获取私聊文件下载链接
	FileDownloadURLGetGroup   Endpoint = "get_group_file_download_url"   // 获取群文件下载链接
	FilesGetGroup             Endpoint = "get_group_files"               // 获取群文件列表
	FileMoveGroup             Endpoint = "move_group_file"               // 移动群文件
	FileRenameGroup           Endpoint = "rename_group_file"             // 重命名群文件
	FileDeleteGroup           Endpoint = "delete_group_file"             // 删除群文件
	FolderCreateGroup         Endpoint = "create_group_folder"           // 创建群文件夹
	FolderRenameGroup         Endpoint = "rename_group_folder"           // 重命名群文件夹
	FolderDeleteGroup         Endpoint = "delete_group_folder"           // 删除群文件夹

	// Messages
	MessageSendPrivate   Endpoint = "send_private_message"   // 发送私聊消息
	MessageSendGroup     Endpoint = "send_group_message"     // 发送群聊消息
	MessageGet           Endpoint = "get_message"            // 获取消息
	MessagesHistoryGet   Endpoint = "get_history_messages"   // 获取历史消息列表
	MessagesForwardedGet Endpoint = "get_forwarded_messages" // 获取合并转发消息内容
	MessageRecallPrivate Endpoint = "recall_private_message" // 撤回私聊消息
	MessageRecallGroup   Endpoint = "recall_group_message"   // 撤回群聊消息
	MessageMarkAsRead    Endpoint = "mark_message_as_read"   // 标记消息为已读

)
