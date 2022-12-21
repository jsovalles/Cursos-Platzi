package com.platzi.market.dao.crud;

import com.platzi.market.dao.entity.Purchase;
import org.springframework.data.repository.CrudRepository;

import java.util.List;
import java.util.Optional;

public interface IPurchaseCrudRepository extends CrudRepository<Purchase, Integer> {

    Optional<List<Purchase>> findByClientId(String clientId);
}
