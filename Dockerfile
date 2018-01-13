FROM debian:8
ADD dialogflow /dialogflow
ADD .env /.env
ENTRYPOINT ["/dialogflow"]
