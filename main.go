package main

import (
	"fmt"

	"teste-git-go/data"
	"teste-git-go/services"
)

func main() {

	alunos := data.ListarAlunos()

	fmt.Println("===== LISTA DE ALUNOS =====")

	for _, aluno := range alunos {
		fmt.Println(services.FormatarAluno(aluno))
	}

	fmt.Println()

	media := services.CalcularMediaTurma(alunos)

	fmt.Printf("Média da turma: %.2f\n", media)
}
