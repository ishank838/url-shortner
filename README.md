# url-shortner

A simple url shortner service for creating a short url from a large one.

## Build

Initialise the env variables in local.env
```
WEB_PORT=8080
REDIS_PORT=6379
REDIS_ADRESS="db:6379"
```

Navigate to api/ and run

```
go run main.go
```

### Docker
```
docker-compose up
```

## API's
Url shortner service has following api's.
```
POST: /shorten
GET: /{short_key}
````

/shorten accepts a url and returns a shortened url.
/{short_key} is used to redirect the short url to actual url.
