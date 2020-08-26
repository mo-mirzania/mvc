package service

import (
	"github.com/mo-mirzania/test/domain"
	"github.com/mo-mirzania/test/utils/resterror"
)

var (
	// UserServices outside var
	UserServices UserService = &userService{}
)

// UserService interface definition
type UserService interface {
	GetUser(int) (*domain.User, *resterror.RestErr)
}

// userService struct
type userService struct{}

// GetUser method
func (s *userService) GetUser(id int) (*domain.User, *resterror.RestErr) {
	var user domain.User
	user.ID = id
	result, err := user.GetUser()
	if err != nil {
		return nil, err
	}
	return result, nil
}
