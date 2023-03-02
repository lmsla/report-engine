# build stage
FROM golang:1.18.6-bullseye AS build-env
ADD . /src
RUN apt-get update -y
RUN apt-get install -y tzdata
RUN apt-get install -y xfonts-75dpi
RUN apt install -y /src/wkhtmltox_0.12.6.1-2.bullseye_amd64.deb
RUN ldconfig
RUN cd /src && go build -buildvcs=false -o app
COPY --from=build-env /src/migrate/migrations /app/migrate/migrations
ENTRYPOINT /src/app

