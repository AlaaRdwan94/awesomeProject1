## how to run

--this project uses postgresql database by default if you want to change to mysql please uncomment 
mysql configration environment variables in `.env` file 
--pre-requests : 

  - install either `postgresql` or `mysql` database in your machine 
  - create the database and user with the same configrations in `.env` file 
  - run `go run main.go`
  - import the postman collection from `promotion-task.postman_collection.json` file

--no-need-to :

  - create the database schema or tables

## frameworks used :

  - Gorilla Mux --> for API 
  - Gorm --> for database

## my code includes : 

  - Database schema implementation using `gorm` to implement tables and relations 
  - CRUD operations for `users` and `companies`
  - Factory design pattern between `Model` and `Controllers`
  - APIs implementation using `gorilla/mux`
