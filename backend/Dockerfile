# Use golang version 1.16
FROM golang:1.16-alpine

# create the directory where all app files will live
RUN mkdir /app

# add contents of root directory into the app directory we created
ADD . /app

# specify /app as our work directory so that all commands run inside there
WORKDIR /app

# use go modules for easier dependency management 
ENV GO111MODULE=on

# copy go files an
COPY go.mod .
COPY go.sum .
RUN go mod download

# build the binary executable
RUN go build -o main .

# start the web server
CMD ["/app/main"]
