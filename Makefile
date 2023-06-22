#golang
.PHONY: build
build:
	go build -v ./cmd/apiserver

.PHONY: test
test:
	go test -v -race -timeout 30s ./...
.DEFAULT_GOAL := build

#********************************************

#docker
.PHONY: check_docker_installation
check_docker_installation:
	@if [ "$$(uname -s)" = "Linux" ]; then \
		DOCKER_COMMAND="docker"; \
	elif [ "$$(uname -s)" = "Darwin" ]; then \
		DOCKER_COMMAND="docker"; \
	else \
		DOCKER_COMMAND="docker.exe"; \
	fi; \
	if ! command -v "$$DOCKER_COMMAND" >/dev/null; then \
		echo "Docker is not installed. Installing Docker..."; \
		$(MAKE) install_docker; \
	else \
		echo "Docker is already installed."; \
	fi

.PHONY: download_docker
download_docker:
	OS := $(shell uname -s)
	ifeq ($(OS),Darwin)
		# Installing Docker on macOS
		@echo "Installing Docker for macOS..."
		@curl -fsSL https://get.docker.com -o get-docker.sh
		@sh get-docker.sh
		@rm get-docker.sh
	endif
	ifeq ($(OS),Linux)
		# Installing Docker on Linux
		@echo "Installing Docker for Linux..."
		@curl -fsSL https://get.docker.com -o get-docker.sh
		@sudo sh get-docker.sh
		@sudo usermod -aG docker $(USER)
		@rm get-docker.sh
		@echo "Please log out and log back in to use Docker without sudo."
	endif
	ifeq ($(OS),Windows_NT)
		# Installing Docker on Windows
		@echo "Please download Docker Desktop for Windows from the official website and follow the installation instructions:"
		@echo "https://www.docker.com/products/docker-desktop"
	endif

.PHONY: download_psql_image
download_psql_image:
	docker pull postgres

.PHONY: create_container
create_container:
	docker run --name db -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -p 5432:5432 -d postgres:latest

.PHONY: exec_container
exec_container:
	docker exec -it db psql -U root

#********************************************

#commands for quick project deployment

.PHONY: prepare_container_with_db
prepare_container_with_db:
