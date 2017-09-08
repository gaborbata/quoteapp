FROM golang:1.9

# Create application directory
RUN mkdir /app
ADD . /app
WORKDIR /app

# Create folder for dependencies
RUN mkdir /dependencies
ENV GOPATH /dependencies

# Get dependencies
RUN go get github.com/boltdb/bolt
RUN go get github.com/gin-gonic/gin

# Build application
RUN go build -o main quoteapp.go

# Expose port where the application will listen
EXPOSE 3030

# Run application
CMD ["/app/main"]
