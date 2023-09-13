#PERSON REST API

The PERSON REST API is a simple API that allows you to perform CRUD operations on the "Person" resource. It provides endpoints for creating, retrieving, updating, and deleting persons. This document provides an overview of the API, how to use it, and how to set it up for local development or deployment.

#Table of Contents
1. Endpoints
2. Usage
3. Setup
4. Contributing
5. License
6. Endpoints

The API offers the following endpoints:

##Create a New Person:

Endpoint: POST /api/
Description: Create a new person.
Request Format:

```json
Copy code
{
  "name": "John",
  "email": "john@gmail.com"
}
Response Format:
```json
{
  "id": 1,
  "name": "John",
  "email": "john@gmail.com"
}

##Get Person by ID:

Endpoint: GET /api/{id}
Description: Retrieve details of a person by their ID.
Response Format:
```json

{
  "id": 1,
  "name": "John",
  "email": "john@gmail.com"
}

##Update Person:

Endpoint: PUT /api/{id}
Description: Update details of an existing person.
Request Format:

```json
{
  "first_name": "John"
}

Response Format:
```json

{
  "id": 1,
  "name": "John",
  "email": ""
}

##Delete Person:

Endpoint: DELETE /api/{id}
Description: Remove a person by their ID.
Response Format:
```json
{
  "message": "Person  has been deleted."
}

#Usage
You can use the PERSON REST API to manage persons in your application. Below are some sample usages:

1. Create a New Person:
http
POST /api/
Content-Type: application/json

{
  "id": 1,
  "name": "John",
  "email": "john@gmail.com"
}
2. Get Person by ID:
http
GET /api/1

3. Update Person:

http
PUT /api/1
Content-Type: application/json

{
  "first_name": "Jane"
}
4. Delete Person:

http
DELETE /api/1
For more detailed information on request and response formats, please refer to the API Documentation.

#Setup
Follow these steps to set up and run the API locally:

Clone the repository:

bash

git clone [Link Text] https://github.com/Mac-5/PersonApi.git
Install dependencies:

bash

cd your-api
go mod tidy
Configure the database:

Update the database configuration in your application code if needed. Ensure that the database server is running.

Run the API:

bash
Copy code
go run main.go
The API will be accessible at [Link Text] http://localhost:PORT, where PORT is the port specified in your configuration.

For deployment instructions and additional information, please refer to the API Documentation.

#Contributing
Contributions to this API are welcome! If you have ideas for improvements or new features, please open an issue or submit a pull request.

#License
This project is licensed under the MIT License. Feel free to use and modify the code as needed.