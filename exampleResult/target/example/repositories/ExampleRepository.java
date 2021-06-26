package com.example.example.repositories; 

import com.example.example.models.entities.Example;

import org.springframework.stereotype.Repository;
import org.springframework.data.repository.CrudRepository;

@Repository
public interface ExampleRepository extends CrudRepository<Example, Long> { 

}