FROM golang:1.21-alpine

# Create app directory
WORKDIR /app

# copy files to the container
COPY . .

# download dependencies
RUN go get -d -v ./...

# build app
RUN go build -o main .

# expose ports
EXPOSE 8080

# run app
CMD ["./main"]
