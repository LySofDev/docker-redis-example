FROM node:8.12.0-alpine
RUN mkdir /node
COPY package*.json /node/
WORKDIR /node
RUN npm i
COPY . /node/
