# Go REST API With JWT Authentication

To create a simple REST API in Go, you can use the net/http package that comes
with the standard library. In this example, I will demonstrate how to create a
basic REST API with two endpoints: one for retrieving all items and another for
adding an item.

To secure access to your Go REST API, you can use JSON Web Tokens (JWT) for 
authentication and authorization. This example demonstrates how to integrate 
JWT with the existing API. We'll use the github.com/dgrijalva/jwt-go package 
to handle JWT creation and validation.

Install the `jwt-go` package:
```bash
go get github.com/dgrijalva/jwt-go
```

The code adds a JWT authentication middleware and a new /login endpoint. The 
middleware checks if the Authorization header is present in the request and 
validates the JWT. If the JWT is valid, the request is passed to the next 
handler. The /login endpoint accepts a POST request with a username and 
password, and if the credentials are valid, it generates and returns a JWT.

Replace the JwtSecretKey value with a strong secret key for your application. 
In a production environment, you should store the secret key securely, for 
example, in environment variables or a secret manager.

Important: The provided login handler uses hardcoded credentials for 
demonstration purposes only. In a real-world application, you should replace 
this with your own authentication logic, such as querying a database to check 
if the provided username and password are correct.

To test the secure access:

Build and run the application (or build and run the Docker container if you prefer).

Request a JWT token using the /login endpoint with valid credentials:
```bash
curl -X POST -d "username=testuser&password=testpassword" http://localhost:8080/login
```

The response will contain the JWT token:
```json
{
  "token": "your_jwt_token_here"
}
```

Use the JWT token to access the `/items` endpoints. To add an item:
```bash
curl -X POST -H "Content-Type: application/json" -H "Authorization: Bearer your_jwt_token_here" -d '{"name": "item1"}' http://localhost:8080/items
```
To get all items:
```bash
curl -X GET -H "Authorization: Bearer your_jwt_token_here" http://localhost:8080/items
```

These requests should now work only when the Authorization header contains a 
valid JWT token. If the token is missing, expired, or invalid, the server will 
return an HTTP 401 Unauthorized status.

You have now added secure access to your Go REST API using JSON Web Tokens. You 
can further improve this solution by implementing role-based access control, 
more advanced token management, and other security best practices based on 
your application's requirements.

