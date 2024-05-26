package com.torabian.fireback.firebackspring;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.domain.EntityScan;

 
@SpringBootApplication
@EntityScan(basePackages = {"com.torabian", "com.fireback"})
public class FirebackSpringApplication  {
 
	public static void main(String[] args) {
		SpringApplication.run(FirebackSpringApplication.class, args); 
	}
}
