# Uyir Backend API Documentation

## Base URL

```
http://localhost:6969
```

---

## Routes

### 1. User Authentication & Info

#### **POST /signup**

- **Description:** Register a new user.
- **Request Body (JSON):**
  ```json
  {
    "firstName": "string",
    "lastName": "string",
    "username": "string",
    "email": "string",
    "password": "string",
    "role": "string",
    "vehicleType": "string",
    "fuelType": "string",
    "vehicleNumber": "string"
  }
  ```

#### **POST /login**

- **Description:** Login user.
- **Request Body (JSON):**
  ```json
  {
    "username": "string",
    "password": "string"
  }
  ```

#### **GET /user**

- **Description:** Get current user's reports and points.
- **Request:** Cookie with `session_token` required.

#### **GET /me**

- **Description:** Get current user session info.
- **Request:** Cookie with `session_token` required.

---

### 2. Reports

#### **POST /new**

- **Description:** Submit a new report.
- **Request Type:** `multipart/form-data`
- **Request Body:**
  - `latitude`: float
  - `longitude`: float
  - `location`: string
  - `type`: string
  - `file`: image file

#### **GET /reports**

- **Description:** Get all reports.

#### **GET /reports/pending/**

- **Description:** Get all pending reports.

#### **GET /reports/hospitals**

- **Description:** Get reports relevant for hospitals.

#### **GET /reports/police**

- **Description:** Get reports relevant for police.

#### **GET /reports/pwd**

- **Description:** Get reports relevant for PWD (Public Works Department).

#### **POST /reports/updateStatus**

- **Description:** Update the status of a report (e.g., mark as resolved).
- **Request Body (JSON):**
  ```json
  {
    "id": "string" // Report ID
  }
  ```

---

### 3. AI & Similarity

#### **POST /similarReports**

- **Description:** Get similar reports based on location.
- **Request Body (JSON):**
  ```json
  {
    "latitude": float,
    "longitude": float
  }
  ```

---

### 4. Miscellaneous

#### **GET /**

- **Description:** Health check or root endpoint.

---

## Notes

- Most endpoints that modify or access user-specific data require a valid `session_token` cookie.
- All POST endpoints expect `Content-Type: application/json` unless otherwise specified.
- File uploads use `multipart/form-data`.

---
