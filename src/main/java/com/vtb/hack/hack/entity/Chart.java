package com.vtb.hack.hack.entity;

import java.sql.Time;

import jakarta.persistence.CascadeType;
import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.ManyToOne;
import jakarta.persistence.Table;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.EqualsAndHashCode;
import lombok.NoArgsConstructor;

@EqualsAndHashCode(callSuper = false)
@Builder
@NoArgsConstructor
@AllArgsConstructor
@Entity
@Table(name = "charts_offices")
@Data
public class Chart extends BaseEntity {
    @Column(nullable = false)
    private String weeksDay;

    private Time chartsTime;

    private int chart;

    @ManyToOne(cascade = CascadeType.ALL)
    @JoinColumn(name = "office_id", nullable = false)
    private Office office;
}
