echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker push gobjserver/gobjserver
docker push gobjserver/gobjserver:0.0.0