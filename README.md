# Call-Cat-API

Call-Cat-API is an API made in Go that queries the [Cat API](https://docs.thecatapi.com/).

Here we have 2 endpoints:
* /login (for authentication)
* /breeds (for query breeds by name, you should pass a name through query param like '/breeds?name=sib')

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

You will need installed:
* Go
* Mysql
* Git

### Installing

First we will clone this project to your local machine:

```bash
git clone https://github.com/rfaelmarini/call-cat-api.git
```

Second in server.go file you will setup env variables:

```
os.Setenv("API_KEY", "your_cat_api_key_here")
os.Setenv("JWT_SECRET", "your_jwt_secret_here")
os.Setenv("DB_NAME", "your_db_name")
os.Setenv("DB_USER", "your_db_user")
os.Setenv("DB_PASSWORD", "your_db_password")
os.Setenv("DB_ADDRESS", "your_db_address_with_port")
```

Third run the server:
```bash
go run server.go
```

## Configuring and running the tests

The tests file is in service folder and named service_test.go, there you may change the env variables, is recommend to use a different database that is set in env variables that was set in server.go.

To run the tests, you should execute this statement in the folder '/service':
```bash
go test
```

## Built With

* [Go](https://golang.org/) - The language used
* [Gin](https://github.com/gin-gonic/gin) - The web framework used