package dictionnaire

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	// premiere valeur est celle retourner
	// second est une valeur boolean qui indique si la valeur existe
	definition, ok := d[word]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}
	return definition, nil
}
