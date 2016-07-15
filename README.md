# Example auth service

An exploration of an authentication microservice in go based on
[JWT](https://jwt.io) and Postgres for storage.

## Build

`go install`

## Configuration

The service is configured by environment variables:

- `DB_URL`
- `SECRET_KEY`

## Testing

Unit tests require a docker environment with the necessary services to run.

`docker-compose run api go test`
