package com.vtb.hack.hack.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import com.vtb.hack.hack.dto.OfficeDTO;
import com.vtb.hack.hack.dto.OfficeFiltRequestTO;
import com.vtb.hack.hack.service.OfficeService;

import io.swagger.v3.oas.annotations.Operation;
import io.swagger.v3.oas.annotations.media.ArraySchema;
import io.swagger.v3.oas.annotations.media.Content;
import io.swagger.v3.oas.annotations.media.Schema;
import io.swagger.v3.oas.annotations.responses.ApiResponse;

@RequestMapping("/data/offices")
@RestController
public class OfficeController {
    
    @Autowired
    private OfficeService officeService;

    @GetMapping
    @Operation(description = "Get all offices")
    @ApiResponse(content = { @Content(
            mediaType = "application/json",
            array = @ArraySchema(schema = @Schema(implementation = OfficeDTO.class)))
    })
    public Iterable<OfficeDTO> getAllOffices() {
        return officeService.findAll();
    }

    @GetMapping("/office")
    @Operation(description = "Get office by id")
    @ApiResponse(content = { @Content(
            mediaType = "application/json",
            schema = @Schema(implementation = OfficeDTO.class))
    })
    public OfficeDTO getOfficeById(@RequestParam() Long id) {
        return officeService.findById(id);
    }

    @PostMapping("/filter")
    @Operation(description = "Get all offices with filter")
    @ApiResponse(content = { @Content(
            mediaType = "application/json",
            array = @ArraySchema(schema = @Schema(implementation = OfficeDTO.class)))
    })
    public Iterable<OfficeDTO> getAllOfficesFiltered(@RequestBody OfficeFiltRequestTO officeFiltRequestTO) {
        return officeService.findAllFiltered(officeFiltRequestTO);
    }
}
