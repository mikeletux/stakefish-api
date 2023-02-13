FROM golang:1.19.0 as build

ENV CGO_ENABLED=0
WORKDIR /app
COPY . .

RUN mkdir /app/bin && \
    go mod download && \
    go build -o /app/bin/stakeapi /app/cmd/main


FROM gcr.io/distroless/static-debian11

ARG stakefish_api_version
ENV STAKEFISH_API_VERSION=$stakefish_api_version

COPY --from=build /app/bin/stakeapi /
CMD ["/stakeapi"]
