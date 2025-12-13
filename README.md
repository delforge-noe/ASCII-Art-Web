# Générateur ASCII-ART-WEB

Ce projet est une application web simple développée en Golang (Go) qui a pour but de prendre du texte en entrée et d'en ressortir la version ASCII art en utilisant différentes polices (bannières) que l'utilisateur peut séléctionner

---

## Fonctionnalités

- **Conversion ASCII Art** : le texte choisi par l'utilisateur sera converti en ASCII Art tant qu'il n'utilise pas de caractères qui ne sont pas compris entre les valeurs 32 (inclus) et 126 (inclus) de la table ASCII.
- **Choix de la bannière (police)** : Supporte les bannières standard, shadow, et thinkertoy.
- **Interface Web (front-end)** : Interface utilisateur intuitive construite en HTML/CSS.
- **Gestion des Erreurs HTTP** : Implémentation des codes de statut HTTP standard selon les spécifications :
  - `200 OK` (Succès)
  - `400 Bad Request` (Requète invalide, caractère non supporté)
  - `404 Not Found` (Route non gérée)
  - `500 Internal Server Error` (Erreur du serveur lors du rendu)

---

## Technologies Utilisées

- **Langage** : Golang (Go)
- **Serveur Web** : Package `net/http` de Go.
- **Front-end** : HTML / Go Templates (`html/template`)

---

## Structure du Projet

``` text
.
├── assets/
│   ├── image1.png      
│   ├── image2.png
│   └── image3.png
├── banners/
│   ├── standard.txt    
│   ├── shadow.txt
│   └── thinkertoy.txt
├── templates/
│   └── index.html
├── gene/
|    └── generateur.go
├── main.go                   
├── README.md 
└── go.mod
```

---

## Comment Utiliser le Projet

- **Prérequis**

Assurez vous d'avoir Go (version 1.18 ou supérieure) installé sur votre machine.

- **Étapes**

1. **Cloner le dépôt :**  Dans votre terminal

``` bash
$ git clone https://github.com/delforge-noe/ASCII-Art-Web.git
```

2. **Installer les dépendances** : Si vous utilisez un module externe (`ascii-art-web/gene` dans ce cas là), assurez vous que les modules sont initialisés :

``` bash
$ cd ASCII-Art-Web
$ go mod tidy
```

3. **Lancer le serveur** : Exécutez le fichier `start.sh`, il lance le serveur et ouvre automatiquement `index.html` dans votre navigateur ouvert :

``` bash
$ bash start.sh
```

---

## Auteurs

Noé Delforge,
Alix Blondel,
Margot Cahard.
Un projet fait dans le cadre de la formation de Zone01.