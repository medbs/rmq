FROM openjdk:8u111-jdk-alpine
#RUN apk add --update \
#    curl \
#    && rm -rf /var/cache/apk/*
COPY ./target/rmq-1.jar /usr/
WORKDIR /usr/
ENTRYPOINT ["java","-jar","rmq-1.jar"]
#EXPOSE 8080
