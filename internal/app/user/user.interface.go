package user

type UserRepository interface {
	GetUser() ([]User, error)
}