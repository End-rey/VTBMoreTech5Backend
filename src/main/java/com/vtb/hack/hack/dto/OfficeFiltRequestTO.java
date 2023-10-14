package com.vtb.hack.hack.dto;

import java.sql.Time;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Builder
@Data
@AllArgsConstructor
@NoArgsConstructor
public class OfficeFiltRequestTO {
    long[] servicesIds;
    double latitude;
    double longitude;
    Time startTime;
    Time endTime;
    int limit; 
}
