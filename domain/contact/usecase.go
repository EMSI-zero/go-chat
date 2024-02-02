package contact

import (
	"context"
	"fmt"
)

type ContactService interface {
	mustImplementBaseService()
	GetContacts(ctx context.Context) ([]*Contact, error)
	AddContact(ctx context.Context, req *AddContectRequest) error
	DeleteContact(ctx context.Context, req *RemoveContactRequest) error
}

type UnImplementedContactService struct{}

func (UnImplementedContactService) mustImplementBaseService() {}

func (UnImplementedContactService) GetContacts(ctx context.Context) ([]*Contact, error) {
	return nil, fmt.Errorf("service not implemented")
}
func (UnImplementedContactService) AddContact(ctx context.Context, req *AddContectRequest) error {
	return fmt.Errorf("service not implemented")
}
func (UnImplementedContactService) DeleteContact(ctx context.Context, req *RemoveContactRequest) error {
	return fmt.Errorf("service not implemented")
}
