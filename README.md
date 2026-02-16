# HelmOs

Monorepo with Go API and Next.js 14 Dashboard for HelmOs.

## Structure

```
helmos/
├── apps/
│   ├── api/          → Go backend (port 3001, SQLite, CORS, JWT-ready)
│   └── dashboard/    → Next.js 14 frontend (App Router, TypeScript, Tailwind)
├── docker/
│   └── traefik/      → Traefik configuration
├── .env.example      → Documented environment variables
├── docker-compose.yml      → Production
├── docker-compose.dev.yml  → Local development
└── README.md
```

## Requirements

- **API:** Go 1.22+
- **Dashboard:** Node.js 18+, npm
- **Optional:** Docker and Docker Compose for deployment

## Setup

1. Copy the environment file:

   ```bash
   cp .env.example .env
   ```

2. Edit `.env` and set at least `JWT_SECRET` and `HELMOS_DATA_DIR` (and the rest as needed).

## Local development

### API (Go)

```bash
cd apps/api
go mod tidy   # if you need to fetch dependencies
go run ./cmd/api
```

The API will be available at `http://localhost:3001`. Endpoints: `GET /health`, `GET /db-check`.

### Dashboard (Next.js)

```bash
cd apps/dashboard
npm install
npm run dev
```

The dashboard will be available at `http://localhost:3000`.

### With Docker (development)

```bash
cp .env.example .env
docker compose -f docker-compose.dev.yml --profile dev up
```

## Production

```bash
cp .env.example .env
# Adjust .env (HELMOS_DOMAIN, JWT_SECRET, etc.)
docker compose up -d
```

## Environment variables

See `.env.example`. Summary:

| Variable          | Description                    |
|-------------------|--------------------------------|
| `HELMOS_DOMAIN`   | Service domain                 |
| `HELMOS_EMAIL`    | Admin email                    |
| `HELMOS_DATA_DIR` | Data directory (SQLite)        |
| `HELMOS_PORT`     | API port (default 3001)        |
| `JWT_SECRET`      | Secret for JWT                 |
| `ADMIN_EMAIL`     | Dashboard admin email          |

## License

See [LICENSE](LICENSE).
