package com.vtb.hack.hack.dto;

import java.util.Date;
import java.util.Set;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Builder
@Data
@AllArgsConstructor
@NoArgsConstructor
public class AtmDTO {
    private Long id;
	private String address;
	private float latitude;
	private float longitude;
	private Date workTimeStart;
	private Date workTimeEnd;
	private Set<ServiceDTO> services;
}
