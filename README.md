# timeService

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/l-const/timeService)

## Run locally
---
#### 1. Install dependencies
---

```code
	go mod download
```
###  2. Run Server
---

```shell
	go run ./cmd/server/main.go <host> <port>
```

e.g

```shell
	go run ./cmd/server/main.go localhost 9998
```

#### 3. Regarding Host/port variables:

 * If not host or port variable is passed during execution the server tries to read ENVIRONMENT VARIABLES **HOST** and **PORT**.
 *  The environment variables are predefined here [config/env.go](https://github.com/l-const/timeService/blob/main/config/.env#L2-L3) as follow: **HOST**=localhost, **PORT**=8080.

## Run with Docker
---