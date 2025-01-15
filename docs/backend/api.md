# API Framework

## Overview
The API follows a **RESTful** architecture, enabling efficient communication between the backend and frontend. All API responses are in JSON format.

## Authentication
- **Method**: JSON Web Tokens (JWT).
- **Token Expiry**: Adjustable based on user role (e.g., 24 hours for admins, 12 hours for regular users).

## Documentation
The APIs are documented using **Swagger** to ensure easy integration and clear endpoint descriptions.

## Key Endpoints
### User Authentication
- **POST /api/auth/login**
  - Description: Logs in a user and returns a JWT token.
  - Payload:
    ```json
    {
      "email": "user@example.com",
      "password": "password123"
    }
    ```
  - Response:
    ```json
    {
      "token": "your-jwt-token"
    }
    ```

### Invoice Management
- **GET /api/invoices**
  - Description: Retrieves a list of all invoices.
- **POST /api/invoices**
  - Description: Creates a new invoice.
  - Payload:
    ```json
    {
      "customer_name": "John Doe",
      "total_amount": 150.50
    }
    ```

