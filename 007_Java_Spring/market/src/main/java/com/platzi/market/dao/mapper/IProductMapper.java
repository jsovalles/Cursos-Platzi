package com.platzi.market.dao.mapper;

import com.platzi.market.dao.entity.Product;
import com.platzi.market.service.dto.DTOProduct;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;

import java.util.List;

@Mapper(componentModel = "spring", uses = {ICategoryMapper.class})
public interface IProductMapper {

    DTOProduct toProduct(Product product);
    List<DTOProduct> toProducts(List<Product> products);
    @Mapping(target = "barsCode", ignore = true)
    Product toDTOProduct(DTOProduct dtoProduct);
}
