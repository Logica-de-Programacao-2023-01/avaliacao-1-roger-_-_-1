package q5

//Pedro começou a frequentar aulas de programação. Na primeira aula, sua tarefa foi escrever um programa simples. O
//programa deveria fazer o seguinte: na sequência de caracteres fornecida, composta por letras latinas maiúsculas e
//minúsculas, ele:
//
//* deleta todas as vogais;
//* insere um caractere "." antes de cada consoante;
//* substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas.
//
//As vogais são as letras "A", "O", "E", "U", "I", e o restante são consoantes. A entrada do programa é exatamente uma
//sequência de caracteres e a saída deve ser uma única sequência de caracteres, resultante após o processamento do
//programa na sequência de caracteres inicial.
//
//Ajude Pedro a lidar com esta tarefa fácil.
import (
	"strings"
)

func ProcessString(s string) string {
	var vogais = []string{"A", "E", "I", "O", "U", "a", "e", "i", "o", "u"}
	consoantes := []string{"B", "C", "D", "F", "G", "H", "J", "K", "L", "M", "N", "P", "Q", "R", "S", "T", "V", "W", "X", "Y", "Z"}
	for _, termo := range vogais {
		s = strings.ReplaceAll(s, termo, "")
	}
	for _, termo := range consoantes {
		s = strings.ReplaceAll(s, strings.ToLower(termo), "."+strings.ToLower(termo))
	}
	for _, termo := range consoantes {
		s = strings.ReplaceAll(s, termo, "."+strings.ToLower(termo))
	}

	return s
}
