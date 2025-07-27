include .env
export $(shell sed 's/=.*//' .env | xargs)

.PHONY: start-postgres
start-postgres:
	docker run --rm -d \
		--name postgres17 \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=password \
		-e POSTGRES_DB=go_blog \
		-p 5432:5432 \
		-v pgdata:/var/lib/postgresql/data \
		postgres:17.5-alpine

stop-postgres:
	docker rm -f postgres17
create-db:
	docker exec -it postgres17 createdb -U $(DB_USER) --owner=$(DB_USER) $(DB_NAME)