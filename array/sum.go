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
	lengthOfNumbers := len(numbersToSum) //on stocke le nombre d'element dans le slice
	// make est une fonction qui permet de preparer un slice
	// la capaciter est le nombre d'element pouvant etre stocker dedans
	sums := make([]int, lengthOfNumbers) // on prepare un nouveau slice avec une capactiter du nombre d'element
	// on itere le slice de slice et pour chaque on fait la somme
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}
