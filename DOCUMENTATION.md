#API Documentation

This document provides documentation for the PERSON REST API, which allows you to perform CRUD operations on the "Person" resource.

##Table of Contents
1. Endpoints
2. Standard Formats
3. Sample Usage
4. Known Limitations
5. Setup Instructions


#Endpoints
###Create a New Person
###Endpoint: POST /api/
###Description: Create a new person.
###Request Format:
```json
{
  "name": "John",
  "email": "john@gmail.com"
}

Response Format:
json
{
  "id": 1,
  "name": "John",
  "email": "john@gmail.com"
}

###Get Person by ID
###Endpoint: GET /api/{id}
###Description: Retrieve details of a person by their ID.
###Response Format:

```json

{
  "id": 1,
  "name": "John",
  "email": "john@gmail"
}


###Update Person
###Endpoint: PUT /api/persons/{id}
###Description: Update details of an existing person.
###Request Format:
```json
Copy code
{
  "name": "Jane"
  "email": "jane@gmail"
}
###Response Format:
```json

{
  "id": 1,
  "name": "Jane",
  "email": "jane@gmail"
}


###Delete Person
###Endpoint: DELETE /api/{id}
###Description: Remove a person by their ID.
###Response Format:
```json

{
  "message": "Person has been deleted."
}


#Standard Formats
Request Format: All request data should be provided in JSON format in the request body.

Response Format: All responses are in JSON format.

##Sample Usage
Create a New Person
Request:

http
POST /api/
Content-Type: application/json

{
  "name": "John",
  "email": "john@gmail"
}
Response:

```json

{
  "id": 1,
  "name": "John",
  "email": "john@gmail"
}

Get Person by ID
Request:

http
GET /api/1
Response:

json

{
  "id": 1,
  "name": "John",
  "email": "john@gmail"
}

Update Person
Request:

http
PUT /api/1
Content-Type: application/json

{
  "name": "Jane"
}
Response:

json

{
  "id": 1,
  "first_name": "Jane",
  "email": ""
}
Delete Person
Request:

http
Copy code
DELETE /api/1
Response:

json

{
  "message": "Person  has been deleted."
}


#Known Limitations

> Only string data types are allowed for first_name and last_name fields.
> The API does not support pagination for listing persons.
> No authentication and authorization mechanisms are implemented in this version.

#Setup Instructions
Follow these steps to set up and deploy the API locally or on a server:

##Clone the Repository:

Clone the API repository from GitHub:

bash
git clone [Link Text] https://github.com/Mac-5/PersonApi.git
Install Dependencies:

Navigate to the API directory and install the required dependencies:

bash
cd your-api
go mod tidy
Database Configuration:

Update the database configuration in your application code (if applicable). Ensure that the database server is running.

Run the API:

Start the API server:

bash
go run main.go
The API should be accessible at [Link Text] http://localhost:PORT, where PORT is the port specified in your configuration.

##Testing:

Test the API using tools like Postman or curl. Refer to the "Sample Usage" section for examples.

##Deployment:

To deploy the API on a server, follow your server hosting provider's deployment instructions. Ensure that the server environment is properly configured and that your API is accessible over the internet.