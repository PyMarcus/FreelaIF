package com.freelaif.personal_chat.controllers;

import com.freelaif.personal_chat.models.ChatMessage;
import com.freelaif.personal_chat.services.ChatService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.messaging.handler.annotation.MessageMapping;
import org.springframework.messaging.simp.annotation.SendToUser;
import org.springframework.stereotype.Controller;

@Controller
public class ChatController {

    @Autowired
    private ChatService chatService;

    /**
     * Cliente envia msg para /app/private-message
     * Entao, recebe de volta para /user/{username}/queue/reply
     * @param message mensagem a ser enviada
     * @return mensagem enviada
     */
    @MessageMapping("/private-message")
    public void sendMessage(ChatMessage message){
        chatService.sendPrivateMessage(message);
    }
}
