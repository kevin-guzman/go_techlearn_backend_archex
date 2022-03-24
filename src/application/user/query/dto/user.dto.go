package dto

import "time"

type UserDto struct {
	Name          string
	Creation_date time.Time
	Email         string
	Role          string
	Id            int
}
