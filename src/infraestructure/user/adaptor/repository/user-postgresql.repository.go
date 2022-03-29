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

func (rup RepositoryUserPostgreSql) Save(user model.User) error {
	entity := entity.User{Name: user.Name, Password: user.Password, Creation_date: time.Now(), Email: user.Email, Role: user.Role}
	result := rup.userRepository.Create(&entity)
	return result.Error
}

func (rup RepositoryUserPostgreSql) ExistUserNameAndEmail(name, email string) (bool, error) {
	var user model.User
	var count int64 = 0
	result := rup.userRepository.Where("name = ?", name).Or("email = ?", email).Find(&user).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}
	return count > 0, nil
}

func (rup RepositoryUserPostgreSql) GetUserByEmail(email string) (user model.User, err error) {
	result := rup.userRepository.Raw("SELECT * FROM users u WHERE email = ? AND u.deleted_at IS NULL ORDER BY u.id LIMIT 1", email).Scan(&user)
	return user, result.Error
}

func (rup RepositoryUserPostgreSql) EditUser(id int, user model.User) error {
	result := rup.userRepository.Exec("UPDATE users SET updated_at = ? ,name = ? ,email = ? WHERE users.deleted_at IS NULL AND id = ?", time.Now(), user.Name, user.Email, id)
	return result.Error
}

func (rup RepositoryUserPostgreSql) Delete(id int) error {
	var entityUser entity.User
	result := rup.userRepository.Where("id = ?", id).Delete(&entityUser)
	return result.Error
}
