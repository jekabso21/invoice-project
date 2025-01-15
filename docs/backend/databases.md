## Overview
The backend for the Invoice Project uses two databases to handle different types of data:
1. **PostgreSQL**: To store structured, lightweight data such as user accounts and team information.
2. **MongoDB**: To store detailed, dynamic invoice data and related properties, with a focus on flexibility and scalability.

---

## PostgreSQL

### Purpose
PostgreSQL is used for storing relational, normalized data, ensuring data integrity and reducing redundancy. It adheres to **Second Normal Form (2NF)**, where all non-key attributes are fully functionally dependent on the primary key.

### Use Cases
- User authentication and role management.
- Team information and membership relationships.

### Example Schema (2NF)
The following schema ensures no partial dependencies and eliminates redundancy:

#### Users Table
```sql
CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    role VARCHAR(20) NOT NULL DEFAULT 'user',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

#### Teams Table
```sql
CREATE TABLE teams (
    team_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_by INT REFERENCES users(user_id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

#### Membership Table
To avoid redundancy, team membership is stored in a separate table:
```sql
CREATE TABLE user_teams (
    user_team_id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(user_id) ON DELETE CASCADE,
    team_id INT REFERENCES teams(team_id) ON DELETE CASCADE,
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

#### Advantages of 2NF
- Reduces redundancy by separating related data into smaller tables.
- Improves scalability and maintainability of the database.

---

## MongoDB

### Purpose
MongoDB stores unstructured and dynamic data, such as invoices with nested properties. It provides flexibility for storing complex documents without requiring a predefined schema.

### Use Cases
- Invoices associated with specific teams.
- Dynamic properties like invoice items, totals, and statuses.

### Example Document Structure
#### Invoice Document
Each invoice is stored as a document under a team, with all related information embedded.

```json
{
  "_id": "invoice-id-12345",
  "team_id": "team-id-67890",
  "invoice_number": "INV-2023-0001",
  "created_by": "user-id-112233",
  "created_at": "2025-01-15T08:30:00Z",
  "due_date": "2025-02-15T08:30:00Z",
  "status": "pending",
  "items": [
    {
      "description": "Consulting Services",
      "quantity": 5,
      "unit_price": 200.0,
      "total": 1000.0
    },
    {
      "description": "Hosting Services",
      "quantity": 1,
      "unit_price": 50.0,
      "total": 50.0
    }
  ],
  "total_amount": 1050.0
}
```

### Key Relationships
Invoices are linked to teams via the `team_id` field:
```json
{
  "team_id": "team-id-67890",
  "invoice_number": "INV-2023-0001",
  ...
}
```

### Query Examples
- Retrieve all invoices for a team:
  ```javascript
  db.invoices.find({ team_id: "team-id-67890" });
  ```

- Update the status of an invoice:
  ```javascript
  db.invoices.updateOne(
    { _id: "invoice-id-12345" },
    { $set: { status: "paid" } }
  );
  ```

- Retrieve invoices created by a specific user:
  ```javascript
  db.invoices.find({ created_by: "user-id-112233" });
  ```

---

## Key Considerations

### PostgreSQL
- **Normalization**: Achieved up to 2NF to ensure minimal redundancy and dependency.
- **Data Integrity**: Enforced through foreign key constraints.
- **Scalability**: Efficient for transactional data like user and team information.

### MongoDB
- **Flexibility**: Ideal for nested, dynamic data like invoices.
- **Query Performance**: Index fields like `team_id`, `created_by`, and `status` to improve query performance.
- **Scalability**: Supports horizontal scaling for handling large datasets.