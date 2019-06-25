FROM openjdk:8u111-jdk-alpine
COPY ./target/rmq-1.jar /usr/
WORKDIR /usr/
ENTRYPOINT ["java","-jar","rmq-1.jar"]
