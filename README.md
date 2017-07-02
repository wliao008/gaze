## Usage

```
$ go run main.go
   _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _ _
| |  _ _ _ _ _ _ _ _ _ _  |  _ _ _|  _ _ _ _ _    |_   _  |_   _ _ _|  _ _ _ _  |
| |_ _ _ _ _ _ _  |_   _| |_ _ _ _ _ _|  _ _ _ _| |  _|  _|  _|  _ _ _ _ _ _|   |
|_ _ _ _ _ _ _  |_  |_ _ _ _ _ _  |  _ _|_ _ _ _ _ _| |_  |_ _ _ _ _|  _ _  | | |
|  _ _ _  |   |_  |_ _  |  _ _ _ _|_   _ _ _ _  |_ _ _  | |  _  |  _ _|  _ _ _|_|
|_ _ _  |_ _|_ _|_ _  | |_ _ _  |   |_ _ _   _|_   _  |_ _ _|  _|_ _  |_ _|  _  |
|    _ _|  _ _ _ _ _ _|  _ _  | | |_ _  |_ _|   |_  |_ _ _  |_ _ _  | |  _ _|  _|
| |_|  _ _|  _ _  |  _ _|  _ _|_ _|   |_ _ _ _| | |  _ _ _|_  |  _| | |_ _  |_  |
|   |_ _ _ _|  _ _|_  | |_ _ _  |  _| | |  _ _ _|_ _|  _ _ _ _|  _ _|_ _ _ _ _| |
| |_   _ _ _| |   |_ _ _ _|  _ _| |_ _ _| |_    |   |_ _ _ _  | |_   _ _ _ _  | |
|_ _|_ _ _ _ _ _|_ _ _ _ _ _|_ _ _ _ _ _ _ _ _|_ _|_ _ _ _ _ _|_ _ _|_ _ _ _ _  |
```

## Benchmarks

Benchmarks are ran on the following hardware: 

MacBook Pro 2014, OS X 10.9, 2.6 GHz Intel Core i5, 16 GB 1600 MHz DDR3.

Algorithm | Size | Iterations | ns/ops | ms/op | allocs/op
----------|-------|------|-----------|---------|-----------
Backtracking | 50x25 | 50 | 12935311 | 12.93 | 52
Backtracking | 100x50 | 20 | 51892698 | 51.89 | 102
Backtracking | 1000x500 | 1 | 5309771223 | 5309.77 | 1002
Prim | 50x25 | 2000 | 288071 | 0.28 | 2626
Prim | 100x50 | 500 | 1147424| 1.14 | 10522
Prim | 1000x500 | 3 | 145765352 | 145.76 | 1055684

While Prim is an order of magnitude faster, the memory allocations is still room for much improvement.

## Deployment

Steps on deploying this to the Amazon AWS server.

1. Build the executable.
```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o mazing *.go
```

2. Then build the docker container
```
$ sudo docker build -f Dockerfile.dev -t mazing:dev .
```

3. Tag the container for use with https://hub.docker.com
```
$ docker tag mazing:dev $DOCKER_ID_USER/mazing:latest
```

4. Login to docker hub with the credentials
```
$ docker login
```
5. Push the container to docker hub
```
$ docker push $DOCKER_ID_USER/mazing:latest
```
6. ssh onto the Amazon AWS server

7. Pull down the container from docker hub
```
$ docker pull wliao/mazing:latest
```
8. Finally, run the container in detach mode
```
$ docker run -d -p 80:8080 wliao/mazing:latest
```
http://mazing.cc should now have the latest code running.
```