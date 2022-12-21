package com.platzi.market.service.dto;

import com.platzi.market.dao.entity.Category;
import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

import java.math.BigDecimal;

@Getter
@Setter
@ToString
public class DTOProduct {

    private int productId;
    private String name;
    private int categoryId;
    private BigDecimal price;
    private int stock;
    private boolean status;
    private Category category;
}
