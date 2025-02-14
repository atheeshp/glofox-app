# glofox-app

## Setup Instructions

### Prerequisites

Ensure you have the following installed:

- [Go](https://go.dev/dl/) (1.23 used in this project)
- [cURL](https://curl.se/) (optional, for command-line API testing)

### Installation & Running the Server

1. Clone the repository:

   ```sh
   git clone github.com/atheeshp/glofox-app
   cd glofox-app

To start:

```sh
go mod tidy
go run cmd/main.go # The server should start at: http://localhost:8080
```

After starting the server run the commands in the CLI to check:

```sh
# to create a class using command line
curl -X POST http://localhost:8080/api/classes/ \
  -H "Content-Type: application/json" \
  -d '{
        "name": "Yoga",
        "start_date": "2024-06-01",
        "end_date": "2024-06-10",
        "capacity": 20
      }'

# to book a class using command line
curl -X POST http://localhost:8080/api/bookings/ \
  -H "Content-Type: application/json" \
  -d '{
        "member": "John Doe",
        "date": "2024-06-10"
      }'


```

To test:

```sh
go test ./...
```
