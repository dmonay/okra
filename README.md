A GO microservice to calculate distance
=======================

This service responds to requests to calculate the straight-line distance between two sets of coordinates. 

##Dependencies
- [gocheck](http://gopkg.in/check.v1)
- [Go-MySQL-Driver](https://github.com/go-sql-driver/mysql)
- [gorp](https://github.com/coopernurse/gorp)

## Data Stores

- MySQL

## Installation/Set up

In your project directory, you can run

``` go
go get github.com/dmonay/worth_my_salt
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


## Documentation

### Getting distance

    POST /distance
    

**Sample Request Body**


```json
{
"lat1":"39.768434", 
"lon1":"-104.901872", 
"lat2":"44.7793732", 
"lon2":"-69.6734886", 
"unit":"km"
}
```


*NOTE*: Units MUST be one of "mi", "miles", "km", or "kilometers".

**Sample Response Body**

```json
{
   "Success":"You have 2927.7229366372153 km to travel."
}
```
