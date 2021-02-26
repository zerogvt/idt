# Part 1 - Web server

A simple server, which provides a RESTful API. 
It should implements next two API calls:

* GET /user/{id}
  Returns user details based on the id. You can return static mock results.
* PUT /user/{id}?{name}
  Modifies a user's name.

# how to run
`go run . [:port_number]`
(This creates users 0-99 in the mocked DB)

# Sample commands-output
```
Retrieve a user:
```
curl http://localhost:8080/user/90
> {"id":90,"name":"000090"}
```
Create a user:
```
curl -X PUT http://localhost:8080/user/101?name=billy
```
Update a user:
```
curl -X PUT http://localhost:8080/user/101?name=vasilis
curl http://localhost:8080/user/101
> {"id":101,"name":"vasilis"}
```
