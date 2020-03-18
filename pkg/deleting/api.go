package deleting

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TaskAPI struct {
	TaskService TaskService
}

func (t *TaskAPI) Delete(c *gin.Context) {
	id, _ :=  strconv.Atoi(c.Param("id"))
	t.TaskService.Delete(uint(id))

	c.Status(http.StatusOK)
}

