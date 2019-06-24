### Basic Messaging with Spring boot & RabbitMQ

#### Prerequisites
* Java
* Docker & docker-compose
* Maven


#### Run the project

`docker-compose up`

`mvn clean package && mvn spring-boot:run`

Go to http://localhost:15672 to access rabbitMQ GUI (login/password: guest/guest)




send message with curl
curl -d "msg=hello" -X POST http://localhost:8080/send
