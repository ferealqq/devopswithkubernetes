FROM golang:1.18-alpine as build

WORKDIR /app

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o main

FROM scratch as runner 
COPY --from=build /app/main /opt/app/main

CMD [ "/opt/app/main" ] 