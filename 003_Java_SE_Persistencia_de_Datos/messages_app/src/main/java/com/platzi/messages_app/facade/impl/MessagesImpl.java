package com.platzi.messages_app.facade.impl;

import com.platzi.messages_app.dao.model.Message;
import com.platzi.messages_app.facade.IMessages;
import com.platzi.messages_app.service.ISrvMessage;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
public class MessagesImpl implements IMessages {

    @Autowired
    ISrvMessage srvMessage;

    @Override
    @PostMapping
    public void createMessage(@RequestBody Message message) {
        srvMessage.createMessage(message);
    }

    @Override
    @GetMapping
    public List<Message> readMessages() {
        return srvMessage.readMessages();
    }

    @Override
    @PutMapping
    public void updateMessage(@RequestBody Message message) {
        srvMessage.updateMessage(message);
    }

    @Override
    @DeleteMapping("/{message-id}")
    public void deleteMessage(@PathVariable(value = "message-id") Integer messageId) {
        srvMessage.deleteMessage(messageId);
    }
}
