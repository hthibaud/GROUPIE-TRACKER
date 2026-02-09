# GROUPIE-TRACKER

Projet web Groupie-Tracker utilisant des API et fait en Go pour consulter des artistes, leurs dates et lieux de concerts.

**Résumé**
- **But :** Fournir une interface web simple listant des artistes, leurs relations (dates/lieux) et des pages de détails.
- **Technologie :** Serveur HTTP en Go, templates HTML dans `web/templates`, ressources statiques dans `web/assets`.

**Prérequis**
- Go
- HTML
- Tailwind

**Installation & Exécution**
1. Cloner le dépôt.
2. Depuis la racine du projet, lancer le serveur : go run ./cmd/server/main.go

```bash
# Groupie Tracker

1. Introduction

API & Application Web — développer une application web interactive de suivi de groupes musicaux en utilisant Go et des technologies web modernes.

Objectif
- Concevoir et livrer une application web fonctionnelle permettant de consulter des artistes, leurs dates et lieux de concerts, et d'explorer les relations entre données.
- Travail en 9 séances de 4 heures, avec une soutenance finale.

Compétences visées
- Structure du code et architecture (endpoints / HTTP)
- Gestion des erreurs HTTP
- Tests unitaires et documentation du code
- Développement Backend : manipulation des 4 endpoints, gestion des relations, parsing JSON
- Intégration API : recherche avec suggestions, filtres, carte interactive des lieux
- Interface Utilisateur : clarté des vues et de la démo
- Soft skills : présentation orale, justification des choix, qualité des réponses

Livrables attendus
- Projet fonctionnel et site complet
- Présentation orale (intro, features, démo technique, retour d'expérience)
- Gestion de projet : README détaillé, guide d'usage, description légère de l'architecture
- Repository GitHub organisé (branches, commits clairs - conventional commits)
- Archive .zip du projet

En bref
- Fournir le projet fini, la documentation (`README.md`) contenant le guide d'usage et indications Git/GDP.

2. Comment démarrer (Guide rapide)

Prérequis
- Go
- HTML
- Tailwind

Lancer le serveur localement

```bash
go run ./cmd/server/main.go
```

Ouvrir : http://localhost:8080

Routes principales
- `/` — page d'accueil
- `/artists` — liste et filtres
- `/artist?id=<id>` — détail artiste
- `/locations` — page des lieux

- Ressources statiques servies depuis `/static/` (dossier `web/assets`)

Points rapides
- Le serveur écoute par défaut sur le port `8080` (voir `cmd/server/main.go`).

---

Fichier principal : `cmd/server/main.go` — voir les handlers dans `internal/handlers` et les templates dans `web/templates`.