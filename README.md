## Installation
Clone project on your local machine
```bash
$ git clone git@github.com:thimovez/service.git
```

## Running Application in production mode
Make sure you have Docker Compose installed.
```bash
docker compose up service
```

## Running Application in develop mode
Run database
```bash
make db
```
Run backend
```bash
make app
```



| Method   | URL Pattern                  | Action                                        | Command                                              |
|:---------|:-----------------------------|:----------------------------------------------|:-----------------------------------------------------|
| POST     | /authorization/login         | Login registered user into service            | curl -i -X POST http://localhost:4000/snippet/create |
| POST     | /authorization/registration  | Register user into our service                | curl -i -X POST http://localhost:4000/snippet/create |
| POST     | /authorization/refresh       | Renew refresh token when access token expired | curl -i -X POST http://localhost:4000/snippet/create |