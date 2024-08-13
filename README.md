# golang-prometheus-task-async
Go application that simulates CPU and memory-intensive tasks and exposes prometheus metrics on the number of tasks processed and HTTP requests received.

## Features

- **Simulate CPU and Memory Tasks:** The application simulates CPU-intensive and memory-intensive tasks based on the input provided via a REST API.
  - This will be useful for Horizontal Pod Autoscaler (HPA) blog post that I will do in the future.
- **Prometheus Metrics:** Exposes Prometheus metrics for tracking the number of tasks processed and HTTP requests.
- **Dockerized Deployment:** Easily deploy the application using Docker and Docker Compose.

## Getting Started

### Prerequisites

- Docker
- Docker Compose

### Building and Running the Application

1. **Clone the repository:**

   ```bash
   git clone https://github.com/ruanbekker/golang-prometheus-task-async.git
   cd golang-prometheus-task-async
   ```

2. **Build and run the application using Docker Compose:**

   ```bash
   make up
   ```

   This will build the Docker image and start the application along with Prometheus and Grafana.

3. **Access the application:**

   - **MyApp:** http://localhost:8080
   - **Prometheus:** http://localhost:9090
   - **Grafana:** http://localhost:3000

### API Endpoints

- **`GET /`**: Returns the hostname of the running container.
- **`POST /task`**: Processes a task. The request body should contain a JSON object with the following structure:

  ```json
  {
    "type": "cpu" | "memory",
    "duration": <int>
  }
  ```

- **`GET /metrics`**: Exposes Prometheus metrics for monitoring.

### Makefile Commands

- **`make build`**: Builds the Docker image.
- **`make up`**: Builds and runs the containers in detached mode.
- **`make clean`**: Stops and removes all containers.
- **`make logs`**: View the logs from the containers.

### Monitoring with Prometheus and Grafana

Prometheus and Grafana are included in the Docker Compose setup. Grafana is configured with default settings to allow anonymous access with admin privileges for easy access.

### Project Structure

- **Dockerfile**: Multi-stage Dockerfile for building and running the Go application.
- **docker-compose.yaml**: Docker Compose configuration to run MyApp, Prometheus, and Grafana.
- **Makefile**: Helper commands for building, running, and managing the Docker containers.
- **main.go**: The main Go application code that handles HTTP requests and processes tasks.

### Resources

- [Prometheus](https://prometheus.io/)
- [Grafana](https://grafana.com/)
- [Docker](https://www.docker.com/)
- [Go](https://golang.org/)


