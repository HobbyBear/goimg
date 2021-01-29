docker stop imgserver

docker rm imgserver

# shellcheck disable=SC2006
# shellcheck disable=SC2046
docker rmi `docker images | grep imgserver | awk '{print $3}'`


docker-compose up