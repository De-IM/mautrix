package event

// MessageFlags is the flags of "message" (see MessageFlags* consts)
// https://discord.com/developers/docs/resources/channel#message-object-message-flags
type MessageFlags int

// Valid MessageFlags values
const (
	// MessageFlagsCrossPosted This message has been published to subscribed channels (via Channel Following).
	MessageFlagsCrossPosted MessageFlags = 1 << 0
	// MessageFlagsIsCrossPosted this message originated from a message in another channel (via Channel Following).
	MessageFlagsIsCrossPosted MessageFlags = 1 << 1
	// MessageFlagsSuppressEmbeds do not include any embeds when serializing this message.
	MessageFlagsSuppressEmbeds MessageFlags = 1 << 2
	// TODO: deprecated, remove when compatibility is not needed
	MessageFlagsSupressEmbeds MessageFlags = 1 << 2
	// MessageFlagsSourceMessageDeleted the source message for this crosspost has been deleted (via Channel Following).
	MessageFlagsSourceMessageDeleted MessageFlags = 1 << 3
	// MessageFlagsUrgent this message came from the urgent message system.
	MessageFlagsUrgent MessageFlags = 1 << 4
	// MessageFlagsHasThread this message has an associated thread, with the same id as the message.
	MessageFlagsHasThread MessageFlags = 1 << 5
	// MessageFlagsEphemeral this message is only visible to the user who invoked the Interaction.
	MessageFlagsEphemeral MessageFlags = 1 << 6
	// MessageFlagsLoading this message is an Interaction Response and the bot is "thinking".
	MessageFlagsLoading MessageFlags = 1 << 7
	// MessageFlagsFailedToMentionSomeRolesInThread this message failed to mention some roles and add their members to the thread.
	MessageFlagsFailedToMentionSomeRolesInThread MessageFlags = 1 << 8
	// MessageFlagsSuppressNotifications this message will not trigger push and desktop notifications.
	MessageFlagsSuppressNotifications MessageFlags = 1 << 12
	// MessageFlagsIsVoiceMessage this message is a voice message.
	MessageFlagsIsVoiceMessage MessageFlags = 1 << 13
)
