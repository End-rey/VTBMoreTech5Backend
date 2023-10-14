package com.vtb.hack.hack.service.implementation;

import java.util.stream.Collectors;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.vtb.hack.hack.dto.OfficeDTO;
import com.vtb.hack.hack.entity.Office;
import com.vtb.hack.hack.repository.OfficeRepo;
import com.vtb.hack.hack.service.OfficeService;
import com.vtb.hack.hack.utils.ConvertToTo;

import jakarta.persistence.EntityNotFoundException;

@Service
public class OfficeServiceImpl implements OfficeService {

    private final OfficeRepo officeRepo;

    @Autowired
    public OfficeServiceImpl(OfficeRepo officeRepo) {
        this.officeRepo = officeRepo;
    }

    @Override
    public Iterable<OfficeDTO> findAll() {
        return officeRepo.findAll().stream()
            .map(ConvertToTo::officeToTO)
            .collect(Collectors.toSet());
    }

    @Override
    public OfficeDTO findById(Long id) {
        Office office = officeRepo.findById(id).orElseThrow(
                () -> new EntityNotFoundException("Office with id " + id + " not found")
        );
        return ConvertToTo.officeToTO(office);
    }

    @Override
    public Long addOffice(OfficeDTO officeDTO) {
        // TODO Auto-generated method stub
        throw new UnsupportedOperationException("Unimplemented method 'addOffice'");
    }
    
}
