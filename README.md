# Go REST API

This is a simple REST API built with Go and the Gin framework. It provides CRUD functionality for managing events and includes user authentication.

## Features

*   **Create, Read, Update, and Delete (CRUD) operations for events.**
*   **User signup and login.**
*   **Authentication using JWT.**
*   **Uses SQLite for data storage.**
*   **Built with the Gin web framework.**

## API Endpoints

### Event Endpoints

| Method | Endpoint      | Description        |
| :--- | :------------ | :----------------- |
| `GET`  | `/events`     | Get all events     |
| `POST` | `/events`     | Create a new event (Authentication required)|
| `GET`  | `/events/:id` | Get an event by ID |
| `PUT`  | `/events/:id` | Update an event (Authentication required)   |
| `DELETE`| `/events/:id` | Delete an event (Authentication required)   |

### User Endpoints

| Method | Endpoint      | Description        |
| :--- | :------------ | :----------------- |
| `POST` | `/signup`     | Create a new user  |
| `POST` | `/login`      | Login a user       |

### Create Event

To create a new event, send a `POST` request to `/events` with the following JSON body and a valid JWT in the `Authorization` header.

```json
{
  "name": "Event Name",
  "description": "Event Description",
  "location": "Event Location",
  "dateTime": "2025-01-01T15:30:00.0000Z"
}
```

### Get Events

To get all events, send a `GET` request to `/events`.

### Get Event by ID

To get a specific event, send a `GET` request to `/events/:id`, where `:id` is the ID of the event.

### Update Event

To update an event, send a `PUT` request to `/events/:id` with the same JSON body as the create request and a valid JWT in the `Authorization` header.

### Delete Event

To delete an event, send a `DELETE` request to `/events/:id` with a valid JWT in the `Authorization` header.

### Signup

To create a new user, send a `POST` request to `/signup` with the following JSON body:

```json
{
  "email": "test@example.com",
  "password": "password"
}
```

### Login

To login, send a `POST` request to `/login` with the following JSON body:

```json
{
  "email": "test@example.com",
  "password": "password"
}
```

The API will return a JWT upon successful login.

## Getting Started

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/popey17/go-rest-api.git
    ```

2.  **Install dependencies:**

    ```bash
    go mod tidy
    ```

3.  **Run the application:**

    ```bash
    go run main.go
    ```

The API will be running at `http://localhost:8080`.

## `api test`

The `api test` directory contains `.http` files that can be used with the [REST Client](https.marketplace.visualstudio.com/items?itemName=humao.rest-client) extension for Visual Studio Code to test the API endpoints.