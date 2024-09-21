# NGL Token Supply API

This application provides a simple REST API to fetch the total supply of NGL tokens.

## Setup

1. Clone the repository.
2. Install Go on your system if it's not installed yet.
3. In order to install dependencies run:
    ```bash
    go mod tidy
    ```

## Configuration

Configuration can be provided in two ways:

1. Through a `config.json` file with the following structure:
    ```json
    {
        "port": "8080",
        "url": "https://someurl/getsomething"
    }
    ```

2. Alternatively, the configuration can be provided via environment variables:
    - `PORT`: The port number to run the server on (e.g., `8080`).
    - `URL`: The URL of the external service providing the NGL token supply.

## Running the Application

1. Make sure the config is set up properly (either by providing a `config.json` file or setting environment variables).
2. Build the application and run:

   ```bash
   go build -o main
   ./main
   ```
