# build stage
FROM golang:1.23.0-bullseye AS build-env
# FROM golang:alpine AS build-env
ADD . /src
# RUN cd /src && go mod tidy && go build -o report_backend


RUN rm /etc/localtime
RUN ln -s /usr/share/zoneinfo/Asia/Taipei /etc/localtime
RUN apt-get update -y
RUN apt-get install -y tzdata wget vim lsb-release
RUN apt-get install -y xfonts-75dpi


# RUN apt install mysql-client
# RUN apt install libmysqlclient-dev

# RUN apt install -y /src/wkhtmltox_0.12.6.1-2.bullseye_amd64.deb
# RUN rm -f /src/wkhtmltox_0.12.6.1-2.bullseye_amd64.deb

RUN apt install -y /src/google-chrome-stable_current_amd64.deb
RUN rm -f /src/google-chrome-stable_current_amd64.deb
RUN apt-get install fonts-wqy-microhei ttf-wqy-microhei

# FROM alpine
WORKDIR /app
RUN mkdir /app/files
RUN mkdir /app/files/screenshot_files
RUN mkdir /app/files/html_files
RUN mkdir /app/files/report_files
RUN mkdir /app/files/logo
RUN mkdir /app/log_record


# RUN apt-get install fonts-wqy-microhei ttf-wqy-microhei fonts-wqy-zenhei ttf-wqy-zenhei
RUN fc-cache -f -v
RUN ldconfig


RUN cd /src && go build -buildvcs=false -o report_backend
ENTRYPOINT /src/report_backend
# fonts-cwtex-yen
 