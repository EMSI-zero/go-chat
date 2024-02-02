package contact

import (
	"context"
	"fmt"

	"github.com/EMSI-zero/go-chat/domain/contact"
	"github.com/EMSI-zero/go-chat/domain/user"
	"github.com/EMSI-zero/go-chat/infra/dbrepo"
	"github.com/EMSI-zero/go-chat/registry"
)

type ContactService struct {
	contact.UnImplementedContactService
	db dbrepo.DBConn
}

func NewContactService(sr registry.ServiceRegistry) {
	sr.RegisterContactService(&ContactService{
		db: sr.GetDB(),
	})
}

func (cs *ContactService) GetContacts(ctx context.Context) ([]*contact.Contact, error) {
	db, err := cs.db.GetConn(ctx)
	if err != nil {
		return nil, err
	}

	uid, err := user.GetUserFromCtx(ctx)
	if err != nil {
		return nil, err
	}

	contacts := make([]*contact.Contact, 0)
	err = db.Model(&contact.Contact{}).Where("user_id = ?", uid).Take(&contacts).Error
	if err != nil {
		return nil, err
	}

	return contacts, nil

}
func (cs *ContactService) AddContact(ctx context.Context, req *contact.AddContectRequest) error {
	db, err := cs.db.GetConn(ctx)
	if err != nil {
		return err
	}

	uid, err := user.GetUserFromCtx(ctx)
	if err != nil {
		return err
	}

	if req.ContactID == uid {
		return fmt.Errorf("cannot add yourself as your contact")
	}

	contact := &contact.Contact{
		UserID:      uid,
		ContactID:   req.ContactID,
		ContactName: req.ContactName,
	}

	err = db.Create(&contact).Error
	if err != nil {
		return err
	}

	return nil
}
func (cs *ContactService) DeleteContact(ctx context.Context, req *contact.RemoveContactRequest) error {
	db, err := cs.db.GetConn(ctx)
	if err != nil {
		return err
	}

	err = db.Where("contact_id = ?", req.ContactID).Delete(&contact.Contact{}).Error
	if err != nil {
		return err
	}

	return nil
}
