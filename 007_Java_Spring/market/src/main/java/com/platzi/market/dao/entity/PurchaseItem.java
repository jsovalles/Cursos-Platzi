package com.platzi.market.dao.entity;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.*;
import java.math.BigDecimal;

@Getter
@Setter
@Entity
@Table(name = "compras_productos")
public class PurchaseItem {

    @EmbeddedId
    private PurchaseItemPK id;
    @Column(name = "cantidad")
    private Integer quantity;
    private BigDecimal total;
    @Column(name = "estado")
    private Boolean status;
    @ManyToOne
    @MapsId(value = "purchaseId")
    @JoinColumn(name = "id_compra", insertable = false, updatable = false)
    private Purchase purchase;
    @ManyToOne
    @JoinColumn(name = "id_producto", insertable = false, updatable = false)
    private Product product;

}
