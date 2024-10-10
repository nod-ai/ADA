FROM busybox

ENV UPDATED="2024-09-24"
ENV DESCRIPTION="Redfish Event Listener/Exporter"
ENV LISTENER_IP="127.0.0.1"
ENV LISTENER_PORT="8080"
ENV USE_SSL="false"
ENV CERTFILE="path/to/certfile"
ENV KEYFILE="path/to/keyfile"

# At least these need to be configured when running this image
ENV REDFISH_SERVERS=
ENV SUBSCRIPTION_PAYLOAD=

COPY amd-redfish-exporter /bin/redfish_exporter

EXPOSE      8080
EXPOSE      2112
USER        nobody
ENTRYPOINT  [ "/bin/redfish_exporter" ]
