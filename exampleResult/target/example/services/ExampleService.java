package com.example.example.services; 

import org.springframework.transaction.annotation.Transactional;
import com.example.example.repositories.ExampleRepository;
import com.example.example.models.entities.Example;

import org.springframework.stereotype.Service;
import lombok.AllArgsConstructor;
import javax.persistence.EntityNotFoundException;

import java.util.List;

@Service
@AllArgsConstructor
public class ExampleService { 

	private final ExampleRepository repository;

	@Transactional
	public Example create(Example example) {
		return repository.save(example);
	}

	public List<Example> getAll() {
		return (List<Example>) repository.findAll();
	}

	public Example getById(Long id) {
		return repository.findById(id).orElseThrow(EntityNotFoundException::new);
	}

	@Transactional
	public Example edit(Example example) {
		return repository.save(example);
	}

	@Transactional
	public void delete(Long id) {
		repository.deleteById(id);
	}
}