# Pismo
> API service for account and transactions
> 
![Test](https://github.com/hameedhub/pismo/actions/workflows/ci.yml/badge.svg)

##  Getting started
> Clone the repo and cd into the file directory

- Using Docker, run the command below to spin up the application

```shell
docker-compose up 
```

- To build binary

```shell
make build
```
- To run all the tests

```shell
make test
```

- To format file

```shell
make fmt
```

- Local development installation

```shell
 go mod tidy
```

## Environmental Variables
The application use the `.env` file on the root folder as variable configuration, currently it is not gitignore for testing purposes.
Below are required variables to start the application
```shell
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=pismo
DB_PORT=5432
SERVER_PORT=:8080
```

## Documentation (Swagger Doc)
The API endpoints are documented using swagger doc. It is automated using Swag and to generate new doc run
```shell
make swag-gen
```
The Swagger doc is available on [swagger]: http://localhost:8080/swagger/index.html after starting the application

## API Response 
On success response example
```shell
{
  "account_id": 1,
  "document_number": "12345678900"
}
```
On error response example
```shell
{
  "message": "invalid document number"
}
```
Status Code response definition
```shell
201 resource created
200 resource responsed
400 bad request
409 resource conflict
500 internal server error 

```
## Architectural Pattern
Repository and service pattern - structure;
```shell
cmd 
  - pismo
docs
  - swagger
internal
  - config 
  - handler
  - model 
  - repository
  - server
  - services
```
### Note
During development process ensure to build the image after changes
```shell
docker-compose up --build
```
To stop the application process, kill the process and run below command
```shell
docker-compose down
```
Also, ensure this following port are made available for application to use
```shell
 8080 - application server
 5432 - postgres database
 
```