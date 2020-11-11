FROM alpine:latest

ARG SERVICE_NAME="txsigner"
ARG DAEMON_PRIVATE_KEY="what3ver"
ARG PORT=8080

ENV PATH=/bin
ENV PORT=${PORT}
ENV SERVICE_NAME=${SERVICE_NAME}
ENV DAEMON_PRIVATE_KEY=${DAEMON_PRIVATE_KEY}

COPY build/bin/$SERVICE_NAME /usr/bin
COPY build/docker/entrypoint.sh /

EXPOSE ${PORT}

CMD [ "/entrypoint.sh" ]