package com.vtb.hack.hack.entity;

import java.sql.Time;

import jakarta.persistence.CascadeType;
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
@Table(name = "works_time_offices")
@Data
public class WorksTime extends BaseEntity {
	private String WeeksDay;
	private Time WorkTimeStart;
	private Time WorkTimeEnd;
	private Time LunchTimeStart;
	private Time LunchTimeEnd;
	private boolean IsLegal;

    @ManyToOne(cascade = CascadeType.ALL)
    @JoinColumn(name = "office_id", nullable = false)
	private Office office;
}
