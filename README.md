# Go REST API

This is a simple REST API built with Go and the Gin framework. It provides CRUD functionality for managing events.

## Features

*   **Create, Read, Update, and Delete (CRUD) operations for events.**
*   **Uses SQLite for data storage.**
*   **Built with the Gin web framework.**

## API Endpoints

The following API endpoints are available:

| Method | Endpoint      | Description        |
| :--- | :------------ | :----------------- |
| `GET`  | `/events`     | Get all events     |
| `POST` | `/events`     | Create a new event |
| `GET`  | `/events/:id` | Get an event by ID |
| `PUT`  | `/events/:id` | Update an event    |
| `DELETE`| `/events/:id` | Delete an event    |

### Create Event

To create a new event, send a `POST` request to `/events` with the following JSON body:

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

To update an event, send a `PUT` request to `/events/:id` with the same JSON body as the create request.

### Delete Event

To delete an event, send a `DELETE` request to `/events/:id`.

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

The `api test` directory contains `.http` files that can be used with the [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) extension for Visual Studio Code to test the API endpoints.
