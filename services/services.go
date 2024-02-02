package services

import (
	"github.com/EMSI-zero/go-chat/registry"
	"github.com/EMSI-zero/go-chat/services/contact"
	"github.com/EMSI-zero/go-chat/services/user"
)

func InitServices(sr registry.ServiceRegistry) {
	user.NewUserService(sr)
	contact.NewContactService(sr)
}
