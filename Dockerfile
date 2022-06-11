####### build CONTAINER #######
FROM golang:alpine AS build
WORKDIR /go/src
COPY ./ .

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN apk add git tzdata ca-certificates
# RUN git config --global url."https://{{GitHubPersonalAccessToken for PrivateRepository}}:x-oauth-basic@github.com/".insteadOf "https://github.com/"
RUN go get -d -v ./...
# RUN go test -v ./...
RUN go build -a -installsuffix cgo -o demo-echo-v4 .

####### production CONTAINER #######
FROM scratch AS runtime

COPY --from=build /go/src/demo-echo-v4 ./

COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENV TZ=Asia/Tokyo

EXPOSE 1323/tcp
ENTRYPOINT ["./demo-echo-v4"]
