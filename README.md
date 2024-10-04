## Prerequisites
- Docker
- Golang

## Setup

1. Clone the repository
2. in project folder to Run the database: docker-compose up -d
3. once container run from docker now make connection to database using program like tableplus, dbeaver, etc ![Uploading image_2024-10-04_155751826.png…]()

4. ![image_2024-10-04_155643931](https://github.com/user-attachments/assets/0096f027-d75b-4d4b-886e-bc345ba4e64c)
5. in project folder to Run the application: go run main.go
6. now you can test the api using program like postman

 ## API Endpoints:
- POST: `http://localhost:8081/animal`
- PUT: `http://localhost:8081/animal/{id}`
- DELETE: `http://localhost:8081/animal/{id}`

- GET: `http://localhost:8081/animals`
- GET: `http://localhost:8081/animal/{id}`

 ## API Documentation:
 Postman Collection : https://documenter.getpostman.com/view/38755939/2sAXxLDEuj
