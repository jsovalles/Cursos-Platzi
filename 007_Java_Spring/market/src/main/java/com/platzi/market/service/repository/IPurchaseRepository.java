package com.platzi.market.service.repository;

import com.platzi.market.service.dto.DTOPurchase;

import java.util.List;
import java.util.Optional;

public interface IPurchaseRepository {

    List<DTOPurchase> getAllPurchases();
    Optional<List<DTOPurchase>> getByClient(String clientId);
    DTOPurchase savePurchase(DTOPurchase purchase);

}
