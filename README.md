# Simple authentication example with JWT

## How to
1. Clone the repo
2. `go mod tidy` to install dependencies
3. Create .envrc file for your environment variables
```
export DSN=postgres://<db_user>:<db_pass>@<db_host>/<db_name>
export JWT_SECRET=<random words>
```
4. `make start` to start the application


## API Reference

####  `GET http://localhost:9000/`

```http
  _response_
  {
      "msg": "index route"
  }
```

#### `POST http://localhost:9000/register`

```http
  _request_
  {
      "username" : "user one",
      "email": "userone@gmail.com",
      "password" : "userone123"
  }

  _response_
  {
    "msg": "user successfully registered"
  }
```

#### `POST http://localhost:9000/login`

```http
  _request_
  {
      "email": "userone@gmail.com",
      "password" : "userone123"
  }

  _response_
  {
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.
      eyJpc3MiOiJqd3QtZXhhbXBsZSIsInN1YiI6InVzZXJvbmVAZ21haWwuY29tIiwiZXhwIjoxNjU2OTQ3MTUyLCJpYXQiOjE2NTY5NDM1NTJ9.
      vHSTfTB1ICllFc4-J2lGFJT7J_Dwodse1WEFYn1JC1M"
  }
```

#### `GET http://localhost:9000/private`

| Header             | Description    |
|:-------------------|:---------------|
| `Authorization`    | Your JWT Token |

```http
  _response_
  {
    "msg": "private route",
    "user": "userone@gmail.com"
  }
```
