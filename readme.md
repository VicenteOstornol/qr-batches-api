# How to run the project

- Clone the repository.
- Run docker-compose up.
- The server will be running on http://127.0.0.1:8080.

# Endpoints

- Prefix: /api (e.g. http://127.0.0.1:8080/api)

## Users

### POST /user/create

Creates a new user.

#### Request Body

- `User` (object, required)
  - `name` (string, required) - The name of the user.

#### Responses

- `200` - Successful operation. Returns the created user.
- `400` - Bad request. The user data is invalid.
- `422` - Unprocessable Entity. The request was well-formed but was unable to be followed due to semantic errors.
- `500` - Internal server error. An error occurred on the server.

### GET /user/{id}/batches

Gets the batches by user ID.

#### Path Parameters

- `id` (string, required) - The ID of the user.

#### Responses

- `200` - Successful operation. Returns the batches of the user.
- `500` - Internal server error. An error occurred on the server.

## Batches

### POST /batch/create

Creates a new batch.

#### Request Body

- `Batch` (object, required)
  - `id` (string, required) - The ID of the batch.
  - `name` (string, required) - The name of the batch.
  - `user_id` (string, required) - The ID of the user.
  - `amount_qrs` (number, required) - The amount of QRs of the batch.

#### Responses

- `200` - Successful operation. Returns the created batch.
- `400` - Bad request. The batch data is invalid.
- `422` - Unprocessable Entity. The request was well-formed but was unable to be followed due to semantic errors.
- `500` - Internal server error. An error occurred on the server.

### GET /batch/{id}/download

Downloads the PDF of the batch by batch ID.

#### Path Parameters

- `id` (string, required) - The ID of the batch.

#### Responses

- `200` - Successful operation. Returns the PDF of the batch.
- `500` - Internal server error. An error occurred on the server.

### GET /batch/{id}/{qr_number}

Get the batch with the QR number. This endpoint is used when the user scans a QR code.

#### Path Parameters

- `id` (string, required) - The ID of the batch.
- `qr_number` (string, required) - The QR number of the batch.
