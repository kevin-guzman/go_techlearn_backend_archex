package command

type CommandLoginUser struct {
	Email    string `json:"Email"`
	Password string `json:"Password" minLength:"6"`
}
