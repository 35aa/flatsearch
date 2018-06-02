# flatsearch

Let's begin ;)

## Options service

Suppose to handle users requests with search data and store them into storage.

### Installation
```
// PROD
docker build --file Dockerfile.options --tag options:0.0.1 .
docker run -p 3000:3000 options:0.0.1
curl 127.0.0.1:3000

// DEV
docker-compose --file docker-compose.options.yaml build
docker-compose --file docker-compose.options.yaml up
```
