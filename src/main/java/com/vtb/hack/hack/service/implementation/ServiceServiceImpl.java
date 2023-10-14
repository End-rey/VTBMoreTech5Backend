package com.vtb.hack.hack.service.implementation;

import java.util.stream.Collectors;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.vtb.hack.hack.dto.ServiceDTO;
import com.vtb.hack.hack.repository.ServiceRepo;
import com.vtb.hack.hack.service.ServiceService;
import com.vtb.hack.hack.utils.ConvertToTo;

@Service
public class ServiceServiceImpl implements ServiceService {
    
    private final ServiceRepo serviceRepo;

    @Autowired
    private ServiceServiceImpl(ServiceRepo serviceRepo) {
        this.serviceRepo = serviceRepo;
    }

    @Override
    public Iterable<ServiceDTO> findAll() {
        return serviceRepo.findAll().stream()
                .map(ConvertToTo::serviceToTO)
                .collect(Collectors.toSet());
    }
}
