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

## Documentation (Swagger Doc)
The API endpoints are documented using swagger doc. It is automated using Swag and to generate new doc run
```shell
make swag-gen
```
The Swagger doc is available on [swagger]: http://localhost:8080/swagger/index.html after starting the application

## Environmental Variables
The application use the `.env` file on the root folder as variable configuration, currently it is not gitignore for testing purposes.



