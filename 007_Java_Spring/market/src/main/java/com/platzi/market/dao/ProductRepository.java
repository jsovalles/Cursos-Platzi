package com.platzi.market.dao;

import com.platzi.market.dao.crud.IProductCrudRepository;
import com.platzi.market.dao.entity.Product;
import com.platzi.market.dao.mapper.IProductMapper;
import com.platzi.market.service.dto.DTOProduct;
import com.platzi.market.service.repository.IProductRepository;
import org.mapstruct.factory.Mappers;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Repository;

import java.util.List;
import java.util.Optional;

@Repository
public class ProductRepository implements IProductRepository {
    @Autowired
    private IProductCrudRepository productCrudRepository;

    @Autowired
    private IProductMapper mapper;

    @Override
    public List<DTOProduct> getAllProducts() {
        List<Product> products = (List<Product>) productCrudRepository.findAll();
        return mapper.toProducts(products);
    }

    public Optional<List<DTOProduct>> getByCategory(int categoryId) {
        List<Product> products = productCrudRepository.findByCategoryIdOrderByNameAsc(categoryId);
        return Optional.of(mapper.toProducts(products));
    }

    @Override
    public Optional<List<DTOProduct>> getScarseProducts(int quantity) {
        //Optional<List<Product>> products = productCrudRepository.findByStockValueLessThanAndStatus(quantity, true);
        //return products.map(product -> productMapper.toProducts(product));
        return null;
    }

    public Optional<DTOProduct> getProduct(int productId) {
        return productCrudRepository.findById(productId).map(product1 -> mapper.toProduct(product1));
    }

    @Override
    public DTOProduct save(DTOProduct product) {
        return mapper.toProduct(productCrudRepository.save(mapper.toDTOProduct(product)));
    }

    @Override
    public void delete(int productId) {
        productCrudRepository.deleteById(productId);
    }

}
