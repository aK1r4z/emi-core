package api

import milky "github.com/aK1r4z/emi-core/types"

type (
	SendPrivateMessageRequest struct {
		UserID  int64                    `json:"user_id"`
		Message []milky.OutgoingSegment `json:"message"`
	}

	SendPrivateMessageResponse struct {
		MessageSeq int64 `json:"message_seq"`
		Time       int64 `json:"time"`
	}

	SendGroupMessageRequest struct {
		GroupID int64                    `json:"group_id"`
		Message []milky.OutgoingSegment `json:"message"`
	}

	SendGroupMessageResponse struct {
		MessageSeq int64 `json:"message_seq"`
		Time       int64 `json:"time"`
	}

	RecallPrivateMessageRequest struct {
		UserID     int64 `json:"user_id"`
		MessageSeq int64 `json:"message_seq"`
	}

	RecallPrivateMessageResponse struct{}

	RecallGroupMessageRequest struct {
		GroupID    int64 `json:"group_id"`
		MessageSeq int64 `json:"message_seq"`
	}

	RecallGroupMessageResponse struct{}

	GetMessageRequest struct {
		MessageScene milky.MessageScene `json:"message_scene"`
		PeerID       int64               `json:"peer_id"`
		MessageSeq   int64               `json:"message_seq"`
	}

	GetMessageResponse struct {
		Message milky.IncomingMessage `json:"message"`
	}

	GetHistoryMessageRequest struct {
		MessageScene    milky.MessageScene `json:"message_scene"`
		PeerID          int64               `json:"peer_id"`
		StartMessageSeq *int64              `json:"start_message_seq"`
		Limit           int32               `json:"limit"`
	}

	GetHistoryMessageResponse struct {
		Messages       []milky.IncomingMessage `json:"messages"`
		NextMessageSeq *int64                   `json:"next_message_seq"`
	}

	GetResourceTempURLRequest struct {
		ResourceID string `json:"resource_id"`
	}

	GetResourceTempURLResponse struct {
		URL string `json:"url"`
	}

	GetForwardedMessagesRequest struct {
		ForwardID string `json:"forward_id"`
	}

	GetForwardedMessagesResponse struct {
		Messages []milky.IncomingForwardedMessage `json:"messages"`
	}

	MarkMessageAsReadRequest struct {
		MessageScene milky.MessageScene `json:"message_scene"`
		PeerID       int64               `json:"peer_id"`
		MessageSeq   int64               `json:"message_seq"`
	}

	MarkMessageAsReadResponse struct{}
)
