package models

import "gorm.io/gorm"

type Aluno struct {
	// O comando abaixo vai adicionar campos que são padrões a um model do gorm (id, datacriação, dataalteração e etc...)
	gorm.Model
	Nome string `json:"nome"`
	CPF  string `json:"cpf"`
	RG   string `json:"rg"`
}
