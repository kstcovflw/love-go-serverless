# Makefile

# Define the default AWS profile to use
AWS_PROFILE ?= default
# The runtime target for Lambda
GOOS ?= linux
GOARCH ?= amd64

# Functions directory
FUNCTIONS_DIR := src/functions
# Binary name for Lambda
BINARY_NAME := main

# Find all function directories
FUNCTIONS := $(notdir $(wildcard $(FUNCTIONS_DIR)/*))

# Default target
.PHONY: all
all: build

# Build all functions
.PHONY: build
build:
	@echo "Building functions..."
	@$(foreach func, $(FUNCTIONS), \
		go build -o $(FUNCTIONS_DIR)/$(func)/$(BINARY_NAME) $(FUNCTIONS_DIR)/$(func)/main.go && \
		echo "Built $(func)/$(BINARY_NAME)"; \
	)






# Deploy stack with AWS CDK
.PHONY: deploy
deploy:
	echo "Deploying stack..."
	cdk deploy --profile $(AWS_PROFILE)

# Remove/Delete Lambda functions and associated resources
.PHONY: destroy
destroy:
	echo "Destroying stack..."
	cdk destroy --profile $(AWS_PROFILE) --force

# Clean up binaries
.PHONY: clean
clean:
	echo "Cleaning up..."
	$(foreach func, $(FUNCTIONS), \
		rm -f $(FUNCTIONS_DIR)/$(func)/$(BINARY_NAME); \
	)
