FROM golang:alpine
WORKDIR /app
ENV GOPROXY=https://goproxy.cn,direct
COPY . .
RUN go build -o main . && ls -l
EXPOSE 8090
CMD ["./main"]