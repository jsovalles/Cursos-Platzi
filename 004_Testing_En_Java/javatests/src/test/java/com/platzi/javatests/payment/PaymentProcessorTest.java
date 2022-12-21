package com.platzi.javatests.payment;

import org.junit.Assert;
import org.junit.Before;
import org.junit.Test;
import org.junit.runner.RunWith;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.Mockito;
import org.mockito.junit.MockitoJUnitRunner;

@RunWith(MockitoJUnitRunner.class)
public class PaymentProcessorTest {

    @Mock
    PaymentGateway paymentGateway;

    @InjectMocks
    PaymentProcessor paymentProcessor;

    @Test
    public void makePaymentSucessful() {
        Mockito.when(paymentGateway.requestPayment(Mockito.any())).thenReturn(new PaymentResponse(PaymentResponse.PaymentStatus.OK));
        boolean result = paymentProcessor.makePayment(1000);
        Assert.assertTrue(result);
    }

    @Test
    public void makePaymentError() {
        Mockito.when(paymentGateway.requestPayment(Mockito.any())).thenReturn(new PaymentResponse(PaymentResponse.PaymentStatus.ERROR));
        boolean result = paymentProcessor.makePayment(1000);
        Assert.assertFalse(result);
    }
}