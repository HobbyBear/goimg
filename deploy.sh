docker stop imgserver

docker rm imgserver

# shellcheck disable=SC2006
# shellcheck disable=SC2046
docker rmi imgserver

docker build -t imgserver .

docker-compose up -d