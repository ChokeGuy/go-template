# Go Template

## Prerequisites

Before getting started, ensure you have the following installed and configured:

- **Go Language**: Make sure Go is installed on your system. You can download it from [golang.org](https://golang.org/).
- **Linux or WSL**: Use a Linux environment or Windows Subsystem for Linux (WSL) for seamless development.

## Getting Started

Follow the steps below to set up and run the project:

### 1. Initialize Kafka using Docker Compose

Start the Kafka services using the provided `docker-compose` file:

```bash
docker compose -f kafka.yml up
```

This command will start the Kafka and Zookeeper containers required for the application.

### 2. Start the Server

Use the `make` command to build and run the server:

```bash
make server
```

This will start the Go server and handle any dependencies or build requirements automatically.

## Project Structure

The repository is structured as follows:

```
.
├── cmd/                # Entry point for the application
├── pkg/                # Package modules
├── internal/           # Internal application logic
├── configs/            # Configuration files
├── kafka.yml           # Docker Compose file for Kafka
├── Makefile            # Makefile with useful commands
├── README.md           # Project documentation
└── ...
```

## Usage

Once the server is running, you can interact with it through the specified API endpoints. Ensure that Kafka is up and running before starting the server to avoid connectivity issues.
