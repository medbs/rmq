package com.rmq;

import org.springframework.amqp.rabbit.core.RabbitTemplate;
import org.springframework.boot.CommandLineRunner;
import org.springframework.stereotype.Component;

import java.util.Scanner;
import java.util.concurrent.TimeUnit;


@Component
public class Producer implements CommandLineRunner {


    private final RabbitTemplate rabbitTemplate;
    private final Consumer receiver;

    public Producer(Consumer receiver, RabbitTemplate rabbitTemplate) {
        this.receiver = receiver;
        this.rabbitTemplate = rabbitTemplate;
    }

    @Override
    public void run(String... args) throws Exception {

        Scanner scanner = new Scanner(System.in);

        System.out.println("Write a message");

        String message = scanner.nextLine();

        System.out.println("Message sent:" + message);
        rabbitTemplate.convertAndSend(RmqApplication.topicExchangeName, "rmq.key.mq", message);
        receiver.getLatch().await(10000, TimeUnit.MILLISECONDS);
    }
}
