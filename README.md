# Api with Go and Couchbase on Docker

This script setup a Go Api linked to Couchbase on Docker

## Dockerfile

The Dockerfile sets up an image with GoLang library on Alpine and installs the following libraries: 

* git
* couchbase/gocb.v1
* gorilla/mux
* joho/godotenv
* satori/go.uuid

After installing, it copies folder content to the app-folder.

Create the Docker container by using these instructions:

```bash
docker build -t go-api:1.0 .

docker run -dit --restart unless-stopped -p 8081:8081 --name go-api go-api:1.0
```

Make sure on the server, the loggedin user is added to group docker:

```bash
sudo usermod -aG docker XXXX 
```

XXXX is the username