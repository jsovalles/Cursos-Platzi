package com.platzi.market.service.dto;

import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

@Getter
@Setter
@ToString
public class DTOCategory {
    private int categoryId;
    private String category;
    private boolean active;
}
