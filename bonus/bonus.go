package bonus

//Joãozinho ganhou um kit de construção de torres. O kit é composto por várias barras de madeira, e seus comprimentos são
//conhecidos. As barras podem ser empilhadas umas sobre as outras, desde que seus comprimentos sejam iguais.
//
//Joãozinho quer construir o menor número possível de torres com as barras que tem. Você deve ajudar Joãozinho a usar as
//barras da melhor maneira possível, determinando a altura da torre mais alta e quantas torres podem ser construídas.
func CalculateTowers(barLengths []int) (int, int) {
	altmax := 0
	ntorre := 0
	for i := 0; i < len(barLengths);{
		altura := 0
		ndestaque := 0
		for j := i; j < len(barLengths) || barLengths[i] == barLengths[j]; j++ {
			altura++
			ndestaque++
		}
		if altura > altmax && altura > 1 {
			altmax = altura
		}
		ntorre++
		i += ndestaque
	}
	return altmax, ntorre
}
