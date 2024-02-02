package chat

import (
	"context"
	"fmt"
)

type ChatService interface {
	mustImplementBaseService()
	CreatePrivateConversation(ctx context.Context, req *NewChatRequest) error
	FetchConversations(ctx context.Context) ([]*ConversationSummary, error)
	FetchContversation(ctx context.Context, id int64) ([]*MessageView, error)
	FetchContversationInfo(ctx context.Context, id int64) (*ConversationInfo, error)
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

func (UnImplementedChatService) CreatePrivateConversation(ctx context.Context, req *NewChatRequest) error {
	return fmt.Errorf("service not implemented")
}
func (UnImplementedChatService) FetchConversations(ctx context.Context) ([]*ConversationSummary, error) {
	return nil, fmt.Errorf("service not implemented")
}
func (UnImplementedChatService) FetchContversation(ctx context.Context, id int64) ([]*MessageView, error) {
	return nil, fmt.Errorf("service not implemented")
}
func (UnImplementedChatService) FetchContversationInfo(ctx context.Context, id int64) (*ConversationInfo, error) {
	return nil, fmt.Errorf("service not implemented")
}
func (UnImplementedChatService) RemovePrivateConversation(ctx context.Context, req *RemoveChatRequest) error {
	return fmt.Errorf("service not implemented")
}
func (UnImplementedChatService) CreateGroupConversation(ctx context.Context, req *NewGroupChatRequest) error {
	return fmt.Errorf("service not implemented")
}
func (UnImplementedChatService) RemoveGroupConversation(ctx context.Context, req *LeaveGroupChatRequest) error {
	return fmt.Errorf("service not implemented")
}
func (UnImplementedChatService) AddParticipants(ctx context.Context, req *AddParticipantsRequest) error {
	return fmt.Errorf("service not implemented")
}
func (UnImplementedChatService) RemoveParticipants(ctx context.Context, req *RemoveParticipantsRequest) error {
	return fmt.Errorf("service not implemented")
}
func (UnImplementedChatService) RemoveMessage(ctx context.Context, req *RemoveMessageRequest) error {
	return fmt.Errorf("service not implemented")
}
func (UnImplementedChatService) EditMessage(ctx context.Context, req *EditMessageRequest) error {
	return fmt.Errorf("service not implemented")
}
