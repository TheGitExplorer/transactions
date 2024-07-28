# transactions REST API Example

This is an example of a RESTful web service implemented in Go using the Gin framework. The service manages transactions and allows users to add, retrieve, and query transactions by type or sum up linked transactions.

## Features

- **Add Transaction**: Add a new transaction with a specified ID, amount, type, and optional parent ID.
- **Get Transaction**: Retrieve details of a transaction by its ID.
- **Get Transactions by Type**: List all transaction IDs that match a specific type.
- **Get Sum of Transactions**: Calculate the total amount for a transaction and all its linked (child) transactions.

## Project Structure
transactions/
│
├── main.go
├── config/
│ └── db.go
├── routes/
│ └── routes.go
├── service/
│ └── service.go
├── entity/
  └── Entity.go

####################

- **main.go**: Entry point of the application.
- **config/db.go**: Handles database connection setup.
- **routes/transaction_routes.go**: Defines the API endpoints and their handlers.
- **services/transaction_service.go**: Contains business logic.
- **entities/transaction.go**: Data models and database interaction methods.

## Getting Started

### Prerequisites

- Go 1.16 or later
- MySQL server

### Setup

1. **Clone the repository**:
    ```bash
    git clone https://github.com/TheGitExploere/transactions.git
    cd transactions
    ```

2. **Initialize Go modules**:
    ```bash
    go mod tidy
    ```

3. **Configure the database**:
    - Update the MySQL connection string in `config/db.go` with your credentials.
    - Create the necessary tables in your MySQL database:

    ```sql
    CREATE DATABASE transaction_service;
    USE transaction_service;

    CREATE TABLE transactions (
        id BIGINT PRIMARY KEY,
        amount DOUBLE,
        type VARCHAR(255),
        parent_id BIGINT
    );
    ```

4. **Run the application**:
    ```bash
    go run main.go
    ```

5. **API Endpoints**:
    - **Add Transaction**: `PUT /transactionservice/transaction/:id`
    - **Get Transaction**: `GET /transactionservice/transaction/:id`
    - **Get Transactions by Type**: `GET /transactionservice/types/:type`
    - **Get Sum of Transactions**: `GET /transactionservice/sum/:id`

## Testing

Use tools like `curl` or Postman to interact with the API endpoints. For example, to add a new transaction:

```bash
curl -X PUT -H "Content-Type: application/json" -d '{"amount": 5000, "type":"cars"}' http://localhost:8080/transactionservice/transaction/10


