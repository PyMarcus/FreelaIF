package com.freelaif.personal_chat.services;

import com.freelaif.personal_chat.models.ChatMessage;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.messaging.simp.SimpMessagingTemplate;
import org.springframework.stereotype.Service;

@Service
public class ChatService {
    @Autowired
    private SimpMessagingTemplate messageTemplate;

    /**
     * Envia mensagem para o usuario que estiver no canal: /queue/messages
     * @param message
     */
    public void sendPrivateMessage(ChatMessage message){
        messageTemplate.convertAndSendToUser(
                message.getTo(), "/queue/messages", message
        );
    }
}
