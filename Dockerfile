# build stage
FROM golang:1.18.6-bullseye AS build-env
ADD . /src
RUN timedatectl set-timezone Asia/Taipei
RUN apt-get update -y
RUN apt-get install -y tzdata
RUN apt-get install -y xfonts-75dpi
RUN apt install -y /src/wkhtmltox_0.12.6.1-2.bullseye_amd64.deb
RUN apt install -y /src/google-chrome-stable_current_amd64.deb
RUN rm -f /src/wkhtmltox_0.12.6.1-2.bullseye_amd64.deb
RUN rm -f /src/google-chrome-stable_current_amd64.deb
RUN apt-get install fonts-wqy-microhei ttf-wqy-microhei fonts-wqy-zenhei ttf-wqy-zenhei
RUN fc-cache -f -v
RUN ldconfig
RUN cd /src && go build -buildvcs=false -o app
ENTRYPOINT /src/app

