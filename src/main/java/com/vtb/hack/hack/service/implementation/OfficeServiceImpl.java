package com.vtb.hack.hack.service.implementation;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.vtb.hack.hack.dto.OfficeDTO;
import com.vtb.hack.hack.dto.OfficeFiltRequestTO;
import com.vtb.hack.hack.entity.Office;
import com.vtb.hack.hack.repository.OfficeRepo;
import com.vtb.hack.hack.service.OfficeService;
import com.vtb.hack.hack.utils.ConvertToTo;

import jakarta.persistence.EntityNotFoundException;
import lombok.AllArgsConstructor;

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
                () -> new EntityNotFoundException("Office with id " + id + " not found"));
        return ConvertToTo.officeToTO(office);
    }

    @Override
    public Long addOffice(OfficeDTO officeDTO) {
        // TODO Auto-generated method stub
        throw new UnsupportedOperationException("Unimplemented method 'addOffice'");
    }

    @Override
    public Iterable<OfficeDTO> findAllFiltered(OfficeFiltRequestTO officeFiltRequestTO) {
        List <Long> servicesIds = Arrays.stream(officeFiltRequestTO.getServicesIds()).boxed().toList();
        Iterable<Office> filteredOfficesItearble = officeRepo.findAll().stream()
                .filter(office -> office.getServices().stream()
                        .allMatch(service -> 
                            servicesIds.contains(service.getId())))
                .collect(Collectors.toSet());


        List<Office> filteredOfficesList = new ArrayList<Office>();
        filteredOfficesItearble.forEach(filteredOfficesList::add);

        filteredOfficesList.sort((a, b) -> {
            return Double.compare(
            calculateDistance(new Pos(a.getLongitude(),  a.getLatitude()), 
                new Pos(officeFiltRequestTO.getLongitude(), officeFiltRequestTO.getLatitude())),
            calculateDistance(new Pos(b.getLongitude(), b.getLatitude()), 
                new Pos(officeFiltRequestTO.getLongitude(), officeFiltRequestTO.getLatitude())));
        });

        if (officeFiltRequestTO.getLimit() != 0) {
            while (filteredOfficesList.size() > officeFiltRequestTO.getLimit()) {
                filteredOfficesList.remove(filteredOfficesList.size());
            }
        }

        return filteredOfficesList.stream()
            .map(ConvertToTo::officeToTO)
            .collect(Collectors.toSet());
    }

    @AllArgsConstructor
    private class Pos {
        double latitude;
        double longitude;
    }

    private static double calculateDistance(Pos myPos, Pos pointPos) {
        return Math.sqrt(
                Math.pow(myPos.latitude - pointPos.latitude, 2) + Math.pow(myPos.longitude - pointPos.longitude, 2));
    }
}
