FROM alpine:3.7
WORKDIR /app

ADD consignment.json .
ADD consignment-cli .
RUN chmod +x /app/consignment-cli

CMD ["/app/consignment-cli"]
