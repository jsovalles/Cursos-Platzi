package com.platzi.messages_app.facade;

import com.platzi.messages_app.dao.model.Message;

import java.util.List;

public interface IMessages {

    void createMessage(Message message);

    List<Message> readMessages();

    void updateMessage(Message message);

    void deleteMessage(Integer messageId);
}
