package com.platzi.market.service.repository;

import com.platzi.market.service.dto.DTOProduct;

import java.util.List;
import java.util.Optional;

public interface IProductRepository {
    List<DTOProduct> getAllProducts();
    Optional<List<DTOProduct>> getByCategory(int categoryId);
    Optional<List<DTOProduct>> getScarseProducts(int quantity);
    Optional<DTOProduct> getProduct(int productId);
    DTOProduct save(DTOProduct product);
    void delete(int productId);

}
