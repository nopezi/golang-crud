# Dockerfile References: https://docs.docker.com/engine/reference/builder/

# Start from the latest golang base image
FROM golang:1.17.1 as builder

RUN apt --yes --force-yes update && apt --yes --force-yes upgrade && \
    apt --yes --force-yes install git \
    make openssh-client

# Add Maintainer Info
LABEL maintainer="Diki Haryadi <diki.haryadi1902@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
RUN go build -o main .


######## Start a new stage from scratch #######
FROM alpine:latest
# FROM alpine:3.14

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app

# RUN mkdir /app

WORKDIR /app

# Expose port 5999 to the outside world
EXPOSE 5999

# Copy the Pre-built binary file from the previous stage
# COPY --from=builder /app .
COPY --from=builder /app /app

RUN pwd
RUN ls /app -a 
RUN ls /
# Command to run the executable
CMD ["./main"] 
# CMD /app/main serve-http
# CMD ["./main"] 
# ENTRYPOINT ["main serve-http"]