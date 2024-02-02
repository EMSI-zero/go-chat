package chat

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/EMSI-zero/go-chat/domain/chat"
	"github.com/EMSI-zero/go-chat/domain/user"
	"github.com/EMSI-zero/go-chat/infra/colrepo"
	"github.com/EMSI-zero/go-chat/infra/dbrepo"
	"github.com/EMSI-zero/go-chat/registry"
	"go.mongodb.org/mongo-driver/bson"
)

type ChatService struct {
	chat.UnImplementedChatService
	db    dbrepo.DBConn
	colDb colrepo.ColDBConn
}

func NewChatService(sr registry.ServiceRegistry) {
	sr.RegisterChatService(&ChatService{
		db:    sr.GetDB(),
		colDb: sr.GetColDB(),
	})
}

func (cs *ChatService) CreatePrivateConversation(ctx context.Context, req *chat.NewChatRequest) error {
	db, err := cs.db.GetConn(ctx)
	if err != nil {
		return err
	}

	uid, err := user.GetUserFromCtx(ctx)
	if err != nil {
		return err
	}

	uniquePrivateID := strconv.Itoa(int(uid)) + "_" + strconv.Itoa(int(req.ReceiverID))

	convUser := new(user.User)

	if err := db.Where("id = ?", req.ReceiverID).First(&convUser).Error; err != nil {
		return err
	}

	newConv := &chat.Conversation{
		IsPrivate:      true,
		PrivateMembers: uniquePrivateID,
		CreatedBy:      uid,
	}

	if err = db.Create(newConv).Error; err != nil {
		return err
	}

	return nil
}

func (cs *ChatService) FetchConversations(ctx context.Context) ([]*chat.ConversationSummary, error) {
	db, err := cs.db.GetConn(ctx)
	if err != nil {
		return nil, err
	}

	uid, err := user.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	convs := make([]*chat.ConversationSummary, 0)

	// TODO: fix tablenames
	err = db.Model(&chat.Conversation{}).
		Joins("JOIN user_conversation AS uc ON uc.conversation_id = conversation.id").
		Where("uc.user_id = ?", uid).
		Select(
			"conversation.id",
			"conversation.Name",
			"conversation.is_private",
			"conversation.image_path",
			"uc.last_read",
		).Scan(&convs).
		Error
	if err != nil {
		return nil, err
	}

	// TODO: Fetch Unread and Last Seens

	return convs, nil
}

func (cs *ChatService) FetchConversation(ctx context.Context, id int64) ([]*chat.Message, error) {
	db, err := cs.db.GetConn(ctx)
	if err != nil {
		return nil, err
	}

	colDb, err := cs.colDb.GetConn()
	if err != nil {
		return nil, err
	}

	conv := new(chat.Conversation)

	if err = db.Where("id =?", id).First(&conv).Error; err != nil {
		return nil, err
	}

	col := colDb.Database("chat_db").Collection("messages")

	filter := bson.D{{Key: "conversation_id", Value: id}}
	cursor, err := col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	messages := make([]*chat.Message, 0)
	if err = cursor.All(ctx, &messages); err != nil {
		return nil, err
	}

	return messages, nil
}

func (cs *ChatService) FetchContversationInfo(ctx context.Context, id int64) (*chat.ConversationInfo, error) {

	return nil, nil
}

func (cs *ChatService) RemovePrivateConversation(ctx context.Context, id int64) error {
	db, err := cs.db.GetConn(ctx)
	if err != nil {
		return err
	}

	uid, err := user.GetUserFromCtx(ctx)
	if err != nil {
		return err
	}

	colDb, err := cs.colDb.GetConn()
	if err != nil {
		return err
	}

	conv := new(chat.Conversation)
	err = db.Find(&conv, id).Error
	if err != nil {
		return err
	}

	col := colDb.Database("chat_db").Collection("messages")

	filter := bson.D{{Key: "conversation_id", Value: id}}

	_, err = col.DeleteMany(ctx, filter)
	if err != nil {
		return err
	}

	if !conv.IsPrivate {
		return fmt.Errorf("not a private conversation")
	}

	if !strings.Contains(conv.PrivateMembers, strconv.Itoa(int(uid))) {
		return fmt.Errorf("unautorized access")
	}

	err = db.Delete(conv, conv.ID).Error
	if err != nil {
		return err
	}

	return nil

}

func (cs *ChatService) CreateGroupConversation(ctx context.Context, req *chat.NewGroupChatRequest) error {
	return nil
}

func (cs *ChatService) RemoveGroupConversation(ctx context.Context, req *chat.LeaveGroupChatRequest) error {
	return nil
}

func (cs *ChatService) AddParticipants(ctx context.Context, req *chat.AddParticipantsRequest) error {
	return nil
}

func (cs *ChatService) RemoveParticipants(ctx context.Context, req *chat.RemoveParticipantsRequest) error {
	return nil
}

func (cs *ChatService) RemoveMessage(ctx context.Context, req *chat.RemoveMessageRequest) error {
	uid, err := user.GetUserFromCtx(ctx)
	if err != nil {
		return err
	}

	colDb, err := cs.colDb.GetConn()
	if err != nil {
		return err
	}

	col := colDb.Database("chat_db").Collection("messages")

	filter := bson.D{{Key: "conversation_id", Value: req.ConversationID},
		{Key: "message_id", Value: req.MessageID},
		{Key: "sender_id", Value: uid}}

	_, err = col.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (cs *ChatService) EditMessage(ctx context.Context, req *chat.EditMessageRequest) error {
	return nil
}
