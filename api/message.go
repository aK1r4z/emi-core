package api

import milky_types "github.com/aK1r4z/emi-core/types"

type (
	// 发送私聊消息
	SendPrivateMessageRequest struct {
		UserID  int64           `json:"user_id"` // 好友 QQ 号
		Message []milky_types.Segment `json:"message"` // 消息内容
	}

	SendPrivateMessageResponse struct {
		MessageSeq int64 `json:"message_seq"` // 消息序列号
		Time       int64 `json:"time"`        // 消息发送时间
	}

	// 发送群聊消息
	SendGroupMessageRequest struct {
		GroupID int64           `json:"group_id"` // 群号
		Message []milky_types.Segment `json:"message"`  // 消息内容
	}

	SendGroupMessageResponse struct {
		MessageSeq int64 `json:"message_seq"` // 消息序列号
		Time       int64 `json:"time"`        // 消息发送时间
	}

	// 撤回私聊消息
	RecallPrivateMessageRequest struct {
		UserID     int64 `json:"user_id"`     // 好友 QQ 号
		MessageSeq int64 `json:"message_seq"` // 消息序列号
	}

	RecallPrivateMessageResponse struct{}

	// 撤回群聊消息
	RecallGroupMessageRequest struct {
		GroupID    int64 `json:"group_id"`    // 群号
		MessageSeq int64 `json:"message_seq"` // 消息序列号
	}

	RecallGroupMessageResponse struct{}

	// 获取消息
	GetMessageRequest struct {
		MessageScene milky_types.MessageScene `json:"message_scene"` // 消息场景，可能值：friend group temp
		PeerID       int64              `json:"peer_id"`       // 好友 QQ 号或群号
		MessageSeq   int64              `json:"message_seq"`   // 消息序列号
	}

	GetMessageResponse struct {
		Message milky_types.IncomingMessage `json:"message"` // 消息内容
	}

	// 获取历史消息列表
	GetHistoryMessagesRequest struct {
		MessageScene    milky_types.MessageScene `json:"message_scene"`     // 消息场景，可能值：friend group temp
		PeerID          int64              `json:"peer_id"`           // 好友 QQ 号或群号
		StartMessageSeq *int64             `json:"start_message_seq"` // 起始消息序列号，由此开始从新到旧查询，不提供则从最新消息开始
		Limit           int32              `json:"limit"`             // 期望获取到的消息数量，最多 30 条，默认值：20
	}

	GetHistoryMessagesResponse struct {
		Messages       []milky_types.IncomingMessage `json:"messages"`         // 获取到的消息（message_seq 升序排列），部分消息可能不存在，如撤回的消息
		NextMessageSeq *int64                  `json:"next_message_seq"` // 下一页起始消息序列号
	}

	// 获取临时资源链接
	GetResourceTempURLRequest struct {
		ResourceID string `json:"resource_id"` // 资源 ID
	}

	GetResourceTempURLResponse struct {
		URL string `json:"url"` // 临时资源链接
	}

	// 获取合并转发消息内容
	GetForwardedMessagesRequest struct {
		ForwardID string `json:"forward_id"` // 转发消息 ID
	}

	GetForwardedMessagesResponse struct {
		Messages []milky_types.IncomingForwardedMessage `json:"messages"` // 转发消息内容
	}

	// 标记消息为已读
	MarkMessageAsReadRequest struct {
		MessageScene milky_types.MessageScene `json:"message_scene"` // 消息场景，可能值：friend group temp
		PeerID       int64              `json:"peer_id"`       // 好友 QQ 号或群号
		MessageSeq   int64              `json:"message_seq"`   // 标为已读的消息序列号，该消息及更早的消息将被标记为已读
	}

	MarkMessageAsReadResponse struct{}
)
