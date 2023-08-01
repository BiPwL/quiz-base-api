#golang
.PHONY: build
build:
	go build -v ./cmd/apiserver;

.PHONY: test
test:
	go test -v -cover -race -timeout 60s ./...;
.DEFAULT_GOAL := build

.PHONY: server
server:
	go run main.go
#********************************************

#docker
.PHONY: check_docker_installation
check_docker_installation:
	@OS="$$(uname -s)"; \
	if [ "$$OS" = "Linux" ]; then \
		DOCKER_COMMAND="docker"; \
	elif [ "$$OS" = "Darwin" ]; then \
		DOCKER_COMMAND="docker"; \
	elif [ "$$OS" = "FreeBSD" ]; then \
		DOCKER_COMMAND="docker"; \
	else \
		DOCKER_COMMAND="docker.exe"; \
	fi; \
	if ! command -v "$$DOCKER_COMMAND" >/dev/null; then \
		echo "Docker is not installed. Installing Docker..."; \
		$(MAKE) download_docker_internal; \
	fi;

.PHONY: download_docker_internal
download_docker_internal:
	@OS="$$(uname -s)"; \
	if [ "$$OS" = "Linux" ]; then \
		echo "Installing Docker for Linux..."; \
		curl -fsSL https://get.docker.com -o get-docker.sh; \
		sudo sh get-docker.sh; \
		sudo usermod -aG docker $$(whoami); \
		rm get-docker.sh; \
		echo "Docker installation completed."; \
		echo "Please log out and log back in to use Docker without sudo."; \
	elif [ "$$OS" = "Darwin" ]; then \
		echo "Installing Docker for macOS..."; \
		curl -fsSL https://get.docker.com -o get-docker.sh; \
		sh get-docker.sh; \
		rm get-docker.sh; \
		echo "Docker installation completed."; \
	elif [ "$$OS" = "FreeBSD" ]; then \
		echo "Installing Docker for FreeBSD..."; \
		pkg install -y docker; \
		echo 'docker_enable="YES"' >> /etc/rc.conf; \
		service docker start; \
		echo "Docker installation completed."; \
	else \
		echo "Please download Docker for your OS from the official website and follow the installation instructions:"; \
		echo "https://www.docker.com"; \
	fi;

.PHONY: download_docker
download_docker: check_docker_installation
	@echo "Docker is already installed.";

.PHONY: download_psql_image
download_psql_image:
	docker pull postgres;

.PHONY: create_container
create_container:
	docker run --name db -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:latest;

.PHONY: start_container
start_container:
	docker start db;

.PHONY: exec_container
exec_container:
	docker exec -it db psql -U root;

.PHONY: create_db
create_db:
	docker exec -it db createdb --username=root --owner=root quiz_base;

.PHONY: drop_db
drop_db:
	docker exec -it db dropdb quiz_base;

.PHONY: exec_db
exec_db:
	docker exec -it db psql quiz_base;

.PHONY: migrate_up
migrate_up:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/quiz_base?sslmode=disable" -verbose up;

.PHONY: migrate_down
migrate_down:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/quiz_base?sslmode=disable" -verbose down;

.PHONY: sqlc
sqlc:
	sqlc generate

#********************************************

#commands for quick project deployment
.PHONY: prepare_container_with_db
prepare_container_with_db:
	$(MAKE) download_docker; \
	$(MAKE) download_psql_image; \
	$(MAKE) create_container; \
	$(MAKE) exec_container;