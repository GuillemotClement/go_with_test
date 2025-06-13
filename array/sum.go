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
