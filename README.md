# Patient Management System (Makerble)

This is a simple Patient Management System built with Go (Golang) for the backend and PostgreSQL as the database. It allows users to register and manage patient details.
The system uses JWT-based authentication to ensure secure access.

## Features

- Register new patients
- JWT Authentication for authorized access
- Handle patient data (name, age, gender, contact)
- Separate roles for doctors and receptionists

  ## UI Link : https://golangpatient.netlify.app/
  ## Video Link : https://drive.google.com/file/d/1ctLRx_6jnTVYmabeP6cMxCQsX_O7zQd1/view?usp=sharing

## Prerequisites

Before you begin, ensure that you have the following installed:

- [Go (Golang)](https://golang.org/dl/) - Version 1.18 or higher
- [PostgreSQL](https://www.postgresql.org/download/) - For the database
- [Git](https://git-scm.com/) - For version control

## Setup Instructions

### 1. Clone the Repository

Clone this repository to your local machine using Git:

bash
git clone https://github.com/your-username/makerble.git


Install Dependencies

Navigate to the project directory and install the required Go dependencies:

cd makerble

go mod tidy

Setup PostgreSQL Database

Install PostgreSQL and create a new database for the project.

psql -U postgres

CREATE DATABASE myapp_db;

Create the users table and other required tables by running the SQL scripts provided in the project.

4. Configure Database Connection

Open the main.go file and update the database connection details (if needed), such as database name, username, and password:

Example of a DB connection string
connStr := "user=postgres password=yourpassword dbname=myapp_db sslmode=disable"

5. Run the Go Server

Start the Go server with the following command:

go run main.go

This will start the server on http://localhost:8080.
