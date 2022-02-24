package command

type CommandRegisterCompany struct {
	Name  string `json:"Name" binding:"required,min=5"`
	Owner string `json:"Owner" binding:"required,min=5"`
	Phone string `json:"Phone" binding:"required,min=7"`
	Email string `json:"Email" binding:"required,min=5"`
}
