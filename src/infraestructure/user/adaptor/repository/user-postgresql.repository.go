package repository

import (
	"golang-gingonic-hex-architecture/src/domain/user/model"
	"golang-gingonic-hex-architecture/src/infraestructure/user/entity"
	"time"

	"gorm.io/gorm"
)

type RepositoryUserPostgreSql struct {
	userRepository *gorm.DB
}

func NewRepositoryUserPostgreSql(conn *gorm.DB) *RepositoryUserPostgreSql {
	return &RepositoryUserPostgreSql{
		userRepository: conn.Model(&entity.User{}),
	}
}

func (rup *RepositoryUserPostgreSql) Save(user model.User) error {
	entity := entity.User{Name: user.Name, CompanyId: user.CompanyId, Password: user.Password, Creation_date: time.Now(), Email: user.Email, Role: user.Role}
	result := rup.userRepository.Create(&entity)
	return result.Error
}

func (rup *RepositoryUserPostgreSql) ExistUserNameAndEmail(name, email string) (bool, error) {
	var user model.User
	var count int64 = 0
	result := rup.userRepository.Where("name = ?", name).Or("email = ?", email).Find(&user).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}

func (rup *RepositoryUserPostgreSql) GetUserByEmail(email string) (user model.User, err error) {
	d := rup.userRepository.Raw("SELECT * FROM users u WHERE email = ? AND u.deleted_at IS NULL ORDER BY u.id LIMIT 1", email).Scan(&user)
	err = d.Error
	return user, d.Error
}

func (rup *RepositoryUserPostgreSql) EditUser(id int, user model.User) error {
	var entityUser entity.User
	entityUser.Name = user.Name
	entityUser.Email = user.Email
	result := rup.userRepository.Where("id = ?", id).Updates(entityUser)
	return result.Error
}

func (rup *RepositoryUserPostgreSql) Delete(id int) error {
	var entityUser entity.User
	entityUser.Id = id
	result := rup.userRepository.Unscoped().Delete(&entityUser)
	return result.Error
}
