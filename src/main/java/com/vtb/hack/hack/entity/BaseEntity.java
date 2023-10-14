package com.vtb.hack.hack.entity;

import java.util.Date;

import org.hibernate.annotations.CreationTimestamp;
import org.hibernate.annotations.UpdateTimestamp;

import jakarta.persistence.Column;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.MappedSuperclass;
import jakarta.persistence.PrePersist;
import jakarta.persistence.PreUpdate;
import lombok.AllArgsConstructor;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@MappedSuperclass
@AllArgsConstructor
@NoArgsConstructor
public class BaseEntity {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private long id;

    @UpdateTimestamp
    @Column(name = "last_update_time")
    private Date lastUpdateTime;

    @CreationTimestamp
    @Column(name = "creation_time")
    private Date creationTime;

    @PrePersist
    protected void onCreate() {
        lastUpdateTime = creationTime = new Date();
    }

    @PreUpdate
    protected void onUpdate() {
        lastUpdateTime = new Date();
    }
}
