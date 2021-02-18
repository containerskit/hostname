FROM golang:1.15-alpine as build

WORKDIR /src
ADD go.mod .
RUN go mod download
RUN apk add upx
ADD *.go ./
RUN CGO_ENABLED=0 go build -ldflags="-s -w"
RUN upx hostname

FROM scratch
COPY --from=build /src/hostname /bin/hostname
ENTRYPOINT [ "/bin/hostname", "/run/config/hostname" ]
