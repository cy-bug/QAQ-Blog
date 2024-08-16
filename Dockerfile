FROM golang:alpine
WORKDIR /app
ENV GOPROXY=https://goproxy.cn,direct
COPY . .
EXPOSE 8090
CMD ['./main']