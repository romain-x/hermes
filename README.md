# Hermes 💬

Un serveur de chat TCP minimaliste en Go 🐹

![Go](https://img.shields.io/badge/Go-1.27-00ADD8?logo=go&logoColor=white)
![Licence](https://img.shields.io/badge/Licence-MIT-green)


> [!CAUTION]
> Ce projet n'a pas vocation à être mis en production : il a été réalisé avant tout dans un but d'apprentissage.


## Prérequis

- [Go](https://go.dev/dl/) 1.27 ou supérieur
- `nc` (netcat) pour se connecter, ou n'importe quel client TCP


## Utilisation

Lancer le serveur :

```sh
go run ./cmd/server
```

Se connecter au chat :

```sh
nc localhost 8090
```

- Choisir un pseudo
- Écrire un message : tout le monde le reçoit
- `/leave` pour quitter le chat sans fermer le serveur pour les autres


## Structure

```
cmd/
└── server/
    └── main.go      démarre le serveur, accepte les connexions
internal/
├── design/
│   └── design.go    bannière ASCII affichée à la connexion
├── identity/
│   ├── identity.go  demande un pseudo et crée l'utilisateur
│   └── user.go      type utilisateur et liste des connectés
└── handlers/
    └── handle.go    lit les messages et les envoie aux autres
```


## Licence

Distribué sous licence MIT. Voir le fichier [LICENSE](LICENSE) pour plus de détails.


## Auteurs

- [@romain-x](https://www.github.com/romain-x)
