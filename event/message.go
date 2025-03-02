// Copyright (c) 2023 Tulir Asokan
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package event

import (
	"encoding/json"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/net/html"

	"github.com/De-IM/mautrix/crypto/attachment"
	"github.com/De-IM/mautrix/id"
)

// MessageType is the sub-type of a m.room.message event.
// https://spec.matrix.org/v1.2/client-server-api/#mroommessage-msgtypes
type MessageType string

func (mt MessageType) IsText() bool {
	switch mt {
	case MsgText, MsgNotice, MsgEmote:
		return true
	default:
		return false
	}
}

func (mt MessageType) IsMedia() bool {
	switch mt {
	case MsgImage, MsgVideo, MsgAudio, MsgFile, MessageType(EventSticker.Type):
		return true
	default:
		return false
	}
}

// Msgtypes
const (
	MsgText     MessageType = "m.text"
	MsgEmote    MessageType = "m.emote"
	MsgNotice   MessageType = "m.notice"
	MsgImage    MessageType = "m.image"
	MsgLocation MessageType = "m.location"
	MsgVideo    MessageType = "m.video"
	MsgAudio    MessageType = "m.audio"
	MsgFile     MessageType = "m.file"
	MsgCommand  MessageType = "m.command"
	MsgCard     MessageType = "m.card"

	MsgVerificationRequest MessageType = "m.key.verification.request"

	MsgBeeperGallery MessageType = "com.beeper.gallery"
)

// Format specifies the format of the formatted_body in m.room.message events.
// https://spec.matrix.org/v1.2/client-server-api/#mroommessage-msgtypes
type Format string

// Message formats
const (
	FormatHTML Format = "org.matrix.custom.html"
)

// RedactionEventContent represents the content of a m.room.redaction message event.
//
// https://spec.matrix.org/v1.8/client-server-api/#mroomredaction
type RedactionEventContent struct {
	Reason string `json:"reason,omitempty"`

	// The event ID is here as of room v11. In old servers it may only be at the top level.
	Redacts id.EventID `json:"redacts,omitempty"`
}

// ReactionEventContent represents the content of a m.reaction message event.
// This is not yet in a spec release, see https://github.com/matrix-org/matrix-doc/pull/1849
type ReactionEventContent struct {
	RelatesTo RelatesTo `json:"m.relates_to"`
}

func (content *ReactionEventContent) GetRelatesTo() *RelatesTo {
	return &content.RelatesTo
}

func (content *ReactionEventContent) OptionalGetRelatesTo() *RelatesTo {
	return &content.RelatesTo
}

func (content *ReactionEventContent) SetRelatesTo(rel *RelatesTo) {
	content.RelatesTo = *rel
}

// MessageEventContent represents the content of a m.room.message event.
//
// It is also used to represent m.sticker events, as they are equivalent to m.room.message
// with the exception of the msgtype field.
//
// https://spec.matrix.org/v1.2/client-server-api/#mroommessage
type MessageEventContent struct {
	// Base m.room.message fields
	MsgType MessageType `json:"msgtype,omitempty"`
	Body    string      `json:"body"`

	// Extra fields for text types
	Format        Format `json:"format,omitempty"`
	FormattedBody string `json:"formatted_body,omitempty"`

	// Extra field for m.location
	GeoURI string `json:"geo_uri,omitempty"`

	// Extra fields for media types
	URL  id.ContentURIString `json:"url,omitempty"`
	Info *FileInfo           `json:"info,omitempty"`
	File *EncryptedFileInfo  `json:"file,omitempty"`

	FileName string `json:"filename,omitempty"`

	Mentions *Mentions `json:"m.mentions,omitempty"`

	// Edits and relations
	NewContent *MessageEventContent `json:"m.new_content,omitempty"`
	RelatesTo  *RelatesTo           `json:"m.relates_to,omitempty"`

	// In-room verification
	To         id.UserID            `json:"to,omitempty"`
	FromDevice id.DeviceID          `json:"from_device,omitempty"`
	Methods    []VerificationMethod `json:"methods,omitempty"`

	replyFallbackRemoved bool

	MessageSendRetry         *BeeperRetryMetadata     `json:"com.beeper.message_send_retry,omitempty"`
	BeeperGalleryImages      []*MessageEventContent   `json:"com.beeper.gallery.images,omitempty"`
	BeeperGalleryCaption     string                   `json:"com.beeper.gallery.caption,omitempty"`
	BeeperGalleryCaptionHTML string                   `json:"com.beeper.gallery.caption_html,omitempty"`
	BeeperPerMessageProfile  *BeeperPerMessageProfile `json:"com.beeper.per_message_profile,omitempty"`

	BeeperLinkPreviews []*BeeperLinkPreview `json:"com.beeper.linkpreviews,omitempty"`

	MSC1767Audio *MSC1767Audio `json:"org.matrix.msc1767.audio,omitempty"`
	MSC3245Voice *MSC3245Voice `json:"org.matrix.msc3245.voice,omitempty"`

	Components [][]MessageComponent `json:"components,omitempty"`
	BotCommand *BotCommand          `json:"bot_command,omitempty"`
	MsgCardV2  *MsgCardV2           `json:"cards,omitempty"`
	// The flags of the message, which describe extra features of a message.
	// This is a combination of bit masks; the presence of a certain permission can
	// be checked by performing a bitwise AND between this int and the flag.
	Flags MessageFlags `json:"flags"`
}

func (content *MessageEventContent) GetFileName() string {
	if content.FileName != "" {
		return content.FileName
	}
	return content.Body
}

func (content *MessageEventContent) GetCaption() string {
	if content.FileName != "" && content.Body != "" && content.Body != content.FileName {
		return content.Body
	}
	return ""
}

func (content *MessageEventContent) GetFormattedCaption() string {
	if content.Format == FormatHTML && content.FormattedBody != "" {
		return content.FormattedBody
	}
	return ""
}

func (content *MessageEventContent) GetRelatesTo() *RelatesTo {
	if content.RelatesTo == nil {
		content.RelatesTo = &RelatesTo{}
	}
	return content.RelatesTo
}

func (content *MessageEventContent) OptionalGetRelatesTo() *RelatesTo {
	return content.RelatesTo
}

func (content *MessageEventContent) SetRelatesTo(rel *RelatesTo) {
	content.RelatesTo = rel
}

func (content *MessageEventContent) SetEdit(original id.EventID) {
	newContent := *content
	content.NewContent = &newContent
	content.RelatesTo = (&RelatesTo{}).SetReplace(original)
	if content.MsgType == MsgText || content.MsgType == MsgNotice {
		content.Body = "* " + content.Body
		if content.Format == FormatHTML && len(content.FormattedBody) > 0 {
			content.FormattedBody = "* " + content.FormattedBody
		}
		// If the message is long, remove most of the useless edit fallback to avoid event size issues.
		if len(content.Body) > 10000 {
			content.FormattedBody = ""
			content.Format = ""
			content.Body = content.Body[:50] + "[edit fallback cut…]"
		}
	}
}

// TextToHTML converts the given text to a HTML-safe representation by escaping HTML characters
// and replacing newlines with <br/> tags.
func TextToHTML(text string) string {
	return strings.ReplaceAll(html.EscapeString(text), "\n", "<br/>")
}

// ReverseTextToHTML reverses the modifications made by TextToHTML, i.e. replaces <br/> tags with newlines
// and unescapes HTML escape codes. For actually parsing HTML, use the format package instead.
func ReverseTextToHTML(input string) string {
	return html.UnescapeString(strings.ReplaceAll(input, "<br/>", "\n"))
}

func (content *MessageEventContent) EnsureHasHTML() {
	if len(content.FormattedBody) == 0 || content.Format != FormatHTML {
		content.FormattedBody = TextToHTML(content.Body)
		content.Format = FormatHTML
	}
}

func (content *MessageEventContent) GetFile() *EncryptedFileInfo {
	if content.File == nil {
		content.File = &EncryptedFileInfo{}
	}
	return content.File
}

func (content *MessageEventContent) GetInfo() *FileInfo {
	if content.Info == nil {
		content.Info = &FileInfo{}
	}
	return content.Info
}

type Mentions struct {
	UserIDs []id.UserID `json:"user_ids,omitempty"`
	Room    bool        `json:"room,omitempty"`
}

func (m *Mentions) Add(userID id.UserID) {
	if userID != "" && !slices.Contains(m.UserIDs, userID) {
		m.UserIDs = append(m.UserIDs, userID)
	}
}

func (m *Mentions) Has(userID id.UserID) bool {
	return m != nil && slices.Contains(m.UserIDs, userID)
}

type EncryptedFileInfo struct {
	attachment.EncryptedFile
	URL id.ContentURIString `json:"url"`
}

type FileInfo struct {
	MimeType      string              `json:"mimetype,omitempty"`
	ThumbnailInfo *FileInfo           `json:"thumbnail_info,omitempty"`
	ThumbnailURL  id.ContentURIString `json:"thumbnail_url,omitempty"`
	ThumbnailFile *EncryptedFileInfo  `json:"thumbnail_file,omitempty"`

	Blurhash     string `json:"blurhash,omitempty"`
	AnoaBlurhash string `json:"xyz.amorgan.blurhash,omitempty"`

	Width    int `json:"-"`
	Height   int `json:"-"`
	Duration int `json:"-"`
	Size     int `json:"-"`
}

type serializableFileInfo struct {
	MimeType      string                `json:"mimetype,omitempty"`
	ThumbnailInfo *serializableFileInfo `json:"thumbnail_info,omitempty"`
	ThumbnailURL  id.ContentURIString   `json:"thumbnail_url,omitempty"`
	ThumbnailFile *EncryptedFileInfo    `json:"thumbnail_file,omitempty"`

	Blurhash     string `json:"blurhash,omitempty"`
	AnoaBlurhash string `json:"xyz.amorgan.blurhash,omitempty"`

	Width    json.Number `json:"w,omitempty"`
	Height   json.Number `json:"h,omitempty"`
	Duration json.Number `json:"duration,omitempty"`
	Size     json.Number `json:"size,omitempty"`
}

func (sfi *serializableFileInfo) CopyFrom(fileInfo *FileInfo) *serializableFileInfo {
	if fileInfo == nil {
		return nil
	}
	*sfi = serializableFileInfo{
		MimeType:      fileInfo.MimeType,
		ThumbnailURL:  fileInfo.ThumbnailURL,
		ThumbnailInfo: (&serializableFileInfo{}).CopyFrom(fileInfo.ThumbnailInfo),
		ThumbnailFile: fileInfo.ThumbnailFile,

		Blurhash:     fileInfo.Blurhash,
		AnoaBlurhash: fileInfo.AnoaBlurhash,
	}
	if fileInfo.Width > 0 {
		sfi.Width = json.Number(strconv.Itoa(fileInfo.Width))
	}
	if fileInfo.Height > 0 {
		sfi.Height = json.Number(strconv.Itoa(fileInfo.Height))
	}
	if fileInfo.Size > 0 {
		sfi.Size = json.Number(strconv.Itoa(fileInfo.Size))

	}
	if fileInfo.Duration > 0 {
		sfi.Duration = json.Number(strconv.Itoa(int(fileInfo.Duration)))
	}
	return sfi
}

func (sfi *serializableFileInfo) CopyTo(fileInfo *FileInfo) {
	*fileInfo = FileInfo{
		Width:         numberToInt(sfi.Width),
		Height:        numberToInt(sfi.Height),
		Size:          numberToInt(sfi.Size),
		Duration:      numberToInt(sfi.Duration),
		MimeType:      sfi.MimeType,
		ThumbnailURL:  sfi.ThumbnailURL,
		ThumbnailFile: sfi.ThumbnailFile,
		Blurhash:      sfi.Blurhash,
		AnoaBlurhash:  sfi.AnoaBlurhash,
	}
	if sfi.ThumbnailInfo != nil {
		fileInfo.ThumbnailInfo = &FileInfo{}
		sfi.ThumbnailInfo.CopyTo(fileInfo.ThumbnailInfo)
	}
}

func (fileInfo *FileInfo) UnmarshalJSON(data []byte) error {
	sfi := &serializableFileInfo{}
	if err := json.Unmarshal(data, sfi); err != nil {
		return err
	}
	sfi.CopyTo(fileInfo)
	return nil
}

func (fileInfo *FileInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal((&serializableFileInfo{}).CopyFrom(fileInfo))
}

func numberToInt(val json.Number) int {
	f64, _ := val.Float64()
	if f64 > 0 {
		return int(f64)
	}
	return 0
}

func (fileInfo *FileInfo) GetThumbnailInfo() *FileInfo {
	if fileInfo.ThumbnailInfo == nil {
		fileInfo.ThumbnailInfo = &FileInfo{}
	}
	return fileInfo.ThumbnailInfo
}
