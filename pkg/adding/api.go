package adding

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type TaskAPI struct {
	TaskService TaskService
}

func (t *TaskAPI) Create(c *gin.Context) {
	var taskDTO TaskDTO
	err := c.BindJSON(&taskDTO)
	if err != nil {
		c.Status(http.StatusBadRequest)
		log.Fatalln(err)
		return
	}

	createdTask := t.TaskService.Save(ToTask(taskDTO))

	c.JSON(http.StatusOK, gin.H{"task": ToTaskDTO(createdTask)})
}

