package chat

import (
	"time"

	"github.com/EMSI-zero/go-chat/domain/user"
)

// ---------------- Models ----------------
type Conversation struct {
	ID             int64
	IsPrivate      bool   `json:"isPrivate"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	PrivateMembers string `json:"-"`
	ImagePath      string `json:"imagePath"`
	CreatedAt      time.Time
	CreatedBy      int64
}

type UserConversation struct {
	ID             int64
	UserID         int64     `json:"userId"`
	ConversationID int64     `json:"conversationID"`
	Role           int32     `json:"role" gorm:"col:role_c"`
	LastRead       time.Time `json:"lastRead"`
	JoinedSince    time.Time `json:"joinedSince"`
}

type Message struct {
	ID             int64             `json:"messageID" bson:"_id"`
	ConversationID int64             `json:"conversationId" bson:"conversation_id"`
	SenderID       int64             `json:"senderId" bson:"sender_id"`
	SenderName     string            `json:"senderName" bson:"sender_name"`
	Content        map[string]string `json:"content" bson:"content"`
	CreatedAt      time.Time         `json:"createdAt" bson:"created_at,inline"`
	EditedAt       time.Time         `json:"editedAt" bson:"edited_at"`
}

// ---------------- DTOs ----------------

type NewChatRequest struct {
	ReceiverID int64
}

type ConversationSummary struct {
	ID               int64     `json:"id"`
	ConversationName string    `json:"conversationName"`
	IsPrivate        bool      `json:"isPrivate"`
	ImagePath        string    `json:"imagePath"`
	LastRead         time.Time `json:"lastRead"`
	UnreadMessages   int       `json:"unreadMessages" gorm:"-"`
	LastMessage      Message   `json:"lastMessage" gorm:"-"`
	LastUpdate       time.Time `json:"lastUpdate" gorm:"-"`
}

type ConversationInfo struct {
	ID               int64      `json:"id"`
	ConversationName string     `json:"conversationName"`
	IsPrivate        bool       `json:"isPrivate"`
	ImagePath        string     `json:"imagePath"`
	Members          []UserView `json:"members"`
}

type UserView struct {
	user.UserSummary
	Role int32 `json:"role"`
}

type RemoveChatRequest struct {
	ConversationID int64
}

type AddParticipantsRequest struct {
	ConversationID int64
	UserID         int64
}

type RemoveParticipantsRequest struct {
	ConversationID int64
	UserID         int64
}

type NewGroupChatRequest struct {
	ImagePath string
	Members   []int64
}

type LeaveGroupChatRequest struct {
	ConversationID int64
	DeleteGroup    bool
}

type RemoveMessageRequest struct {
	ConversationID int64
	MessageID      int64
}

type EditMessageRequest struct {
	ConversationID int64
	MessageID      int64
}
