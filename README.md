# Run Instructions

## 1. Start a Fresh DB

```bash
docker compose down -v
docker compose up -d db
```

## 2. Build App

```bash
go build -o /tmp/api-embed ./cmd/api
```

## 3. Run Binary from `/tmp`

This runs the binary outside the repo (for example, without a local `sql/migrations` path).

```bash
cd /tmp
DB_URL="postgres://user:password@localhost:5432/mediaapp?sslmode=disable" \
SERVER_PORT=":8100" \
SERVER_ENV="DEV" \
./api-embed
```
