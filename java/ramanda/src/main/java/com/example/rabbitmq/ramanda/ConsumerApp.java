package com.example.rabbitmq.ramanda;

import com.rabbitmq.client.*;

public class ConsumerApp {

    public static void main(String[] args) throws Exception{
        ConnectionFactory factory = new ConnectionFactory();
        factory.setUri("amqp://guest:guest@localhost:5672/");
        factory.setVirtualHost("/");

        Connection connection = factory.newConnection();
        Channel channel = connection.createChannel();

        DeliverCallback deliverCallback = (consumerTag, message) -> {
            System.out.println(message.getEnvelope().getRoutingKey());
            System.out.println(new String(message.getBody()));
        };

        CancelCallback cancelCallback = consumerTag -> {
            System.out.println("Consumer is canceled");
        };

        channel.basicConsume("whatsapp", true, deliverCallback, cancelCallback);
//
//        channel.close();
//        connection.close();
    }

}
