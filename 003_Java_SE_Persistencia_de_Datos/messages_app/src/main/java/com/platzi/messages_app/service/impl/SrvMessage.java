package com.platzi.messages_app.service.impl;

import com.platzi.messages_app.dao.IMessagesDAO;
import com.platzi.messages_app.dao.model.Message;
import com.platzi.messages_app.service.ISrvMessage;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

@Service
public class SrvMessage implements ISrvMessage {

    @Autowired
    IMessagesDAO messagesDAO;

    @Override
    public void createMessage(Message message) {
        messagesDAO.createMessage(message);
    }

    @Override
    public List<Message> readMessages() {
        return messagesDAO.readMessages();
    }

    @Override
    public void updateMessage(Message message) {
        messagesDAO.updateMessage(message);
    }

    @Override
    public void deleteMessage(Integer messageId) {
        messagesDAO.deleteMessage(messageId);
    }
}
