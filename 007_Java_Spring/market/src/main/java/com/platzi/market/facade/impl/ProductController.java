package com.platzi.market.facade.impl;

import com.platzi.market.service.dto.DTOProduct;
import com.platzi.market.service.impl.ProductService;
import io.swagger.annotations.ApiOperation;
import io.swagger.annotations.ApiParam;
import io.swagger.annotations.ApiResponse;
import io.swagger.annotations.ApiResponses;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;
import java.util.Optional;

@RestController
@RequestMapping("/products")
public class ProductController {

    @Autowired
    private ProductService productService;

    @ApiOperation("Get all supermarket products")
    @ApiResponse(code = 200, message = "OK")
    @GetMapping
    public ResponseEntity<List<DTOProduct>> getAllProducts() {
        return new ResponseEntity<>(productService.getAllproducts(), HttpStatus.OK);
    }

    @ApiOperation("Get a product according to an id")
    @ApiResponses({
            @ApiResponse(code = 200, message = "OK"),
            @ApiResponse(code = 404, message = "Not found")
    })
    @GetMapping("/{product-id}")
    public ResponseEntity<DTOProduct> getProduct(@ApiParam(value = "id of the product", required = true) @PathVariable("product-id") int productId) {
        return productService.getProduct(productId)
                .map(dtoProduct -> new ResponseEntity<>(dtoProduct, HttpStatus.OK))
                .orElse(new ResponseEntity<>(HttpStatus.NOT_FOUND));
    }

    @GetMapping("category/{category-id}")
    public ResponseEntity<List<DTOProduct>> getByCategory(@PathVariable("category-id") int categoryId) {
        return productService.getByCategory(categoryId)
                .map(dtoProducts -> new ResponseEntity<>(dtoProducts, HttpStatus.OK))
                .orElse(new ResponseEntity<>(HttpStatus.NOT_FOUND));
    }

    public Optional<List<DTOProduct>> getScarseProducts(int quantity) {
        return productService.getScarseProducts(quantity);
    }

    @PostMapping
    public ResponseEntity<DTOProduct> save(@RequestBody DTOProduct product) {
        return new ResponseEntity<>(productService.save(product), HttpStatus.CREATED);
    }

    @DeleteMapping(("/{product-id}"))
    public ResponseEntity delete(@PathVariable("product-id") int productId) {
        if (productService.delete(productId))
            return new ResponseEntity(HttpStatus.OK);

        return new ResponseEntity(HttpStatus.NOT_FOUND);
    }
}
