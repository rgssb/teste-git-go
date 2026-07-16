package data

import "teste-git-go/models"

func ListarAlunos() []models.Aluno {

	return []models.Aluno{
		{"João", 20, 8.5},
		{"Maria", 19, 9.4},
		{"Carlos", 22, 7.8},
		{"Ana", 21, 10},
		{"Pedro", 18, 6.9},
	}
}
