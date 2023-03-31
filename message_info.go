package mirai

type MessageInfoType string

const (
	MessageInfoTypeSource         MessageInfoType = "Source"
	MessageInfoTypeQuote          MessageInfoType = "Quote"
	MessageInfoTypeAt             MessageInfoType = "At"
	MessageInfoTypeAtAll          MessageInfoType = "AtAll"
	MessageInfoTypeFace           MessageInfoType = "Face"
	MessageInfoTypePlain          MessageInfoType = "Plain"
	MessageInfoTypeImage          MessageInfoType = "Image"
	MessageInfoTypeFlashImage     MessageInfoType = "FlashImage"
	MessageInfoTypeVoice          MessageInfoType = "Voice"
	MessageInfoTypeXml            MessageInfoType = "Xml"
	MessageInfoTypeJson           MessageInfoType = "Json"
	MessageInfoTypeApp            MessageInfoType = "App"
	MessageInfoTypePoke           MessageInfoType = "Poke"
	MessageInfoTypeDice           MessageInfoType = "Dice"
	MessageInfoTypeMarketFace     MessageInfoType = "MarketFace"
	MessageInfoTypeMusicShare     MessageInfoType = "MusicShare"
	MessageInfoTypeForwardMessage MessageInfoType = "ForwardMessage"
	MessageInfoTypeFile           MessageInfoType = "File"
	MessageInfoTypeMiraiCode      MessageInfoType = "MiraiCode"
)

type MessageInfoInterface interface {
	setType()
}

type MessageInfoSource struct {
	Type MessageInfoType `json:"type"`
	Id   int64           `json:"id"`
	Time int64           `json:"time"`
}

func (m *MessageInfoSource) setType() { m.Type = MessageInfoTypeSource }

type MessageInfoAt struct {
	Type    MessageInfoType `json:"type"`
	Target  int64           `json:"target"`
	Display string          `json:"display,omitempty"`
}

func (m *MessageInfoAt) setType() { m.Type = MessageInfoTypeAt }

type MessageInfoAtAll struct {
	Type MessageInfoType `json:"type"`
}

func (m *MessageInfoAtAll) setType() { m.Type = MessageInfoTypeAtAll }

type MessageInfoPlain struct {
	Type MessageInfoType `json:"type"`
	Text string          `json:"text"`
}

func (m *MessageInfoPlain) setType() { m.Type = MessageInfoTypePlain }
