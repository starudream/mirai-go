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

type MessageInfoPlain struct {
	Type MessageInfoType `json:"type"`
	Text string          `json:"text"`
}

func (m *MessageInfoPlain) setType() { m.Type = MessageInfoTypePlain }
