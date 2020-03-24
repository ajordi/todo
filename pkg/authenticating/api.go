package authenticating

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AuthenticateAPI struct {
	AuthenticateService AuthenticateService
}

func (a *AuthenticateAPI) Create(c *gin.Context) {
	var userDTO UserDTO
	err := c.BindJSON(&userDTO)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Fatalln(err)
		return
	}

	createdUser := a.AuthenticateService.Save(ToUser(userDTO))

	c.JSON(http.StatusOK, gin.H{"user": ToUserDTO(createdUser)})
}
