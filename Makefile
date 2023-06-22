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
	elif [ "$$(uname -s)" = "FreeBSD" ]; then \
    		DOCKER_COMMAND="docker"; \
	else \
		DOCKER_COMMAND="docker.exe"; \
	fi; \
	if ! command -v "$$DOCKER_COMMAND" >/dev/null; then \
		echo "Docker is not installed. Installing Docker..."; \
		echo "aaaaaaaaaa"; \
		uname -s; \
		$(MAKE) download_docker_internal; \
	else \
		echo "Docker is already installed."; \
	fi

.PHONY: download_docker_internal
download_docker_internal:
	@OS="$$(uname -s)"; \
	if [ "$$OS" = "Darwin" ]; then \
		echo "Installing Docker for macOS..."; \
		curl -fsSL https://get.docker.com -o get-docker.sh; \
		sh get-docker.sh; \
		rm get-docker.sh; \
		echo "Docker installation completed."; \
	elif [ "$$OS" = "Linux" ]; then \
		echo "Installing Docker for Linux..."; \
		curl -fsSL https://get.docker.com -o get-docker.sh; \
		sudo sh get-docker.sh; \
		sudo usermod -aG docker $$(whoami); \
		rm get-docker.sh; \
		echo "Docker installation completed."; \
		echo "Please log out and log back in to use Docker without sudo."; \
	elif [ "$$OS" = "FreeBSD" ]; then \
		echo "Installing Docker for FreeBSD..."; \
		pkg install -y docker; \
		echo 'docker_enable="YES"' >> /etc/rc.conf; \
		service docker start; \
		echo "Docker installation completed."; \
	else \
		echo "Please download Docker Desktop for Windows from the official website and follow the installation instructions:"; \
		echo "https://www.docker.com/products/docker-desktop"; \
	fi

.PHONY: download_docker
download_docker: check_docker_installation download_docker_internal
	@echo "Downloading Docker..."




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
