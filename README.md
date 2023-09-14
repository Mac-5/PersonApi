# PERSON REST API

The PERSON REST API is a simple API that allows you to perform CRUD operations on the "Person" resource. It provides endpoints for creating, retrieving, updating, and deleting persons. This document provides an overview of the API, how to use it, and how to set it up for local development or deployment.

## Table of Contents
- [PERSON REST API](#person-rest-api)
  - [Table of Contents](#table-of-contents)
  - [Endpoints](#endpoints)
    - [Create a New Person](#create-a-new-person)
    - [Get All Persons](#get-all-persons)
  - [Setup](#setup)
  - [Contributing](#contributing)
  - [License](#license)

## Endpoints

The API offers the following endpoints:

### Create a New Person

- **Endpoint:** POST  [https://person-crud-api-bdzs.onrender.com/api](https://person-crud-api-bdzs.onrender.com/api)
- **Description:** Create a new person.
- **Request Format:**


```json
{
  "name": "John",
  "email": "john@gmail.com"
}
```

- **Response Format:**
  
```json
{
  "id": 1,
  "name": "John",
  "email": "john@gmail.com"
}
```
### Get All Persons

- **Endpoint:** GET  [https://person-crud-api-bdzs.onrender.com/api](https://person-crud-api-bdzs.onrender.com/api)
- **Description:** Create a new person.
- **Request Format:**


```json
{
  "name": "John",
  "email": "john@gmail.com"
}
```

- **Response Format:**
  
```json
{
  "id": 1,
  "name": "John",
  "email": "john@gmail.com"
}
```
**Get Person by ID**
-**Endpoint:** GET  [https://person-crud-api-bdzs.onrender.com/api/{id}](https://person-crud-api-bdzs.onrender.com/api/{id})
-**Description**: Retrieve details of a person by their ID.
-**Response Format**
```json
{
  "id": 1,
  "name": "John",
  "email": "john@gmail.com"
}
```
**Update Person**
-**Endpoint**: PUT  [https://person-crud-api-bdzs.onrender.com/api/{id}](https://person-crud-api-bdzs.onrender.com/api/{id})
-**Description**: Update details of an existing person.
-**Request Format**:
```json
{
  "first_name": "John"
}
```
**Response Format**
```json
{
  "id": 1,
  "name": "John",
  "email": ""
}
```
**Delete Person**
**Endpoint:** DELETE  [https://person-crud-api-bdzs.onrender.com/api/{id}](https://person-crud-api-bdzs.onrender.com/api/{id})
Description:** Remove a person by their ID.
**Response Format:**

```json
{
  "message": "Person has been deleted."
}
```

**Usage**
You can use the PERSON REST API to manage persons in your application. Below are some sample usages:

*Create a New Person:*

POST  [https://person-crud-api-bdzs.onrender.com/api](https://person-crud-api-bdzs.onrender.com/api)
Content-Type: application/json

```json
{
  "id": 1,
  "name": "John",
  "email": "john@gmail.com"
}
```

*Get all Persons:*
GET  [https://person-crud-api-bdzs.onrender.com/api](https://person-crud-api-bdzs.onrender.com/api)
*Get Person by ID:*
GET  [https://person-crud-api-bdzs.onrender.com/api/{id}](https://person-crud-api-bdzs.onrender.com/api/{id})

*Update Person:*
PUT  [https://person-crud-api-bdzs.onrender.com/api/{id}](https://person-crud-api-bdzs.onrender.com/api/{id})
Content-Type: application/json
```json

{
  "first_name": "Jane"
}

```
*Delete Person:*
DELETE  [https://person-crud-api-bdzs.onrender.com/api/{id}](https://person-crud-api-bdzs.onrender.com/api/{id})

For more detailed information on request and response formats, please refer to the API Documentation.

## Setup
Follow these steps to set up and run the API locally:

**Clone the repository:**
git clone [https://github.com/Mac-5/PersonApi.git
](https://)

**Install dependencies:**
> cd your-api
go mod tidy

**Configure the database:**
Update the database configuration in your application code if needed. Ensure that the database server is running.

**Run the API:**
> go run main.go

The API will be accessible at http://localhost:PORT, where PORT is the port specified in your configuration.

For deployment instructions and additional information, please refer to the API Documentation.

## Contributing
Contributions to this API are welcome! If you have ideas for improvements or new features, please open an issue or submit a pull request.

## License
This project is licensed under the MIT License. Feel free to use and modify the code as needed.




