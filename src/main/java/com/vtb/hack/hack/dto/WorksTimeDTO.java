package com.vtb.hack.hack.dto;

import java.util.Date;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Builder
@Data
@AllArgsConstructor
@NoArgsConstructor
public class WorksTimeDTO {
    private String weeksDay;
	private Date workTimeStart;
	private Date workTimeEnd;
	private Boolean isLegal;
}
