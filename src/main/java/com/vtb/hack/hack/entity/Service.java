package com.vtb.hack.hack.entity;

import java.util.HashSet;
import java.util.Set;

import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.ManyToMany;
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
@Table(name = "services")
@Data
public class Service extends BaseEntity {
    @Column(nullable = false)
    private String serviceName;

    @EqualsAndHashCode.Exclude
    @Builder.Default
    @ManyToMany(mappedBy = "services")
    private Set<Office> office = new HashSet<>();

    @EqualsAndHashCode.Exclude
    @Builder.Default
    @ManyToMany(mappedBy = "services")
    private Set<Atm> atm = new HashSet<>();
}
