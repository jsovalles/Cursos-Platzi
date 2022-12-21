package com.platzi.market.service.impl;

import com.platzi.market.service.dto.DTOPurchase;
import com.platzi.market.service.repository.IPurchaseRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class PurchaseService {

    @Autowired
    IPurchaseRepository purchaseRepository;

    public List<DTOPurchase> getAllPurchases() {
        return purchaseRepository.getAllPurchases();
    }

    public Optional<List<DTOPurchase>> getByClient(String clientId) {
        return purchaseRepository.getByClient(clientId);
    }

    public DTOPurchase savePurchase(DTOPurchase purchase) {
        return purchaseRepository.savePurchase(purchase);
    }
}
