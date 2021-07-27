FROM golang:1.16.6-alpine AS GO_BUILD
ADD ./service /service/service
WORKDIR /service/service
RUN go build -o /go/bin/service/service

FROM nickgryg/alpine-pandas
RUN pip3 install mpu
# RUN pip3 install pandas
# WORKDIR /usr/src/app
ADD ./scripts /usr/src/app/scripts
COPY --from=GO_BUILD /go/bin/service/service /usr/src/app/service/service
EXPOSE 8081
WORKDIR /usr/src/app/service
CMD ./service