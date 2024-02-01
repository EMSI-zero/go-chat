package user

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Salt      string `json:"-"`
	Bio       string `json:"bio"`
	ImagePath string `json:"imagePath"`
}

type RegisterRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	Bio       string `json:"bio"`
	ImagePath string `json:"imagePath"`
}

type LoginRequest struct {
}

type LoginResponse struct {
}

type GetCurrentUserResponse struct {
}

type GetUserResponse struct{}

type UpdateUserRequest struct{}

type UpdateUserResponse struct{}

type DeleteUserRequest struct{}
