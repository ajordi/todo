package user

import (
	"github.com/ajordi/todo/pkg/authenticating"
	"github.com/jinzhu/gorm"
)

type Storage struct {
	GORM *gorm.DB
}

func (s *Storage) Save(user authenticating.User) authenticating.User {
	s.GORM.Save(&user)

	return user
}


func (s *Storage) FindByUsername(username string) authenticating.User {
	var user authenticating.User
	s.GORM.Where(authenticating.User{UserName: username}).Find(&user)

	return user
}
