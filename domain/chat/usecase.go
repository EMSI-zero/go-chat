package chat

import "context"

type ChatService interface {
	mustImplementBaseService()
	CreatePrivateConversation(ctx context.Context, req *NewChatRequest) error
	FetchConversations(ctx context.Context) ([]*ConversationSummary, error)
	FetchContversation(ctx context.Context, id int64) ([]*MessageView, error)
	FetchContversationInfo(ctx context.Context, id int64) (ConversationInfo, error)
	RemovePrivateConversation(ctx context.Context, req *RemoveChatRequest) error

	CreateGroupConversation(ctx context.Context, req *NewGroupChatRequest) error
	RemoveGroupConversation(ctx context.Context, req *LeaveGroupChatRequest) error
	AddParticipants(ctx context.Context, req *AddParticipantsRequest) error
	RemoveParticipants(ctx context.Context, req *RemoveParticipantsRequest) error

	RemoveMessage(ctx context.Context, req *RemoveMessageRequest) error
	EditMessage(ctx context.Context, req *EditMessageRequest) error
}

type UnImplementedChatService struct{}

func (UnImplementedChatService) mustImplementBaseService() {}
