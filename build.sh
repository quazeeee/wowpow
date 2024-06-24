
docker build --no-cache -t powservice -f server.Dockerfile .
docker build --no-cache -t powclinet -f client.Dockerfile .