# Go Library Stats Service

This project provides an API service to fetch statistics about books and authors from the Gutendex library. It offers endpoints to retrieve book counts for given languages, potential readership data based on the language, and service status diagnostics.

## Features

- **Book Count Endpoint**: Returns the count of books and unique authors for specified languages.
- **Readership Endpoint**: Provides potential readership numbers for books in a given language, based on the population of countries where the language is official.
- **Diagnostics Endpoint**: Offers a status overview of the service and its dependent external APIs.

## Getting Started

### Prerequisites

- Go 1.15 or later
- Access to the internet for fetching data from external APIs

### Installation

Clone the repository and navigate to the project directory:

```bash
git clone https://github.com/britotiago03/cloud-tech-assignment1.git
cd cloud-tech-assignment1
```

Download the necessary Go modules:

```bash
go mod tidy
```

### Running Locally

To start the server on your local machine, run:

```bash
go run cmd/myservice/main.go
```

The server will start, and you can access the API endpoints at `http://localhost:8080`.

## API Endpoints

- **Book Count**: `GET /bookcount/?language=<language_codes>`
- **Readership**: `GET /readership/{language_code}?limit={limit}`
- **Diagnostics**: `GET /status/`

Replace `<language_codes>` with comma-separated 2-letter ISO 639-1 language codes and `{language_code}` with a single 2-letter ISO 639-1 language code. The `{limit}` parameter is optional.

## Deployment on Render

Follow [Render's documentation](https://render.com/docs) to deploy this Go application. You'll need to create a new Web Service, choose the repository, and configure it based on Render's Go environment.
