![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/l-const/timeService) 
# timeService

## Run locally



#### 1. **Install dependencies**



```code
 go mod download
```
####  2. **Run Server**



```shell
 go run ./cmd/server/main.go <host> <port>
```

e.g

```shell
 go run ./cmd/server/main.go localhost 9998
```

#### 3. **Regarding Host/port variables**:




 * If the host or port variable is not passed during execution the server tries to read the ENVIRONMENT vars **HOST** and **PORT**.
 * The environment variables are predefined in [config/.env](https://github.com/l-const/timeService/blob/main/config/.env#L2-L3), as follow: **HOST**=localhost, **PORT**=8080.

e.g running without specifying anything: 


```shell
  go run ./cmd/server/main.go 
```

or alternatively:

```shell
 HOST=localhost PORT=9993 go run ./cmd/server/main.go
```

## Run with Docker


#### 1. Building the image

```shell
  docker build -t timeservice .
```

#### 2. Running the container


How to run:

One simple way using host networking is the following:

```shell
  docker run -d  --netwrok=host timeservice 
```
Now, the service is accessible through **http://localhost:8080**

```shell
curl -X GET  "http://localhost:8080/ptlist?period=1h&tz=Europe/Athens&t1=20210714T204603Z&t2=20210715T123456Z"
```