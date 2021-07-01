# builder image
FROM golang:latest as builder
LABEL maintainer="Andika Ahmad Ramadhan <andikahmadr@gmail.com>"
RUN mkdir /build /data
WORKDIR /build
COPY go.mod go.sum /build/
RUN apt update -y && apt install libpcap-dev -y
RUN go mod download
COPY *.go /build/
RUN GOOS=linux GOARCH=arm GOARM=7 go build -a -o depulso .
CMD ["/build/depulso"]

#FROM openjdk:8-jre-slim
FROM adoptopenjdk:8-jre-hotspot
WORKDIR /app
RUN apt update && apt install -y libpcap-dev libatomic1
RUN mkdir /data
COPY --from=kmdr7/cicflowmeter:armv7 /src/target/CICFlowMeterV3-0.0.4-SNAPSHOT.jar /app/CICFlowmeter.jar
COPY lib/* /lib/
COPY --from=builder /build/depulso /app/depulso
CMD ["/app/depulso"]
