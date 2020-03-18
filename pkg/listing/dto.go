package listing

type TaskDTO struct {
	ID		uint 	`json:"id,string,omitempty"`
	Name	string	`json:"name"`
}
