# wip file
FROM alpine:3.6

# install nginx
RUN apk update && \
    apk add --no-cache nginx

# document root
ADD app /app
ADD default.conf /etc/nginx/conf.d/default.conf

# port setting
EXPOSE 80

RUN mkdir -p /run/nginx

# run nginx in foreground
CMD nginx -g "daemon off;"