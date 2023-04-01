# Go REST API With JWT Authentication

To create a simple REST API in Go, you can use the net/http package that comes
with the standard library. In this example, I will demonstrate how to create a
basic REST API with two endpoints: one for retrieving all items and another for
adding an item. Note that for this demo application, these items are not
persisted to a database, so each time the application is run, you will
need to add items (see below).

# Create Go Module
To create a Go module, follow these steps:

1. Ensure you have Go installed on your system. You can download Go from the official website: https://golang.org/dl/

2. Open your terminal or command prompt and navigate to the directory where you want to create your new Go module.

3. Run the following command to create a new Go module:
```
go mod init github.com/jimnewpower/gorest
```
This command creates a new file called go.mod in the current directory. The go.mod file contains the module name and the Go version used. It should look something like this:
```go
module github.com/jimnewpower/gorest

go 1.20
```
Now you can start adding Go source files to your module.

# JSON Web Tokens
To secure access to your Go REST API, you can use JSON Web Tokens (JWT) for 
authentication and authorization. This example demonstrates how to integrate 
JWT with the existing API. We'll use the `github.com/dgrijalva/jwt-go` package 
to handle JWT creation and validation.

Install the `jwt-go` package:
```bash
go get github.com/dgrijalva/jwt-go
```

The code adds a JWT authentication middleware and a new `/login` endpoint. The 
middleware checks if the Authorization header is present in the request and 
validates the JWT. If the JWT is valid, the request is passed to the next 
handler. The `/login` endpoint accepts a `POST` request with a username and 
password, and if the credentials are valid, it generates and returns a JWT.

Replace the JwtSecretKey value with a strong secret key for your application. 
In a production environment, you should store the secret key securely, for 
example, in environment variables or a secret manager (e.g. Conjur).

Important: The provided login handler uses hardcoded credentials for 
demonstration purposes only. In a real-world application, you should replace 
this with your own authentication logic, such as querying a database to check 
if the provided username and password are correct.

# Build and Run the Application
```
go build
```

To test the secure access:

Build and run the application (or build and run the Docker container if you prefer).
```bash
build main.go
docker build -t gorest .
docker run -d -p 8080:8080 --name gorest-container gorest
```

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

# Dockerizing the Application
To deploy the Go REST API as a Docker image, you'll need to create a Dockerfile and then build and run the image using Docker. Follow these steps:

Install Docker on your system if you haven't already. You can find the installation instructions for your platform on the official Docker website: https://docs.docker.com/get-docker/

In your project directory (go-rest-api), create a new file called Dockerfile:
```bash
touch Dockerfile
```

Open the Dockerfile and add the following content:
```dockerfile
# Start from the official Go image
FROM golang:1.17-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy the Go modules and build cache
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Create a minimal image for deployment
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary from the builder image
COPY --from=builder /app/main .

# Expose the port on which the API will run
EXPOSE 8080

# Start the application
CMD ["./main"]
```

This Dockerfile uses a multi-stage build process to create a lightweight Docker image. The first stage (builder) starts with the official Go image, sets the working directory, copies the Go modules and source code, and builds the application. The second stage uses the alpine base image, copies the compiled binary from the builder stage, exposes the port, and starts the application.

Build the Docker image:
```bash
docker build -t gorest .
```

Run the Docker container:
```bash
docker run -d -p 8080:8080 --name gorest-container gorest
```

This command runs the Docker container, maps port 8080 on the host to port 8080 on the container, and assigns the container a name (go-rest-api-container).

Test the API using curl or a tool like Postman, just like you did before:
To add an item:
```bash
curl -X POST -H "Content-Type: application/json" -d '{"name": "item1"}' http://localhost:8080/items
```

To get all items:
```bash
curl -X GET http://localhost:8080/items
```

# Deploying to Kubernetes
To deploy a Docker image on Kubernetes, you can follow these step-by-step instructions. We'll use the Docker image of the Go REST API created in previous steps.

1. Set up a Kubernetes cluster:

You can use a managed Kubernetes service like Google Kubernetes Engine (GKE), Amazon Elastic Kubernetes Service (EKS), or Azure Kubernetes Service (AKS). Alternatively, you can set up a self-managed Kubernetes cluster using tools like kubeadm, kops, or minikube.

For this example, we'll assume you have a Kubernetes cluster up and running with kubectl configured to access it.

2. Push the Docker image to a container registry:

You need to push your Docker image to a container registry like Docker Hub, Google Container Registry (GCR), or Amazon Elastic Container Registry (ECR).

For this example, let's assume you have pushed the image to Docker Hub with the name yourusername/gorest:latest. Replace this with the actual image name and tag.

3. Create a Kubernetes deployment:

Create a file named gorest-deployment.yaml with the following content:
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: gorest
spec:
  replicas: 1
  selector:
    matchLabels:
      app: gorest
  template:
    metadata:
      labels:
        app: gorest
    spec:
      containers:
      - name: gorest-container
        image: yourusername/gorest:latest
        ports:
        - containerPort: 8080
```
Replace yourusername/gorest:latest with the actual image name and tag.

Apply the deployment to the cluster:
```bash
kubectl apply -f gorest-deployment.yaml
```

4. Create a Kubernetes service to expose the deployment:

Create a file named gorest-service.yaml with the following content:
```yaml
apiVersion: v1
kind: Service
metadata:
  name: gorest
spec:
  selector:
    app: gorest
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8080
  type: LoadBalancer
```
Apply the service to the cluster:
```
kubectl apply -f gorest-service.yaml
```
This service configuration creates a LoadBalancer that routes external traffic to the Go REST API deployment.

5. Access the deployed API:

Get the external IP address of the service:
```bash
kubectl get services gorest
```
Once the EXTERNAL-IP is assigned, you can access the API using the external IP and the specified port (80):
```bash
curl -X GET http://EXTERNAL-IP/items
```
Your Go REST API is now deployed on a Kubernetes cluster. You can further customize the deployment by adding environment variables, configuring resource limits, setting up autoscaling, and more.



