FROM golang

ENV TZ=Asia/Shanghai

WORKDIR /go/src/123_hao_dai

# Copy the current code into our WORKDIR
COPY . .

RUN go build main.go

EXPOSE 8080

CMD ["./main"]
