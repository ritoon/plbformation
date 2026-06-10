## US Timbrée

**Objectif du TP :** Créer un _worker_ chargé de télécharger des images et de leur ajouter un **watermark**.

**Contexte :** Vous êtes le nouveau développeur chez **Roumet**, la plus vieille maison de vente philatélique, fondée en 1896.

Julia, en charge de la comptabilité et de la refonte du site,
a remarqué que, depuis quelque temps, sur [Catawiki](https://www.catawiki.com/fr), des vendeurs s’approprient leur base d’images pour réaliser des ventes commissionnées à leurs dépens.

Un lundi matin, vos paupières se lèvent à peine après un troisième café serré quand Julia entre dans votre champ de vision avec un sourire tendre et lance :

> « C’est compliqué d’ajouter un _trademark_ sur les images ? »

Jean-Charles, développeur vétéran et ardent défenseur de PHP 5, ne vous a pas laissé voir son code source, de peur que vous ne copiiez son “beau travail”.
`
C’est à vous de jouer : créez un **outil** permettant de récupérer localement les images d’une vente pour les traiter. En l’absence de _backend_, Julia se chargera de les renvoyer sur le FTP une fois qu’elles auront été traitées.

**Exemple d’URL :**

```sh
https://www.roumet.com/photos/574/1.jpg
```

Soit : `https://www.roumet.com/photos/{numero_de_la_vente}/{numero_du_lot}.jpg`

**Attention :** lorsqu’il y a plusieurs images pour un même lot, l’URL peut être :

```sh
https://www.roumet.com/photos/574/1-1.jpg
```

Les numéros de lots se suivent et s’incrémentent jusqu’à **1000**.

Si il y a plusieurs images sur un même lot, leur nombre ne dépasse pas 10.

### Tâches

1. **Téléchargement local des images**

   - Générer les URLs en fonction des lots et des sous-images éventuelles.
   - Enregistrer les images dans un premier dossier `tmp/`.

2. **Ajout un watermark**

   - Lire les images dans le dossier `tmp/`.
   - Récupérer l'image du logo Roumet pour l'apposer en bas à droite de l'image du timbre.
   - Écrire les versions traitées dans le dossier `processed/` avec le watermark ajouté.

### Contraintes

- Le code doit utiliser le **design pattern Worker**.
- Les images doivent conserver **exactement** le même nom qu’auparavant.
- Vous ne pouvez pas utiliser plus de **trois workers** à la fois, au risque de faire tomber le service du site. Laissez-le souffler à chaques requêtes.
- Pour le TP, choisissez **une seule vente** (entre **527** et **574**).
- Le logo Roumet se trouve ici : `img/logo.png`.

#### Packages recommandés

- [Package `os` — gestion des fichiers et dossiers](https://pkg.go.dev/os)
- [Package `net/http` — téléchargement des images](https://pkg.go.dev/net/http)
- [Package `mergi` — ajout du watermark](https://pkg.go.dev/github.com/noelyahan/mergi#Watermark)
