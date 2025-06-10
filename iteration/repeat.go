package iteration

func Repeat(character string) string {
	var repeated string // on declare une variable sans lui affected de valeur

	// on creer une boucle pour repeter le character, et pour chaque iteration l'ajouter a la string qui est ensuite retourner
	for i := 0; i < 5; i++ {
		repeated = repeated + character
	}
	return repeated
}
