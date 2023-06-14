# Golang OMDB Movie Search

This project provides a lightweight Golang server that can search for movies and fetch movie details from the OMDB API. It offers two endpoints: /search and /detail/:id to perform the respective operations.

The server employs environment variables for storing API access credentials and supports HTTP frameworks like Gin for building the API.

---

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Ensure you have the following installed on your system:

- Go 1.20 or later
- Gin framework
- Godotenv
- An IDE or terminal

### Installation

1. Clone this repository to your local machine using

   git clone `https://github.com/jameschandra/golang-omdb-api.git`

2. Navigate to the project directory.

   cd `<project-dir>`

3. Install the necessary go dependencies.

   go mod tidy

4. Run the server locally.

   go run cmd/server/main.go

---

## Usage

- Search for a movie: `GET /search?s=<movie_title>`
- Get a movie's details: `GET /detail/:id`

Replace <movie_title> and :id with the movie title and movie ID you're interested in, respectively.

---

## Testing

We have unit tests for important functionalities. Run the tests with the following command:

    `go test ./test/movie`

---

## References

Readability of code: https://thehosk.medium.com/why-code-readability-is-important-e0c228a238a

Separation of concerns for codes: https://github.com/oralordos/separation

Unit tests for critical files: https://go.dev/doc/tutorial/add-a-test

Implementation of Clean Architecture: https://medium.com/@MTrax/clean-architecture-in-golang-a-practical-approach-to-building-maintainable-software-19e4d16c635b

Environment variables for API credentials: -

---

## Images

<img width="685" alt="image" src="https://github.com/jameschandra/golang-omdb-api/assets/53634665/fcd62e63-1e41-46c0-8d73-44f8db161e42">

Image 1) API Status

<img width="685" alt="image" src="https://github.com/jameschandra/golang-omdb-api/assets/53634665/47a4eee1-8c0e-4c74-bcd0-e5357b3a64a0">

Image 2) API Curl Response
