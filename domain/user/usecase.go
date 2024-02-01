package user

type UserService interface{
	mustImplementBaseService()
}

type UnImplementedUserService struct{}


func (UnImplementedUserService) mustImplementBaseService(){}