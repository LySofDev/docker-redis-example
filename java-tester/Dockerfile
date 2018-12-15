FROM maven:3.6.0-jdk-8-alpine
RUN mkdir /java
COPY . /java/
WORKDIR /java/bid.vicero.redis.example
RUN mvn clean install
