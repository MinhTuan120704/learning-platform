# Learning Platform

Backend microservices cho nền tảng học trực tuyến. Repo gồm 3 service Go chạy độc lập, dùng PostgreSQL riêng từng service, Redis dùng chung, Docker Compose để chạy local.

## Kiến trúc

```text
learning-platform
├── services
│   ├── identity   # đăng ký, đăng nhập, JWT, role, permission
│   ├── course     # category, course, section, lesson
│   └── learning   # enrollment
├── docker-compose.yml
├── Makefile
└── go.work
```

```text
                        +----------------+
                        |     Client     |
                        +--------+-------+
                                 |
                        (API Gateway - chưa có)
                                 |
        --------------------------------------------------
        |                    |                    |
+---------------+   +----------------+   +----------------+
| identity      |   | course         |   | learning       |
| :8081         |   | :8082          |   | :8083          |
+---------------+   +----------------+   +----------------+
        |                    |                    |
   postgres              postgres              postgres
  (identity_db)         (course_db)          (learning_db)
        |                    |                    |
        --------------------- redis -----------------
                (DB0: identity, DB1: course permission
                 cache, DB2: learning — kết nối sẵn,
                 chưa dùng để cache)
```

**Giao tiếp giữa các service:**
- `course` → `identity`: gọi `GET /internal/permissions?user_id=` (header `X-Internal-Api-Key`) để lấy role/permission, cache kết quả trong Redis DB1 (TTL 2 phút).
- `learning` → `course`: gọi `GET /api/v1/courses/:courseId` (public) để lấy thông tin course khi cần hiển thị chi tiết. `learning` không gọi `identity`.
- `course` và `learning` không tự verify JWT — nhận danh tính qua header `X-User-ID`, giả định đã được một tầng phía trước (Gateway, hiện tại là client tự set khi test) xác thực và forward xuống.

Service chính:

| Service | Port | Database | Vai trò |
| --- | ---: | --- | --- |
| `identity` | `8081` | `identity_db` | Auth, user, role, permission, internal permission API |
| `course` | `8082` | `course_db` | Course catalog, category, section, lesson |
| `learning` | `8083` | `learning_db` | User enrollment |
| `redis` | `6379` | DB `0/1/2` | Refresh token (dự phòng), permission cache, cache dùng chung |

## Tech stack

- Go `1.26.5`
- Gin
- PostgreSQL `17`
- Redis `8`
- pgx
- golang-migrate
- Docker Compose

## Chạy local bằng Docker

Yêu cầu:

- Docker + Docker Compose
- `make`
- `golang-migrate` nếu muốn chạy migration bằng Makefile

### 1. Tạo file env cho từng service và root

```bash
cp .env.example .env
cp services/identity/.env.example services/identity/.env
cp services/course/.env.example services/course/.env
cp services/learning/.env.example services/learning/.env
```

File `.env` ở root chứa `REDIS_PASSWORD` để Docker Compose interpolate vào `REDIS_URL` của từng service.


### 2. Chạy toàn bộ stack

```bash
make build
```

Kiểm tra container:

```bash
make ps
```

### 3. Health check

```bash
curl http://localhost:8081/health
curl http://localhost:8082/health
curl http://localhost:8083/health
```

## Migration

Sau khi database container healthy, chạy migration:

```bash
make migrate-up-identity
make migrate-up-course
make migrate-up-learning
```

Rollback 1 version:

```bash
make migrate-down-identity
make migrate-down-course
make migrate-down-learning
```

Xem version:

```bash
make migrate-version-identity
make migrate-version-course
make migrate-version-learning
```

Tạo migration mới:

```bash
make create-migration DIR=services/course/migrations NAME=create_course_table
```

## Dữ liệu seed

`identity` seed:

- Admin: `admin@test.com` / `Password123!`
- Instructor: `instructor1@test.com` ... `instructor10@test.com` / `Password123!`
- Student: `student1@test.com` ... `student10@test.com` / `Password123!`

Role có sẵn: `admin`, `instructor`, `student`

Permission có sẵn:

- `category.create`, `category.manage`
- `course.create`, `course.manage`
- `section.create`, `section.manage`
- `lesson.create`, `lesson.manage`
- `enrollment.manage`, `enrollment.self`
- `user.manage`, `role.manage`

`course` seed:

- Category: Lập trình Web, Lập trình Mobile, Khoa học dữ liệu
- Course mẫu: Golang, React.js, Flutter, Machine Learning
- Section + lesson mẫu cho Golang, React.js

## API

### Identity service

Base URL: `http://localhost:8081`

| Method | Path | Auth | Mô tả |
| --- | --- | --- | --- |
| `GET` | `/health` | Không | Health check |
| `POST` | `/api/v1/auth/register` | Không | Tạo user |
| `POST` | `/api/v1/auth/login` | Không | Đăng nhập |
| `POST` | `/api/v1/auth/refresh` | Không | Refresh token |
| `POST` | `/api/v1/auth/logout` | Không | Revoke refresh token |
| `GET` | `/internal/permissions?user_id={uuid}` | `X-Internal-Api-Key` | Lấy role + permission user |

```bash
curl -X POST http://localhost:8081/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@test.com","password":"Password123!"}'
```

```bash
curl -X POST http://localhost:8081/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"name":"Test User","email":"user@test.com","password":"Password123!"}'
```

### Course service

Base URL: `http://localhost:8082`

Public endpoints:

| Method | Path | Mô tả |
| --- | --- | --- |
| `GET` | `/health` | Health check |
| `GET` | `/api/v1/categories` | Danh sách category |
| `GET` | `/api/v1/categories/:id` | Chi tiết category |
| `GET` | `/api/v1/courses` | Danh sách course |
| `GET` | `/api/v1/courses/:courseId` | Chi tiết course |
| `GET` | `/api/v1/courses/:courseId/sections` | Section trong course |
| `GET` | `/api/v1/sections/:sectionId/lessons` | Lesson trong section |
| `GET` | `/api/v1/lessons/:id` | Chi tiết lesson |

Protected endpoints dùng header `X-User-ID`. Service gọi `identity` qua `/internal/permissions` và cache permission trong Redis 2 phút.

| Method | Path | Permission |
| --- | --- | --- |
| `POST` | `/api/v1/categories` | `category.create` |
| `PATCH` | `/api/v1/categories/:id` | `category.manage` |
| `DELETE` | `/api/v1/categories/:id` | `category.manage` |
| `POST` | `/api/v1/courses` | `course.create` |
| `PATCH` | `/api/v1/courses/:courseId` | `course.manage` |
| `DELETE` | `/api/v1/courses/:courseId` | `course.manage` |
| `POST` | `/api/v1/courses/:courseId/sections` | `section.create` |
| `PATCH` | `/api/v1/sections/:id` | `section.manage` |
| `DELETE` | `/api/v1/sections/:id` | `section.manage` |
| `POST` | `/api/v1/sections/:sectionId/lessons` | `lesson.create` |
| `PATCH` | `/api/v1/lessons/:id` | `lesson.manage` |
| `DELETE` | `/api/v1/lessons/:id` | `lesson.manage` |

```bash
curl -X POST http://localhost:8082/api/v1/courses \
  -H "Content-Type: application/json" \
  -H "X-User-ID: <admin-or-instructor-user-id>" \
  -d '{
    "category_id": "11111111-1111-1111-1111-111111111111",
    "title": "Go REST API",
    "slug": "go-rest-api",
    "description": "Xay dung REST API bang Go"
  }'
```

### Learning service

Base URL: `http://localhost:8083`

Enrollment endpoints dùng header `X-User-ID`.

| Method | Path | Mô tả |
| --- | --- | --- |
| `GET` | `/health` | Health check |
| `POST` | `/api/v1/enrollments` | Ghi danh course |
| `GET` | `/api/v1/enrollments/me` | Danh sách course đã ghi danh |
| `DELETE` | `/api/v1/enrollments/:courseId` | Hủy ghi danh |

```bash
curl -X POST http://localhost:8083/api/v1/enrollments \
  -H "Content-Type: application/json" \
  -H "X-User-ID: <student-user-id>" \
  -d '{"course_id":"a1111111-1111-1111-1111-111111111111"}'
```

## Environment

Các file mẫu nằm tại:

- `services/identity/.env.example`
- `services/course/.env.example`
- `services/learning/.env.example`

Biến quan trọng:

| Biến | Service | Mô tả |
| --- | --- | --- |
| `APP_ENV` | Tất cả | Môi trường chạy |
| `HTTP_HOST`, `HTTP_PORT` | Tất cả | Host/port HTTP |
| `POSTGRES_HOST`, `POSTGRES_PORT`, `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB`, `POSTGRES_SSLMODE` | Tất cả | Kết nối PostgreSQL |
| `REDIS_URL` | Tất cả | Redis connection URL, Compose inject theo DB `0/1/2` |
| `JWT_SECRET`, `JWT_ACCESS_EXPIRE_MINUTES`, `JWT_REFRESH_EXPIRE_DAYS`, `JWT_ISSUER` | `identity` | JWT config |
| `IDENTITY_SERVICE_URL` | `course`, `learning` | URL gọi identity service |
| `INTERNAL_API_KEY` | `identity`, `course`, `learning` | Shared key cho internal API |
| `REDIS_PASSWORD` | Root `.env` | Password Redis cho Docker Compose |

## Lệnh Makefile

```bash
make up        # docker compose up -d
make build     # build image rồi chạy
make rebuild   # down rồi build lại
make down      # dừng container
make restart   # restart container
make ps        # xem container
make logs      # follow logs
make clean     # down -v + docker system prune -f
make fmt       # go fmt toàn bộ service
make tidy      # go mod tidy toàn bộ service
```

Chạy riêng service:

```bash
make identity
make course
make learning
```

## Development

```bash
make fmt      # format code
make tidy     # tidy module