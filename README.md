# HelmOs

Monorepo con API (Go) y Dashboard (Next.js 14) para HelmOs.

## Estructura

```
helmos/
├── apps/
│   ├── api/          → Backend en Go (puerto 3001, SQLite, CORS, JWT listo)
│   └── dashboard/    → Frontend Next.js 14 (App Router, TypeScript, Tailwind)
├── docker/
│   └── traefik/      → Configuración de Traefik
├── .env.example      → Variables de entorno documentadas
├── docker-compose.yml      → Producción
├── docker-compose.dev.yml  → Desarrollo local
└── README.md
```

## Requisitos

- **API:** Go 1.22+
- **Dashboard:** Node.js 18+, npm
- **Opcional:** Docker y Docker Compose para despliegue

## Configuración

1. Copia las variables de entorno:

   ```bash
   cp .env.example .env
   ```

2. Edita `.env` y define al menos `JWT_SECRET` y `HELMOS_DATA_DIR` (y el resto según necesites).

## Desarrollo local

### API (Go)

```bash
cd apps/api
go mod tidy   # si hace falta descargar dependencias
go run ./cmd/api
```

La API estará en `http://localhost:3001`. Endpoints: `GET /health`, `GET /db-check`.

### Dashboard (Next.js)

```bash
cd apps/dashboard
npm install
npm run dev
```

El dashboard estará en `http://localhost:3000`.

### Con Docker (desarrollo)

```bash
cp .env.example .env
docker compose -f docker-compose.dev.yml --profile dev up
```

## Producción

```bash
cp .env.example .env
# Ajusta .env (HELMOS_DOMAIN, JWT_SECRET, etc.)
docker compose up -d
```

## Variables de entorno

Ver `.env.example`. Resumen:

| Variable          | Descripción                    |
|-------------------|--------------------------------|
| `HELMOS_DOMAIN`   | Dominio del servicio           |
| `HELMOS_EMAIL`    | Email del administrador        |
| `HELMOS_DATA_DIR` | Directorio de datos (SQLite)   |
| `HELMOS_PORT`     | Puerto de la API (default 3001)|
| `JWT_SECRET`      | Secreto para JWT               |
| `ADMIN_EMAIL`     | Email del admin del dashboard  |

## Licencia

Ver [LICENSE](LICENSE).
