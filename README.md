# Golang Restful API using GORM ORM (MySQL) Gorilla Mux

## Getting Started
Run the following Golang commands to install all the necessary packages.

`go get -u github.com/gorilla/mux` for serving the api

`go get -u github.com/jinzhu/gorm` ORM supports MySQL, SQLite, MSSQL, Postgres

`go get -u github.com/go-sql-driver/mysql` MySQL driver to enable SQL connection

`go get -u github.com/joho/godotenv` Loads your environment variables (database, AWS, Redis configurations etc.)

### Running documentation locally (Only documentation of packages your have installed)
For offline documentation on the following packages run `godoc -http :6060` and then visit `http://localhost:6060`. Note that you can change the port to your preferred port number.

## Setting configuration file
Create a .env file in the root of the project and set the following parameters

`db_name = database_name` Name of database

`db_user = user`  // Database username

`db_pass = secret` // Database password

`db_type = mysql`   // MySQL driver

`db_host = localhost` // Database host

`db_port = 3306`  // Database port

`charset = utf8` // Database charset

`parse_time = True` // Database parse time

`web_port = 8085`   // Port to serve api

`prefix = /api/v1`  // API route sub route prefix

## API Endpoints

* GET `api/v1/users` retrieve all users
* GET `api/v1/users/1` retrieve user with id = 1
* POST `api/v1/users` create a new user
* PUT `api/v1/users/1` update the record with id = 1
* DELETE `api/v1/users/1` delete the user with id = 1