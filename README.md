![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/l-const/timeService) 
# timeService

## Run locally



#### 1. **Install dependencies**

---

```code
 go mod download
```
####  2. **Run Server**

---

```shell
 go run ./cmd/server/main.go <host> <port>
```

e.g

```shell
 go run ./cmd/server/main.go localhost 9998
```

#### 3. **Regarding Host/port variables**:

---


 * If the host or port variable is not passed during execution the server tries to read the ENVIRONMENT vars **HOST** and **PORT**.
 * The environment variables are predefined in [config/env.go](https://github.com/l-const/timeService/blob/main/config/.env#L2-L3) as follow: **HOST**=localhost, **PORT**=8080.

e.g running without specifying anything: 


```shell
  go run ./cmd/server/main.go 
```

or

```shell
 HOST=localhost PORT=9993 go run ./cmd/server/main.go
```

## Run with Docker
---

#### 1. Building the image

```shell
  docker build 
```

