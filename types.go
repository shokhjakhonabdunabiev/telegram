package telegram

type User struct {
	ID                        int64  `json:"id"`
	IsBot                     bool   `json:"is_bot"`
	FirstName                 string `json:"first_name"`
	LastName                  string `json:"last_name,omitempty"`
	Username                  string `json:"username,omitempty"`
	LanguageCode              string `json:"language_code,omitempty"`
	IsPremium                 bool   `json:"is_premium,omitempty"`
	AddedToAttachmentMenu     bool   `json:"added_to_attachment_menu,omitempty"`
	CanJoinGroups             bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages   bool   `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries     bool   `json:"supports_inline_queries,omitempty"`
	CanConnectToBusiness      bool   `json:"can_connect_to_business,omitempty"`
	HasMainWebApp             bool   `json:"has_main_web_app,omitempty"`
	HasTopicsEnabled          bool   `json:"has_topics_enabled,omitempty"`
	AllowsUsersToCreateTopics bool   `json:"allows_users_to_create_topics,omitempty"`
}

type Chat struct {
	ID               int64  `json:"id"`
	Type             string `json:"type"`
	Title            string `json:"title,omitempty"`
	Username         string `json:"username,omitempty"`
	FirstName        string `json:"first_name,omitempty"`
	LastName         string `json:"last_name,omitempty"`
	IsForum          bool   `json:"is_forum,omitempty"`
	IsDirectMessages bool   `json:"is_direct_messages,omitempty"`
}

type ChatFullInfo struct {
	ID                                 int64                 `json:"id"`
	Type                               string                `json:"type"`
	Title                              string                `json:"title,omitempty"`
	Username                           string                `json:"username,omitempty"`
	FirstName                          string                `json:"first_name,omitempty"`
	LastName                           string                `json:"last_name,omitempty"`
	IsForum                            bool                  `json:"is_forum,omitempty"`
	IsDirectMessages                   bool                  `json:"is_direct_messages,omitempty"`
	AccentColorID                      int                   `json:"accent_color_id"`
	MaxReactionCount                   int                   `json:"max_reaction_count"`
	Photo                              *ChatPhoto            `json:"photo,omitempty"`
	ActiveUsernames                    []string              `json:"active_usernames,omitempty"`
	Birthdate                          *Birthdate            `json:"birthdate,omitempty"`
	BusinessIntro                      *BusinessIntro        `json:"business_intro,omitempty"`
	BusinessLocation                   *BusinessLocation     `json:"business_location,omitempty"`
	BusinessOpeningHours               *BusinessOpeningHours `json:"business_opening_hours,omitempty"`
	PersonalChat                       *Chat                 `json:"personal_chat,omitempty"`
	ParentChat                         *Chat                 `json:"parent_chat,omitempty"`
	Available_reactions                []*ReactionType       `json:"available_reactions,omitempty"`
	Background_custom_emoji_id         string                `json:"background_custom_emoji_id,omitempty"`
	ProfileAccentColorID               int                   `json:"profile_accent_color_id,omitempty"`
	ProfileBackgroundCustomEmojiID     string                `json:"profile_background_custom_emoji_id,omitempty"`
	EmojiStatusCustomEmojiID           string                `json:"emoji_status_custom_emoji_id,omitempty"`
	EmojiStatusExpirationDate          int                   `json:"emoji_status_expiration_date,omitempty"`
	Bio                                string                `json:"bio,omitempty"`
	HasPrivateForwards                 bool                  `json:"has_private_forwards,omitempty"`
	HasRestrictedVoiceAndVideoMessages bool                  `json:"has_restricted_voice_and_video_messages,omitempty"`
	JoinToSendMessages                 bool                  `json:"join_to_send_messages,omitempty"`
	JoinByRequest                      bool                  `json:"join_by_request,omitempty"`
	Description                        string                `json:"description,omitempty"`
	InviteLink                         string                `json:"invite_link,omitempty"`
	PinnedMessage                      *Message              `json:"pinned_message,omitempty"`
	Permissions                        *ChatPermissions      `json:"permissions,omitempty"`
	AcceptedGiftTypes                  *AcceptedGiftTypes    `json:"accepted_gift_types,omitempty"`
	CanSendPaidMedia                   bool                  `json:"can_send_paid_media,omitempty"`
	SlowModeDelay                      int                   `json:"slow_mode_delay,omitempty"`
	UnrestrictBoostCount               int                   `json:"unrestrict_boost_count,omitempty"`
	MessageAutoDeleteTime              int                   `json:"message_auto_delete_time,omitempty"`
	HasAggressiveAntiSpamEnabled       bool                  `json:"has_aggressive_anti_spam_enabled,omitempty"`
	HasHiddenMembers                   bool                  `json:"has_hidden_members,omitempty"`
	HasProtectedContent                bool                  `json:"has_protected_content,omitempty"`
	HasVisibleHistory                  bool                  `json:"has_visible_history,omitempty"`
	StickerSetName                     string                `json:"sticker_set_name,omitempty"`
	CanSetStickerSet                   bool                  `json:"can_set_sticker_set,omitempty"`
	CustomEmojiStickerSetName          string                `json:"custom_emoji_sticker_set_name,omitempty"`
	LinkedChatID                       int                   `json:"linked_chat_id,omitempty"`
	Location                           *ChatLocation         `json:"location,omitempty"`
	Rating                             *UserRating           `json:"rating,omitempty"`
	FirstProfileAudio                  *Audio                `json:"first_profile_audio,omitempty"`
	UniqueGiftColors                   *UniqueGiftColors     `json:"unique_gift_colors,omitempty"`
	PaidMessageStarCount               int                   `json:"paid_message_star_count,omitempty"`
}

type Message struct {
	MessageID                     int                            `json:"message_id"`
	MessageThreadID               int                            `json:"message_thread_id,omitempty"`
	DirectMessagesTopic           *DirectMessagesTopic           `json:"direct_messages_topic,omitempty"`
	From                          *User                          `json:"from,omitempty"`
	SenderChat                    *Chat                          `json:"sender_chat,omitempty"`
	SenderBoostCount              int                            `json:"sender_boost_count,omitempty"`
	SenderBusinessBot             *User                          `json:"sender_business_bot,omitempty"`
	SenderTag                     string                         `json:"sender_tag,omitempty"`
	Date                          int64                          `json:"date"`
	BusinessConnectionID          string                         `json:"business_connection_id,omitempty"`
	Chat                          *Chat                          `json:"chat"`
	ForwardOrigin                 *MessageOrigin                 `json:"forward_origin,omitempty"`
	IsTopicMessage                bool                           `json:"is_topic_message,omitempty"`
	IsAutomaticForward            bool                           `json:"is_automatic_forward,omitempty"`
	ReplyToMessage                *Message                       `json:"reply_to_message,omitempty"`
	ExternalReply                 *ExternalReplyInfo             `json:"external_reply,omitempty"`
	Quote                         *TextQuote                     `json:"quote,omitempty"`
	ReplyToStory                  *Story                         `json:"reply_to_story,omitempty"`
	ReplyToChecklistTaskID        int                            `json:"reply_to_checklist_task_id,omitempty"`
	ViaBot                        *User                          `json:"via_bot,omitempty"`
	EditDate                      int64                          `json:"edit_date,omitempty"`
	HasProtectedContent           bool                           `json:"has_protected_content,omitempty"`
	IsFromOffline                 bool                           `json:"is_from_offline,omitempty"`
	IsPaidPost                    bool                           `json:"is_paid_post,omitempty"`
	MediaGroupID                  string                         `json:"media_group_id,omitempty"`
	AuthorSignature               string                         `json:"author_signature,omitempty"`
	PaidStarCount                 int                            `json:"paid_star_count,omitempty"`
	Text                          string                         `json:"text,omitempty"`
	Entities                      []*MessageEntity               `json:"entities,omitempty"`
	LinkPreviewOptions            *LinkPreviewOptions            `json:"link_preview_options,omitempty"`
	SuggestedPostInfo             *SuggestedPostInfo             `json:"suggested_post_info,omitempty"`
	EffectID                      string                         `json:"effect_id,omitempty"`
	Animation                     *Animation                     `json:"animation,omitempty"`
	Audio                         *Audio                         `json:"audio,omitempty"`
	Document                      *Document                      `json:"document,omitempty"`
	PaidMedia                     *PaidMediaInfo                 `json:"paid_media,omitempty"`
	Photo                         []*PhotoSize                   `json:"photo,omitempty"`
	Sticker                       *Sticker                       `json:"sticker,omitempty"`
	Story                         *Story                         `json:"story,omitempty"`
	Video                         *Video                         `json:"video,omitempty"`
	VideoNote                     *VideoNote                     `json:"video_note,omitempty"`
	Voice                         *Voice                         `json:"voice,omitempty"`
	Caption                       string                         `json:"caption,omitempty"`
	CaptionEntities               []*MessageEntity               `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia         bool                           `json:"show_caption_above_media,omitempty"`
	HasMediaSpoiler               bool                           `json:"has_media_spoiler,omitempty"`
	Checklist                     *Checklist                     `json:"checklist,omitempty"`
	Contact                       *Contact                       `json:"contact,omitempty"`
	Dice                          *Dice                          `json:"dice,omitempty"`
	Game                          *Game                          `json:"game,omitempty"`
	Poll                          *Poll                          `json:"poll,omitempty"`
	Venue                         *Venue                         `json:"venue,omitempty"`
	Location                      *Location                      `json:"location,omitempty"`
	NewChatMembers                []*User                        `json:"new_chat_members,omitempty"`
	LeftChatMember                *User                          `json:"left_chat_member,omitempty"`
	ChatOwnerLeft                 *ChatOwnerLeft                 `json:"chat_owner_left,omitempty"`
	ChatOwnerChanged              *ChatOwnerChanged              `json:"chat_owner_changed,omitempty"`
	NewChatTitle                  string                         `json:"new_chat_title,omitempty"`
	NewChatPhoto                  []*PhotoSize                   `json:"new_chat_photo,omitempty"`
	DeleteChatPhoto               bool                           `json:"delete_chat_photo,omitempty"`
	GroupChatCreated              bool                           `json:"group_chat_created,omitempty"`
	SupergroupChatCreated         bool                           `json:"supergroup_chat_created,omitempty"`
	ChannelChatCreated            bool                           `json:"channel_chat_created,omitempty"`
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	MigrateToChatID               int                            `json:"migrate_to_chat_id,omitempty"`
	MigrateFromChatID             int                            `json:"migrate_from_chat_id,omitempty"`
	PinnedMessage                 *MaybeInaccessibleMessage      `json:"pinned_message,omitempty"`
	Invoice                       *Invoice                       `json:"invoice,omitempty"`
	SuccessfulPayment             *SuccessfulPayment             `json:"successful_payment,omitempty"`
	RefundedPayment               *RefundedPayment               `json:"refunded_payment,omitempty"`
	UsersShared                   *UsersShared                   `json:"users_shared,omitempty"`
	ChatShared                    *ChatShared                    `json:"chat_shared,omitempty"`
	Gift                          *GiftInfo                      `json:"gift,omitempty"`
	UniqueGift                    *UniqueGiftInfo                `json:"unique_gift,omitempty"`
	GiftUpgradeSent               *GiftInfo                      `json:"gift_upgrade_sent,omitempty"`
	ConnectedWebsite              string                         `json:"connected_website,omitempty"`
	WriteAccessAllowed            *WriteAccessAllowed            `json:"write_access_allowed,omitempty"`
	PassportData                  *PassportData                  `json:"passport_data,omitempty"`
	ProximityAlertTriggered       *ProximityAlertTriggered       `json:"proximity_alert_triggered,omitempty"`
	BoostAdded                    *ChatBoostAdded                `json:"boost_added,omitempty"`
	ChatBackgroundSet             *ChatBackground                `json:"chat_background_set,omitempty"`
	ChecklistTasksDone            *ChecklistTasksDone            `json:"checklist_tasks_done,omitempty"`
	ChecklistTasksAdded           *ChecklistTasksAdded           `json:"checklist_tasks_added,omitempty"`
	DirectMessagePriceChanged     *DirectMessagePriceChanged     `json:"direct_message_price_changed,omitempty"`
	ForumTopicCreated             *ForumTopicCreated             `json:"forum_topic_created,omitempty"`
	ForumTopicEdited              *ForumTopicEdited              `json:"forum_topic_edited,omitempty"`
	ForumTopicClosed              *ForumTopicClosed              `json:"forum_topic_closed,omitempty"`
	ForumTopicReopened            *ForumTopicReopened            `json:"forum_topic_reopened,omitempty"`
	GeneralForumTopicHidden       *GeneralForumTopicHidden       `json:"general_forum_topic_hidden,omitempty"`
	GeneralForumTopicUnhidden     *GeneralForumTopicUnhidden     `json:"general_forum_topic_unhidden,omitempty"`
	GiveawayCreated               *GiveawayCreated               `json:"giveaway_created,omitempty"`
	Giveaway                      *Giveaway                      `json:"giveaway,omitempty"`
	GiveawayWinners               *GiveawayWinners               `json:"giveaway_winners,omitempty"`
	GiveawayCompleted             *GiveawayCompleted             `json:"giveaway_completed,omitempty"`
	PaidMessagePriceChanged       *PaidMessagePriceChanged       `json:"paid_message_price_changed,omitempty"`
	SuggestedPostApproved         *SuggestedPostApproved         `json:"suggested_post_approved,omitempty"`
	SuggestedPostApprovalFailed   *SuggestedPostApprovalFailed   `json:"suggested_post_approval_failed,omitempty"`
	SuggestedPostDeclined         *SuggestedPostDeclined         `json:"suggested_post_declined,omitempty"`
	SuggestedPostPaid             *SuggestedPostPaid             `json:"suggested_post_paid,omitempty"`
	SuggestedPostRefunded         *SuggestedPostRefunded         `json:"suggested_post_refunded,omitempty"`
	VideoChatScheduled            *VideoChatScheduled            `json:"video_chat_scheduled,omitempty"`
	VideoChatStarted              *VideoChatStarted              `json:"video_chat_started,omitempty"`
	VideoChatEnded                *VideoChatEnded                `json:"video_chat_ended,omitempty"`
	VideoChatParticipantsInvited  *VideoChatParticipantsInvited  `json:"video_chat_participants_invited,omitempty"`
	WebAppData                    *WebAppData                    `json:"web_app_data,omitempty"`
	ReplyMarkup                   *InlineKeyboardMarkup          `json:"reply_markup,omitempty"`
}

// This object represents a unique message identifier.
type MessageId struct {
	// Unique message identifier.
	// In specific instances (e.g., message containing a video sent to a big chat),
	// the server might automatically schedule a message instead of sending it immediately.
	// In such cases, this field will be 0 and the relevant message will be unusable until it is actually sent
	MessageID int `json:"message_id"`
}

type InaccessibleMessage struct {
	Chat      *Chat `json:"chat"`
	MessageID int   `json:"message_id"`
	Date      int64 `json:"date"`
}

// TODO
type MaybeInaccessibleMessage struct{}

type MessageEntity struct {
	Type           string `json:"type"`
	Offset         int    `json:"offset"`
	Length         int    `json:"length"`
	URL            string `json:"url,omitempty"`
	User           *User  `json:"user,omitempty"`
	Language       string `json:"language,omitempty"`
	CustomEmojiID  string `json:"custom_emoji_id,omitempty"`
	UnixTime       int64  `json:"unix_time,omitempty"`
	DateTimeFormat string `json:"date_time_format,omitempty"`
}

type TextQuote struct {
	Text     string           `json:"text"`
	Entities []*MessageEntity `json:"entities,omitempty"`
	Position int              `json:"position"`
	IsManual bool             `json:"is_manual,omitempty"`
}

type ExternalReplyInfo struct {
	Origin             *MessageOrigin      `json:"origin"`
	Chat               *Chat               `json:"chat,omitempty"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	Animation          *Animation          `json:"animation,omitempty"`
	Audio              *Audio              `json:"audio,omitempty"`
	Document           *Document           `json:"document,omitempty"`
	PaidMedia          *PaidMediaInfo      `json:"paid_media,omitempty"`
	Photo              []*PhotoSize        `json:"photo,omitempty"`
	Sticker            *Sticker            `json:"sticker,omitempty"`
	Story              *Story              `json:"story,omitempty"`
	Video              *Video              `json:"video,omitempty"`
	VideoNote          *VideoNote          `json:"video_note,omitempty"`
	Voice              *Voice              `json:"voice,omitempty"`
	HasMediaSpoiler    bool                `json:"has_media_spoiler,omitempty"`
	Checklist          *Checklist          `json:"checklist,omitempty"`
	Contact            *Contact            `json:"contact,omitempty"`
	Dice               *Dice               `json:"dice,omitempty"`
	Game               *Game               `json:"game,omitempty"`
	Giveaway           *Giveaway           `json:"giveaway,omitempty"`
	GiveawayWinners    *GiveawayWinners    `json:"giveaway_winners,omitempty"`
	Invoice            *Invoice            `json:"invoice,omitempty"`
	Location           *Location           `json:"location,omitempty"`
	Poll               *Poll               `json:"poll,omitempty"`
	Venue              *Venue              `json:"venue,omitempty"`
}

// TODO
type MessageOrigin struct {
	Type string `json:"type"`
	Date int64  `json:"date"`

	// user
	SenderUser *User `json:"sender_user,omitempty"`

	// hidden_user
	SenderUserName string `json:"sender_user_name,omitempty"`

	// chat
	SenderChat *Chat `json:"sender_chat,omitempty"`

	// channel
	Chat            *Chat  `json:"chat,omitempty"`
	MessageID       int    `json:"message_id,omitempty"`
	AuthorSignature string `json:"author_signature,omitempty"`
}

type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int64  `json:"file_size,omitempty"`
}

type Animation struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int64      `json:"file_size,omitempty"`
}

type Audio struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Duration     int        `json:"duration"`
	Performer    string     `json:"performer,omitempty"`
	Title        string     `json:"title,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int64      `json:"file_size,omitempty"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
}

type Document struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int64      `json:"file_size,omitempty"`
}

type Story struct {
	Chat *Chat `json:"chat"`
	ID   int   `json:"id"`
}

type VideoQuality struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	Codec        string `json:"codec"`
	FileSize     int64  `json:"file_size,omitempty"`
}

type Video struct {
	FileID         string          `json:"file_id"`
	FileUniqueID   string          `json:"file_unique_id"`
	Width          int             `json:"width"`
	Height         int             `json:"height"`
	Duration       int             `json:"duration"`
	Thumbnail      *PhotoSize      `json:"thumbnail,omitempty"`
	Cover          []*PhotoSize    `json:"cover,omitempty"`
	StartTimestamp int             `json:"start_timestamp,omitempty"`
	Qualities      []*VideoQuality `json:"qualities,omitempty"`
	FileName       string          `json:"file_name,omitempty"`
	MimeType       string          `json:"mime_type,omitempty"`
	FileSize       int64           `json:"file_size,omitempty"`
}

type VideoNote struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Length       int        `json:"length"`
	Duration     int        `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileSize     int64      `json:"file_size,omitempty"`
}

type Voice struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Duration     int    `json:"duration"`
	MimeType     string `json:"mime_type,omitempty"`
	FileSize     int64  `json:"file_size,omitempty"`
}

type PaidMediaInfo struct {
	StarCount int          `json:"star_count"`
	PaidMedia []*PaidMedia `json:"paid_media"`
}

// TODO
type PaidMedia struct{}

type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	UserID      int64  `json:"user_id,omitempty"`
	Vcard       string `json:"vcard,omitempty"`
}

type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

type PollOption struct {
	Text         string           `json:"text"`
	TextEntities []*MessageEntity `json:"text_entities,omitempty"`
	VoterCount   int              `json:"voter_count"`
}

type Poll struct {
	ID                    string           `json:"id"`
	Question              string           `json:"question"`
	QuestionEntities      []*MessageEntity `json:"question_entities,omitempty"`
	Options               []*PollOption    `json:"options"`
	TotalVoterCount       int              `json:"total_voter_count"`
	IsClosed              bool             `json:"is_closed"`
	IsAnonymous           bool             `json:"is_anonymous"`
	Type                  string           `json:"type"`
	AllowsMultipleAnswers bool             `json:"allows_multiple_answers"`
	CorrectOptionID       int              `json:"correct_option_id,omitempty"`
	Explanation           string           `json:"explanation,omitempty"`
	ExplanationEntities   []*MessageEntity `json:"explanation_entities,omitempty"`
	OpenPeriod            int              `json:"open_period,omitempty"`
	CloseDate             int64            `json:"close_date,omitempty"`
}

type ChecklistTask struct {
	ID              string           `json:"id"`
	Text            string           `json:"text"`
	TextEntities    []*MessageEntity `json:"text_entities,omitempty"`
	CompletedByUser *User            `json:"completed_by_user,omitempty"`
	CompletedByChat *Chat            `json:"completed_by_chat,omitempty"`
	CompletionDate  int64            `json:"completion_date,omitempty"`
}

type Checklist struct {
	Title                    string           `json:"title"`
	TitleEntities            []*MessageEntity `json:"title_entities,omitempty"`
	Tasks                    []*ChecklistTask `json:"tasks"`
	OthersCanAddTasks        bool             `json:"others_can_add_tasks,omitempty"`
	OthersCanMarkTasksAsDone bool             `json:"others_can_mark_tasks_as_done,omitempty"`
}

type ChecklistTasksDone struct {
	ChecklistMessage       *Message `json:"checklist_message"`
	MarkedAsDoneTaskIDs    []int    `json:"marked_as_done_task_ids"`
	MarkedAsNotDoneTaskIDs []int    `json:"marked_as_not_done_task_ids"`
}

type ChecklistTasksAdded struct {
	ChecklistMessage *Message         `json:"checklist_message"`
	Tasks            []*ChecklistTask `json:"tasks"`
}

type Location struct {
	Latitude             float32 `json:"latitude"`
	Longitude            float32 `json:"longitude"`
	HorizontalAccuracy   float32 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int     `json:"live_period,omitempty"`
	Heading              int     `json:"heading,omitempty"`
	ProximityAlertRadius int     `json:"proximity_alert_radius,omitempty"`
}

type Venue struct {
	Location        *Location `json:"location"`
	Title           string    `json:"title"`
	Address         string    `json:"address"`
	FoursquareID    string    `json:"foursquare_id,omitempty"`
	FoursquareType  string    `json:"foursquare_type,omitempty"`
	GooglePlaceID   string    `json:"google_place_id,omitempty"`
	GooglePlaceType string    `json:"google_place_type,omitempty"`
}

type WebAppData struct {
	Data       string `json:"data"`
	ButtonText string `json:"button_text"`
}

type ProximityAlertTriggered struct {
	Traveler *User `json:"traveler"`
	Watcher  *User `json:"watcher"`
	Distance int   `json:"distance"`
}

type MessageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int `json:"message_auto_delete_time,omitempty"`
}

type ChatBoostAdded struct {
	BoostCount int `json:"boost_count"`
}

// TODO
type BackgroundType struct{}

type ChatBackground struct {
	Type *BackgroundType `json:"type"`
}

type ForumTopicCreated struct {
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
	IsNameImplicit    bool   `json:"is_name_implicit,omitempty"`
}

// This object represents a service message about a forum topic closed in the chat.
// Currently holds no information.
type ForumTopicClosed struct{}

type ForumTopicEdited struct {
	Name              string `json:"name,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// This object represents a service message about a forum topic reopened in the chat.
// Currently holds no information.
type ForumTopicReopened struct{}

// This object represents a service message about General forum topic hidden in the chat.
// Currently holds no information.
type GeneralForumTopicHidden struct{}

// This object represents a service message about General forum topic unhidden in the chat.
// Currently holds no information.
type GeneralForumTopicUnhidden struct{}

type SharedUser struct {
	UserID    int64        `json:"user_id"`
	FirstName string       `json:"first_name,omitempty"`
	LastName  string       `json:"last_name,omitempty"`
	Username  string       `json:"username,omitempty"`
	Photo     []*PhotoSize `json:"photo,omitempty"`
}

type UsersShared struct {
	RequestID int           `json:"request_id"`
	Users     []*SharedUser `json:"users"`
}

type ChatShared struct {
	RequestID int          `json:"request_id"`
	ChatID    int64        `json:"chat_id"`
	Title     string       `json:"title,omitempty"`
	Username  string       `json:"username,omitempty"`
	Photo     []*PhotoSize `json:"photo,omitempty"`
}

type WriteAccessAllowed struct {
	FromRequest        bool   `json:"from_request,omitempty"`
	WebAppName         string `json:"web_app_name,omitempty"`
	FromAttachmentMenu bool   `json:"from_attachment_menu,omitempty"`
}

type VideoChatScheduled struct {
	StartDate int64 `json:"start_date"`
}

// This object represents a service message about a video chat started in the chat.
// Currently holds no information.
type VideoChatStarted struct{}

type VideoChatEnded struct {
	Duration int `json:"duration"`
}
type VideoChatParticipantsInvited struct {
	Users []*User `json:"users"`
}

type PaidMessagePriceChanged struct {
	PaidMessageStarCount int `json:"paid_message_star_count"`
}

type DirectMessagePriceChanged struct {
	AreDirectMessagesEnabled bool `json:"are_direct_messages_enabled"`
	DirectMessageStarCount   int  `json:"direct_message_star_count,omitempty"`
}

type SuggestedPostApproved struct {
	SuggestedPostMessage *Message            `json:"suggested_post_message,omitempty"`
	Price                *SuggestedPostPrice `json:"price,omitempty"`
	SendDate             int64               `json:"send_date"`
}

type SuggestedPostApprovalFailed struct {
	SuggestedPostMessage *Message            `json:"suggested_post_message,omitempty"`
	Price                *SuggestedPostPrice `json:"price"`
}

type SuggestedPostDeclined struct {
	SuggestedPostMessage *Message `json:"suggested_post_message,omitempty"`
	Comment              string   `json:"comment,omitempty"`
}

type SuggestedPostPaid struct {
	SuggestedPostMessage *Message    `json:"suggested_post_message,omitempty"`
	Currency             string      `json:"currency"`
	Amount               int         `json:"amount,omitempty"`
	StarAmount           *StarAmount `json:"star_amount,omitempty"`
}

type SuggestedPostRefunded struct {
	SuggestedPostMessage *Message `json:"suggested_post_message,omitempty"`
	Reason               string   `json:"reason"`
}

type GiveawayCreated struct {
	PrizeStarCount int `json:"prize_star_count,omitempty"`
}

type Giveaway struct {
	Chats                         []*Chat  `json:"chats"`
	WinnersSelectionDate          int      `json:"winners_selection_date"`
	WinnerCount                   int      `json:"winner_count"`
	OnlyNewMembers                bool     `json:"only_new_members,omitempty"`
	HasPublicWinners              bool     `json:"has_public_winners,omitempty"`
	PrizeDescription              string   `json:"prize_description,omitempty"`
	CountryCodes                  []string `json:"country_codes,omitempty"`
	PrizeStarCount                int      `json:"prize_star_count,omitempty"`
	PremiumSubscriptionMonthCount int      `json:"premium_subscription_month_count,omitempty"`
}

type GiveawayWinners struct {
	Chat                          *Chat   `json:"chat"`
	GiveawayMessageID             int     `json:"giveaway_message_id"`
	WinnersSelectionDate          int     `json:"winners_selection_date"`
	WinnerCount                   int     `json:"winner_count"`
	Winners                       []*User `json:"winners"`
	AdditionalChatCount           int     `json:"additional_chat_count,omitempty"`
	PrizeStarCount                int     `json:"prize_star_count,omitempty"`
	PremiumSubscriptionMonthCount int     `json:"premium_subscription_month_count,omitempty"`
	UnclaimedPrizeCount           int     `json:"unclaimed_prize_count,omitempty"`
	OnlyNewMembers                bool    `json:"only_new_members,omitempty"`
	WasRefunded                   bool    `json:"was_refunded,omitempty"`
	PrizeDescription              string  `json:"prize_description,omitempty"`
}

type GiveawayCompleted struct {
	WinnerCount         int      `json:"winner_count"`
	UnclaimedPrizeCount int      `json:"unclaimed_prize_count,omitempty"`
	GiveawayMessage     *Message `json:"giveaway_message,omitempty"`
	IsStarGiveaway      bool     `json:"is_star_giveaway,omitempty"`
}

type LinkPreviewOptions struct {
	IsDisabled       bool   `json:"is_disabled,omitempty"`
	URL              string `json:"url,omitempty"`
	PreferSmallMedia bool   `json:"prefer_small_media,omitempty"`
	PreferLargeMedia bool   `json:"prefer_large_media,omitempty"`
	ShowAboveText    bool   `json:"show_above_text,omitempty"`
}

type SuggestedPostPrice struct {
	Currency string `json:"currency"`
	Amount   int64  `json:"amount"`
}

type SuggestedPostInfo struct {
	State    string              `json:"state"`
	Price    *SuggestedPostPrice `json:"price,omitempty"`
	SendDate int64               `json:"send_date,omitempty"`
}

type DirectMessagesTopic struct {
	TopicID int64 `json:"topic_id"`
	User    *User `json:"user,omitempty"`
}

type File struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size,omitempty"`
	FilePath     string `json:"file_path,omitempty"`
}

type WebAppInfo struct {
	URL string `json:"url"`
}

type InlineKeyboardMarkup struct {
	InlineKeyboard []*InlineKeyboardButton `json:"inline_keyboard"`
}

type InlineKeyboardButton struct {
	Text                         string                       `json:"text"`
	IconCustomEmojiID            string                       `json:"icon_custom_emoji_id,omitempty"`
	Style                        string                       `json:"style,omitempty"`
	URL                          string                       `json:"url,omitempty"`
	CallbackData                 string                       `json:"callback_data,omitempty"`
	WebApp                       *WebAppInfo                  `json:"web_app,omitempty"`
	LoginURL                     *LoginURL                    `json:"login_url,omitempty"`
	SwitchInlineQuery            string                       `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string                       `json:"switch_inline_query_current_chat,omitempty"`
	SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
	CopyText                     *CopyTextButton              `json:"copy_text,omitempty"`
	CallbackGame                 *CallbackGame                `json:"callback_game,omitempty"`
	Pay                          bool                         `json:"pay,omitempty"`
}

type LoginURL struct {
	URL                string `json:"url"`
	ForwardText        string `json:"forward_text,omitempty"`
	BotUsername        string `json:"bot_username,omitempty"`
	RequestWriteAccess bool   `json:"request_write_access,omitempty"`
}

type SwitchInlineQueryChosenChat struct {
	Query             string `json:"query,omitempty"`
	Allow_user_chats  bool   `json:"allow_user_chats,omitempty"`
	AllowBotChats     bool   `json:"allow_bot_chats,omitempty"`
	AllowGroupChats   bool   `json:"allow_group_chats,omitempty"`
	AllowChannelChats bool   `json:"allow_channel_chats,omitempty"`
}

type CopyTextButton struct {
	Text string `json:"text"`
}

type ChatPhoto struct {
	SmallFileID       string `json:"small_file_id"`
	SmallFileUniqueID string `json:"small_file_unique_id"`
	BigFileID         string `json:"big_file_id"`
	BigFileUniqueID   string `json:"big_file_unique_id"`
}

type ChatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages,omitempty"`
	CanSendAudios         bool `json:"can_send_audios,omitempty"`
	CanSendDocuments      bool `json:"can_send_documents,omitempty"`
	CanSendPhotos         bool `json:"can_send_photos,omitempty"`
	CanSendVideoNotes     bool `json:"can_send_video_notes,omitempty"`
	CanSendVideos         bool `json:"can_send_videos,omitempty"`
	CanSendVoiceNotes     bool `json:"can_send_voice_notes,omitempty"`
	CanSendPolls          bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
	CanEditTag            bool `json:"can_edit_tag,omitempty"`
	CanChangeInfo         bool `json:"can_change_info,omitempty"`
	CanInviteUsers        bool `json:"can_invite_users,omitempty"`
	CanPinMessages        bool `json:"can_pin_messages,omitempty"`
	CanManageTopics       bool `json:"can_manage_topics,omitempty"`
}

type Birthdate struct {
	Day   int `json:"day"`
	Month int `json:"month"`
	Year  int `json:"year,omitempty"`
}

type BusinessIntro struct {
	Title   string   `json:"title,omitempty"`
	Message string   `json:"message,omitempty"`
	Sticker *Sticker `json:"sticker,omitempty"`
}

type BusinessLocation struct {
	Address  string    `json:"address"`
	Location *Location `json:"location,omitempty"`
}

type BusinessOpeningHoursInterval struct {
	OpeningMinute int `json:"opening_minute"`
	ClosingMinute int `json:"closing_minute"`
}

type BusinessOpeningHours struct {
	TimeZoneName *Location                       `json:"time_zone_name"`
	OpeningHours []*BusinessOpeningHoursInterval `json:"opening_hours,omitempty"`
}

type UserRating struct {
	Level              int `json:"level"`
	Rating             int `json:"rating"`
	CurrentLevelRating int `json:"current_level_rating"`
	NextLevelRating    int `json:"next_level_rating,omitempty"`
}

type ChatLocation struct {
	Location *Location `json:"location"`
	Address  string    `json:"address"`
}

// TODO
type ReactionType struct {
	Type          string `json:"type"`
	Emoji         string `json:"emoji,omitempty"`
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
}

type GiftBackground struct {
	CenterColor int `json:"center_color"`
	EdgeColor   int `json:"edge_color"`
	TextColor   int `json:"text_color"`
}

type Gift struct {
	ID                     string          `json:"id"`
	Sticker                *Sticker        `json:"sticker"`
	StarCount              int             `json:"star_count"`
	UpgradeStarCount       int             `json:"upgrade_star_count,omitempty"`
	IsPremium              bool            `json:"is_premium,omitempty"`
	HasColors              bool            `json:"has_colors,omitempty"`
	TotalCount             int             `json:"total_count,omitempty"`
	RemainingCount         int             `json:"remaining_count,omitempty"`
	PersonalTotalCount     int             `json:"personal_total_count,omitempty"`
	PersonalRemainingCount int             `json:"personal_remaining_count,omitempty"`
	Background             *GiftBackground `json:"background,omitempty"`
	UniqueGiftVariantCount int             `json:"unique_gift_variant_count,omitempty"`
	PublisherChat          *Chat           `json:"publisher_chat,omitempty"`
}

type Gifts struct {
	Gifts []*Gift `json:"gifts"`
}

type UniqueGiftModel struct {
	Name           string   `json:"name"`
	Sticker        *Sticker `json:"sticker"`
	RarityPerMille int      `json:"rarity_per_mille"`

	// Rarity of the model if it is a crafted model.
	// Currently, can be “uncommon”, “rare”, “epic”, or “legendary”.
	Rarity string `json:"rarity,omitempty"`
}

type UniqueGiftSymbol struct {
	Name           string   `json:"name"`
	Sticker        *Sticker `json:"sticker"`
	RarityPerMille int      `json:"rarity_per_mille"`
}

type UniqueGiftBackdropColors struct {
	CenterColor int `json:"center_color"`
	EdgeColor   int `json:"edge_color"`
	SymbolColor int `json:"symbol_color"`
	TextColor   int `json:"text_color"`
}

type UniqueGiftBackdrop struct {
	Name           string                    `json:"name"`
	Colors         *UniqueGiftBackdropColors `json:"colors"`
	RarityPerMille int                       `json:"rarity_per_mille"`
}

type UniqueGiftColors struct {
	ModelCustomEmojiID    string `json:"model_custom_emoji_id"`
	SymbolCustomEmojiID   string `json:"symbol_custom_emoji_id"`
	LightThemeMainColor   int    `json:"light_theme_main_color"`
	LightThemeOtherColors []int  `json:"light_theme_other_colors"`
	DarkThemeMainColor    int    `json:"dark_theme_main_color"`
	DarkThemeOtherColors  []int  `json:"dark_theme_other_colors"`
}

type UniqueGift struct {
	GiftID           string              `json:"gift_id"`
	BaseName         string              `json:"base_name"`
	Name             string              `json:"name"`
	Number           int                 `json:"number"`
	Model            *UniqueGiftModel    `json:"model"`
	Symbol           *UniqueGiftSymbol   `json:"symbol"`
	Backdrop         *UniqueGiftBackdrop `json:"backdrop"`
	IsPremium        bool                `json:"is_premium,omitempty"`
	IsBurned         bool                `json:"is_burned,omitempty"`
	IsFromBlockchain bool                `json:"is_from_blockchain,omitempty"`
	Colors           *UniqueGiftColors   `json:"colors,omitempty"`
	PublisherChat    *Chat               `json:"publisher_chat,omitempty"`
}

type GiftInfo struct {
	Gift                    *Gift            `json:"gift"`
	OwnedGiftID             string           `json:"owned_gift_id,omitempty"`
	ConvertStarCount        int              `json:"convert_star_count,omitempty"`
	PrepaidUpgradeStarCount int              `json:"prepaid_upgrade_star_count,omitempty"`
	IsUpgradeSeparate       bool             `json:"is_upgrade_separate,omitempty"`
	CanBeUpgraded           bool             `json:"can_be_upgraded,omitempty"`
	Text                    string           `json:"text,omitempty"`
	Entities                []*MessageEntity `json:"entities,omitempty"`
	IsPrivate               bool             `json:"is_private,omitempty"`
	UniqueGiftNumber        int              `json:"unique_gift_number,omitempty"`
}

type UniqueGiftInfo struct {
	Gift               *UniqueGift `json:"gift"`
	Origin             string      `json:"origin"`
	LastResaleCurrency string      `json:"last_resale_currency,omitempty"`
	LastResaleAmount   int         `json:"last_resale_amount,omitempty"`
	OwnedGiftID        string      `json:"owned_gift_id,omitempty"`
	TransferStarCount  int         `json:"transfer_star_count,omitempty"`
	NextTransferDate   int         `json:"next_transfer_date,omitempty"`
}

type AcceptedGiftTypes struct {
	UnlimitedGifts      bool `json:"unlimited_gifts"`
	LimitedGifts        bool `json:"limited_gifts"`
	UniqueGifts         bool `json:"unique_gifts"`
	PremiumSubscription bool `json:"premium_subscription"`
	GiftsFromChannels   bool `json:"gifts_from_channels"`
}

type StarAmount struct {
	Amount         int `json:"amount"`
	NanostarAmount int `json:"nanostar_amount,omitempty"`
}

type ChatOwnerLeft struct {
	NewOwner *User `json:"new_owner"`
}

type ChatOwnerChanged struct {
	NewOwner *User `json:"new_owner"`
}

// Stickers related
type Sticker struct {
	FileID           string        `json:"file_id"`
	FileUniqueID     string        `json:"file_unique_id"`
	Type             string        `json:"type"`
	Width            int           `json:"width"`
	Height           int           `json:"height"`
	IsAnimated       bool          `json:"is_animated"`
	IsVideo          bool          `json:"is_video"`
	Thumbnail        *PhotoSize    `json:"thumbnail,omitempty"`
	Emoji            string        `json:"emoji,omitempty"`
	SetName          string        `json:"set_name,omitempty"`
	PremiumAnimation *File         `json:"premium_animation,omitempty"`
	MaskPosition     *MaskPosition `json:"mask_position,omitempty"`
	CustomEmojiID    string        `json:"custom_emoji_id,omitempty"`
	NeedsRepainting  bool          `json:"needs_repainting,omitempty"`
	FileSize         int64         `json:"file_size,omitempty"`
}

type StickerSet struct {
	Name  string `json:"name"`
	Title string `json:"title"`

	// Type of stickers in the set, currently one of “regular”, “mask”, “custom_emoji”
	StickerType string `json:"sticker_type"`

	Stickers  []*Sticker `json:"stickers"`
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
}

type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float32 `json:"x_shift"`
	YShift float32 `json:"y_shift"`
	Scale  float32 `json:"scale"`
}

// Payment related
type Invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int    `json:"total_amount"`
}

type ShippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

type OrderInfo struct {
	Name            string           `json:"name,omitempty"`
	PhoneNumber     string           `json:"phone_number,omitempty"`
	Email           string           `json:"email,omitempty"`
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

type SuccessfulPayment struct {
	Currency                   string     `json:"currency"`
	TotalAmount                int        `json:"total_amount"`
	InvoicePayload             string     `json:"invoice_payload"`
	SubscriptionExpirationDate int64      `json:"subscription_expiration_date,omitempty"`
	IsRecurring                bool       `json:"is_recurring,omitempty"`
	IsFirstRecurring           bool       `json:"is_first_recurring,omitempty"`
	ShippingOptionID           string     `json:"shipping_option_id,omitempty"`
	OrderInfo                  *OrderInfo `json:"order_info,omitempty"`
	TelegramPaymentChargeID    string     `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID    string     `json:"provider_payment_charge_id"`
}

type RefundedPayment struct {
	Currency                string `json:"currency"`
	TotalAmount             int    `json:"total_amount"`
	InvoicePayload          string `json:"invoice_payload"`
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeID string `json:"provider_payment_charge_id"`
}

// Telegram Passport related
type PassportData struct {
	Data        []*EncryptedPassportElement `json:"data"`
	Credentials *EncryptedCredentials       `json:"credentials"`
}

type PassportFile struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size"`
	FileDate     int64  `json:"file_date"`
}

type EncryptedPassportElement struct {
	Type        string          `json:"type"`
	Data        string          `json:"data,omitempty"`
	PhoneNumber string          `json:"phone_number,omitempty"`
	Email       string          `json:"email,omitempty"`
	Files       []*PassportFile `json:"files,omitempty"`
	FrontSide   *PassportFile   `json:"front_side,omitempty"`
	ReverseSide *PassportFile   `json:"reverse_side,omitempty"`
	Selfie      *PassportFile   `json:"selfie,omitempty"`
	Translation []*PassportFile `json:"translation,omitempty"`
	Hash        string          `json:"hash"`
}

type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

// Game related
type Game struct {
	Title        string           `json:"title"`
	Description  string           `json:"description"`
	Photo        []*PhotoSize     `json:"photo"`
	Text         string           `json:"text,omitempty"`
	TextEntities []*MessageEntity `json:"text_entities,omitempty"`
	Animation    *Animation       `json:"animation,omitempty"`
}

// A placeholder, currently holds no information.
// Use BotFather to set up your game.
type CallbackGame struct{}
