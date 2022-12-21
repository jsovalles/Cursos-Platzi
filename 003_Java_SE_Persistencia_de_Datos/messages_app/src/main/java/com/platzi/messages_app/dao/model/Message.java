package com.platzi.messages_app.dao.model;

import lombok.Data;

@Data
public class Message {

    private int messageId;
    private String message;
    private String messageAuthor;
    private String messageDate;
}
