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

