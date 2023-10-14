package com.vtb.hack.hack.service;

import com.vtb.hack.hack.dto.AtmDTO;

public interface AtmService {
    Iterable<AtmDTO> findAll();
    AtmDTO findById(Long id);
    Long addAtm(AtmDTO atmDTO);
}
