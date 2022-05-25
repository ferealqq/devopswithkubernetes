FROM golang:1.18-alpine as build

WORKDIR /app

COPY . ./
RUN go mod download

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o server
RUN ./build_wasm

FROM scratch as runner 
COPY --from=build /app/server /opt/app/server
COPY --from=build /app/static /opt/app/static

WORKDIR /opt/app

ENV PORT=$PORT
ENV GIN_MODE=release

EXPOSE $PORT

CMD [ "./server" ]