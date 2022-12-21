package com.platzi.market.dao.mapper;

import com.platzi.market.dao.entity.PurchaseItem;
import com.platzi.market.service.dto.DTOPurchaseItem;
import org.mapstruct.InheritInverseConfiguration;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;

@Mapper(componentModel = "spring", uses = {IProductMapper.class})
public interface IPurchaseItemMapper {

    @Mapping(source = "id.productId", target = "productId")
    DTOPurchaseItem toPurchaseItem(PurchaseItem purchaseItem);

    @InheritInverseConfiguration
    @Mapping(target = "purchase", ignore = true)
    @Mapping(target = "product", ignore = true)
    @Mapping(target = "id.purchaseId", ignore = true)
    PurchaseItem toDTOPurchaseItem(DTOPurchaseItem purchaseItem);
}
