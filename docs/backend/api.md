## Overview
The Invoice Project API is designed for seamless frontend interaction. All endpoints, except for `register` and `login`, are secured behind an **authentication middleware** that validates a 24-hour session ID. The session ID contains critical user data such as team memberships, permissions, and roles, enabling efficient authorization checks for every request.

---

## Authentication Endpoints

### Register
- **Endpoint**: `POST /api/auth/register`
- **Purpose**: Register a new user.
- **Request Body**:
  ```json
  {
    "username": "example_user",
    "email": "user@example.com",
    "password": "securepassword"
  }
  ```
- **Response**:
  - **201**: User created successfully.
  - **400**: Validation error.

---

### Login
- **Endpoint**: `POST /api/auth/login`
- **Purpose**: Log in a user and generate a session ID.
- **Request Body**:
  ```json
  {
    "email": "user@example.com",
    "password": "securepassword"
  }
  ```
- **Response**:
  ```json
  {
    "session_id": "unique-session-id",
    "expires_at": "2025-01-16T12:00:00Z"
  }
  ```
  - **200**: Login successful.
  - **401**: Invalid credentials.

---

## Team Management Endpoints

### Create Team
- **Endpoint**: `POST /api/teams`
- **Purpose**: Create a new team.
- **Authorization**: Requires a valid session.
- **Request Body**:
  ```json
  {
    "team_name": "Team Alpha",
    "owner_id": "user-id-123"
  }
  ```
- **Response**:
  - **201**: Team created successfully.
  - **400**: Validation error.

---

### Invite Someone to Team
- **Endpoint**: `POST /api/teams/:team_id/invite`
- **Purpose**: Send an invite to a user for a specific team.
- **Authorization**: Requires a valid session and team admin permissions.
- **Request Body**:
  ```json
  {
    "email": "invitee@example.com",
    "expiration": "2h" // Duration for invite validity
  }
  ```
- **Response**:
  - **200**: Invite sent successfully.
  - **400**: Validation error.

---

### Accept Invitation
- **Endpoint**: `POST /api/teams/invite/accept`
- **Purpose**: Accept an invitation to join a team.
- **Authorization**: Requires a valid session.
- **Request Body**:
  ```json
  {
    "invite_token": "example-invite-token"
  }
  ```
- **Response**:
  - **200**: Invitation accepted, and user added to the team.
  - **400**: Invalid or expired token.

---

### Remove Someone from Team
- **Endpoint**: `DELETE /api/teams/:team_id/users/:user_id`
- **Purpose**: Remove a user from a team.
- **Authorization**: Requires a valid session and team admin permissions.
- **Response**:
  - **200**: User removed from the team.
  - **403**: Not authorized to perform this action.

---

## Invoice Management Endpoints

### Recover All Invoices
- **Endpoint**: `GET /api/invoices`
- **Purpose**: Retrieve all invoices for teams the user belongs to.
- **Authorization**: Requires a valid session.
- **Response**:
  ```json
  [
    {
      "invoice_id": "invoice-id-123",
      "team_id": "team-id-123",
      "invoice_number": "INV-2023-0001",
      "status": "pending",
      "total_amount": 250.00
    }
  ]
  ```
  - **200**: Success.

---

### Recover Specific Invoice
- **Endpoint**: `GET /api/invoices/:invoice_id`
- **Purpose**: Retrieve a specific invoice by ID.
- **Authorization**: Requires a valid session.
- **Response**:
  ```json
  {
    "invoice_id": "invoice-id-123",
    "team_id": "team-id-123",
    "invoice_number": "INV-2023-0001",
    "items": [
      {
        "description": "Service A",
        "quantity": 2,
        "unit_price": 100.0,
        "total": 200.0
      }
    ],
    "total_amount": 250.00
  }
  ```
  - **200**: Success.
  - **404**: Invoice not found.

---

### Create New Invoice
- **Endpoint**: `POST /api/invoices`
- **Purpose**: Create a new invoice under a team.
- **Authorization**: Requires a valid session and team permissions.
- **Request Body**:
  ```json
  {
    "team_id": "team-id-123",
    "invoice_number": "INV-2023-0002",
    "items": [
      {
        "description": "Service B",
        "quantity": 3,
        "unit_price": 50.0
      }
    ],
    "total_amount": 150.00,
    "due_date": "2025-02-15"
  }
  ```
- **Response**:
  - **201**: Invoice created successfully.
  - **400**: Validation error.

---

### Update Invoice
- **Endpoint**: `PUT /api/invoices/:invoice_id`
- **Purpose**: Update details of an invoice.
- **Authorization**: Requires a valid session and team permissions.
- **Request Body**:
  ```json
  {
    "due_date": "2025-03-01",
    "status": "paid"
  }
  ```
- **Response**:
  - **200**: Invoice updated successfully.
  - **404**: Invoice not found.

---

### Delete Invoice
- **Endpoint**: `DELETE /api/invoices/:invoice_id`
- **Purpose**: Delete an invoice.
- **Authorization**: Requires a valid session and team admin permissions.
- **Response**:
  - **200**: Invoice deleted successfully.
  - **404**: Invoice not found.

---

## Key Points on Session and Authorization
1. **Session Validation**:
   - All endpoints except `register` and `login` validate the session ID.
   - The session ID holds user data such as `team_id`, `permissions`, and `role`.

2. **Authorization Checks**:
   - The backend verifies if the user has appropriate permissions (e.g., admin, member).
   - For actions like creating or deleting invoices, team-specific roles are checked.

3. **Session Expiry**:
   - Sessions are valid for 24 hours. After expiry, users must log in again.
