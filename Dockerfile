FROM golang:1.20 as builder

ENV PATH /etc/community_grpc

WORKDIR "$PATH"
COPY . .

RUN go mod download
RUN go mod verify
RUN go build -o community



FROM golang:1.20

ENV PATH /etc/community_grpc
RUN mkdir -p "$PATH"
WORKDIR "$PATH"

COPY . .
COPY --from=builder "$PATH"/community $PATH

EXPOSE 5001
CMD ["./community"]