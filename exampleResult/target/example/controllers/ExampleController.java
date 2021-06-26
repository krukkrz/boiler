package com.example.example.controllers; 

import com.example.example.services.ExampleService;
import com.example.example.models.dtos.ExampleDto;
import com.example.example.models.entities.Example;
import com.example.example.models.mapping.ExampleMapper;
import static com.example.example.models.mapping.ExampleMapper.toDto;
import static com.example.example.models.mapping.ExampleMapper.toModel;

import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.RequestMapping;
import lombok.AllArgsConstructor;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PutMapping;
import org.springframework.web.bind.annotation.DeleteMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.PathVariable;
import static org.springframework.http.HttpStatus.OK;

import java.util.List;
import java.util.stream.Collectors;

@RestController
@RequestMapping("/examples")
@AllArgsConstructor
public class ExampleController { 

	private final ExampleService service;

	@PostMapping()
	public ResponseEntity<ExampleDto> create(@RequestBody ExampleDto exampleDto) {
		Example example = toModel(exampleDto);
		Example savedExample = service.create(example);
		return ResponseEntity.ok(toDto(savedExample));
	}

	@GetMapping()
	public ResponseEntity<List<ExampleDto>> getAll() {
		return ResponseEntity.ok(
			service.getAll()
				.stream()
				.map(ExampleMapper::toDto)
				.collect(Collectors.toList()));
	}

	@GetMapping("/{id}")
	public ResponseEntity<ExampleDto> getById(@PathVariable("id") Long id) {
		return ResponseEntity.ok(
			toDto(service.getById(id))
		);
	}

	@PutMapping()
	public ResponseEntity<ExampleDto> edit(@RequestBody ExampleDto exampleDto) {
		Example example = toModel(exampleDto);
		Example savedExample = service.edit(example);
		return ResponseEntity.ok(toDto(savedExample));
	}

	@DeleteMapping("/{id}")
	@ResponseStatus(OK)
	public void delete(@PathVariable("id")Long id) {
		service.delete(id);
	}
}