package com.torabian.fireback.firebackspring;

import java.util.List;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import com.torabian.fireback.project.Book;
import com.torabian.fireback.project.BookRepository;

@SpringBootApplication
public class FirebackSpringApplication  {

	public static void main(String[] args) {
		SpringApplication.run(FirebackSpringApplication.class, args);
		System.out.println("Hi this is a new spring boot application");

		int[] numbers = {1,2,3,4,5,6,7,8};

		BinarySearch n = new BinarySearch();

		System.out.println(n.binarySearch(numbers, 5));
		System.out.println("Finding books:");

	}
 

}
