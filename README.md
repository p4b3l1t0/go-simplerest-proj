# go-simplerest-proj
Go REST API CRUD Simple Project

## This is a simple go api rest CRUD project

### This code uses PostGreSQL as Database Engine.

    docker run --name some-postgres -e POSTGRES_USER=p4b3l -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres

### Install following modules

    go get github.com/p4b3l1t0/go-simplerest-proj
  
    go get -u gorm.io/drivers/posgres
 
    go get -u gorm.io/gorm
  
    go get github.com/gorilla/mux 
