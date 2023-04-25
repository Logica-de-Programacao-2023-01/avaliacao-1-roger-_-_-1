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

import "strings"

func ProcessString(s string) string {
	var vogais = []string{"A","E","I","O","U","a","e","i","o","u"}
	consoantes := []string{"BCDFGHJKLMNPQRSTVWXYZbcdfghjklmnpqrstvwxyz"}
	for i := 0; i < len(vogais); i++ {
		if strings.ContainsAny(s, vogais[i]) {
			s = strings.ReplaceAll(s, vogais[i], "")
		}
	}
	for i := 0; i < len(consoantes); i++ {
		if strings.ContainsAny(s, consoantes[i]) {
			s = strings.ReplaceAll(s, consoantes[i], "."+strings.ToLower(consoantes[i]))
		}
	}

	return s
}
