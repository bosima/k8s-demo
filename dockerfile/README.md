## Print 'Hello docker'

```
go build hellodocker.go

docker build -f Dockerfile.hellodocker -t hellodocker:1.0 .

docker run hellodocker:1.0
```

## Print 'Hello {Arg}'

```
go build hellodocker-args.go

docker build -f Dockerfile.hellodocker-args -t hellodocker:2.0 .
docker run -e hellodocker:2.0
docker run -e greetname="bosima" hellodocker:2.0

docker build --build-arg greetname=world  -f Dockerfile.hellodocker-args -t hellodocker:2.0 .
docker run -e hellodocker:2.0
docker run -e greetname="bosima" hellodocker:2.0
```

## Print 'Hello docker' using http server

```
go build hellodocker-server.go

docker build --build-arg port=8080 -f Dockerfile.hellodocker-server -t hellodocker:3.0 .
docker run  -p 8080:8080 -it --rm hellodocker:3.0

docker build -f Dockerfile.hellodocker-server -t hellodocker:3.0 .
docker run -e port=8080 -p 8080:8080 -it --rm hellodocker:3.0
```