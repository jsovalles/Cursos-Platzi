package com.platzi.market.dao;

import com.platzi.market.dao.crud.IPurchaseCrudRepository;
import com.platzi.market.dao.entity.Purchase;
import com.platzi.market.dao.mapper.IPurchaseMapper;
import com.platzi.market.service.dto.DTOPurchase;
import com.platzi.market.service.repository.IPurchaseRepository;
import org.mapstruct.factory.Mappers;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

@Repository
public class PurchaseRepository implements IPurchaseRepository {

    @Autowired
    private IPurchaseCrudRepository purchaseCrudRepository;

    @Autowired
    private IPurchaseMapper mapper;

    @Override
    public List<DTOPurchase> getAllPurchases() {
        List<Purchase> out = (List<Purchase>) purchaseCrudRepository.findAll();
        return mapper.toPurchases(out);
        //return mapper.toPurchases((List<Purchase>) purchaseCrudRepository.findAll());
    }

    @Override
    public Optional<List<DTOPurchase>> getByClient(String clientId) {
        return purchaseCrudRepository.findByClientId(clientId)
                .map(purchases -> mapper.toPurchases(purchases));
    }

    @Override
    public DTOPurchase savePurchase(DTOPurchase purchase) {
        Purchase purchaseValue = mapper.toDTOPurchase(purchase);
        purchaseValue.getPurchases().forEach(purchaseItem -> purchaseItem.setPurchase(purchaseValue));

        return mapper.toPurchase(purchaseCrudRepository.save(purchaseValue));
    }
}
