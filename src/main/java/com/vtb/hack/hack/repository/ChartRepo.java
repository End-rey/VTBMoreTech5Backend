package com.vtb.hack.hack.repository;

import org.springframework.data.jpa.repository.JpaRepository;

import com.vtb.hack.hack.entity.Chart;

public interface ChartRepo extends JpaRepository<Chart, Long> {
    
}
