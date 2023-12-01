# Examples

The examples use a sample database with data from AirBnb (see below). You can start the
database with the [docker-compose](docker-compose.yml) file in this directory.

## Setup the start the database

Run the following command in this directory to setup a sample database with 
AirBnb data (see https://www.mongodb.com/docs/atlas/sample-data/sample-airbnb/#std-label-sample-airbnb).

```bash
docker-compose up --build --force-recreate --remove-orphan
```