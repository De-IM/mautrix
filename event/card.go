package event

type CardHeader struct {
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle"`
	ImageUrl     string `json:"image_url"`
	ImageType    string `json:"image_type"`
	ImageAltText string `json:"image_alt_text"`
}

type DecoratedText struct {
	StartIcon Icon   `json:"start_icon"`
	Text      string `json:"text"`
}

type Icon struct {
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
	Header   CardHeader `json:"header"`
	Sections []Section  `json:"sections"`
}

type CardV2 struct {
	CardId string `json:"card_id"`
	Card   Card   `json:"card"`
}
