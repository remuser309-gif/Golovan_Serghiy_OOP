# IoT Device Management API

Серверна частина курсового проєкту з ООП — REST API для управління IoT-пристроями.

## Технології

- Go 1.26
- Chi router
- PostgreSQL + upper/db
- JWT аутентифікація
- Docker / Render

## Запуск локально

### Вимоги

- Go 1.26+
- PostgreSQL 17+

### Налаштування

Створіть БД:

```sql
CREATE DATABASE iot_db;
```

### Запуск

```bash
cd boilerplate-go-back
$env:DB_NAME="iot_db"
$env:DB_PASSWORD="postgres"
go run ./cmd/server/
```

Або двічі клікніть `start.bat`.

Сервер запуститься на `http://localhost:8080`.

## API Endpoints

| Метод | Шлях | Опис | Auth |
|---|---|---|---|
| `GET` | `/api/ping` | Health check | - |
| `POST` | `/api/v1/auth/register` | Реєстрація | - |
| `POST` | `/api/v1/auth/login` | Вхід | - |
| `POST` | `/api/v1/auth/logout` | Вихід | + |
| `GET` | `/api/v1/users` | Інформація про себе | + |
| `PUT` | `/api/v1/users` | Оновити профіль | + |
| `DELETE` | `/api/v1/users` | Видалити акаунт | + |
| `GET` | `/api/v1/organizations` | Список організацій | + |
| `POST` | `/api/v1/organizations` | Створити організацію | + |
| `GET` | `/api/v1/organizations/{id}` | Організація за ID | + |
| `PUT` | `/api/v1/organizations/{id}` | Оновити організацію | + |
| `DELETE` | `/api/v1/organizations/{id}` | Видалити організацію | + |
| `GET` | `/api/v1/organizations/{orgId}/rooms` | Кімнати організації | + |
| `POST` | `/api/v1/organizations/{orgId}/rooms` | Створити кімнату | + |
| `GET` | `/api/v1/organizations/{orgId}/rooms/{id}` | Кімната за ID | + |
| `PUT` | `/api/v1/organizations/{orgId}/rooms/{id}` | Оновити кімнату | + |
| `DELETE` | `/api/v1/organizations/{orgId}/rooms/{id}` | Видалити кімнату | + |
| `GET` | `/api/v1/organizations/{orgId}/devices` | Пристрої організації (фільтр: `?category=SENSOR`) | + |
| `POST` | `/api/v1/organizations/{orgId}/devices` | Створити пристрій | + |
| `GET` | `/api/v1/organizations/{orgId}/devices/rooms/{roomId}` | Пристрої кімнати | + |
| `GET` | `/api/v1/organizations/{orgId}/devices/{id}` | Пристрій за ID | + |
| `PUT` | `/api/v1/organizations/{orgId}/devices/{id}` | Оновити пристрій | + |
| `DELETE` | `/api/v1/organizations/{orgId}/devices/{id}` | Видалити пристрій | + |
| `GET` | `/api/v1/devices/{devId}/measurements` | Виміри пристрою (фільтр: `?from=...&to=...`) | + |
| `POST` | `/api/v1/devices/{devId}/measurements` | Створити вимір | + |
| `GET` | `/api/v1/devices/{devId}/measurements/{id}` | Вимір за ID | + |
| `PUT` | `/api/v1/devices/{devId}/measurements/{id}` | Оновити вимір | + |
| `DELETE` | `/api/v1/devices/{devId}/measurements/{id}` | Видалити вимір | + |
| `GET` | `/api/v1/devices/{devId}/events` | Події пристрою (фільтр: `?action=...`, `?from=...&to=...`) | + |
| `POST` | `/api/v1/devices/{devId}/events` | Створити подію | + |
| `GET` | `/api/v1/devices/{devId}/events/{id}` | Подія за ID | + |
| `PUT` | `/api/v1/devices/{devId}/events/{id}` | Оновити подію | + |
| `DELETE` | `/api/v1/devices/{devId}/events/{id}` | Видалити подію | + |

## Postman

У папці `.postman/` — готова колекція для тестування:
- `IoT Device Management API.postman_collection.json`
- `IoT Device Management API.postman_environment.json`

Імпортуйте обидва файли в Postman, виберіть середовище **"IoT Device Management API (Local)"** і виконуйте запити по порядку.

## Деплой на Render

1. Створіть PostgreSQL у Render (безкоштовно)
2. Створіть Web Service з Docker-репозиторію
3. Пропишіть змінні середовища:
   - `DB_HOST`, `DB_NAME`, `DB_USER`, `DB_PASSWORD`
   - `JWT_SECRET`
4. Після деплою отримаєте публічну URL-адресу

## Структура проєкту

```
cmd/server/               # Точка входу
config/                   # Конфігурація (БД, JWT)
config/container/         # DI-контейнер
internal/
├── domain/               # Моделі (User, Organization, Room, Device, Measurement, Event)
├── app/                  # Сервіси бізнес-логіки
├── infra/
│   ├── database/         # Репозиторії + SQL-міграції
│   └── http/
│       ├── controllers/  # HTTP-контролери
│       ├── middlewares/   # JWT-мідлвара
│       ├── requests/     # Валідація вхідних даних
│       └── resources/    # DTO для відповідей
.postman/                 # Postman колекція
```
