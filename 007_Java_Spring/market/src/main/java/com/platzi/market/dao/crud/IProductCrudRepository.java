package com.platzi.market.dao.crud;

import com.platzi.market.dao.entity.Product;
import org.springframework.data.repository.CrudRepository;

import java.util.List;

public interface IProductCrudRepository extends CrudRepository<Product, Integer> {

    //@Query(value = "SELECT * FROM productos WHERE id_categoria = ?", nativeQuery = true) --Metodo nativo
    List<Product> findByCategoryIdOrderByNameAsc(int categoryId);

    //Optional<List<Product>> findByStockValueLessThanAndStatus(int stockValue, boolean status);
}
