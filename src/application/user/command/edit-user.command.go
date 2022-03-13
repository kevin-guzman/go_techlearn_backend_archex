package command

type CommandEditUser struct {
	Name   string `json:"Name"`
	Email  string `json:"Email"`
	UserId int
}
