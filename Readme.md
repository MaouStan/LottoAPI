# How to use this workspace
## 0. Setup .env file
  ```bash
    echo DB_HOST= > .env
    echo DB_USER= >> .env
    echo DB_PASSWORD= >> .env
    echo DB_NAME= >> .env
    echo DB_PORT= >> .env
    echo JWT_SECRET= >> .env
  ```


## 1. RUN Progres SQL server
  ```bash
    docker-compose up --build
  ```

## 2. RUN Go server
  ```bash
    go run cmd/api/main.go
  ```


# Progres SQL Management
## using [pgAdmin4](https://www.pgadmin.org/download/)
- 1. Open pgAdmin4
- 2. Register >> Server
  - Tab: General
    - Name: LottoDB
  - Tab: Connection
    - Host name/address: localhost
    - Port: 5432
    - Maintenance database: postgres
    - Username: postgres
    - Password: lotto2347
- 3. Save
- 4. Connect
- 5. Open Databases >> lottery_aem >> Schemas >> Tables

## pgAdmin4 Example
- 1. Create Table
  - Right click Tables >> Create >> Table
  - Name: users
  - Columns: id, username, password, created_at, updated_at
  - Save
  - Right click Tables >> Refresh
  - Right click users >> View/Edit Data >> All Rows
  - Add data
  - Save
- 2. Query
  - Right click Tables >> Query Tool
  - Query: SELECT * FROM users;
  - Execute
  - Result
