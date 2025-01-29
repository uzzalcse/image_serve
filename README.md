# Image Server

A simple local image server that serves images through HTTP.

## Description

This project provides a basic HTTP server that serves image files from a local directory. Images can be accessed via the URL pattern: `http://localhost:8082/images/image_name.type`

## Setup

1. Place your images in the `images` directory
2. Run the server

## Usage

Access your images through:
```
http://localhost:8082/images/example.jpg
http://localhost:8082/images/photo.png
```

## Supported Image Types

- JPG/JPEG
- PNG
- GIF
- SVG
- WebP

## Running the Server

```bash
# Start the server
# Clone the git repository
https://github.com/uzzalcse/image_serve.git
go mod tidy
# Go the directory then rum 
go run main.go
```

Server will be available at `http://localhost:8082`