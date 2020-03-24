package authenticating

type UserRepository interface {
	Save(user User) User
	FindByUsername(username string) User
}
