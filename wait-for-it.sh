#!/bin/sh
# wait-for-it.sh
GOLANG_HOST=$1
GOLANG_MANAGEMENT_PORT=8080
shift
cmd="$@"



while ! curl -s http://${GOLANG_HOST}:8080/health; do
  >&2 echo "Waiting for Golang Service to become healthy..."
  sleep 5
done

>&2 echo "Golang Service is healthy - starting API Gateway"

exec $cmd


