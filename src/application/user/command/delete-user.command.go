package command

type CommandDeleteUser struct {
	Password string `json:"Password" binding:"required,min=6"`
	Email    string `json:"Email" binding:"required,min=5"`
}
