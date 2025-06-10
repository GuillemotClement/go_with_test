# go_with_test

## Workflow

1. Ecriture des tests
2. Faire en sorte que la compilation se fasse
3. Lancer le test et regarder ce qui fail
4. Ecrire le code minimal pour que le test passe
5. Refactorisation

On travail avec ce workflow. Pour chaque nouveau comportement que l'on souhaite ajouter, on viens proceder de la sorte. De cette maniere le code est tester en continu, et on limite les risque de regression.

Cela permet egalement de bien refecnhir a ce qui est attendu du code.

## Package

Chaque dossier dans un projet Go correspond a un package,

Pour exporter une methode, son nom doit commencer par une majuscule.
(Setting a Go Project)[https://dave.cheney.net/2014/12/01/five-suggestions-for-setting-up-a-go-project]

### Declarer un nouveau package

Pour definir le nom du package, on viens definir la premiere ligne dans le fichier

```go
package <file_name>
```

## Go Doc

On peut venir ajouter de la documentation pour les methodes en ajoutant un commentaire sur la ligne au dessus de la declaration de la fonction.

### Exemple

Dans le fichier de test, on peut ajouter une fonction d'exemple. Le nom de la fonction commence avec `Example<functionName>`.
