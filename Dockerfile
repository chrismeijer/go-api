# SPECIFY BASE IMAGE
    FROM golang:1.12.0-alpine3.9

# INSTALL PACKAGES/LIBRARIES
    RUN apk add git
    RUN go get gopkg.in/couchbase/gocb.v1
    RUN go get github.com/gorilla/mux
    RUN go get github.com/joho/godotenv
    RUN go get github.com/satori/go.uuid

# CREATE DIR WHERE THE APP WILL BE COPIED TO
    RUN mkdir /app
    # COPY TO DIR
        ADD . /app

# SET WORKDIR
    WORKDIR /app

# RUN APP
    RUN go build -o bin/main ./src

# DOCUMENT THAT CONTAINER USE PORT 8081
    EXPOSE 8081

# RUN THE EXECUTABLE
    CMD ["/app/bin/main"]