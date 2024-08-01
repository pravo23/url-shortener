# URL Shortener Service

This is a simple URL Shortener service built using the Gin framework for the web server. The service allows users to shorten URLs and redirect shortened URLs to the original URLs.

## Features

- Shorten a given URL.
- Redirect to the original URL using the shortened URL.

## Requirements

- Go 1.16+
- Gin framework
- URL Shortener handler and store packages

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/pravo23/url-shortener.git
   ```

2. Navigate to the project directory:

   ```sh
   cd url-shortener
   ```

3. Install the required dependencies:

   ```sh
   go mod tidy
   ```

## Usage

1. Start the web server:

   ```sh
   go run main.go
   ```

   The server will start on port `9808`.

2. Endpoints:

   - **GET /**

     Returns a welcome message for the URL Shortener service.

     ```sh
     curl http://localhost:9808/
     ```

     Response:

     ```json
     {
       "message": "The URL Shortener Service!"
     }
     ```

   - **POST /shorten**

     Shortens a given URL.

     ```sh
     curl -X POST -H "Content-Type: application/json" -d '{"url":"http://example.com"}' http://localhost:9808/shorten
     ```

     Response:

     ```json
     {
       "shortUrl": "http://localhost:9808/abcd1234"
     }
     ```

   - **GET /:shortUrl**

     Redirects to the original URL using the shortened URL.

     ```sh
     curl http://localhost:9808/abcd1234
     ```

     This will redirect to `http://example.com`.

## Project Structure

- **main.go**: The main entry point of the application.
- **handler/**: Contains the logic for handling URL shortening and redirection.
- **store/**: Contains the logic for storing and retrieving URL mappings.

## Error Handling

If the server fails to start, it will panic and print the error message.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

---

Feel free to customize this README file according to your project's specifics and requirements.