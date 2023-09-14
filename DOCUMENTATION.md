# PERSON REST API

The PERSON REST API is a simple API that allows you to perform CRUD operations on the "Person" resource. It provides endpoints for creating, retrieving, updating, and deleting persons. This document provides an overview of the API, how to use it, and how to set it up for local development or deployment.

## Table of Contents
- [PERSON REST API](#person-rest-api)
  - [Table of Contents](#table-of-contents)
  - [Endpoints](#endpoints)
    - [Create a New Person](#create-a-new-person)
  - [Standard Formats](#standard-formats)
  - [Known Limitations](#known-limitations)
  - [Setup](#setup)
  - [Testing:](#testing)
  - [Deployment:](#deployment)


## Endpoints

The API offers the following endpoints:

### Create a New Person

- **Endpoint:** POST /api/
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
-**Endpoint:** GET /api/{id}
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
-**Endpoint**: PUT /api/{id}
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
**Endpoint:** DELETE /api/{id}**
Description:** Remove a person by their ID.
**Response Format:**

```json
{
  "message": "Person has been deleted."
}
```

## Standard Formats
Request Format: All request data should be provided in JSON format in the request body.
Response Format: All responses are in JSON format.

**Usage**
You can use the PERSON REST API to manage persons in your application. Below are some sample usages:

*Create a New Person:*

Request
POST /api/
Content-Type: application/json

```json
{
  "name": "John",
  "email": "john@gmail"
}

```


Response
```json
{
  "name": "John",
  "email": "john@gmail"
}

```


*Get Person by ID:*
Request
GET /api/1

Response
```json
{
  "name": "John",
  "email": "john@gmail"
}

```

*Update Person:*
PUT /api/1
Content-Type: application/json
```json

{
  "first_name": "Jane"
}

```
Response
```json
{
  "id": 1,
  "name": "Jane",
  "email": ""
}
```

*Delete Person:*
Request
DELETE /api/1
Response
```json
{
  "message": "Person has been deleted."
}
```

## Known Limitations
1. Only string data types are allowed for name and email fields.
2. The API does not support pagination for listing persons.
3. No authentication and authorization mechanisms are implemented in this version.

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

## Testing:

Test the API using tools like Postman or curl. Refer to the "Sample Usage" section for examples.

## Deployment:

To deploy the API on a server, follow your server hosting provider's deployment instructions. Ensure that the server environment is properly configured and that your API is accessible over the internet.




