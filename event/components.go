package event

import (
	"encoding/json"
	"fmt"
)

// ComponentType is type of component.
type ComponentType uint

// MessageComponent types.
const (
	ActionsRowComponent            ComponentType = 1
	ButtonComponent                ComponentType = 2
	SelectMenuComponent            ComponentType = 3
	TextInputComponent             ComponentType = 4
	UserSelectMenuComponent        ComponentType = 5
	RoleSelectMenuComponent        ComponentType = 6
	MentionableSelectMenuComponent ComponentType = 7
	ChannelSelectMenuComponent     ComponentType = 8
)

// MessageComponent is a base interface for all message components.
type MessageComponent interface {
	json.Marshaler
	Type() ComponentType
}

type unmarshalableMessageComponent struct {
	MessageComponent
}

// UnmarshalJSON is a helper function to unmarshal MessageComponent object.
func (umc *unmarshalableMessageComponent) UnmarshalJSON(src []byte) error {
	var v struct {
		Type ComponentType `json:"type"`
	}
	err := json.Unmarshal(src, &v)
	if err != nil {
		return err
	}

	switch v.Type {
	case ActionsRowComponent:
		umc.MessageComponent = &ActionsRow{}
	case ButtonComponent:
		umc.MessageComponent = &Button{}
	case SelectMenuComponent, ChannelSelectMenuComponent, UserSelectMenuComponent,
		RoleSelectMenuComponent, MentionableSelectMenuComponent:
		umc.MessageComponent = &SelectMenu{}
	case TextInputComponent:
		umc.MessageComponent = &TextInput{}
	default:
		return fmt.Errorf("unknown component type: %d", v.Type)
	}
	return json.Unmarshal(src, umc.MessageComponent)
}

// MessageComponentFromJSON is a helper function for unmarshaling message components
func MessageComponentFromJSON(b []byte) (MessageComponent, error) {
	var u unmarshalableMessageComponent
	err := u.UnmarshalJSON(b)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal into MessageComponent: %w", err)
	}
	return u.MessageComponent, nil
}

// ActionsRow is a container for components within one row.
type ActionsRow struct {
	Components []MessageComponent `json:"components"`
}

// MarshalJSON is a method for marshaling ActionsRow to a JSON object.
func (r ActionsRow) MarshalJSON() ([]byte, error) {
	type actionsRow ActionsRow

	return Marshal(struct {
		actionsRow
		Type ComponentType `json:"type"`
	}{
		actionsRow: actionsRow(r),
		Type:       r.Type(),
	})
}

// UnmarshalJSON is a helper function to unmarshal Actions Row.
func (r *ActionsRow) UnmarshalJSON(data []byte) error {
	var v struct {
		RawComponents []unmarshalableMessageComponent `json:"components"`
	}
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	r.Components = make([]MessageComponent, len(v.RawComponents))
	for i, v := range v.RawComponents {
		r.Components[i] = v.MessageComponent
	}

	return err
}

// Type is a method to get the type of a component.
func (r ActionsRow) Type() ComponentType {
	return ActionsRowComponent
}

// ButtonStyle is style of button.
type ButtonStyle uint

// Button styles.
const (
	// PrimaryButton is a button with blurple color.
	PrimaryButton ButtonStyle = 1
	// SecondaryButton is a button with grey color.
	SecondaryButton ButtonStyle = 2
	// SuccessButton is a button with green color.
	SuccessButton ButtonStyle = 3
	// DangerButton is a button with red color.
	DangerButton ButtonStyle = 4
	// LinkButton is a special type of button which navigates to a URL. Has grey color.
	LinkButton ButtonStyle = 5
)

// ComponentEmoji represents button emoji, if it does have one.
type ComponentEmoji struct {
	Name     string `json:"name,omitempty"`
	ID       string `json:"id,omitempty"`
	Animated bool   `json:"animated,omitempty"`
}

// Button represents button component.
type Button struct {
	// Text label text on the button
	Text string `json:"text,omitempty"`
	// LoginURL is an HTTP URL used to automatically authorize the user. Can be
	// used as a replacement for the Telegram Login Widget
	//
	// optional
	LoginURL *LoginURL `json:"login_url,omitempty"`
	// CallbackData data to be sent in a callback query to the bot when button is pressed, 1-64 bytes.
	//
	// optional
	CallbackData *string `json:"callback_data,omitempty"`
	// WebApp if specified, the described Web App will be launched when the button
	// is pressed. The Web App will be able to send a “web_app_data” service
	// message. Available in private chats only.
	//
	// optional
	WebApp *WebAppInfo `json:"web_app,omitempty"`
	// SwitchInlineQuery if set, pressing the button will prompt the user to select one of their chats,
	// open that chat and insert the bot's username and the specified inline query in the input field.
	// Can be empty, in which case just the bot's username will be inserted.
	//
	// This offers an easy way for users to start using your bot
	// in inline mode when they are currently in a private chat with it.
	// Especially useful when combined with switch_pm… actions – in this case
	// the user will be automatically returned to the chat they switched from,
	// skipping the chat selection screen.
	//
	// optional
	SwitchInlineQuery *string `json:"switch_inline_query,omitempty"`
	// SwitchInlineQueryCurrentChat if set, pressing the button will insert the bot's username
	// and the specified inline query in the current chat's input field.
	// Can be empty, in which case only the bot's username will be inserted.
	//
	// This offers a quick way for the user to open your bot in inline mode
	// in the same chat – good for selecting something from multiple options.
	//
	// optional
	SwitchInlineQueryCurrentChat *string `json:"switch_inline_query_current_chat,omitempty"`
	// CallbackGame description of the game that will be launched when the user presses the button.
	//
	// optional
	CallbackGame *CallbackGame `json:"callback_game,omitempty"`
	// Pay specify True, to send a Pay button.
	//
	// NOTE: This type of button must always be the first button in the first row.
	//
	// optional
	Pay   bool            `json:"pay,omitempty"`
	Style ButtonStyle     `json:"style,omitempty"`
	Emoji *ComponentEmoji `json:"emoji,omitempty"`

	// NOTE: Only button with LinkButton style can have link. Also, URL is mutually exclusive with CustomID.
	URL string `json:"url,omitempty"`
}

type WebAppInfo struct {
	// URL is the HTTPS URL of a Web App to be opened with additional data as
	// specified in Initializing Web Apps.
	URL string `json:"url"`
}

// CallbackGame is for starting a game in an inline keyboard button.
type CallbackGame struct{}

// LoginURL represents a parameter of the inline keyboard button used to
// automatically authorize a user. Serves as a great replacement for the
// Telegram Login Widget when the user is coming from Telegram. All the user
// needs to do is tap/click a button and confirm that they want to log in.
type LoginURL struct {
	// URL is an HTTP URL to be opened with user authorization data added to the
	// query string when the button is pressed. If the user refuses to provide
	// authorization data, the original URL without information about the user
	// will be opened. The data added is the same as described in Receiving
	// authorization data.
	//
	// NOTE: You must always check the hash of the received data to verify the
	// authentication and the integrity of the data as described in Checking
	// authorization.
	URL string `json:"url"`
	// ForwardText is the new text of the button in forwarded messages
	//
	// optional
	ForwardText string `json:"forward_text,omitempty"`
	// BotUsername is the username of a bot, which will be used for user
	// authorization. See Setting up a bot for more details. If not specified,
	// the current bot's username will be assumed. The url's domain must be the
	// same as the domain linked with the bot. See Linking your domain to the
	// bot for more details.
	//
	// optional
	BotUsername string `json:"bot_username,omitempty"`
	// RequestWriteAccess if true requests permission for your bot to send
	// messages to the user
	//
	// optional
	RequestWriteAccess bool `json:"request_write_access,omitempty"`
}

// MarshalJSON is a method for marshaling Button to a JSON object.
func (b Button) MarshalJSON() ([]byte, error) {
	type button Button

	if b.Style == 0 {
		b.Style = PrimaryButton
	}

	return Marshal(struct {
		button
		Type ComponentType `json:"type"`
	}{
		button: button(b),
		Type:   b.Type(),
	})
}

// Type is a method to get the type of a component.
func (Button) Type() ComponentType {
	return ButtonComponent
}

// SelectMenuOption represents an option for a select menu.
type SelectMenuOption struct {
	Label       string          `json:"label,omitempty"`
	Value       string          `json:"value"`
	Description string          `json:"description"`
	Emoji       *ComponentEmoji `json:"emoji,omitempty"`
	// Determines whenever option is selected by default or not.
	Default bool `json:"default"`
}

// SelectMenuDefaultValueType represents the type of an entity selected by default in auto-populated select menus.
type SelectMenuDefaultValueType string

// SelectMenuDefaultValue types.
const (
	SelectMenuDefaultValueUser    SelectMenuDefaultValueType = "user"
	SelectMenuDefaultValueRole    SelectMenuDefaultValueType = "role"
	SelectMenuDefaultValueChannel SelectMenuDefaultValueType = "channel"
)

// SelectMenuDefaultValue represents an entity selected by default in auto-populated select menus.
type SelectMenuDefaultValue struct {
	// ID of the entity.
	ID string `json:"id"`
	// Type of the entity.
	Type SelectMenuDefaultValueType `json:"type"`
}

// SelectMenuType represents select menu type.
type SelectMenuType ComponentType

// SelectMenu types.
const (
	StringSelectMenu      = SelectMenuType(SelectMenuComponent)
	UserSelectMenu        = SelectMenuType(UserSelectMenuComponent)
	RoleSelectMenu        = SelectMenuType(RoleSelectMenuComponent)
	MentionableSelectMenu = SelectMenuType(MentionableSelectMenuComponent)
	ChannelSelectMenu     = SelectMenuType(ChannelSelectMenuComponent)
)

// SelectMenu represents select menu component.
type SelectMenu struct {
	// Type of the select menu.
	MenuType SelectMenuType `json:"type,omitempty"`
	// CustomID is a developer-defined identifier for the select menu.
	CustomID string `json:"custom_id,omitempty"`
	// The text which will be shown in the menu if there's no default options or all options was deselected and component was closed.
	Placeholder string `json:"placeholder"`
	// This value determines the minimal amount of selected items in the menu.
	MinValues *int `json:"min_values,omitempty"`
	// This value determines the maximal amount of selected items in the menu.
	// If MaxValues or MinValues are greater than one then the user can select multiple items in the component.
	MaxValues int `json:"max_values,omitempty"`
	// List of default values for auto-populated select menus.
	// NOTE: Number of entries should be in the range defined by MinValues and MaxValues.
	DefaultValues []SelectMenuDefaultValue `json:"default_values,omitempty"`

	Options  []SelectMenuOption `json:"options,omitempty"`
	Disabled bool               `json:"disabled"`

	// NOTE: Can only be used in SelectMenu with Channel menu type.
	ChannelTypes []ChannelType `json:"channel_types,omitempty"`
}

// Type is a method to get the type of a component.
func (s SelectMenu) Type() ComponentType {
	if s.MenuType != 0 {
		return ComponentType(s.MenuType)
	}
	return SelectMenuComponent
}

// MarshalJSON is a method for marshaling SelectMenu to a JSON object.
func (s SelectMenu) MarshalJSON() ([]byte, error) {
	type selectMenu SelectMenu

	return Marshal(struct {
		selectMenu
		Type ComponentType `json:"type"`
	}{
		selectMenu: selectMenu(s),
		Type:       s.Type(),
	})
}

// TextInput represents text input component.
type TextInput struct {
	CustomID    string         `json:"custom_id"`
	Label       string         `json:"label"`
	Style       TextInputStyle `json:"style"`
	Placeholder string         `json:"placeholder,omitempty"`
	Value       string         `json:"value,omitempty"`
	Required    bool           `json:"required"`
	MinLength   int            `json:"min_length,omitempty"`
	MaxLength   int            `json:"max_length,omitempty"`
}

// Type is a method to get the type of a component.
func (TextInput) Type() ComponentType {
	return TextInputComponent
}

// MarshalJSON is a method for marshaling TextInput to a JSON object.
func (m TextInput) MarshalJSON() ([]byte, error) {
	type inputText TextInput

	return Marshal(struct {
		inputText
		Type ComponentType `json:"type"`
	}{
		inputText: inputText(m),
		Type:      m.Type(),
	})
}

// TextInputStyle is style of text in TextInput component.
type TextInputStyle uint

// Text styles
const (
	TextInputShort     TextInputStyle = 1
	TextInputParagraph TextInputStyle = 2
)
