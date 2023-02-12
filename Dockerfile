FROM golang:1.20.0

WORKDIR /app

COPY . .

ARG stakefish_api_version
ENV STAKEFISH_API_VERSION=$stakefish_api_version

RUN mkdir /app/bin && \
    go mod download && \
    go build -o /app/bin/stakeapi /app/cmd/main

ENTRYPOINT ["/app/bin/stakeapi"]