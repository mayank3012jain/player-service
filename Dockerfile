FROM alpine:3.14

RUN mkdir -p /opt/
WORKDIR /opt/
COPY build/linux/playerService ./playerService
COPY resources/players.csv ./players.csv

RUN chmod +x ./playerService

EXPOSE 8080

ENTRYPOINT ./playerService