package emi_core

import (
	"context"

	"github.com/aK1r4z/emi-core/api"
	milky_types "github.com/aK1r4z/emi-core/types"
)

type Logger interface {
	Tracef(format string, args ...any)
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)

	Trace(args ...any)
	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any)
}

type APIClient interface {

	// SystemAPI

	GetLoginInfo(context.Context, api.GetLoginInfoRequest) (*api.GetLoginInfoResponse, error)                         // 获取登录信息
	GetImplInfo(context.Context, api.GetImplInfoRequest) (*api.GetImplInfoResponse, error)                            // 获取协议端信息
	GetUserProfile(context.Context, api.GetUserProfileRequest) (*api.GetUserProfileResponse, error)                   // 获取用户个人信息
	GetFriendList(context.Context, api.GetFriendListRequest) (*api.GetFriendListResponse, error)                      // 获取好友列表
	GetFriendInfo(context.Context, api.GetFriendInfoRequest) (*api.GetFriendInfoResponse, error)                      // 获取好友信息
	GetGroupList(context.Context, api.GetGroupListRequest) (*api.GetGroupListResponse, error)                         // 获取群列表
	GetGroupInfo(context.Context, api.GetGroupInfoRequest) (*api.GetGroupInfoResponse, error)                         // 获取群信息
	GetGroupMemberList(context.Context, api.GetGroupMemberListRequest) (*api.GetGroupMemberListResponse, error)       // 获取群成员列表
	GetGroupMemberInfo(context.Context, api.GetGroupMemberInfoRequest) (*api.GetGroupMemberInfoResponse, error)       // 获取群成员信息
	SetAvatar(context.Context, api.SetAvatarRequest) (*api.SetAvatarResponse, error)                                  // 设置 QQ 账号头像
	SetNickname(context.Context, api.SetNicknameRequest) (*api.SetNicknameResponse, error)                            // 设置 QQ 账号昵称
	SetBio(context.Context, api.SetBioRequest) (*api.SetBioResponse, error)                                           // 设置 QQ 账号个性签名
	GetCustomFaceURLList(context.Context, api.GetCustomFaceURLListRequest) (*api.GetCustomFaceURLListResponse, error) // 获取自定义表情 URL 列表
	GetCookies(context.Context, api.GetCookiesRequest) (*api.GetCookiesResponse, error)                               // 获取 Cookies
	GetCSRFToken(context.Context, api.GetCSRFTokenRequest) (*api.GetCSRFTokenResponse, error)                         // 获取 CSRF Token

	// MessageAPI

	SendPrivateMessage(context.Context, api.SendPrivateMessageRequest) (*api.SendPrivateMessageResponse, error)       // 发送私聊消息
	SendGroupMessage(context.Context, api.SendGroupMessageRequest) (*api.SendGroupMessageResponse, error)             // 发送群聊消息
	RecallPrivateMessage(context.Context, api.RecallPrivateMessageRequest) (*api.RecallPrivateMessageResponse, error) // 撤回私聊消息
	RecallGroupMessage(context.Context, api.RecallGroupMessageRequest) (*api.RecallGroupMessageResponse, error)       // 撤回群聊消息
	GetMessage(context.Context, api.GetMessageRequest) (*api.GetMessageResponse, error)                               // 获取消息
	GetHistoryMessages(context.Context, api.GetHistoryMessagesRequest) (*api.GetHistoryMessagesResponse, error)       // 获取历史消息列表
	GetResourceTempURL(context.Context, api.GetResourceTempURLRequest) (*api.GetResourceTempURLResponse, error)       // 获取临时资源链接
	GetForwardedMessages(context.Context, api.GetForwardedMessagesRequest) (*api.GetForwardedMessagesResponse, error) // 获取合并转发消息内容
	MarkMessageAsRead(context.Context, api.MarkMessageAsReadRequest) (*api.MarkMessageAsReadResponse, error)          // 标记消息为已读

	// FriendAPI

	SendFriendNudge(context.Context, api.SendFriendNudgeRequest) (*api.SendFriendNudgeResponse, error)             // 发送好友戳一戳
	SendProfileLike(context.Context, api.SendProfileLikeRequest) (*api.SendProfileLikeResponse, error)             // 发送名片点赞
	DeleteFriend(context.Context, api.DeleteFriendRequest) (*api.DeleteFriendResponse, error)                      // 删除好友
	GetFriendRequests(context.Context, api.GetFriendRequestsRequest) (*api.GetFriendRequestsResponse, error)       // 获取好友请求列表
	AcceptFriendRequest(context.Context, api.AcceptFriendRequestRequest) (*api.AcceptFriendRequestResponse, error) // 同意好友请求
	RejectFriendRequest(context.Context, api.RejectFriendRequestRequest) (*api.RejectFriendRequestResponse, error) // 拒绝好友请求

	// GroupAPI

	SetGroupName(context.Context, api.SetGroupNameRequest) (*api.SetGroupNameResponse, error)                                           // 设置群名称
	SetGroupAvatar(context.Context, api.SetGroupAvatarRequest) (*api.SetGroupAvatarResponse, error)                                     // 设置群头像
	SetGroupMemberCard(context.Context, api.SetGroupMemberCardRequest) (*api.SetGroupMemberCardResponse, error)                         // 设置群名片
	SetGroupMemberSpecialTitle(context.Context, api.SetGroupMemberSpecialTitleRequest) (*api.SetGroupMemberSpecialTitleResponse, error) // 设置群成员专属头衔
	SetGroupMemberAdmin(context.Context, api.SetGroupMemberAdminRequest) (*api.SetGroupMemberAdminResponse, error)                      // 设置群管理员
	SetGroupMemberMute(context.Context, api.SetGroupMemberMuteRequest) (*api.SetGroupMemberMuteResponse, error)                         // 设置群成员禁言
	SetGroupMemberWholeMute(context.Context, api.SetGroupMemberWholeMuteRequest) (*api.SetGroupMemberWholeMuteResponse, error)          // 设置群全员禁言
	KickGroupMember(context.Context, api.KickGroupMemberRequest) (*api.KickGroupMemberResponse, error)                                  // 踢出群成员
	GetGroupAnnouncements(context.Context, api.GetGroupAnnouncementsRequest) (*api.GetGroupAnnouncementsResponse, error)                // 获取群公告列表
	SendGroupAnnouncement(context.Context, api.SendGroupAnnouncementRequest) (*api.SendGroupAnnouncementResponse, error)                // 发送群公告
	DeleteGroupAnnouncement(context.Context, api.DeleteGroupAnnouncementRequest) (*api.DeleteGroupAnnouncementResponse, error)          // 删除群公告
	GetGroupEssenceMessages(context.Context, api.GetGroupEssenceMessagesRequest) (*api.GetGroupEssenceMessagesResponse, error)          // 获取群精华消息列表
	SetGroupEssenceMessage(context.Context, api.SetGroupEssenceMessageRequest) (*api.SetGroupEssenceMessageResponse, error)             // 设置群精华消息
	QuitGroup(context.Context, api.QuitGroupRequest) (*api.QuitGroupResponse, error)                                                    // 退出群
	SendGroupMessageReaction(context.Context, api.SendGroupMessageReactionRequest) (*api.SendGroupMessageReactionResponse, error)       // 发送群消息表情回应
	SendGroupNudge(context.Context, api.SendGroupNudgeRequest) (*api.SendGroupNudgeResponse, error)                                     // 发送群戳一戳
	GetGroupNotifications(context.Context, api.GetGroupNotificationsRequest) (*api.GetGroupNotificationsResponse, error)                // 获取群通知列表
	AcceptGroupRequest(context.Context, api.AcceptGroupRequestRequest) (*api.AcceptGroupRequestResponse, error)                         // 同意入群/邀请他人入群请求
	RejectGroupRequest(context.Context, api.RejectGroupRequestRequest) (*api.RejectGroupRequestResponse, error)                         // 拒绝入群/邀请他人入群请求
	AcceptGroupInvitation(context.Context, api.AcceptGroupInvitationRequest) (*api.AcceptGroupInvitationResponse, error)                // 同意他人邀请自身入群
	RejectGroupInvitation(context.Context, api.RejectGroupInvitationRequest) (*api.RejectGroupInvitationResponse, error)                // 拒绝他人邀请自身入群

	// FileAPI

	UploadPrivateFile(context.Context, api.UploadPrivateFileRequest) (*api.UploadPrivateFileResponse, error)                         // 上传私聊文件
	UploadGroupFile(context.Context, api.UploadGroupFileRequest) (*api.UploadGroupFileResponse, error)                               // 上传群文件
	GetPrivateFileDownloadURL(context.Context, api.GetPrivateFileDownloadURLRequest) (*api.GetPrivateFileDownloadURLResponse, error) // 获取私聊文件下载链接
	GetGroupFileDownloadURL(context.Context, api.GetGroupFileDownloadURLRequest) (*api.GetGroupFileDownloadURLResponse, error)       // 获取群文件下载链接
	GetGroupFiles(context.Context, api.GetGroupFilesRequest) (*api.GetGroupFilesResponse, error)                                     // 获取群文件列表
	MoveGroupFile(context.Context, api.MoveGroupFileRequest) (*api.MoveGroupFileResponse, error)                                     // 移动群文件
	RenameGroupFile(context.Context, api.RenameGroupFileRequest) (*api.RenameGroupFileResponse, error)                               // 重命名群文件
	DeleteGroupFile(context.Context, api.DeleteGroupFileRequest) (*api.DeleteGroupFileResponse, error)                               // 删除群文件
	CreateGroupFolder(context.Context, api.CreateGroupFolderRequest) (*api.CreateGroupFolderResponse, error)                         // 创建群文件夹
	RenameGroupFolder(context.Context, api.RenameGroupFolderRequest) (*api.RenameGroupFolderResponse, error)                         // 重命名群文件夹
	DeleteGroupFolder(context.Context, api.DeleteGroupFolderRequest) (*api.DeleteGroupFolderResponse, error)                         // 删除群文件夹
}

type Bot interface {
	Open(context.Context) error
	Close() error

	Wait()

	APIClient

	Handle(EventHandler)

	Logger() Logger
}

type EventSource interface {
	Open() (chan milky_types.RawEvent, error)
	Close() error
}

type EventHandler interface {
	milky_types.EventDescriber
	Handle(Bot, milky_types.Event, milky_types.RawEvent)
}
