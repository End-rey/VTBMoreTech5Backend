package com.vtb.hack.hack.service;

import com.vtb.hack.hack.dto.OfficeDTO;
import com.vtb.hack.hack.dto.OfficeFiltRequestTO;

public interface OfficeService {
    Iterable<OfficeDTO> findAll();
    OfficeDTO findById(Long id);
    Long addOffice(OfficeDTO officeDTO);
    Iterable<OfficeDTO> findAllFiltered(OfficeFiltRequestTO officeFiltRequestTO);
}
