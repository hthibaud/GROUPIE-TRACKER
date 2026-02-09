# Groupie Tracker

Application web interactive de suivi de groupes musicaux utilisant Go et une API externe.

---

## 1. Objectif du Projet

**Groupie Tracker** est une application web développée en Go permettant de consulter les informations sur les groupes musicaux, leurs dates de concert et leurs lieux de tournée. Le projet combine :

- **Backend** : Serveur HTTP en Go avec récupération et parsing de données JSON depuis une API externe
- **Frontend** : Interface web responsive avec filtres et recherche en temps réel
- **Intégration API** : Consommation de l'API [Groupie Trackers](https://groupietrackers.herokuapp.com/api)

**Compétences développées :**
- Architecture HTTP et routage
- Gestion des erreurs HTTP
- Parsing JSON et structures Go
- Templating avec `html/template`
- Recherche et filtres côté serveur
- Conception responsive avec Tailwind CSS

---

## 2. Installation et Lancement

### Prérequis
- **Go** 1.18 ou supérieur
- **Git** (pour cloner le dépôt)

### Cloner le projet
```bash
git clone https://github.com/hthibaud/GROUPIE-TRACKER
cd GROUPIE-TRACKER
```

### Lancer le serveur
```bash
go run ./cmd/server
```

Le serveur démarre alors sur **http://localhost:8080**

### Architecture générale
```
GROUPIE-TRACKER/
├── cmd/server/           # Point d'entrée de l'application
├── internal/api/         # Appels API et structures de données
├── internal/handlers/    # Routeurs et handlers
└── web/
    ├── templates/        # Fichiers HTML (index, artists, artist, locations, etc.)
    └── assets/          # Ressources statiques (CSS, JS, images, vidéos)

---

## 3. Routes Principales et Fonctions

| Route | Méthode | Description |
|-------|---------|-------------|
| `/` | GET | **Page d'accueil** — Affiche une présentation générale et liens vers les sections principales |
| `/artists` | GET | **Liste des artistes** — Affiche tous les artistes avec filtres avancés |
| `/artist?id=<id>` | GET | **Détail d'un artiste** — Page détaillée avec info artiste, dates et lieux de concerts |
| `/locations` | GET | **Carte des lieux** — Affiche les lieux de concerts de manière interactive |
| `/static/*` | GET | **Ressources statiques** — CSS, JS, images, vidéos depuis `web/assets` |

---

## 4. Fonctionnalités Implémentées

### Fonctionnalités Obligatoires

- [x] **Récupération données API** — Consommation des 4 endpoints de l'API Groupie Trackers (artists, locations, dates, relations)
- [x] **Affichage artistes** — Liste dynamique de tous les artistes
- [x] **Détail artiste** — Page complète avec informations, membres, dates et lieux
- [x] **Filtres avancés** — Recherche par nom, année création/album, nombre de membres, lieux
- [x] **Gestion erreurs HTTP** — Codes 404, 500 avec pages d'erreur appropriées
- [x] **Serveur web fonctionnel** — Port 8080, gestion des requêtes GET
- [x] **Interface responsive** — Utilisation de Tailwind CSS pour esthétique cohérente
- [x] **Routage dynamique** — Routes paramétrées pour les artistes individuels

### Fonctionnalités Bonus

- [x] **Recherche temps réel** — Filtres côté serveur avec traitement instantané
- [x] **Page locations** — Vue des lieux de tournée
- [x] **Gestion ressources statiques** — Serveur de fichiers pour CSS/JS/images/vidéos
- [x] **Pages d'erreur personnalisées** — 404.html et 500.html stylisés

---

## 5. Gestion de Projet

### Méthodologie
- **Type de versioning** : Git avec GitHub
- **Branches** : Organisation par feature/fix/docs
- **Commits** : Messages clairs et descriptifs (Conventional Commits)

### Lien GitHub
> **https://github.com/hthibaud/GROUPIE-TRACKER**

### Lien Notion
> **https://www.notion.so/GROUPIE-TRACKER-2c56c25064e880c189e2cb2690b26094**

### Convention Commits
```
feat: ajoute filtres de recherche avancée
fix: corrige erreur 404 sur route artist
docs: met à jour README avec exemples
refactor: simplifie la gestion des templates
```

---

## 6. Utilisation et Tests

### Vérifier l'installation
```bash
go mod download
go run ./cmd/server
# Le serveur doit démarrer sans erreur
```
---

**Développé par :** Romain Carrere et Thibaud Herry  
**Date de création :** Décembre 2025  
**Dernière mise à jour :** Février 2026