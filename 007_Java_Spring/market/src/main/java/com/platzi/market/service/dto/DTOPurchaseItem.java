package com.platzi.market.service.dto;

import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

@Getter
@Setter
@ToString
public class DTOPurchaseItem {
    private int productId;
    private int quantity;
    private double total;
    private boolean status;
}
