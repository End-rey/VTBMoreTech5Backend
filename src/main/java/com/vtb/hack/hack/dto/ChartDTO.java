package com.vtb.hack.hack.dto;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Builder
@Data
@AllArgsConstructor
@NoArgsConstructor
public class ChartDTO {
    private String weeksDay;
	private int[] chartByTime;
}
