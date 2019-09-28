package utils

import "github.com/YaroslavRozum/go-boilerplate/models"

func DumpUser(user *models.User) models.User {
	return models.User{
		ID:       user.ID,
		Name:     user.Name,
		Surname:  user.Surname,
		Email:    user.Email,
		UserName: user.UserName,
	}
}
