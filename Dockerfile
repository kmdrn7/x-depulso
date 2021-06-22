# builder image
FROM golang:latest as builder
LABEL maintainer="Andika Ahmad Ramadhan <andikahmadr@gmail.com>"
RUN mkdir /build /data
WORKDIR /build
COPY go.mod go.sum /build/
RUN apt update -y && apt install libpcap-dev -y
RUN go mod download
COPY *.go /build/
RUN GOOS=linux go build -a -o depulso .
CMD ["/build/depulso"]

FROM openjdk:8-jre-slim
WORKDIR /app
RUN apt update && apt install -y libpcap-dev
RUN mkdir /data
COPY --from=kmdr7/cicflowmeter /src/target/CICFlowMeterV3-0.0.4-SNAPSHOT.jar /app/CICFlowmeter.jar
COPY --from=kmdr7/cicflowmeter /src/jnetpcap-1.4.r1425/*.so /lib64/
COPY --from=builder /build/depulso /app/depulso
CMD ["/app/depulso"]