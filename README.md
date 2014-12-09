Okra
=======================

An OKR app to manage and facilitate your company's objectives and key results between team members.

##Dependencies
- [gin](github.com/gin-gonic/gin)
- [Go-MySQL-Driver](https://github.com/go-sql-driver/mysql)
- [gorm](github.com/jinzhu/gorm)
- [mgo](gopkg.in/mgo.v2)
- [mgo/bson](gopkg.in/mgo.v2/bson)
- [yaml](gopkg.in/yaml.v1)
- [cli](github.com/codegangsta/cli)

## Data Stores

- MySQL
- MongoDB

## Installation/Set up

1. Install MySql and MongoDB.

2. In your project directory, you can run

``` go
go get github.com/dmonay/do-work-api
```


## Tests and Benchmarks

To run tests, run 

```go
go test
```



To run benchmarks, run

```go 
go test -check.b
```

NOTE: no tests as of yet. To be added. 
## Documentation

1. Migrate the db. Run:

```go 
go run server.go migratedb
```

2. Start the server. Run: 

```go 
go run server.go server
```

### Create your organization

    POST /create/organization
    

**Sample Request Body**


```json
{
"organization":"DopeStartup"
}
```


**Sample Response Body**

```json
{
   "You have successfully created an organization"
}
```


### Add an OKR tree to the organization

    POST /create/tree/:organization
    

**Sample Request Body**

*NOTE*: Must pass in a parameter for the timeframe. One of "annual" or "monthly".

```json
{
"timeframe":"weekly"
}
```


**Sample Response Body**

```json
{
   "You have successfully created a tree with the annual timeframe for the DopeStartup organization"
}
```

### Add or update mission in your tree

    POST /update/mission/:organization
    

**Sample Request Body**

*NOTE*: Must pass in a parameter for the mission.

```json
{
"mission":"To thrive amidst enemies"
}
```


**Sample Response Body**

```json
{
   "You have successfully added a mission"
}
```
 	