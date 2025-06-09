# Test

- `import "testing"` est la package pour realiser des tests

- un test est simplement une foncton qui respect quelque regles :
  - le fichierr s'appelle `xxx_test.go`
  - le nom de la fonction commence par `Test`
  - la fonction prend 1 argument `t *testing.T`

`t.Errorf()` est une methode qui permet d'affocher des message quand le test echoue

- `%q` permet d'afficher la valeur de la variable entre guillemet
