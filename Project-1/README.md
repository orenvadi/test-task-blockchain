# NGL Token Supply API

This application provides a simple REST API to fetch the total supply of NGL tokens.

## Setup

1. Clone this repo.
2. Install Go on your system if it's not installed.
3. Install dependencies:
    ```bash
    go mod tidy
    ```

## Configuration

Configuration can be provided in two ways:

1. With a `config.json`:
    ```json
    {
        "port": "8080",
        "url": "https://someurl/getsomething"
    }
    ```

2. Alternatively, the configuration can be provided via environment variables:
    - `PORT`: The port number to run the server(`8080`).
    - `URL`: The URL of the external service providing the NGL token.

## Running the Application

1. Config must be properly setted up (either by a `config.json` or environment variables .env).
2. Build the application and run:

   ```bash
   go build -o main
   ./main
   ```
