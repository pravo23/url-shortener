# URL Shortener Service

A simple URL Shortener service built with the Gin framework.

## Features

- Shorten URLs
- Redirect shortened URLs to original URLs

## Requirements

- Go 1.16+
- Gin framework

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/pravo23/url-shortener.git
   ```
2. Navigate to the project directory:
   ```sh
   cd url-shortener
   ```
3. Install dependencies:
   ```sh
   go mod tidy
   ```

## Usage

1. Start the server:
   ```sh
   go run main.go
   ```
   The server will run on port `9808`.

2. Endpoints:

   - **GET /**: Welcome message
     ```sh
     curl http://localhost:9808/
     ```
     Response:
     ```json
     {
       "message": "The URL Shortener Service!"
     }
     ```

   - **POST /shorten**: Shorten a URL
     ```sh
     curl -X POST -H "Content-Type: application/json" -d '{"url":"http://example.com"}' http://localhost:9808/shorten
     ```
     Response:
     ```json
     {
       "shortUrl": "http://localhost:9808/abcd1234"
     }
     ```

   - **GET /:shortUrl**: Redirect to original URL
     ```sh
     curl http://localhost:9808/abcd1234
     ```
     This will redirect to `http://example.com`.

## Project Structure

- `main.go`: Main entry point
- `handler/`: URL handling logic
- `store/`: URL storage logic
