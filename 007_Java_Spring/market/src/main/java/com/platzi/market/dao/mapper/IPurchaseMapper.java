package com.platzi.market.dao.mapper;

import com.platzi.market.dao.entity.Purchase;
import com.platzi.market.service.dto.DTOPurchase;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;

import java.util.List;

@Mapper(componentModel = "spring", uses = {IPurchaseItemMapper.class})
public interface IPurchaseMapper {

    DTOPurchase toPurchase(Purchase purchase);

    List<DTOPurchase> toPurchases(List<Purchase> purchaseList);

    @Mapping(target = "client", ignore = true)
    Purchase toDTOPurchase(DTOPurchase purchase);

}
