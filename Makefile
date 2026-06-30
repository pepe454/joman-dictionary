include .env
export


# Database / Postgres stuff

psql-connect:
	psql -h localhost -p 5432 -U postgresadmin -d joman

init-db:
	docker compose up -d postgres

delete-db:
	docker compose down postgres && docker volume prune --all

seed-db:
	cd backend && go run cmd/seed/main.go

reset-db:
	$(MAKE) delete-db && $(MAKE) init-db && $(MAKE) seed-db

copy-from-pgadmin:
	docker cp joman-pgadmin-1:/var/lib/pgadmin/storage/$(user)/* ./backend/db/


# Go stuff

test:
	cd backend && go test ./...

install-go:
	cd $$HOME && \
	curl -LO https://go.dev/dl/go1.26.3.linux-amd64.tar.gz && \
	sudo rm -rf /usr/local/go && \
	sudo tar -C /usr/local -xzf go1.26.3.linux-amd64.tar.gz && \
	export PATH=$$PATH:/usr/local/go/bin && \
	go version

source-env:
	set -o allexport && source .env && set +o allexport

sqlc-generate:
	cd backend && sqlc generate

tidy:
	cd backend && go mod tidy
