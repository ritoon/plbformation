# Array dans le monde réel

![ants](img/ip.png)

L’array, ou tableau, est un type de collection bas niveau permettant d'interagir avec des éléments finis.

Par exemple, je connais à l’avance le nombre d’éléments que va contenir un ensemble, comme des coordonnées GPS, une IPv4 ou un UUID.

Dans la majorité des cas, on va en premier lieu utiliser un slice (tableau dynamique). Et si un jour il y a une contrainte de volumétrie ou un temps de réponse à optimiser, il faudra réaliser une conversion vers un array.

Ce jour est venu et, désolé, mais c’est vous qui allez vous y coller !

**Objectif :** refactorer la structure en slice de l’IPv4 vers un array afin d’optimiser le nombre d’opérations par seconde.

**Étape 1 :** exécuter les tests de benchmark dans votre terminal

```sh
go test -bench=.
```

**Étape 2 :** refactorer le code pour basculer d’un slice vers un array

```go
// remplacer
type IPv4 = []byte
// par
type IPv4 = [4]byte
```

**Étape 3 :** adapter le code pour qu’il puisse fonctionner à nouveau ainsi que les tests.

**Étape 4 :** exécuter à nouveau les tests de benchmark dans votre terminal

```sh
go test -bench=.
```

## Que constatez-vous ?

Le nombre d’opérations par seconde est plus important, avec une empreinte mémoire moindre grâce à l’utilisation d’un array.

## Aller plus loin

- [Implémentation de l’IPv4 dans la std lib de Go](https://pkg.go.dev/net#IPv4)
- [Bench Slice vs Array](https://go-benchmarks.com/array-vs-slice)
