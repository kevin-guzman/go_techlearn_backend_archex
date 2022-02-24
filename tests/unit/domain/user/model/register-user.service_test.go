package model_test

import (
	"golang-gingonic-hex-architecture/src/domain/user/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	It("Should fail with a password less than 6", func() {
		_, err := model.NewUser("JUan", "1234", "role", "email", 1)
		Expect(err).Error()
		Expect(err.Error()).To(Equal("The leng of the password is incorrect"))
	})

	It("Should create a user", func() {
		expectUser := model.User{
			Name:      "JUan",
			Password:  "123ss3r4",
			Role:      "role",
			Email:     "email",
			CompanyId: 1,
		}
		usr, err := model.NewUser(expectUser.Name, expectUser.Password, expectUser.Role, expectUser.Email, expectUser.CompanyId)
		Expect(err).To(BeNil())
		Expect(&expectUser).To(Equal(usr))
	})
})
