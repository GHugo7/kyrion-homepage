<h1 align="center">Kyrion Homepage</h1>

<p align="center">
  Dashboard auto-généré pour homelab — découvre les containers Docker et les expose dans une interface web.
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Go-00ADD8?style=flat-square&logo=go&logoColor=white">
  <img src="https://img.shields.io/badge/React_19-61DAFB?style=flat-square&logo=react&logoColor=black">
  <img src="https://img.shields.io/badge/TypeScript-3178C6?style=flat-square&logo=typescript&logoColor=white">
  <img src="https://img.shields.io/badge/SQLite-003B57?style=flat-square&logo=sqlite&logoColor=white">
  <a href="https://kyrion.ovh"><img src="https://img.shields.io/badge/live-kyrion.ovh-success?style=flat-square"></a>
</p>

---

## Principe

La plupart des dashboards de homelab demandent de déclarer chaque service à la main dans un fichier YAML. Ici, la liste se construit toute seule : l'API interroge le daemon Docker, récupère les containers en cours, et va chercher leur URL publique dans Nginx Proxy Manager. Une base SQLite ne sert qu'à stocker ce qu'on veut surcharger — catégorie, URL manuelle, masquage d'un service.

Résultat : un nouveau container apparaît sur le dashboard dès qu'il démarre, sans toucher à la configuration.

## Architecture

```
┌──────────────┐     GET /api/services     ┌──────────────┐
│  React + TS  │ ────────────────────────► │   API Go     │
│    (Vite)    │ ◄──────────────────────── │   (chi)      │
└──────────────┘          JSON             └──────┬───────┘
                                                  │
                            ┌─────────────────────┼─────────────────────┐
                            ▼                     ▼                     ▼
                    ┌───────────────┐   ┌──────────────────┐   ┌──────────────┐
                    │ Docker socket │   │  Nginx Proxy Mgr │   │   SQLite     │
                    │  (containers) │   │  (URLs publiques)│   │ (overrides)  │
                    └───────────────┘   └──────────────────┘   └──────────────┘
```

```
.
├── backend/     # API Go — chi, SQLite
│   ├── main.go       # routeur et démarrage
│   ├── docker.go     # récupération des containers
│   ├── npm.go        # client Nginx Proxy Manager
│   ├── bdd.go        # accès SQLite et schéma
│   ├── handlers.go   # handlers HTTP
│   └── struct.go     # types partagés
└── frontend/    # interface React + TypeScript (Vite)
```

## Prérequis

- Go — version indiquée dans `backend/go.mod`
- Node.js 20+
- Accès au socket Docker (`/var/run/docker.sock`)
- Une instance Nginx Proxy Manager joignable

## Installation

```bash
git clone https://github.com/GHugo7/kyrion-homepage.git
cd kyrion-homepage
```

### Backend

Créer `backend/.env` :

```env
NPM_IDENTITY=admin@example.com
NPM_SECRET=votre_mot_de_passe
ALLOWED_ORIGINS=http://localhost:5173
```

Puis :

```bash
cd backend
go mod download
go run .
```

L'API écoute sur le port `1818`.

### Frontend

```bash
cd frontend
npm install
npm run dev
```

L'interface est servie sur `http://localhost:5173`.

## Configuration

| Variable | Description | Défaut |
|---|---|---|
| `NPM_IDENTITY` | Identifiant Nginx Proxy Manager | — |
| `NPM_SECRET` | Mot de passe Nginx Proxy Manager | — |
| `ALLOWED_ORIGINS` | Origines CORS autorisées, séparées par des virgules | `http://localhost:5173` |

## API

| Méthode | Route | Description |
|---|---|---|
| `GET` | `/api/services` | Liste des services détectés, avec catégorie et URL |

## Base de données

SQLite, un seul table. Le fichier est créé automatiquement au démarrage s'il n'existe pas.

```sql
CREATE TABLE service_overrides (
    container_name TEXT PRIMARY KEY,
    category       TEXT,
    url            TEXT,
    enabled        BOOLEAN DEFAULT 1
);
```

Un container absent de cette table est affiché avec ses valeurs par défaut. Passer `enabled` à `0` le masque du dashboard.

## Licence

MIT