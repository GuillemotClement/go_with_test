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
