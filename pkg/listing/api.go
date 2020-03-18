package listing

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TaskAPI struct {
	TaskService TaskService
}

func (p *TaskAPI) FindAll(c *gin.Context) {
	products := p.TaskService.FindAll()

	c.JSON(http.StatusOK, gin.H{"task": ToTaskDTOs(products)})
}

func (p *TaskAPI) FindByID(c *gin.Context) {
	id, _ :=  strconv.Atoi(c.Param("id"))
	product := p.TaskService.FindByID(uint(id))

	c.JSON(http.StatusOK, gin.H{"task": ToTaskDTO(product)})
}
