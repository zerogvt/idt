# Part 1 - Web server

A simple server, which provides a RESTful API. 
It should implements next two API calls:

* GET /user/{id}
  Returns user details based on the id. You can return static mock results.
* PUT /user/{id}?{name}
  Modifies a user's name.

# how to run
`go run . [:port_number]`

# Sample commands-output
```
Retrieve a user:
```
> curl http://localhost:8080/user/90
{"id":90,"name":"000090"}
```
Create a user:
```
> curl -X PUT http://localhost:8080/user/101?name=billy -vv
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 8080 (#0)
> PUT /user/101?name=billy HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> 
< HTTP/1.1 201 Created
< Date: Fri, 26 Feb 2021 20:51:23 GMT
< Content-Length: 0
< 
* Connection #0 to host localhost left intact
* Closing connection 0
```
Update a user:
```
> curl -X PUT http://localhost:8080/user/101?name=vasilis

> curl http://localhost:8080/user/101
{"id":101,"name":"vasilis"}
```
