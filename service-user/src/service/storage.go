package service

import (
	"github.com/payments/service-user/src"
)

type Service struct {
	Repository src.RepositoryList
}

type User struct {
	Name      string `json:"name"`
	Cpf_cnpj  string `json:"cpf_cnpj"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	User_type int    `json:"user_type"`
}
