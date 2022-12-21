package com.platzi.javatests.payment;

import lombok.AllArgsConstructor;
import lombok.Data;

@Data
@AllArgsConstructor
public class PaymentResponse {

    enum PaymentStatus {
        OK, ERROR
    }

    private PaymentStatus status;

}
