package com.example.example.models.entities; 

import javax.persistence.Id;
import javax.persistence.GeneratedValue;
import javax.persistence.Entity;
import lombok.Builder;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.AllArgsConstructor;
import lombok.NoArgsConstructor;

@Entity
@Builder
@Data
@EqualsAndHashCode
@AllArgsConstructor
@NoArgsConstructor
public class Example { 

	@Id
	@GeneratedValue
	private Long id;
	//TODO implement parameters here
}