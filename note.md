# Test

- `import "testing"` est la package pour realiser des tests

- un test est simplement une foncton qui respect quelque regles :
  - le fichierr s'appelle `xxx_test.go`
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
