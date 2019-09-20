FROM golang:1.13.0-alpine AS build-env
ADD . /build
RUN cd /build && go build -mod=vendor -o app

# final stage
FROM alpine
WORKDIR /opt/service
COPY --from=build-env /build/app /opt/service/
ENTRYPOINT ./app