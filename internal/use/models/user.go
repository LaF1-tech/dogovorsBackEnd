package models

import "errors"

var (
	ErrNoPermission = errors.New("no permission")
)

type Permission string

var (
	PermissionAdmin Permission = "CreateUsers"
)

type User struct {
	UserID      int
	Username    string
	FirstName   string
	LastName    string
	Permissions []Permission
}

func (u *User) AssertPermission(p Permission) error {
	for _, permission := range u.Permissions {
		if permission == p {
			return nil
		}
	}

	return ErrNoPermission
}
