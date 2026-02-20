package types

import "encoding/json"

type SegmentType string

const (
	TextSegmentType       SegmentType = "text"        // Incoming, Outgoing,
	MentionSegmentType    SegmentType = "mention"     // Incoming, Outgoing,
	MentionAllSegmentType SegmentType = "mention_all" // Incoming, Outgoing,
	FaceSegmentType       SegmentType = "face"        // Incoming, Outgoing,
	ReplySegmentType      SegmentType = "reply"       // Incoming, Outgoing,
	ImageSegmentType      SegmentType = "image"       // Incoming, Outgoing,
	RecordSegmentType     SegmentType = "record"      // Incoming, Outgoing,
	VideoSegmentType      SegmentType = "video"       // Incoming, Outgoing,
	ForwardSegmentType    SegmentType = "forward"     // Incoming, Outgoing,
	FileSegmentType       SegmentType = "file"        // Incoming,
	MarketFaceSegmentType SegmentType = "market_face" // Incoming,
	LightAppSegmentType   SegmentType = "light_app"   // Incoming,
	XMLSegmentType        SegmentType = "xml"         // Incoming,
)

type Segment interface {
	json.Marshaler
	json.Unmarshaler
	Type() string
}

type TextSegment struct {
	Text string `json:"text"`
}

func (s *TextSegment) Type() string {
	return string(TextSegmentType)
}

func (s *TextSegment) Marshal() ([]byte, error) {
	jsonBytes, err := json.Marshal(s)
	if err != nil {
		return json.RawMessage{}, err
	}
	return jsonBytes, nil
}

func (s *TextSegment) Unmarshal(data []byte) error {
	return json.Unmarshal(data, s)
}

// [TODO] 实现更多消息段
// [TODO] 实现消息段数据和实际对象的转换
