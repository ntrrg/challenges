FROM golang:1.12-alpine3.9
WORKDIR /src
COPY . .
ENTRYPOINT ["hrrun.sh"]

