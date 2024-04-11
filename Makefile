init:
	docker-compose up -d --build --force-recreate
	docker exec -it postgres_banner_db createdb --username=postgres --owner=postgres banner_db
	docker exec -it avito_banner ./migrate -path db/migration -database "postgres://postgres:postgres@db:5432/banner_db?sslmode=disable" -verbose up

migrateup:
	docker exec -it avito_banner ./migrate -path db/migration -database "postgres://postgres:postgres@db:5432/banner_db?sslmode=disable" -verbose up

migratedown:
	docker exec -it avito_banner ./migrate -path db/migration -database "postgres://postgres:postgres@db:5432/banner_db?sslmode=disable" -verbose down

run: 
	docker-compose up --build

down: 
	docker-compose down