# Build the application
all: build test

build:
	@echo "Building..."
	@go build -o bin/mines main.go

# Run the application
run:
	@go run main.go

# Run the binary
preview:
	@go run bin/mines

# Clean the binary
clean:
	@echo "Cleaning..."
	@rm -f bin/mines

# Live Reload
dev:
	@echo "Watching..."
	@if command -v air > /dev/null; then \
		air; \
	else \
		read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
		if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
			go install github.com/air-verse/air@latest; \
			air; \
		else \
			echo "You chose not to install air. Exiting..."; \
			exit 1; \
		fi; \
	fi
