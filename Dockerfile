FROM golang:1-alpine AS build

RUN apk update && apk add make git gcc musl-dev

ADD . /app/src/github_actions

WORKDIR /app/src/github_actions

RUN mv github_actions /github_actions

FROM alpine:latest

ARG SERVICE

ENV APP=github_actions

RUN apk add --no-cache ca-certificates && mkdir /app
COPY --from=build /github_actions /app/github_actions

ENTRYPOINT exec /app/hello
