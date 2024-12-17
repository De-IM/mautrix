package event

type BotCommand struct {
	BotId        string          `json:"bot_id"`
	Command      string          `json:"command"`
	Description  string          `json:"description"`
	Type         string          `json:"type"` // 命令类型，值可为: chat、user、message。
	Options      []CommandOption `json:"options"`
	Scope        BotCommandScope `json:"scope"` // default、all_private_chats、all_group_chats、all_chat_administrators、chat、chat_administrators、chat_member
	LanguageCode string          `json:"language_code"`
}

type CommandOption struct {
	Name        string `json:"name"`
	Type        string `json:"type"` // 如: string, integer等
	Required    bool   `json:"required"`
	Description string `json:"description"`
	Value       string `json:"value"`
}

type BotCommandScope struct {
	Type   string `json:"type" form:"type"`
	ChatId string `json:"chat_id" form:"chat_id"`
	UserId string `json:"user_id" form:"user_id"`
}
