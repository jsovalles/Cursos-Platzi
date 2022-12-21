package com.platzi.market.dao.entity;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;
import java.time.LocalDateTime;
import java.util.List;

@Getter
@Setter
@Entity
@Table(name = "compras")
public class Purchase {

    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    @Column(name = "id_compra")
    private Integer purchaseId;
    @Column(name = "id_cliente")
    private String clientId;
    @Column(name = "fecha")
    private LocalDateTime date;
    @Column(name = "medio_pago")
    private String paymentMethod;
    @Column(name = "comentario")
    private String comment;
    @Column(name = "estado")
    private String status;
    @ManyToOne
    @JoinColumn(name = "id_cliente", insertable = false, updatable = false)
    private Client client;
    @OneToMany(mappedBy = "purchase", cascade = CascadeType.ALL)
    private List<PurchaseItem> purchases;

}
