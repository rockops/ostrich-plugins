package com.example.demo.controller;

import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RestController;

import java.util.Random;

@RestController
public class RandomController {

    private final Random random = new Random();

    @GetMapping("/random")
    public int getRandom() {
        return random.nextInt(101); // Returns random integer between 0 and 100
    }
}
