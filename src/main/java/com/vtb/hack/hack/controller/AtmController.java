package com.vtb.hack.hack.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import com.vtb.hack.hack.dto.AtmDTO;
import com.vtb.hack.hack.service.AtmService;

import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.media.ArraySchema;
import io.swagger.v3.oas.annotations.media.Content;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;

@RequestMapping("/data/atms")
@RestController
public class AtmController {
    
    @Autowired
    private AtmService atmService;

    @GetMapping
    @Operation(description = "Get all atms")
    @ApiResponse(content = { @Content(
            mediaType = "application/json",
            array = @ArraySchema(schema = @Schema(implementation = AtmDTO.class)))
    })
    public Iterable<AtmDTO> getAllAtms() {
        return atmService.findAll();
    }

    @GetMapping("/atm")
    @Operation(description = "Get atm by id")
    @ApiResponse(content = { @Content(
            mediaType = "application/json",
            array = @ArraySchema(schema = @Schema(implementation = AtmDTO.class)))
    })
    public AtmDTO getAtmById(@RequestParam() Long id) {
        return atmService.findById(id);
    }
}
