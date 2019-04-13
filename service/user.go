package service

//BO User
type User struct {
	ID    uint
	Name  string
	Email string
}

type UserService interface {
	GetUser(id uint) (User, error)
}

type userService struct {
}

func (s *userService) GetUser(id uint) (User, error) {
	return User{
		ID:    id,
		Name:  "Vitamin",
		Email: "3485537725@qq.com",
	}, nil
}

func NewUserService() UserService {
	return &userService{}
}
