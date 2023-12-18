# Use the official Golang image as the base image
FROM golang:1.18

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the Go module dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
# RUN go build -o main .

# Expose the port on which the API will listen
EXPOSE 8000

# Specify the command to run when the container starts
# CMD ["go", "run", "apigatway.go"]
RUN chmod +x /app/entrypoint.sh
RUN chmod +x /app/wait-for-it.sh

ENTRYPOINT ["/app/wait-for-it.sh" , "golang-app" , "/app/entrypoint.sh"]

