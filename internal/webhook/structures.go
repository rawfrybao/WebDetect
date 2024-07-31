package webhook

type TgMessage struct {
	ChatID          int64                     `json:"chat_id"`
	Text            string                    `json:"text"`
	ParseMode       string                    `json:"parse_mode,omitempty"`
	ReplyParameters *TgMessageReplyParameters `json:"reply_parameters,omitempty"`
}

type TgMessageReplyParameters struct {
	MessageID int64 `json:"message_id"`
	ChatID    int64 `json:"chat_id"`
}
