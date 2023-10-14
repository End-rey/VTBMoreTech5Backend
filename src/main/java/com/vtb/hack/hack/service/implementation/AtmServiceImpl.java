package com.vtb.hack.hack.service.implementation;

import java.util.stream.Collectors;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.vtb.hack.hack.dto.AtmDTO;
import com.vtb.hack.hack.entity.Atm;
import com.vtb.hack.hack.repository.AtmRepo;
import com.vtb.hack.hack.service.AtmService;
import com.vtb.hack.hack.utils.ConvertToTo;

import jakarta.persistence.EntityNotFoundException;

@Service
public class AtmServiceImpl implements AtmService {

    private final AtmRepo atmRepo;

    @Autowired
    public AtmServiceImpl(AtmRepo atmRepo) {this.atmRepo = atmRepo;}

    @Override
    public AtmDTO findById(Long id){
        Atm atm = atmRepo.findById(id).orElseThrow(
                () -> new EntityNotFoundException("Atm with id " + id + " not found")
        );
        return ConvertToTo.atmToTO(atm);
    }

    @Override
    public Iterable<AtmDTO> findAll() {
        return atmRepo.findAll().stream()
                .map(ConvertToTo::atmToTO)
                .collect(Collectors.toSet());
    }

    @Override
    public Long addAtm(AtmDTO atmDTO) {
        return 0l;
    }
    
}
