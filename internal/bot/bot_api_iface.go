package bot

import (
	"strings"
	"unicode/utf8"

	"github.com/google/uuid"

	"be/internal/datastore/address"
	"be/internal/model"
)

type ChatAPI interface {
	Login(req *LoginBotRequest) (*LoginBotResponse, error)
	Logout()
	CreateStream(req *CreateStreamRequest) (*Stream, error)
	PublishStream(req *CreateMessageRequest) (*Message, error)
	GetSubscribedStreams() (*ListSubscribesResponse, error)
	AddUserToStream(streamID string, userUUIDs []string) error
	CreateProposalViewLink(id, ip, secret string, addressType address.TypeAddress, uuid uuid.UUID) (string, error)
}

type LoginBotRequest struct {
	ID     string `json:"id"`
	IP     string `json:"ip"`
	OS     OS     `json:"os"`
	Secret string `json:"secret"`
}

type LoginBotResponse struct {
	Token string `json:"token"`
}

type OS struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type CreateStreamRequest struct {
	Description string   `json:"description"`
	Invites     []string `json:"invites"`
	Title       string   `json:"title"`
	Type        int      `json:"type"`
	Visible     int      `json:"visible"`
}

//nolint:tagliatelle
type Stream struct {
	CID         string `json:"cID"`
	CreatedAt   int64  `json:"createdAt"`
	Description string `json:"description"`
	ID          string `json:"id"`
	Logo        Avatar `json:"logo"`
	UserName    string `json:"username"`
	Title       string `json:"title"`
	Type        int    `json:"type"`
	Visible     int    `json:"visible"`
}

type Avatar struct {
	FileName string `json:"filename"`
	ID       string `json:"id"`
	Mime     string `json:"mime"`
	Original string `json:"original"`
	Size     int    `json:"size"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

//nolint:tagliatelle
type User struct {
	AliasName   string `json:"aliasName"`
	Avatar      Avatar `json:"avatar"`
	Department  string `json:"department"`
	Description string `json:"description"`
	FirstName   string `json:"firstName"`
	ID          string `json:"id"`
	IsBot       bool   `json:"isBot"`
	LastName    string `json:"lastName"`
	MiddleName  string `json:"middleName"`
	Position    string `json:"position"`
	Status      string `json:"status"`
	UserName    string `json:"username"`
}

//nolint:tagliatelle
type MessageEntity struct {
	Language string `json:"language"`
	Length   int    `json:"length"`
	Offset   int    `json:"offset"`
	StreamID string `json:"streamID"`
	Type     int32  `json:"type"`
	UserID   string `json:"userID"`
	Value    string `json:"value"`
}

type MessageEntityType int32

// MessageEntityType values.
const (
	MessageEntityMention       MessageEntityType = 0
	MessageEntityHashtag       MessageEntityType = 1
	MessageEntityURL                             = 2
	MessageEntityEmail                           = 3
	MessageEntityPhoneNumber                     = 4
	MessageEntityBotCommand                      = 5
	MessageEntityBold                            = 6
	MessageEntityItalic                          = 7
	MessageEntityUnderline                       = 8
	MessageEntityStrikethrough                   = 9
	MessageEntityCode                            = 10
	MessageEntityPre                             = 11
)

//nolint:tagliatelle
type AttachFile struct {
	Caption  string          `json:"caption"`
	Entities []MessageEntity `json:"entities"`
	FileID   string          `json:"fileID"`
}

type Attachment struct {
	Caption  string          `json:"caption"`
	Entities []MessageEntity `json:"entities"`
	File     File            `json:"file"`
}

type File struct {
	FileName string `json:"filename"`
	ID       string `json:"id"`
	Meta     Meta   `json:"meta"`
	Mime     string `json:"mime"`
	Size     int    `json:"size"`
	Type     int    `json:"type"`
}

type Meta struct {
	Thumb    string `json:"thumb"`
	Waveform string `json:"waveform"`
	Duration int    `json:"duration"`
	Height   int    `json:"height"`
	Width    int    `json:"width"`
}

type AttachGallery struct {
	Attachment AttachFile `json:"attachments"`
}

type Gallery struct {
	Attachment Attachment `json:"attachments"`
}

type Location struct {
	Latitude  int `json:"latitude"`
	Longitude int `json:"longitude"`
}

//nolint:tagliatelle
type PollOption struct {
	Text       string `json:"text"`
	VoterCount int    `json:"voterCount"`
}

//nolint:tagliatelle
type Poll struct {
	AllowsMultipleAnswers bool       `json:"allowsMultipleAnswers"`
	ClosePeriod           int        `json:"closePeriod"`
	ID                    string     `json:"id"`
	IsAnonymous           bool       `json:"isAnonymous"`
	IsClosed              bool       `json:"isClosed"`
	OpenPeriod            int        `json:"openPeriod"`
	Options               PollOption `json:"options"`
	Question              string     `json:"question"`
	TotalVoterCount       int        `json:"totalVoterCount"`
}

type Venue struct {
	Address  string   `json:"address"`
	Location Location `json:"location"`
	Title    string   `json:"title"`
}

//nolint:tagliatelle
type Sticker struct {
	Emoji      string `json:"emoji"`
	File       File   `json:"file"`
	IsAnimated bool   `json:"isAnimated"`
	SetName    string `json:"setName"`
}

//nolint:tagliatelle
type CreateMessageRequest struct {
	Contact          *User            `json:"contact,omitempty"`
	Content          string           `json:"content"`
	Entities         []*MessageEntity `json:"entities,omitempty"`
	File             *AttachFile      `json:"file,omitempty"`
	Gallery          *AttachGallery   `json:"gallery,omitempty"`
	Location         *Location        `json:"location,omitempty"`
	PID              string           `json:"pID"`
	Poll             *Poll            `json:"poll,omitempty"`
	ReplyToMessageID string           `json:"replyToMessageID"`
	Silently         bool             `json:"silently"`
	Sticker          *string          `json:"sticker,omitempty"`
	StreamID         string           `json:"streamID"`
	Type             int              `json:"type"`
	Venue            *Venue           `json:"venue,omitempty"`
}

//nolint:tagliatelle
type Message struct {
	ID                string           `json:"id"`
	PID               string           `json:"pID"`
	CID               string           `json:"cID"`
	StreamID          string           `json:"streamID"`
	UserID            string           `json:"userID"`
	Content           string           `json:"content"`
	CreatedAt         int              `json:"createdAt"`
	UpdatedAt         int              `json:"updatedAt"`
	Deleted           bool             `json:"deleted"`
	Entities          []*MessageEntity `json:"entities"`
	Type              int              `json:"type"`
	ServiceType       int              `json:"serviceType"`
	ReplyToMessageID  string           `json:"replyToMessageID"`
	ReplyToMessage    interface{}      `json:"replyToMessage"`
	ForwardFrom       *User            `json:"forwardFrom"`
	ForwardSenderName string           `json:"forwardSenderName"`
	ForwardFromStream *Stream          `json:"forwardFromStream"`
	Document          *Attachment      `json:"document"`
	Gallery           *Gallery         `json:"gallery"`
	Sticker           *Sticker         `json:"sticker"`
	Contact           *User            `json:"contact"`
	Poll              *Poll            `json:"poll"`
	Location          *Location        `json:"location"`
	Venue             *Venue           `json:"venue"`
}

type ChatAPIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ListSubscribesResponse struct {
	Streams []*SubscribedStream `json:"streams"`
}

//nolint:tagliatelle
type SubscribedStream struct {
	InterlocutorReadMessage     string  `json:"interlocutorReadMessage"`
	InterlocutorReceivedMessage string  `json:"interlocutorReceivedMessage"`
	MyReadMessage               string  `json:"myReadMessage"`
	MyReceivedMessage           string  `json:"myReceivedMessage"`
	LastMessage                 string  `json:"lastMessage"`
	LastMessageID               string  `json:"lastMessageID"`
	Stream                      *Stream `json:"stream"`
	Users                       []*User `json:"users"`
}

func BuildMessageEntities(message string, forEntities []model.ForEntity) (string, []*MessageEntity) {
	messageEntities := make([]*MessageEntity, 0, len(forEntities))

	for _, entity := range forEntities {
		if entity.Pattern == model.LinkPattern {
			if entity.Link == "" {
				message = strings.Replace(message, entity.Pattern, entity.Value, 1)
				continue
			}

			messageEntities = append(messageEntities, &MessageEntity{
				Length: utf8.RuneCountInString(entity.Value),
				Offset: utf8.RuneCountInString(message[:strings.Index(message, entity.Pattern)]), //nolint:gocritic
				Type:   int32(MessageEntityURL),
				Value:  entity.Link,
			})

			message = strings.Replace(message, entity.Pattern, entity.Value, 1)
		}
	}

	return message, messageEntities
}
