# builder image
FROM golang:1.16.5 as builder
LABEL maintainer="Andika Ahmad Ramadhan <andikahmadr@gmail.com>"
RUN mkdir /build
WORKDIR /build
COPY go.mod go.sum /build/
RUN apt update -y && apt install libpcap-dev -y
RUN GOMAXPROCS=1 go mod download
COPY *.go /build/
RUN GOMAXPROCS=1 GOOS=linux GOARCH=arm GOARM=7 go build -a -o depulso .

# target image
FROM adoptopenjdk:8-jre-hotspot
ARG TARGETPLATFORM
WORKDIR /app
RUN apt update && apt install -y libpcap-dev libatomic1
RUN mkdir /data
COPY --from=kmdr7/cicflowmeter:1.0 /app/CICFlowmeter.jar /app/CICFlowmeter.jar
COPY lib/$TARGETPLATFORM/* /lib/
COPY --from=builder /build/depulso /app/depulso
CMD ["/app/depulso"]
