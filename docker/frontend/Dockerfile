FROM node:17.5-alpine

WORKDIR /frontend

RUN apk update && apk add git

COPY frontend/package*.json .

RUN yarn install

CMD ["yarn", "dev"]