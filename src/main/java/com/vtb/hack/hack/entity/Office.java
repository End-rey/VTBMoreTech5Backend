package com.vtb.hack.hack.entity;

import java.util.HashSet;
import java.util.Set;

import jakarta.persistence.CascadeType;
import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.JoinTable;
import jakarta.persistence.ManyToMany;
import jakarta.persistence.OneToMany;
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
@Table(name = "offices")
@Data
public class Office extends BaseEntity {
	private String SalePointName;
	private String Address;
	private boolean Status;
	private String OfficeType;
	private String MetroStation;

    @Column(nullable = false)
	private double latitude;

    @Column(nullable = false)
	private double longitude;
    
    @EqualsAndHashCode.Exclude
    @Builder.Default
    @ManyToMany(cascade = { CascadeType.PERSIST, CascadeType.MERGE })
    @JoinTable(name = "offices_services", joinColumns = @JoinColumn(name = "office_id"), inverseJoinColumns = @JoinColumn(name = "service_id"))
	private Set<Service> services = new HashSet<>();

    @EqualsAndHashCode.Exclude
    @Builder.Default
    @OneToMany(cascade = CascadeType.ALL, mappedBy = "office")
    private Set<WorksTime> worksTimes = new HashSet<>();

    @EqualsAndHashCode.Exclude
    @Builder.Default
    @OneToMany(cascade = CascadeType.ALL, mappedBy = "office")
    private Set<Chart> charts = new HashSet<>();
}
