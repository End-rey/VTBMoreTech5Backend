package com.vtb.hack.hack;

import java.io.BufferedReader;
import java.io.InputStreamReader;

import org.springframework.beans.factory.annotation.Value;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class HackApplication implements CommandLineRunner {

	public static void main(String[] args) {
		SpringApplication.run(HackApplication.class, args);
	}

	@Value("${parse.by.python.file}")
	private String isParse;

	@Override
	public void run(String... args) {
		System.out.println(isParse);
		if (isParse.equalsIgnoreCase("True")) {
			try {
				ProcessBuilder processBuilder = new ProcessBuilder("./python/parse_data/venv/Scripts/python.exe", "./python/parse_data/allparse.py");
				Process process = processBuilder.start();
				BufferedReader reader = new BufferedReader(new InputStreamReader(process.getInputStream()));
				String line;
				while ((line = reader.readLine()) != null) {
					System.out.println(line);
				}

				int exitCode = process.waitFor();
				if (exitCode == 0) {
					System.out.println("Python script executed successfully.");
				} else {
					System.err.println("Python script execution failed with exit code " + exitCode);
				}
			} catch (Exception e) {
				e.printStackTrace();
			}
		}
	}

}
