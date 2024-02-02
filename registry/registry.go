package registry

import (
	"github.com/EMSI-zero/go-chat/domain/chat"
	"github.com/EMSI-zero/go-chat/domain/contact"
	"github.com/EMSI-zero/go-chat/domain/user"
	"github.com/EMSI-zero/go-chat/infra/colrepo"
	"github.com/EMSI-zero/go-chat/infra/dbrepo"
)

type ServiceRegistry interface {
	mustImplementBaseRegistry()
	GetDB() dbrepo.DBConn
	GetColDB() colrepo.ColDBConn

	GetUserService() user.UserService
	RegisterUserService(user.UserService)

	GetContactService() contact.ContactService
	RegisterContactService(contact.ContactService)

	GetChatService() chat.ChatService
	RegisterChatService(chat.ChatService)
}

type serviceRegistry struct {
	services serviceMap
	db       dbrepo.DBConn
	colDB    colrepo.ColDBConn
}

func mustImplementBaseRegistry() {}

func (sr *serviceRegistry) mustImplementBaseRegistry() {}

type serviceMap struct {
	userService    user.UserService
	contactService contact.ContactService
	chatService    chat.ChatService
}

func NewServiceRegistry(db dbrepo.DBConn, colDB colrepo.ColDBConn) *serviceRegistry {
	//create an empty service registry
	sr := new(serviceRegistry)
	sr.db = db
	sr.colDB = colDB
	return sr
}

func (sr *serviceRegistry) GetDB() dbrepo.DBConn {
	return sr.db
}

func (sr *serviceRegistry) GetColDB() colrepo.ColDBConn {
	return sr.colDB
}

func (sr *serviceRegistry) GetUserService() user.UserService {
	return sr.services.userService
}

func (sr *serviceRegistry) RegisterUserService(us user.UserService) {
	sr.services.userService = us
}

func (sr *serviceRegistry) GetContactService() contact.ContactService {
	return sr.services.contactService
}

func (sr *serviceRegistry) RegisterContactService(cs contact.ContactService) {
	sr.services.contactService = cs
}

func (sr *serviceRegistry) GetChatService() chat.ChatService {
	return sr.services.chatService
}

func (sr *serviceRegistry) RegisterChatService(cs chat.ChatService) {
	sr.services.chatService = cs
}
