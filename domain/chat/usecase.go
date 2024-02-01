package chat

type ChatService interface{
	mustImplementBaseService()
	// TODO: Add methods
}

type UnImplementedChatService struct{}

func (UnImplementedChatService) mustImplementBaseService(){}