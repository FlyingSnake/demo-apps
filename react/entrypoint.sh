#!/bin/sh

# Replace placeholders in nginx.conf.template with environment variable values
envsubst '${JAVA_SERVICE} ${GOLANG_SERVICE} ${NODEJS_SERVICE} ${PYTHON_SERVICE} ${DOTNET_SERVICE} ${PHP_SERVICE}' < /etc/nginx/nginx.conf.template > /etc/nginx/nginx.conf

# Execute CMD
exec "$@"