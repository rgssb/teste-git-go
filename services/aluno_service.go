package services

import (
	"fmt"

	"teste-git-go/models"
	"teste-git-go/utils"
)

func FormatarAluno(a models.Aluno) string {

	status := "Reprovado"

	if utils.Aprovado(a.Nota) {
		status = "Aprovado"
	}

	return fmt.Sprintf(
		"%s | %d anos | Nota %.1f | %s",
		a.Nome,
		a.Idade,
		a.Nota,
		status,
	)
}

func CalcularMediaTurma(alunos []models.Aluno) float64 {

	total := 0.0

	for _, aluno := range alunos {
		total += aluno.Nota
	}

	return total / float64(len(alunos))
}
