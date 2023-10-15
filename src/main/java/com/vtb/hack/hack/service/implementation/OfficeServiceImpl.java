package com.vtb.hack.hack.service.implementation;

import java.sql.Time;
import java.time.LocalDateTime;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Comparator;
import java.util.Date;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.Set;
import java.util.stream.Collectors;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import com.vtb.hack.hack.dto.OfficeDTO;
import com.vtb.hack.hack.dto.OfficeFiltRequestTO;
import com.vtb.hack.hack.entity.Chart;
import com.vtb.hack.hack.entity.Office;
import com.vtb.hack.hack.repository.OfficeRepo;
import com.vtb.hack.hack.service.OfficeService;
import com.vtb.hack.hack.utils.ConvertToTo;
import com.vtb.hack.hack.utils.Predict;

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
        List<Long> servicesIds = Arrays.stream(officeFiltRequestTO.getServicesIds()).boxed().toList();

        List<Office> filteredOfficesList = officeRepo.findAll().stream()
            .filter(office -> office.getServices().stream()
                .mapToLong(service -> service.getId())
                .anyMatch(serviceId -> servicesIds.contains(serviceId))
            )
            .collect(Collectors.toList());

        filteredOfficesList.sort(Comparator.comparing(office -> {
            double distanceForA = calculateDistance(
                new Pos(office.getLongitude(), office.getLatitude()), 
                new Pos(officeFiltRequestTO.getLongitude(), officeFiltRequestTO.getLatitude())
            );
            float forA = Predict.optimize(
                distanceForA,
                1 / Math.max(0.001, 19 - LocalDateTime.now().getHour()),
                calculateChartByTime(office.getCharts())
            );
            return -forA;
        }));

        if (officeFiltRequestTO.getLimit() > 0 && filteredOfficesList.size() > officeFiltRequestTO.getLimit()) {
            filteredOfficesList = filteredOfficesList.subList(0, officeFiltRequestTO.getLimit());
        }

        return filteredOfficesList.stream()
            .map(ConvertToTo::officeToTO)
            .collect(Collectors.toList()); // Convert to List<OfficeDTO>
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

    private static double calculateChartByTime(Set<Chart> charts) {
        Map<String, int[]> mapWeekDays = new HashMap<>();

        for (Chart chart : charts) {
            if (!mapWeekDays.containsKey(chart.getWeeksDay())) {
                mapWeekDays.put(chart.getWeeksDay(), new int[24]);
            }

            int ind = (int) chart.getChartsTime().toLocalTime().getHour();
            mapWeekDays.get(chart.getWeeksDay())[ind] = chart.getChart();
        }

        String dayOfWeek = "" + (LocalDateTime.now().getDayOfWeek().getValue() - 1);
        int res = mapWeekDays.get(dayOfWeek)[LocalDateTime.now().getHour()];
        return res;
    }
    
}
