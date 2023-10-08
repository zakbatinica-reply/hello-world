# Hello World

## Description

This is a simple web application written in Go using the Gin framework. It serves a web page that displays a greeting message ("Hello, World!") and a table containing details about the HTTP request headers.

## Prerequisites

- Go 1.16 or later

## Directory Structure

- `main.go`: The main Go application file.
- `views/index.html`: The HTML template for the web page.
- `static/css/stylesheet.css`: The CSS file to style the web page.

## How to Build and Run

1. Clone this repository to your local machine.
2. Navigate to the directory containing the application.
3. Build the application with the `go build` command. This will create an executable file in the same directory.

    ```bash
    go build
    ```

4. Run the application with the `./<executable>` command, where `<executable>` is the name of the executable file created by the `go build` command. Set the `PORT` environment variable to `3000`.

    ```bash
    PORT=3000 ./<executable>
    ```

5. Open a web browser and navigate to `http://localhost:3000` to see the web page.

## Docker

You can also run this application inside a Docker container.

### Prerequisites

- Docker

### How to Build and Run

1. Build the Docker image from the Dockerfile:

    ```bash
    docker build -t go-web-app .
    ```

    This command builds a Docker image named `go-web-app` from the Dockerfile in the current directory.

2. Run the Docker container from the image:

    ```bash
    docker run -p 3000:3000 go-web-app
    ```

    This command runs a Docker container from the `go-web-app` image and maps port 3000 inside the container to port 3000 on the host machine.

3. Open a web browser and navigate to `http://localhost:3000` to see the web page.

## How to Test

Use the `curl` command to send a request to the application and see the raw HTML response.

```bash
curl http://localhost:3000
```

Or visit the URL in your browser of choice.

## Note

The application serves static files from the static directory at the /static path. The CSS file is loaded from this path in the HTML file. If you add more static files (like images or JavaScript files), you should put them in the static directory and update the HTML file to load them from the correct path.