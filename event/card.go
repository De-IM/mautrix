package event

type CardHeader struct {
	Type           string   `json:"type,omitempty"`
	Title          string   `json:"title,omitempty"`
	Subtitle       string   `json:"subtitle,omitempty"`
	SubtitleAction *OnClick `json:"subtitleAction,omitempty"`
	ImageUrl       string   `json:"imageUrl,omitempty"`
	ImageType      string   `json:"imageType,omitempty"`
	ImageAltText   string   `json:"imageAltText,omitempty"`
}

type DecoratedText struct {
	StartIcon *CardIcon `json:"startIcon,omitempty"`
	Text      string    `json:"text,omitempty"`
	Texts     []string  `json:"texts,omitempty"`
}

type CardTitle struct {
	Title     string    `json:"title,omitempty"`
	StartIcon *CardIcon `json:"startIcon,omitempty"`
	EndIcon   *CardIcon `json:"endIcon,omitempty"`
}

type CardIcon struct {
	KnownIcon string `json:"knownIcon,omitempty"`
	Url       string `json:"url,omitempty"`
}

type CardButton struct {
	Text    string   `json:"text,omitempty"`
	OnClick *OnClick `json:"onClick,omitempty"`
}

type OnClick struct {
	OpenLink *OpenLink `json:"openLink,omitempty"`
	Action   *Action   `json:"action,omitempty"`
}

type OpenLink struct {
	URL string `json:"url,omitempty"`
}

type Action struct {
	Function   string       `json:"function,omitempty"`
	Parameters []*Parameter `json:"parameters,omitempty"`
}

type Parameter struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type ButtonList struct {
	Buttons []*CardButton `json:"buttons,omitempty"`
}

type Widget struct {
	DecoratedText *DecoratedText `json:"decoratedText,omitempty"`
	ButtonList    *ButtonList    `json:"buttonList,omitempty"`
}

// BorderStyle represents the border style in the image structure
type BorderStyle struct {
	Type string `json:"type"`
}

// CropStyle represents the crop style in the image structure
type CropStyle struct {
	Type string `json:"type"`
}

// Image represents the image structure with URI and styles
type Image struct {
	ImageURI    string       `json:"imageUri"`
	CropStyle   *CropStyle   `json:"cropStyle,omitempty"`
	BorderStyle *BorderStyle `json:"borderStyle,omitempty"`
}

// Item represents a single item in the grid
type GridItem struct {
	Image         Image  `json:"image"`
	Title         string `json:"title"`
	TextAlignment string `json:"textAlignment"`
}

// Grid represents the grid structure
type Grid struct {
	Title       string     `json:"title"`
	ColumnCount int        `json:"columnCount"`
	Items       []GridItem `json:"items"`
	OnClick     OnClick    `json:"onClick"`
}

type Section struct {
	Header                    string    `json:"header,omitempty"`
	Collapsible               bool      `json:"collapsible,omitempty"`
	UncollapsibleWidgetsCount int       `json:"uncollapsibleWidgetsCount,omitempty"`
	Widgets                   []*Widget `json:"widgets,omitempty"`
	Type                      string    `json:"type,omitempty"`
	Footer                    string    `json:"footer,omitempty"`
}

type Card struct {
	Header   *CardHeader `json:"header,omitempty"`
	Sections []*Section  `json:"sections,omitempty"`
	Type     string      `json:"type,omitempty"`
	Body     *CardBody   `json:"body,omitempty"`
	OnClick  *OnClick    `json:"onClick,omitempty"`
	Title    *CardTitle  `json:"title,omitempty"`
	Footer   *CardIcon   `json:"footer,omitempty"`
}

type CardBody struct {
	Description string `json:"description,omitempty"`
	BigImageUrl string `json:"bigImageUrl,omitempty"`
}

type CardV2 struct {
	CardId string `json:"cardId,omitempty"`
	Card   *Card  `json:"card,omitempty"`
}

type MsgCardV2 struct {
	CardsV2 []*CardV2 `json:"cardsV2,omitempty"`
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
