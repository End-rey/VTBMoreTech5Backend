package com.vtb.hack.hack.dto;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Builder
@Data
@AllArgsConstructor
@NoArgsConstructor
public class ServiceDTO {
    private Long serviceId;
    private String serviceName;
}
