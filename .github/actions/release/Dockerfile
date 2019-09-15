FROM alpine:latest

RUN	apk add --no-cache \
  bash \
  ca-certificates \
  curl \
  jq

COPY upload-to-release /bin/upload-to-release

RUN chmod +x /bin/upload-to-release

ENTRYPOINT ["/bin/upload-to-release"]
