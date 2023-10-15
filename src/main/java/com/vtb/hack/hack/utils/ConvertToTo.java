package com.vtb.hack.hack.utils;

import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;
import java.util.stream.Collectors;

import com.vtb.hack.hack.dto.AtmDTO;
import com.vtb.hack.hack.dto.ChartDTO;
import com.vtb.hack.hack.dto.OfficeDTO;
import com.vtb.hack.hack.dto.ServiceDTO;
import com.vtb.hack.hack.dto.WorksTimeDTO;
import com.vtb.hack.hack.entity.Atm;
import com.vtb.hack.hack.entity.Chart;
import com.vtb.hack.hack.entity.Office;
import com.vtb.hack.hack.entity.Service;
import com.vtb.hack.hack.entity.WorksTime;

public class ConvertToTo {
    private ConvertToTo() {
    }

    public static AtmDTO atmToTO(Atm atm) {
        Set<ServiceDTO> services = atm.getServices().stream()
            .map(ConvertToTo::serviceToTO)
            .collect(Collectors.toSet());

        return AtmDTO.builder()
                .id(atm.getId())
                .address(atm.getAddress())
                .latitude(atm.getLatitude())
                .longitude(atm.getLongitude())
                .workTimeStart(atm.getWorkTimeStart())
                .workTimeEnd(atm.getWorkTimeEnd())
                .services(services)
                .build();
    }

    public static OfficeDTO officeToTO(Office office) {
        Set<ServiceDTO> services = office.getServices().stream()
            .map(ConvertToTo::serviceToTO)
            .collect(Collectors.toSet());

        Set<ChartDTO> charts = chartsToTO(office.getCharts());

        Set<WorksTimeDTO> worksTime = office.getWorksTimes().stream()
            .map(ConvertToTo::worksTimeToTO)
            .collect(Collectors.toSet());

        OfficeDTO officeDTO = OfficeDTO.builder()
                .id(office.getId())
                .salePointName(office.getSalePointName())
                .address(office.getAddress())
                .latitude(office.getLatitude())
                .longitude(office.getLongitude())
                .status(office.isStatus())
                .officeType(office.getOfficeType())
                .metroStation(office.getMetroStation())
                .services(services)
                .charts(charts)
                .worksTime(worksTime)
                .build();

        return officeDTO;

    }

    public static ServiceDTO serviceToTO(Service service) {
        return ServiceDTO.builder()
                    .serviceId(service.getId())
                    .serviceName(service.getServiceName())
                    .build();
    }

    public static WorksTimeDTO worksTimeToTO(WorksTime worksTime) {
        return WorksTimeDTO.builder()
                    .weeksDay(worksTime.getWeeksDay())
                    .workTimeStart(worksTime.getWorkTimeStart())
                    .workTimeEnd(worksTime.getWorkTimeEnd())
                    .isLegal(worksTime.getIsLegal())
                    .build();
    }

    public static Set<ChartDTO> chartsToTO(Set<Chart> charts) {
        Map<String, int[]> mapWeekDays = new HashMap<>();

        for (Chart chart : charts) {
            if (!mapWeekDays.containsKey(chart.getWeeksDay())) {
                mapWeekDays.put(chart.getWeeksDay(), new int[24]);
            }

            int ind = (int) chart.getChartsTime().toLocalTime().getHour();
            mapWeekDays.get(chart.getWeeksDay())[ind] = chart.getChart();
        }

        Set<ChartDTO> chartsDTO = new HashSet<>();
        for (Map.Entry<String, int[]> entry : mapWeekDays.entrySet()) {
            String key = entry.getKey();
            int[] value = entry.getValue();
            chartsDTO.add(new ChartDTO(key, value));
        }

        return chartsDTO;
    }
}
