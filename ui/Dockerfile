FROM ubuntu:16.04 as dev

RUN apt-get update && \
  apt-get install -y  curl && \
  curl -fsSL https://deb.nodesource.com/setup_14.x | bash - && \
  apt-get install -y nodejs

RUN mkdir -p /app

COPY . /app/ui
WORKDIR /app/ui

RUN npm install && npm install -g expo-cli
CMD [ "npm", "start" ]