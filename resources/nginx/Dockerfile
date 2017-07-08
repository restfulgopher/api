FROM nginx:1.13.1-alpine

MAINTAINER Alessandro Resta <mail@alesr.me> @alesresta

# Remove default nginx
RUN rm -v /etc/nginx/nginx.conf

# Add custom nginx config
ADD ./resources/nginx/nginx.conf /etc/nginx/nginx.conf

# Add self-signed certifcate and key
ADD ./resources/nginx/ssl/crt/server.crt /etc/nginx/ssl/crt/server.crt
ADD ./resources/nginx/ssl/key/server.key /etc/nginx/ssl/key/server.key

EXPOSE 443

ENTRYPOINT ["nginx"]
