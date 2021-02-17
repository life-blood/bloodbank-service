## Microservice for donations exposing REST-API with MySQL database

Supports CRUD operations for Donation entity

See project docs: https://github.com/life-blood/documentation



#### Run locally

`$ git clone https://github.com/life-blood/bloodbank-service/`

Configure Local MySql Server and setup credentials in .env file
or have configured environment varibales

`DB_PORT=3306`
`DB_USER=root`
`DB_PASS=password`
`DB_NAME=databasename`

Run `db-init.sql` script and initialize the database

Start the server
`$ go run ./main.g`

Run tests
`$go test ./...`


Follow the instruction to configure UI and LifeBlood microservice

- https://github.com/life-blood/lifeblood-ui
- https://github.com/life-blood/accounts-service



## LifeBlood Project Architecture
![alt text](https://i.ibb.co/M7C45Wv/Architecture.png)
