# Online Shopping System - Pissaze

This project focuses on designing and implementing the backend and frontend for **Pissaze**, an online shopping platform. The system includes **user management, shopping cart functionality, referral systems, purchase history tracking, and activity logging**. 

The project is structured into two main parts:
- **Backend**: Developed using **Go (Golang)** with the **Gin framework** and **PostgreSQL** as the database.
- **Frontend**: Developed using **Next.js**, providing a modern UI.

## Database Design

The project was initially developed as part of a **Database Systems course**, with a strong focus on **database design and implementation**. An **Enhanced Entity-Relationship (EER) model** was used in the initial design phase to ensure **clarity, scalability, and optimized database performance**. The database uses **PostgreSQL**, with `pg_cron` for scheduled tasks.

## Docker

Run the full stack (frontend, backend, PostgreSQL) with Docker Compose:

```bash
docker compose up --build
```

- **Frontend**: http://localhost:3000  
- **Backend API**: http://localhost:8082  
- **Swagger**: http://localhost:8082/swagger/index.html  

On first run, the **db-seed** service creates the `pissaze_system` database, applies the schema (`createdDB.sql`) and **Persian test data** (`persian_dataset.sql`) automatically. No manual steps required.

For manual re-seeding (e.g. against an existing Postgres container):

```bash
chmod +x scripts/seed-db.sh
./scripts/seed-db.sh
```

**If the client image fails with `npm error network read ETIMEDOUT`** (slow or restricted network):

- Build the client image using the host network, then start the stack:
  ```bash
  docker build --network=host -f client/Dockerfile -t pissaze-client:latest ./client
  docker compose up -d
  ```
- Or use an npm registry mirror when building:
  ```bash
  docker compose build --build-arg NPM_REGISTRY=https://registry.npmmirror.com
  docker compose up -d
  ```

## Learn More
  - [Gin Documentation](https://gin-gonic.com/) - Learn about Gin framework features.
  - [PostgreSQL Documentation](https://www.postgresql.org/docs/) - Learn about database configuration.
  - [Next.js Documentation](https://nextjs.org/docs) - Learn about Next.js frontend development.
  - [Swagger API Docs](https://swagger.io/) - Learn how to use API documentation.

