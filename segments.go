package emi_core

import (
	"encoding/json"

	milky_types "github.com/aK1r4z/emi-core/types"
)

type ImageSubtype string

const (
	ImageNormal  ImageSubtype = "normal"
	ImageSticker ImageSubtype = "sticker"
)

type Element interface {
	Type() milky_types.SegmentType
	json.Marshaler
	json.Unmarshaler
}

func NewSegment(segmentType milky_types.SegmentType, element Element) (*milky_types.Segment, error) {
	b, err := element.MarshalJSON()
	if err != nil {
		return nil, err
	}

	return &milky_types.Segment{
		Type: segmentType,
		Data: json.RawMessage(b),
	}, nil
}

// 文本消息段
type TextElement struct {
	Text string `json:"text"` // 文本内容
}

func (s *TextElement) Type() milky_types.SegmentType { return milky_types.SegmentText }
func (s *TextElement) MarshalJSON() ([]byte, error)  { return json.Marshal(*s) }
func (s *TextElement) UnmarshalJSON(b []byte) error  { return json.Unmarshal(b, s) }

func NewTextSegment(text string) (*milky_types.Segment, error) {
	return NewSegment(milky_types.SegmentText, &TextElement{Text: text})
}

// 提及消息段
type MentionElement struct {
	UserID int64 `json:"user_id"` // 提及的 QQ 号
}

func (s *MentionElement) Type() milky_types.SegmentType { return milky_types.SegmentMention }
func (s *MentionElement) MarshalJSON() ([]byte, error)  { return json.Marshal(*s) }
func (s *MentionElement) UnmarshalJSON(b []byte) error  { return json.Unmarshal(b, s) }

func NewMentionSegment(userID int64) (*milky_types.Segment, error) {
	return NewSegment(milky_types.SegmentText, &MentionElement{UserID: userID})
}

// 提及全体消息段
type MentionAllElement struct{}

func (s *MentionAllElement) Type() milky_types.SegmentType { return milky_types.SegmentMentionAll }
func (s *MentionAllElement) MarshalJSON() ([]byte, error)  { return json.Marshal(*s) }
func (s *MentionAllElement) UnmarshalJSON(b []byte) error  { return json.Unmarshal(b, s) }

func NewMentionAllSegment() (*milky_types.Segment, error) {
	return NewSegment(milky_types.SegmentText, &MentionAllElement{})
}

// 表情消息段
type FaceElement struct {
	FaceID  string `json:"face_id"`  // 表情 ID
	IsLarge bool   `json:"is_large"` // 是否为超级表情
}

func (s *FaceElement) Type() milky_types.SegmentType { return milky_types.SegmentFace }
func (s *FaceElement) MarshalJSON() ([]byte, error)  { return json.Marshal(*s) }
func (s *FaceElement) UnmarshalJSON(b []byte) error  { return json.Unmarshal(b, s) }

func NewFaceSegment(faceID string, isLarge bool) (*milky_types.Segment, error) {
	return NewSegment(milky_types.SegmentText, &FaceElement{FaceID: faceID, IsLarge: isLarge})
}

// 回复消息段
type ReplyElement struct {
	MessageSeq int64 `json:"message_seq"` // 被引用的消息序列号
}

func (s *ReplyElement) Type() milky_types.SegmentType { return milky_types.SegmentReply }
func (s *ReplyElement) MarshalJSON() ([]byte, error)  { return json.Marshal(*s) }
func (s *ReplyElement) UnmarshalJSON(b []byte) error  { return json.Unmarshal(b, s) }

func NewReplySegment(messageSeq int64) (*milky_types.Segment, error) {
	return NewSegment(milky_types.SegmentText, &ReplyElement{MessageSeq: messageSeq})
}

// 图片消息段
type ImageElement struct {
	ResourceID string       `json:"resource_id"` // 资源 ID
	TempURL    string       `json:"temp_url"`    // 临时 URL
	Width      int32        `json:"width"`       // 图片宽度
	Height     int32        `json:"height"`      // 图片高度
	Summary    string       `json:"summary"`     // 图片预览文本
	SubType    ImageSubtype `json:"sub_type"`    // 图片类型，可能值：normal sticker
}

func (s *ImageElement) Type() milky_types.SegmentType { return milky_types.SegmentImage }
func (s *ImageElement) MarshalJSON() ([]byte, error)  { return json.Marshal(*s) }
func (s *ImageElement) UnmarshalJSON(b []byte) error  { return json.Unmarshal(b, s) }

func NewImageSegment(
	resourceID string,
	tempURL string,
	width int32,
	height int32,
	summary string,
	subType ImageSubtype,
) (*milky_types.Segment, error) {
	return NewSegment(milky_types.SegmentText, &ImageElement{
		ResourceID: resourceID,
		TempURL:    tempURL,
		Width:      width,
		Height:     height,
		Summary:    summary,
		SubType:    subType,
	})
}

// 语音消息段
type RecordElement struct {
	ResourceID string `json:"resource_id"` // 资源 ID
	TempURL    string `json:"temp_url"`    // 临时 URL
	Duration   int32  `json:"duration"`    // 语音时长（秒）
}

func (s *RecordElement) Type() milky_types.SegmentType { return milky_types.SegmentRecord }
func (s *RecordElement) MarshalJSON() ([]byte, error)  { return json.Marshal(*s) }
func (s *RecordElement) UnmarshalJSON(b []byte) error  { return json.Unmarshal(b, s) }

func NewRecordSegment(
	resourceID string,
	tempURL string,
	duration int32,
) (*milky_types.Segment, error) {
	return NewSegment(milky_types.SegmentText, &RecordElement{
		ResourceID: resourceID,
		TempURL:    tempURL,
		Duration:   duration,
	})
}

// 视频消息段
type VideoElement struct {
	ResourceID string `json:"resource_id"` // 资源 ID
	TempURL    string `json:"temp_url"`    // 临时 URL
	Width      int32  `json:"width"`       // 视频宽度
	Height     int32  `json:"height"`      // 视频高度
	Duration   int32  `json:"duration"`    // 视频时长（秒）
}

func (s *VideoElement) Type() milky_types.SegmentType { return milky_types.SegmentVideo }
func (s *VideoElement) MarshalJSON() ([]byte, error)  { return json.Marshal(*s) }
func (s *VideoElement) UnmarshalJSON(b []byte) error  { return json.Unmarshal(b, s) }

func NewVideoSegment(
	resourceID string,
	tempURL string,
	width int32,
	height int32,
	duration int32,
) (*milky_types.Segment, error) {
	return NewSegment(milky_types.SegmentText, &VideoElement{
		ResourceID: resourceID,
		TempURL:    tempURL,
		Width:      width,
		Height:     height,
		Duration:   duration,
	})
}

// 合并转发消息段
type ForwardElement struct {
	ForwardID string   `json:"forward_id"` // 合并转发 ID
	Title     string   `json:"title"`      // 合并转发标题
	Preview   []string `json:"preview"`    // 合并转发预览文本
	Summary   string   `json:"summary"`    // 合并转发摘要
}

func (s *ForwardElement) Type() milky_types.SegmentType { return milky_types.SegmentForward }
func (s *ForwardElement) MarshalJSON() ([]byte, error)  { return json.Marshal(*s) }
func (s *ForwardElement) UnmarshalJSON(b []byte) error  { return json.Unmarshal(b, s) }

func NewForwardSegment(
	forwardID string,
	title string,
	preview []string,
	summary string,
) (*milky_types.Segment, error) {
	return NewSegment(milky_types.SegmentText, &ForwardElement{
		ForwardID: forwardID,
		Title:     title,
		Preview:   preview,
		Summary:   summary,
	})
}
