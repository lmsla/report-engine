# build stage
FROM golang:1.22.0-bullseye AS build-env

# 複製源碼
ADD . /src

RUN cd /src && go mod tidy && go build -o report_backend


# 編譯 Go 應用程式
RUN go build -buildvcs=false -o report_backend

# final stage
FROM golang:1.22.0-bullseye

# RUN rm /etc/localtime
# RUN ln -s /usr/share/zoneinfo/Asia/Taipei /etc/localtime
RUN apt-get update -y
RUN apt-get install -y tzdata wget vim lsb-release
RUN apt-get install -y xfonts-75dpi
RUN apt-get install -y fonts-wqy-microhei ttf-wqy-microhei


# 設置時區
RUN ln -fs /usr/share/zoneinfo/Asia/Taipei /etc/localtime && \
    dpkg-reconfigure --frontend noninteractive tzdata


RUN apt install -y /src/google-chrome-stable_current_amd64.deb
RUN rm -f /src/google-chrome-stable_current_amd64.deb


# 複製編譯好的應用程式
COPY --from=build-env /src/report_backend /app/

WORKDIR /app
RUN mkdir /app/files
RUN mkdir /app/files/screenshot_files
RUN mkdir /app/files/html_files
RUN mkdir /app/files/report_files
RUN mkdir /app/files/logo
RUN mkdir /app/log_record

RUN fc-cache -f -v
RUN ldconfig

ENTRYPOINT /app/report_backend

 


 docker run --name mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=1qaz2wsx -d mysql:8.0.1