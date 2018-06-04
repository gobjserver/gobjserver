CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main.out .
docker build -t gobjserver/gobjserver .
docker tag  gobjserver/gobjserver gobjserver/gobjserver:0.0.0