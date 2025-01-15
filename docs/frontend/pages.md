# Pages Documentation

## Overview
The frontend for the Invoice Project is organized into modular pages, each designed for a specific feature or functionality. This document outlines the purpose, structure, and key features of each page.

---

## Authentication Pages

### **Login Page**
- **Path**: `/auth/login`
- **Purpose**: Allow users to log in and initiate a session.
- **Features**:
  - Form for entering email and password.
  - Validation for required fields.
  - Error handling for incorrect credentials.
  - Redirects authenticated users to the dashboard.
- **API Integration**:
  - `POST /api/auth/login` for authentication.

### **Register Page**
- **Path**: `/auth/register`
- **Purpose**: Allow new users to create an account.
- **Features**:
  - Form for entering username, email, and password.
  - Validation for required fields and password strength.
  - Error handling for duplicate email or username.
  - Redirects users to login after successful registration.
- **API Integration**:
  - `POST /api/auth/register` for user creation.

---

### **Dashboard Pages**

- **Path**: `/dashboard`
- **Purpose**: Provide an interactive and statistical overview of the user’s invoices and financial data.
- **Features**:
  1. **Statistics and Graphs**:
     - A visual representation of invoices categorized by status:
       - **Paid Invoices**: Count and percentage.
       - **Unpaid Invoices**: Count and percentage.
     - Graph showing:
       - Total income over time (e.g., by month or week).
       - Comparison of paid vs. unpaid invoices.

  2. **Recently Created Invoices**:
     - A scrollable list displayed as cards.
     - Each card includes:
       - Invoice number.
       - Creation date.
       - Total amount.
       - Invoice status (e.g., pending, paid).
     - Cards are pressable:
       - Clicking a card opens the full invoice details in a modal or redirects to the **Invoice Details** page.

  3. **Recently Paid Invoices**:
     - Similar to recently created invoices:
       - List displayed as cards with key details (e.g., invoice number, payment date, total amount).
       - Cards are pressable to view full details.

- **API Integration**:
  - `GET /api/invoices`:
    - Fetches all invoices.
    - Data is filtered and grouped to calculate paid/unpaid statistics, total income, and generate recent lists.

- **Design and Layout**:
  - **Header**: Displays a quick summary (e.g., "Total Income: $XXX").
  - **Graphs**: Positioned prominently at the top of the page.
  - **Lists**: Recently created and recently paid invoices are shown below the graphs in separate sections.

---

## Team Management Pages

### **Teams List**
- **Path**: `/teams`
- **Purpose**: Display all teams the user belongs to and allow the creation of new teams.
- **Features**:
  - List of user’s teams with role information.
  - Button to create a new team.
  - Link to manage individual teams.
- **API Integration**:
  - `GET /api/teams` to fetch team data.
  - `POST /api/teams` to create a new team.

### **Team Details**
- **Path**: `/teams/:team_id`
- **Purpose**: Manage team details and members.
- **Features**:
  - View team name and member list with roles.
  - Invite new members by email.
  - Remove existing members (admin-only).
- **API Integration**:
  - `GET /api/teams/:team_id` to fetch team details.
  - `POST /api/teams/:team_id/invite` to invite a member.
  - `DELETE /api/teams/:team_id/users/:user_id` to remove a member.

---

## Invoice Management Pages

### **Invoices List**
- **Path**: `/invoices`
- **Purpose**: Display all invoices for the user’s teams.
- **Features**:
  - List of invoices with status, due date, and total amount.
  - Links to view or edit specific invoices.
  - Button to create a new invoice.
- **API Integration**:
  - `GET /api/invoices` to fetch invoices.

### **Invoice Details**
- **Path**: `/invoices/:invoice_id`
- **Purpose**: View and edit details of a specific invoice.
- **Features**:
  - Display invoice items, status, and totals.
  - Form for updating invoice details (e.g., status, due date).
  - Button to delete the invoice (admin-only).
- **API Integration**:
  - `GET /api/invoices/:invoice_id` to fetch invoice details.
  - `PUT /api/invoices/:invoice_id` to update the invoice.
  - `DELETE /api/invoices/:invoice_id` to delete the invoice.

### **New Invoice**
- **Path**: `/invoices/new`
- **Purpose**: Create a new invoice for a team.
- **Features**:
  - Form for entering invoice details (team, items, totals, due date).
  - Validation for required fields.
  - Redirect to the invoice list upon successful creation.
- **API Integration**:
  - `POST /api/invoices` to create a new invoice.

---

## Error Pages

### **404 - Not Found**
- **Path**: `/404`
- **Purpose**: Displayed when a user tries to access a non-existent page.
- **Features**:
  - User-friendly message indicating the page could not be found.
  - Button to navigate back to the dashboard.

### **500 - Server Error**
- **Path**: `/500`
- **Purpose**: Displayed when an internal server error occurs.
- **Features**:
  - User-friendly message indicating a problem with the server.
  - Option to retry the action or contact support.

---

## Protected Routes
All pages except `login` and `register` are protected and require a valid session to access. Unauthorized users are redirected to the login page.

---

## Page Navigation
- Navigation is managed through a global header that includes links to:
  - Dashboard
  - Teams
  - Invoices
  - User Profile (optional)

- Conditional rendering is used to show/hide certain links based on user roles and permissions.

---

Let me know if there’s anything more you’d like to add or adjust!
