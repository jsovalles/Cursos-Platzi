package com.platzi.market.dao.mapper;

import com.platzi.market.dao.entity.Category;
import com.platzi.market.service.dto.DTOCategory;
import org.mapstruct.InheritInverseConfiguration;
import org.mapstruct.Mapper;
import org.mapstruct.Mapping;
import org.mapstruct.Mappings;

@Mapper(componentModel = "spring")
public interface ICategoryMapper {
    @Mappings({
            @Mapping(source = "id", target = "categoryId"),
            @Mapping(source = "description", target = "category"),
    })
    DTOCategory toCategory(Category category);

    @InheritInverseConfiguration
    @Mapping(target = "products", ignore = true)
    Category toDTOCategory(DTOCategory dtoCategory);
}
