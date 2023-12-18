#!/bin/sh
# wait-for-it.sh
RABBITMQ_HOST=$1
RABBITMQ_MANAGEMENT_PORT=15672
shift
cmd="$@"



while ! curl -s http://${RABBITMQ_HOST}:${RABBITMQ_MANAGEMENT_PORT}/api/healthchecks/node; do
  >&2 echo "Waiting for RabbitMQ to become healthy..."
  sleep 5
done

>&2 echo "RabbitMQ is healthy - starting go service"

exec $cmd


