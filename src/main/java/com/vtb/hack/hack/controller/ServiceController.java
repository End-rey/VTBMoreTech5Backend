package com.vtb.hack.hack.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

import com.vtb.hack.hack.dto.ServiceDTO;
import com.vtb.hack.hack.service.ServiceService;

import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.media.ArraySchema;
import io.swagger.v3.oas.annotations.media.Content;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;

@RequestMapping("/data/services")
@RestController
public class ServiceController {
    
    @Autowired
    private ServiceService serviceService;

    @GetMapping
    @Operation(description = "Get all services")
    @ApiResponse(content = { @Content(
            mediaType = "application/json",
            array = @ArraySchema(schema = @Schema(implementation = ServiceDTO.class)))
    })
    public Iterable<ServiceDTO> getAllOffices() {
        return serviceService.findAll();
    }
}
