package com.platzi.market.dao.entity;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;
import java.util.List;

@Getter
@Setter
@Entity
@Table(name = "clientes")
public class Client {

    @Id
    private String id;

    @Column(name = "nombre")
    private String firstName;

    @Column(name = "apellidos")
    private String lastName;

    @Column(name = "celular")
    private Long cellphone;

    @Column(name = "direccion")
    private String address;

    @Column(name = "correo_electronico")
    private String email;

    @OneToMany(mappedBy = "client")
    private List<Purchase> purchases;
}
