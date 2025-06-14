# Test

- `import "testing"` est la package pour realiser des tests

- un test est simplement une fonction qui respect quelque regles :
  - le fichier s'appelle `xxx_test.go`
  - le nom de la fonction commence par `Test`
  - la fonction prend 1 argument `t *testing.T`

`t.Errorf()` est une methode qui permet d'affocher des message quand le test echoue

- `%q` permet d'afficher la valeur de la variable entre guillemet

## Workflow

Une fois que le test est ecrit et que le code passe le test avec succes, on peut faire un commit pour conserver un code fonctionnel.

On peut ensuite lancer une refacto mais sans push sur la main afin d'eviter des probleme si on reprend le code plus tard.

## Constante

```go
// Declarer une constante
const maConstante = "Hello, "
```

## Test

### Sous test

On peut venir grouper plusieurs test pour rasssembler tout les test d'une seule fonction.

Pour cela, on definit la methode que l'on veux tester puis on definis les differents sous test permettant de tester les differents comportement attendu.

### Ecrire des tests

- `testing.TB` est une interface qui satisfait `*testing.T` et `*testing.B`
- `t.Helper()` est utiliser pour indiquer au test d'utilser cette methode comme un helper

## Benchmark

```shell
go test -bench=.
```

## Coverage

```shell
go test -cover
```

## Fonction

### Variadic Function
C'est une fonction qui peut accepter un nombre indefinie de parametre d'entree. On utilise `...` devant le type de l'argument.
```go
// la fonction peut prendre plusieurs slice de int et retourne in slice de int
func SumAll(numbersToSum ...[]int) []int {
	return nil
}
```


## Buffer
Un buffer est une zone de memoire temporaire ou on stocke des donnees avant de les utiliser, les envoyer ou les afficher. 

## Writter
L'interface `io.Writter` represente n'importe quoi qui sait ecrire des donnees.
```go
type Writer interface {
    Write(p []byte) (n int, err error)
}
```
Si un type a une methode `Write([]byte) (int, error)` alors il implemente l'interface.

Cet interface permet d'ecrire du code generatique qui peut ecrire dans :
- un fichier
- un buffer memoire
- la console
- un socket reseau
- etc
Sans que le code ait besoin de savoir ou ecrire
```go
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// fonction generique qui permet d'ecrire un message
// elle n'as pas besoin de savoir ou elle doit ecrire
// elle permet uniquement d'ecrire dans quelque chose qui implemente `io.Writter`
// c'est l'injection de dependance
func Greet(writer io.Writer, name string) {
	// le writter permet de definis ou ou veux ecrire le message
	fmt.Fprintf(writer, "Hello, %s", name)
}
// la fonction permet d'ecrire un message dans une requete HTTP
// elle appelle Greet qui ecrit rellement le message
// Cette fonction permet de definir ou le message sera ecrit => ici une requete HTTP
func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	// on demarre un serveur qui ecoute sur le port 5001
	// a chaque requete on appelle la fonction qui ecrit le message
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
```

```go
Client navigateur ==> Serveur HTTP (port 5001)
       ==> MyGreeterHandler(w, r)
           ==> Greet(w, "world")
               ==> fmt.Fprintf(w, "Hello, world")
                   ==> La réponse HTTP contient "Hello, world"
```


## Mocking 
