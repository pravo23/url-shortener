# URL Shortener

A simple and efficient URL shortening service built in Go using Gin. This service shortens long URLs and allows you to retrieve the original URL.

## Features

- Shorten long URLs
- Retrieve original URLs from short codes
- Lightweight and fast

## Getting Started

### Using Docker Compose

1. Clone the repository:
   ```bash
   git clone https://github.com/pravo23/url-shortener.git
   cd url-shortener
   ```

2. Build and start the services:
   ```bash
   docker-compose up --build
   ```

3. Access the application at `http://localhost:9808`.

### API Endpoints

- **GET /**  
  Returns a service message.  
  **Response:** `{"message": "The URL Shortener Service!"}`

- **POST /shorten**  
  Shortens a URL.  
  **Request Body:** `{"url": "http://example.com"}`  
  **Response:** `{"shortened_url": "http://short.url/abcd1234"}`

- **GET /{shortUrl}**  
  Redirects to the original URL.  
  **Request URL:** `http://localhost:9808/{shortUrl}`

## Running Locally (Without Docker)

1. Build the application:
   ```bash
   go build -o main .
   ```

2. Ensure Redis is running locally on port 6379.

3. Set Redis environment variables:
   ```bash
   export REDIS_HOST=localhost
   export REDIS_PORT=6379
   ```

4. Run the application:
   ```bash
   ./main
   ```

## Acknowledgements

- [Golang](https://golang.org/)
- [Gin](https://gin-gonic.com/)
- [Redis](https://redis.io/)
- [Alpine Linux](https://www.alpinelinux.org/)