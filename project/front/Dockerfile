FROM golang:1.18-alpine as build

WORKDIR /app

RUN apk --update add --no-cache ca-certificates openssl git tzdata && \
update-ca-certificates

COPY . ./
RUN go mod download

WORKDIR /app/front

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o server
RUN ./build_wasm

FROM scratch as runner 
COPY --from=build /app/front/server /opt/app/server
COPY --from=build /app/front/static /opt/app/static
# To enable downloading a image from the https://picsum.photos/1200
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /opt/app

ENV GIN_MODE=release

CMD [ "./server" ]