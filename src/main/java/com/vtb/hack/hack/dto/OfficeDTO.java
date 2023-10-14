package com.vtb.hack.hack.dto;

import java.util.Set;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Builder
@Data
@AllArgsConstructor
@NoArgsConstructor
public class OfficeDTO {
    private Long id;
	private String salePointName;
	private String address;
	private Float latitude;
	private Float longitude;
	private Boolean status;
	private String officeType;
	private String metroStation;
	private Set<ServiceDTO> services;
	private Set<ChartDTO> charts;
	private Set<WorksTimeDTO> worksTime;
}
