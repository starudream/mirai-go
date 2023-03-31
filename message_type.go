package mirai

import (
	"encoding/json"
)

type MessageType string

const (
	MessageTypeFriend   MessageType = "FriendMessage"
	MessageTypeGroup    MessageType = "GroupMessage"
	MessageTypeTemp     MessageType = "TempMessage"
	MessageTypeStranger MessageType = "StrangerMessage"
	MessageTypeOther    MessageType = "OtherClientMessage"

	MessageTypeFriendSync   MessageType = "FriendSyncMessage"
	MessageTypeGroupSync    MessageType = "GroupSyncMessage"
	MessageTypeTempSync     MessageType = "TempSyncMessage"
	MessageTypeStrangerSync MessageType = "StrangerSyncMessage"
)

var messageTypeMap = map[string]MessageType{
	string(MessageTypeFriend):   MessageTypeFriend,
	string(MessageTypeGroup):    MessageTypeGroup,
	string(MessageTypeTemp):     MessageTypeTemp,
	string(MessageTypeStranger): MessageTypeStranger,
	string(MessageTypeOther):    MessageTypeOther,

	string(MessageTypeFriendSync):   MessageTypeFriendSync,
	string(MessageTypeGroupSync):    MessageTypeGroupSync,
	string(MessageTypeTempSync):     MessageTypeTempSync,
	string(MessageTypeStrangerSync): MessageTypeStrangerSync,
}

type Message struct {
	raw  []byte
	data map[string]any

	Type MessageType `json:"-"`
}

type Messages []Message

func (ms Messages) Len() int {
	return len(ms)
}

func (ms Messages) Empty() bool {
	return len(ms) == 0
}

var _ json.Marshaler = (*Message)(nil)

func (v *Message) MarshalJSON() ([]byte, error) {
	return v.raw, nil
}

var _ json.Unmarshaler = (*Message)(nil)

func (v *Message) UnmarshalJSON(raw []byte) error {
	v.raw = raw
	if err := json.Unmarshal(v.raw, &v.data); err != nil {
		return err
	}
	if s, ok1 := v.data["type"].(string); ok1 {
		if t, ok2 := messageTypeMap[s]; ok2 {
			v.Type = t
		}
	}
	// if v.Type == "" {
	// 	return fmt.Errorf("unknown message type")
	// }
	return nil
}
