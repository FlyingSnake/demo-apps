package com.example.demo.controller;

import com.example.demo.model.User;
import com.example.demo.repository.UserRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RestController;

import java.time.LocalDateTime;
import java.util.List;
import java.util.Random;
import java.util.logging.Logger;

@RestController
public class DemoController {

    private static final Logger logger = Logger.getLogger(DemoController.class.getName());

    @Autowired
    private UserRepository userRepository;

    @GetMapping("/")
    public String hello() {
        return "Hello World (Java / Springboot)";
    }

    @GetMapping("/sleep/{seconds}")
    public String sleep(@PathVariable int seconds) throws InterruptedException {
        Thread.sleep(seconds * 1000);
        return "sleep " + seconds + "s";
    }

    @GetMapping("/status/random")
    public ResponseEntity<String> randomStatus() {
        int[] statusCodes = {200, 201, 202, 204, 400, 401, 403, 404, 500, 501, 502, 503};
        String[] statusMessages = {"OK", "Created", "Accepted", "No Content", "Bad Request", "Unauthorized", "Forbidden", "Not Found", "Internal Server Error", "Not Implemented", "Bad Gateway", "Service Unavailable"};

        Random random = new Random();
        int index = random.nextInt(statusCodes.length);

        int statusCode = statusCodes[index];
        String statusMessage = statusMessages[index];

        return ResponseEntity.status(statusCode).body(statusMessage);
    }

    @GetMapping("/exception")
    public String exception() {
        String errMessage = "Internal Server Error - Manual Exception";
        logger.severe(errMessage);
        throw new RuntimeException(errMessage);
    }
}
