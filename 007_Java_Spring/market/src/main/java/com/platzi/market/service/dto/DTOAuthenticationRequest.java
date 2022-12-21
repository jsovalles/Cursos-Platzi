package com.platzi.market.service.dto;

import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

@Getter
@Setter
@ToString
public class DTOAuthenticationRequest {
    private String username;
    private String password;
}
