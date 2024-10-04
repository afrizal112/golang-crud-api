## Prerequisites
- Docker
- Golang

## Setup

1. Clone the repository
2. in project folder to run Run the database: docker-compose up -d
3. in project folder to Run the application: go run main.go

 ## API Endpoints:
- POST: `http://localhost:8081/animal`
- PUT: `http://localhost:8081/animal/{id}`
- DELETE: `http://localhost:8081/animal/{id}`
- GET: `http://localhost:8081/animals`
- GET: `http://localhost:8081/animal/{id}`
