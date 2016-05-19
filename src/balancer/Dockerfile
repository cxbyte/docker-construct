FROM cxbyte/supervisor

MAINTAINER Andrey Petrov cxbyte@mail.ru

RUN apt-get update && apt-get install nginx -y
RUN echo "\ndaemon off;" >> /etc/nginx/nginx.conf

RUN echo "\n[program:nginx]" >> /etc/supervisor/conf.d/supervisord.conf
RUN echo "\ncommand=/usr/sbin/nginx" >> /etc/supervisor/conf.d/supervisord.conf

# Nginx apps config
COPY nginx.conf /etc/nginx/sites-enabled/nginx.conf

EXPOSE 80
