package adding

import "github.com/jinzhu/gorm"

// Task defines the properties of a task to be added
type Task struct {
	gorm.Model
	Name string
}
