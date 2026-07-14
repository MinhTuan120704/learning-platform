# ==========================================
# Learning Platform - Makefile
# ==========================================

COMPOSE = docker compose

IDENTITY_DB = postgres://postgres:postgres@localhost:5432/identity_db?sslmode=disable
COURSE_DB = postgres://postgres:postgres@localhost:5433/course_db?sslmode=disable
LEARNING_DB = postgres://postgres:postgres@localhost:5434/learning_db?sslmode=disable

.PHONY: \
	up down restart build rebuild ps logs clean \
	identity course learning \
	migrate-up-identity migrate-down-identity migrate-version-identity \
	migrate-up-course migrate-down-course migrate-version-course \
	migrate-up-learning migrate-down-learning migrate-version-learning \
	create-migration \
	fmt tidy

# ==========================================
# Docker
# ==========================================

up:
	$(COMPOSE) up -d

build:
	$(COMPOSE) up --build -d

rebuild:
	$(COMPOSE) down
	$(COMPOSE) up --build -d

down:
	$(COMPOSE) down

restart:
	$(COMPOSE) restart

ps:
	$(COMPOSE) ps

logs:
	$(COMPOSE) logs -f

clean:
	$(COMPOSE) down -v
	docker system prune -f

# ==========================================
# Run single service
# ==========================================

identity:
	$(COMPOSE) up identity

course:
	$(COMPOSE) up course

learning:
	$(COMPOSE) up learning

# ==========================================
# Identity Migration
# ==========================================

migrate-up-identity:
	migrate -path services/identity/migrations -database "$(IDENTITY_DB)" up

migrate-down-identity:
	migrate -path services/identity/migrations -database "$(IDENTITY_DB)" down 1

migrate-version-identity:
	migrate -path services/identity/migrations -database "$(IDENTITY_DB)" version

# ==========================================
# Course Migration
# ==========================================

migrate-up-course:
	migrate -path services/course/migrations -database "$(COURSE_DB)" up

migrate-down-course:
	migrate -path services/course/migrations -database "$(COURSE_DB)" down 1

migrate-version-course:
	migrate -path services/course/migrations -database "$(COURSE_DB)" version

# ==========================================
# Learning Migration
# ==========================================

migrate-up-learning:
	migrate -path services/learning/migrations -database "$(LEARNING_DB)" up

migrate-down-learning:
	migrate -path services/learning/migrations -database "$(LEARNING_DB)" down 1

migrate-version-learning:
	migrate -path services/learning/migrations -database "$(LEARNING_DB)" version

# ==========================================
# Create migration
#
# Example:
# make create-migration DIR=services/course/migrations NAME=create_course_table
# ==========================================

create-migration:
	migrate create -ext sql -dir $(DIR) -seq $(NAME)

# ==========================================
# Go
# ==========================================

fmt:
	cd services/identity && go fmt ./...
	cd services/course && go fmt ./...
	cd services/learning && go fmt ./...

tidy:
	cd services/identity && go mod tidy
	cd services/course && go mod tidy
	cd services/learning && go mod tidy
