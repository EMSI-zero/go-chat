package contact

type ContactService interface {
	mustImplementBaseService()
	// TODO: Add methods
}


type UnImplementedContactService struct{}

func (UnImplementedContactService) mustImplementBaseService(){}