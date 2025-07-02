package com.freelaif.personal_chat.controllers;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.GetMapping;

@Controller
public class PageController {
    @GetMapping("/chat")
    public String chat(){
        return "chat";
    }
}
