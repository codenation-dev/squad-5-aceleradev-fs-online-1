FROM nginx

COPY nginx/mysite.template /etc/nginx/conf.d/default.conf
COPY uati-react/build /data/www

RUN chown -R nginx:nginx /data/www
