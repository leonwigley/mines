# Simple Makefile for a Go project

# Build the application
all: build test

tailwind-install:
	@mkdir -p bin/tools
	@if [ ! -f bin/tools/tailwindcss ]; then curl -sL https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 -o bin/tools/tailwindcss; fi
	
	@chmod +x bin/tools/tailwindcss

build: tailwind-install
	@echo "Building..."
	@bin/tools/tailwindcss -i app/views/css/input.css -o public/css/style.css
	@go build -o bin/mines bin/mines.go

# Run the application
run:
	@go run bin/mines.go

# Run the binary
preview:
	@go run bin/mines

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f bin/mines

# Live Reload
dev: tailwind-install
	@echo "Watching..."
	@if command -v air > /dev/null; then \
		bin/tools/tailwindcss -i app/views/css/input.css -o public/css/style.css --watch & \
		air; \
	else \
		read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/air-verse/air@latest; \
			bin/tools/tailwindcss -i app/views/css/input.css -o public/css/style.css --watch & \
			air; \
		else \
			echo "You chose not to install air. Exiting..."; \
			exit 1; \
		fi; \
	fi