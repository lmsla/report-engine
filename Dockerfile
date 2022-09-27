# build stage
FROM golang:alpine AS build-env
ADD . /src
# RUN go env -w GOFLAGS=-buildvcs=false
RUN cd /src && go build -o app

# final stage
FROM alpine
WORKDIR /app
RUN apk update && apk add tzdata
COPY --from=build-env /src/app /app/
ENTRYPOINT ./app