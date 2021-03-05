# wip file
FROM alpine:latest

# install nginx
RUN apk update && \
    apk add --no-cache nginx

COPY entrypoint.sh /var/tmp
CMD bash -E /var/tmp/entrypoint.sh && /bin/bash