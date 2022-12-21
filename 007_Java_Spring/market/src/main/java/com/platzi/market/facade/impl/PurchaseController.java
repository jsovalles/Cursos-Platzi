package com.platzi.market.facade.impl;

import com.platzi.market.service.dto.DTOPurchase;
import com.platzi.market.service.impl.PurchaseService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.List;

@RestController
@RequestMapping("/purchases")
public class PurchaseController {

    @Autowired
    private PurchaseService service;

    @GetMapping
    public ResponseEntity<List<DTOPurchase>> getAllPurchases() {
        return new ResponseEntity<>(service.getAllPurchases(), HttpStatus.OK);
    }

    @GetMapping("client/{client-id}")
    public ResponseEntity<List<DTOPurchase>> getByClient(@PathVariable("client-id") String clientId) {
        return service.getByClient(clientId).map(dtoPurchases -> new ResponseEntity<>(dtoPurchases, HttpStatus.OK))
                .orElse(new ResponseEntity<>(HttpStatus.NOT_FOUND));
    }

    @PostMapping
    public ResponseEntity<DTOPurchase> savePurchase(@RequestBody DTOPurchase purchase) {
        return new ResponseEntity(service.savePurchase(purchase), HttpStatus.CREATED);
    }
}
