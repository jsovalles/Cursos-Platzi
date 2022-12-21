package com.platzi.market.dao.entity;

import com.fasterxml.jackson.annotation.JsonIgnore;
import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;
import java.util.List;

@Getter
@Setter
@Entity
@Table(name = "categorias")
public class Category {

    @Id
    @Column(name = "id_categoria")
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Integer id;

    @Column(name = "descripcion")
    private String description;

    @Column(name = "estado")
    private Boolean active;

    @JsonIgnore
    @OneToMany(mappedBy = "category")
    private List<Product> products;

}
