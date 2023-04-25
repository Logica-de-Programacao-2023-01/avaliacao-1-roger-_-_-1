package bonus

//Joãozinho ganhou um kit de construção de torres. O kit é composto por várias barras de madeira, e seus comprimentos são
//conhecidos. As barras podem ser empilhadas umas sobre as outras, desde que seus comprimentos sejam iguais.
//
//Joãozinho quer construir o menor número possível de torres com as barras que tem. Você deve ajudar Joãozinho a usar as
//barras da melhor maneira possível, determinando a altura da torre mais alta e quantas torres podem ser construídas.
func CalculateTowers(barLengths []int) (int, int) {
	segura := 0
	for i := 0; i < len(barLengths); i++ {
		for j := 0; j < len(barLengths); j++ {
			if barLengths[i] < barLengths[j] {
				segura = barLengths[i]
				barLengths[i] = barLengths[j]
				barLengths[j] = segura
			}
		}
	}
	altmax := 0
	ntorre := 1
	altura := 1
	for c := 0; c < len(barLengths)-1; c++ {
		if barLengths[c] == barLengths[c+1] {
			altura++
			ntorre--
		} else {
			altura = 1
		}
		if altmax != altura && altura > altmax {
			altmax = altura
		}
		ntorre++
	}
	return altmax, ntorre
}
