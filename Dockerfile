FROM alpine:latest

#RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
#        apk add --no-cache --update tzdata ca-certificates

WORKDIR /app
ADD ./bin/ /app/
CMD ["/app/main"]
EXPOSE 8001 8002
