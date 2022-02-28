package command

type CommandRegisterUser struct {
	Name      string `json:"Name" binding:"required,min=5"`
	Password  string `json:"Password" binding:"required,min=6"`
	Email     string `json:"Email" binding:"required,min=5"`
	Role      string `json:"Role" binding:"required"`
	CompanyId int    `json:"CompanyId" binding:"required"`
}
