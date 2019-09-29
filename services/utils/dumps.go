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

func DumpProduct(product *models.Product) models.Product {
	return models.Product{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}
}
