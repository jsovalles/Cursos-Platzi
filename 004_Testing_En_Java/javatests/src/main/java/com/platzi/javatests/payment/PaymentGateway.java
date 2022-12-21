package com.platzi.javatests.payment;

public interface PaymentGateway {

    PaymentResponse requestPayment(PaymentRequest request);
}
