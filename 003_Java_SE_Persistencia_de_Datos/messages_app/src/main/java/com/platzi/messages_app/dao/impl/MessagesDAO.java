package com.platzi.messages_app.dao.impl;

import com.platzi.messages_app.dao.IMessagesDAO;
import com.platzi.messages_app.dao.model.Message;
import com.platzi.messages_app.utils.ConnectionDatabase;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;

import java.sql.Connection;
import java.sql.PreparedStatement;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.util.ArrayList;
import java.util.List;

@Repository
public class MessagesDAO implements IMessagesDAO {

    @Autowired
    ConnectionDatabase connectionDatabase;


    @Override
    public void createMessage(Message message) {
        try {
            Connection connection = connectionDatabase.getConnection();
            PreparedStatement preparedStatement;
            String query = "INSERT INTO mensajes (mensaje, autor_mensaje) VALUES (?,?)";
            preparedStatement = connection.prepareStatement(query);
            preparedStatement.setString(1, message.getMessage());
            preparedStatement.setString(2, message.getMessageAuthor());
            preparedStatement.executeUpdate();
        } catch (SQLException throwables) {
            throwables.printStackTrace();
        }

    }

    @Override
    public List<Message> readMessages() {
        List<Message> messages = new ArrayList<>();
        try {
            Connection connection = connectionDatabase.getConnection();
            PreparedStatement preparedStatement;
            String query = "SELECT * FROM mensajes";
            preparedStatement = connection.prepareStatement(query);
            ResultSet result = preparedStatement.executeQuery();

            while (result.next()) {
                Message message = new Message();
                message.setMessageId(result.getInt("id_mensaje"));
                message.setMessage(result.getString("mensaje"));
                message.setMessageAuthor(result.getString("autor_mensaje"));
                message.setMessageDate(result.getString("fecha_mensaje"));
                messages.add(message);
            }
        } catch (SQLException throwables) {
            throwables.printStackTrace();
        }
        return messages;
    }

    @Override
    public void updateMessage(Message message) {
        try {
            Connection connection = connectionDatabase.getConnection();
            PreparedStatement preparedStatement;
            String query = "UPDATE mensajes SET mensaje = ? WHERE id_mensaje = ?";
            preparedStatement = connection.prepareStatement(query);
            preparedStatement.setString(1, message.getMessage());
            preparedStatement.setInt(2, message.getMessageId());
            preparedStatement.executeUpdate();
        } catch (SQLException throwables) {
            throwables.printStackTrace();
        }
    }

    @Override
    public void deleteMessage(Integer messageId) {
        try {
            Connection connection = connectionDatabase.getConnection();
            PreparedStatement preparedStatement;
            String query = "DELETE FROM mensajes WHERE id_mensaje = ?";
            preparedStatement = connection.prepareStatement(query);
            preparedStatement.setInt(1, messageId);
            preparedStatement.executeUpdate();

        } catch (SQLException throwables) {
            throwables.printStackTrace();
        }
    }
}
