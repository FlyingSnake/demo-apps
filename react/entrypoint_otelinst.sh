#!/bin/sh
if [ -z "$OTEL_EXPORTER_OTLP_ENDPOINT" ]; then
  export NGINX_MODULE_ENABLED="OFF"
else
  export NGINX_MODULE_ENABLED="ON"
fi
export NGINX_MODULE_SERVICE_NAMESPACE=$(echo "$OTEL_RESOURCE_ATTRIBUTES" | tr ',' '\n' | grep 'k8s.namespace.name' | cut -d '=' -f 2)
export NGINX_MODULE_RESOLVE_BACKENDS="${NGINX_MODULE_RESOLVE_BACKENDS:-ON}"
export NGINX_MODULE_TRACE_AS_ERROR="${NGINX_MODULE_TRACE_AS_ERROR:-OFF}"

envsubst '${NGINX_MODULE_ENABLED} ${OTEL_EXPORTER_OTLP_ENDPOINT} ${OTEL_TRACES_EXPORTER} ${NGINX_MODULE_RESOLVE_BACKENDS} ${OTEL_RESOURCE_ATTRIBUTES_POD_NAME} ${OTEL_SERVICE_NAME} ${NGINX_MODULE_SERVICE_NAMESPACE} ${NGINX_MODULE_TRACE_AS_ERROR}' < /etc/nginx/conf.d/opentelemetry_agent.conf.template > /etc/nginx/conf.d/opentelemetry_agent.conf
envsubst '${JAVA_SERVICE} ${GOLANG_SERVICE} ${NODEJS_SERVICE} ${PYTHON_SERVICE} ${DOTNET_SERVICE} ${PHP_SERVICE}' < /etc/nginx/nginx.conf.template > /etc/nginx/nginx.conf

# Execute CMD
exec "$@"