package com.platzi.messages_app.dao;

import com.platzi.messages_app.dao.model.Message;

import java.util.List;

public interface IMessagesDAO {

    void createMessage(Message message);

    List<Message> readMessages();

    void updateMessage(Message message);

    void deleteMessage(Integer messageId);


}
