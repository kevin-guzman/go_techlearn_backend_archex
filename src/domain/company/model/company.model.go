package model

type Company struct {
	Id    int
	Name  string
	Owner string
	Phone string
	Email string
}

func NewCompany(name, owner, phone, email string) (*Company, error) {
	return &Company{
		Name:  name,
		Owner: owner,
		Phone: phone,
		Email: email,
	}, nil
}
