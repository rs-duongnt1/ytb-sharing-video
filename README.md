# Funny Movie

## Introduction:

Funny Movie is a simple project for sharing videos from youtube another users.

## Demo

Web: http://rsdn.site

Api: http://api.rsdn.site

Websocket: ws://websocket.rsdn.site/ws

## Key features:

- User registration and login
- Sharing YouTube videos
- Viewing a list of shared videos realtime
- Real-time notifications for new video shares

## Prerequisites

- docker
- nodejs >= 16
- golang 1.19
- protobuf (https://buf.build/)
- postgresql

## Installation & Configuration

Clone repository

```sh
git clone https://github.com/rs-duongnt1/youtube-video-share.git && cd youtube-video-share
```

Install dependencies

```sh
yarn install
```

Create file .env from .env.example

```sh
cp .env.example .env
```

> You can edit or update app or database info inside file .env

Create a docker network & docker volume

```sh
docker network create ytb-network --attachable
docker volume create data
```

## Running the Application

```sh
docker compose up
```

After docker start success, navigate to http://localhost:3000/. The app will automatically reload if you change any of the source files.

## Unit testing & integration testing

App have used `cypress` for unit testing and integration testing. You can access to cypress UI with command and start test:

```sh
npx cypress open
```

## Docker Deployment

Login to docker hub (https://hub.docker.com/)

```sh
docker login
```

Build an image

```sh
docker build -t duongnt/ytb_sharing:1.0 -f app.Dockerfile .
docker build -t duongnt/ytb_sharing_api:1.0 -f Dockerfile .
```

Push to deploy docker hub

```sh
docker push duongnt/ytb_sharing:1.0
docker push duongnt/ytb_sharing_api:1.0
```

You can change the docker tag name `duongnt/ytb_sharing:1.0` & `duongnt/ytb_sharing_api:1.0` for your tag in docker hub.

## Usage

It's really simple, you just need to visit the web and experience.

## Troubleshooting

N/A
