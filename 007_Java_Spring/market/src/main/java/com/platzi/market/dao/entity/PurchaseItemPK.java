package com.platzi.market.dao.entity;

import lombok.Getter;
import lombok.Setter;

import javax.persistence.Column;
import javax.persistence.Embeddable;
import java.io.Serializable;

@Getter
@Setter
@Embeddable
public class PurchaseItemPK implements Serializable {
    @Column(name = "id_compra")
    private Integer purchaseId;
    @Column(name = "id_producto")
    private Integer productId;
}
