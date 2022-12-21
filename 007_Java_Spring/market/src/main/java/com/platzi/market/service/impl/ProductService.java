package com.platzi.market.service.impl;

import com.platzi.market.service.dto.DTOProduct;
import com.platzi.market.service.repository.IProductRepository;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;
import java.util.Optional;

@Service
public class ProductService {

    @Autowired
    private IProductRepository productRepository;

    public List<DTOProduct> getAllproducts() {
        return productRepository.getAllProducts();
    }

    public Optional<DTOProduct> getProduct(int productId) {
        return productRepository.getProduct(productId);
    }

    public Optional<List<DTOProduct>> getByCategory(int categoryId) {
        return productRepository.getByCategory(categoryId);
    }

    public Optional<List<DTOProduct>> getScarseProducts(int quantity) {
        return productRepository.getScarseProducts(quantity);
    }

    public DTOProduct save(DTOProduct product) {
        return productRepository.save(product);
    }

    public boolean delete(int productId) {
        return getProduct(productId).map(dtoProduct -> {
            productRepository.delete(productId);
            return true;
        }).orElse(false);
    }
}
