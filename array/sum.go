package array

func Sum(numbers []int) int {
	sum := 0
	// on parcours le tableau de nombre
	// _ permet d'ignorer l'index de l'element en cours d'iteration
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// ... permet de passer un nombre infini de parametre a la fonction.

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		// append permet d'ajouter dans le slice la sum
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			// sums contiens la valeur precedente
			sums = append(sums, 0)
		} else {
			tail := numbers[1:] // on recupere a partir du 2 eme element du slice
			sums = append(sums, Sum(tail))
		}

	}

	return sums
}
