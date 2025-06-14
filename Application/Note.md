# HTTP Server

Creation d'une application Web ou les utilisateurs peuvent tracker combien de partie ils ont gagne.

- `GET /players/{name}` : retourne le nombre total de victoire
- `POST /players/{name}` : enregistre une victoire pour le joueur

## Construction TDD

On ne peut pas `GET` un joueur sans avoir stoker de la data, et il est difficile de tester si `POST` fonctionne sans le `GET`.

On mets en place un mock pour simuler cela.
    - On creer une interface pour recuperer les donnees d'un joueur
    - On creer un espions sur les appel a `PlayerStore` pour s'assurer du stocker des donnees

## Lancer un app Go 

```shell
# build de l'application
go build 

# lancer le programme
./<name_build>
```
