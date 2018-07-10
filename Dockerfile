FROM golang

ENV TZ=Asia/Shanghai

WORKDIR /go/src/123_hao_dai_be

# Copy the current code into our WORKDIR
COPY . .

RUN go get

RUN go build main.go

EXPOSE 8080

CMD ["./main"]
