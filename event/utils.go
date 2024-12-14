package event

import (
	"encoding/json"
	"errors"
)

// All error constants
var (
	ErrJSONUnmarshal           = errors.New("json unmarshal")
	ErrStatusOffline           = errors.New("You can't set your Status to offline")
	ErrVerificationLevelBounds = errors.New("VerificationLevel out of bounds, should be between 0 and 3")
	ErrPruneDaysBounds         = errors.New("the number of days should be more than or equal to 1")
	ErrGuildNoIcon             = errors.New("guild does not have an icon set")
	ErrGuildNoSplash           = errors.New("guild does not have a splash set")
	ErrUnauthorized            = errors.New("HTTP request was unauthorized. This could be because the provided token was not a bot token. Please add \"Bot \" to the start of your token. https://discord.com/developers/docs/reference#authentication-example-bot-token-authorization-header")
)

var (
	// Marshal defines function used to encode JSON payloads
	Marshal func(v interface{}) ([]byte, error) = json.Marshal
	// Unmarshal defines function used to decode JSON payloads
	Unmarshal func(src []byte, v interface{}) error = json.Unmarshal
)

// ChannelType is the type of a Channel
type ChannelType int

// Block contains known ChannelType values
const (
	ChannelTypeGuildText          ChannelType = 0
	ChannelTypeDM                 ChannelType = 1
	ChannelTypeGuildVoice         ChannelType = 2
	ChannelTypeGroupDM            ChannelType = 3
	ChannelTypeGuildCategory      ChannelType = 4
	ChannelTypeGuildNews          ChannelType = 5
	ChannelTypeGuildStore         ChannelType = 6
	ChannelTypeGuildNewsThread    ChannelType = 10
	ChannelTypeGuildPublicThread  ChannelType = 11
	ChannelTypeGuildPrivateThread ChannelType = 12
	ChannelTypeGuildStageVoice    ChannelType = 13
	ChannelTypeGuildDirectory     ChannelType = 14
	ChannelTypeGuildForum         ChannelType = 15
	ChannelTypeGuildMedia         ChannelType = 16
)
