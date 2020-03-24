package authenticating

type AuthenticateService interface {
	Save(user User) User
	FindByUsername(username string) (User, error)
}

type authenticateService struct {
	UserRepository UserRepository
}

func ProviderAuthenticatingService(u UserRepository) AuthenticateService {
	return &authenticateService{u}
}

func (u *authenticateService) Save(user User) User {
	u.UserRepository.Save(user)

	return user
}

func (u *authenticateService) FindByUsername(username string) (User, error) {
	return u.UserRepository.FindByUsername(username), nil
}
