FROM node:10 as build-deps
WORKDIR /usr/src/app
COPY uati-react/. ./
RUN npm install
RUN npm run build

FROM nginx
COPY nginx/mysite.template /etc/nginx/conf.d/default.conf
COPY --from=build-deps /usr/src/app/build /data/www
RUN chown -R nginx:nginx /data/www
