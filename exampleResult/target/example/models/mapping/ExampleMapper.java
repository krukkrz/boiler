package com.example.example.models.mapping; 

import com.example.example.models.entities.Example;
import com.example.example.models.dtos.ExampleDto;

import org.springframework.stereotype.Component;

@Component
public class ExampleMapper { 

	public static Example toModel(ExampleDto example) {
		return Example.builder().build();
	}

	public static ExampleDto toDto(Example example) {
		return ExampleDto.builder().build();
	}
}