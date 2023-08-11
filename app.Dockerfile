FROM node:16.14 as build-stage

WORKDIR /app

COPY . .

RUN yarn

RUN yarn build:prod

CMD [ "yarn", "dev" ]

FROM node:16.14 as build-release-stage

WORKDIR /app

COPY --from=build-stage /app /app

CMD [ "yarn", "preview" ]