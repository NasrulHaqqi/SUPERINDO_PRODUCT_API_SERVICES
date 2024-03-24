# SUPERINDO_PRODUCT_API_SERVICES

This is a Go application for managing products in the SuperIndo store.

## Prerequisites

Before running this application, ensure you have the following:

- Go installed on your system
- PostgreSQL installed and running
- Redis installed and running (optional, for caching)

## Installation

1. Clone this repository to your local machine.
2. Navigate to the project directory.
3. Update the database and Redis configurations in `config/config.go` as per your environment.
4. Install the necessary dependencies by running:
5. Build the application:


## Configuration

You can configure the application settings in `config/config.go`. Update the database and Redis connection details accordingly.

## Database Setup

Ensure that PostgreSQL is running and create a database named `superindo_db`.

You can create the necessary tables by running the SQL migrations provided in the `migrations` directory.

## Running the Application

To start the application, run the following command: go run main.go



The application will start on port 3000 by default. You can access the API endpoints as described below.

## API Endpoints

- `POST /api/products`: Add a new product.
- `GET /api/products`: Get all products.
- `GET /api/products/:id`: Get product by ID.
- `POST /api/products/:name`: Get product by Name.
- `GET /api/products/search`: Search for products.
- `GET /api/products/type/:type`: Filter products by type.
- `GET /api/products/sort/:field`: Sort products by a specific field.

Refer to the API documentation or code comments for detailed information on request and response formats.
