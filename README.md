### Basic Messaging with Spring boot, Go & RabbitMQ

#### Prerequisites
* Docker & docker-compose

#### Run the project

Start RabbitMQ
`docker-compose up rabbitmq`

Start Producer
`docker-compose up producer`

Start Consumer
`docker-compose up consumer`

Send message
`curl -d "msg=message" -X POST http://localhost:8080/send`
