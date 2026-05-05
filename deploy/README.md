# Eater Service - Docker Deployment

Bu folder Docker bilan eater-service ni deploy qilish uchun kerakli fayllarni o'z ichiga oladi.

## Fayllar tavsifi

- **Dockerfile** - Go applicationni multi-stage build bilan containerga build qiladi
- **docker-compose.yml** - Eater service va PostgreSQL databaseni birga run qiladi
- **.dockerignore** - Docker build contextdan chiqarib yuborilishi kerak bo'lgan fayllar
- **.env.example** - Environment variables namunasi

## Talab qilingan dasturlar

- Docker Desktop (yoki Docker + Docker Compose)
- Port 8080 va 5432 bo'sh bo'lishi kerak

## Ishlatish

### 1. Environment variables sozlaymiz

```bash
cd deployment
cp .env.example .env
# .env faylini o'zingizning config bilan tahrirlang
```

### 2. Docker compose ishga tushurish

```bash
cd deployment
docker-compose up -d
```

### 3. Service ishlayotganini tekshirish

```bash
# Logs ko'rish
docker-compose logs -f eater-service

# Health check
curl http://localhost:8080/health
```

### 4. Database migratsiyalarini qo'llash (agar kerak bo'lsa)

```bash
docker-compose exec eater-service ./migrations/run.sh
```

### 5. Service ni to'xtatish

```bash
docker-compose down
```

### 6. Barcha data (volumes) ni o'chirish

```bash
docker-compose down -v
```

## Production sozlamalari

Production da ishlatish uchun:

1. `.env` faylida `JWT_SECRET` ning kuchli qiymatini o'rnating
2. `DB_PASSWORD` ni complex parolga o'zgarting
3. `APP_ENV=production` ga o'rnating
4. SSL/TLS sertifikatlarini qo'shing (nginx yoki load balancer orqali)
5. Resource limitlarini qo'shing `docker-compose.yml` ga

## Migratsiyalar

Migratsiyalar avtomatik `postgres` container ishga tushdirilaganda `/docker-entrypoint-initdb.d` dan qo'llaniladi.

Custom migratsiya uchun:

```bash
docker-compose exec postgres psql -U eater -d eater_db -f /migrations/000001_init.up.sql
```

## Container nomi va network

- **Service container**: `eater-service`
- **Database container**: `eater-postgres`
- **Network**: `eater-network`

Container ichidan database ga ulanish: `postgres://eater:eater123@postgres:5432/eater_db`

## Troubleshooting

### Container ishga tushmasa:

```bash
docker-compose logs eater-service
```

### Database ulanishining xato bo'lsa:

```bash
docker-compose exec postgres psql -U eater -d eater_db -c "SELECT 1"
```

### Port band bo'lsa:

```bash
# 8080 portga javob berayotgan processni topish
lsof -i :8080
# yoki Windows da
netstat -ano | findstr :8080
```

## Build qilish o'zgartirilgan kod bilan

```bash
docker-compose build --no-cache
docker-compose up -d
```
