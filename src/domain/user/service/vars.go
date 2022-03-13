package service

var (
	errTrace       string = ""
	internalError  string = "Internal server error"
	successMessage string = ""
)

const (
	SERVICE_LOGIN    = "SERVICE_LOGIN_USER"
	SERVICE_REGISTER = "SERVICE_REGISTER"
	SERVICE_DELETE   = "SERVICE_DELETE"
	SERVICE_EDIT     = "SERVICE_EDIT"
)

func LoadStringsFromService(serviceName string) {
	switch serviceName {
	case SERVICE_LOGIN:
		successMessage = "User has succesfully logged-in!"
		errTrace = "This error has ocurred in login-user.service.go"
	case SERVICE_REGISTER:
		successMessage = "User has succesfully created!"
		errTrace = "This error has ocurred in register-user.service.go"
	case SERVICE_EDIT:
		successMessage = "User has succesfully updated!"
		errTrace = "This error has ocurred in edit-user.service.go"
	case SERVICE_DELETE:
		successMessage = "User has succesfully deleted!"
		errTrace = "This error has ocurred in delete-user.service.go"
	}

}
