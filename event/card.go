package event

type CardHeader struct {
	Type         string `json:"type"`
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle"`
	ImageUrl     string `json:"image_url"`
	ImageType    string `json:"image_type"`
	ImageAltText string `json:"image_alt_text"`
}

type DecoratedText struct {
	StartIcon CardIcon `json:"start_icon"`
	Text      string   `json:"text"`
}

type CardIcon struct {
	KnownIcon string `json:"known_icon"`
}

type CardButton struct {
	Text    string  `json:"text"`
	OnClick OnClick `json:"on_click"`
}

type OnClick struct {
	OpenLink *OpenLink `json:"open_link,omitempty"`
	Action   *Action   `json:"action,omitempty"`
}

type OpenLink struct {
	URL string `json:"url"`
}

type Action struct {
	Function   string      `json:"function"`
	Parameters []Parameter `json:"parameters"`
}

type Parameter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ButtonList struct {
	Buttons []CardButton `json:"buttons"`
}

type Widget struct {
	DecoratedText *DecoratedText `json:"decorated_text,omitempty"`
	ButtonList    *ButtonList    `json:"button_list,omitempty"`
}

type Section struct {
	Header                    string   `json:"header"`
	Collapsible               bool     `json:"collapsible"`
	UncollapsibleWidgetsCount int      `json:"uncollapsible_widgets_count"`
	Widgets                   []Widget `json:"widgets"`
}

type Card struct {
	Header    CardHeader `json:"header"`
	Sections  []Section  `json:"sections"`
	Type      string     `json:"type"`
	Body      CardBody   `json:"body"`
	OnClick   OnClick    `json:"on_click"`
	TitleIcon CardIcon   `json:"title_icon"`
	Footer    CardIcon   `json:"footer"`
}

type CardBody struct {
	Description string `json:"description"`
	BigImageUrl string `json:"big_image_url"`
}

type CardV2 struct {
	CardId string `json:"card_id"`
	Card   Card   `json:"card"`
}

const (
	CardTypeApp     = "APP"
	CardTypeTrading = "TRADING"
)

const (
	CardHeaderTypeDefault = "DEFAULT"
	CardHeaderTypeCard    = "CARD"
)

const (
	CardSectionTypeTable   = "TABLE"
	CardSectionTypeDEFAULT = "DEFAULT"
)
