## Order Management Microservices

A microservices-based order management system built with Go, demonstrating the use of gRPC for inter-service communication and HTTP for client-server interaction.

### Features

- Microservices: Order Service and Kitchen Service
- gRPC for inter-service communication
- HTTP endpoints for client-server interaction
- Shared protocol buffer definitions
- Simple HTML templating for kitchen display

### Project Structure

- `proto/`: Protocol buffer definition files
- `services/`
  - `common/`: Shared protocol buffer definitions and generated code
  - `kitchen/`: Kitchen service implementation
  - `order/`: Order service implementation

### Getting Started

1. Clone this repository:
    ```sh
    git clone https://github.com/neerrrajj/go-grpc.git
    cd go-grpc
    ```

2. Install the dependencies:
    ```sh
    go mod tidy
    ```

3. Generate protocol buffer code:
    ```sh
    make gen
    ```

4. Run the Order Service:
    ```sh
    make run-orders
    ```

5. In a new terminal, run the Kitchen Service:
    ```sh
    make run-kitchen
    ```

## Usage

- Access the kitchen display at `http://localhost:2000`
- Use the provided HTTP endpoints to create orders
- Orders will be displayed on the kitchen screen