package com.rmq;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class RmqController {

    @Autowired
    Producer producer;

    @RequestMapping("/send")
    public String sendMessage(@RequestParam("msg") String msg) {
        System.out.println("*****" + msg);
        producer.produceMsg(msg);

        return "Successfully Msg Sent";
    }
}
