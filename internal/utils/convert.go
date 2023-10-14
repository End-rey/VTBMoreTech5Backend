package utils

import (
	"github.com/End-rey/VTBMoreTech5Backend/internal/dto"
	"github.com/End-rey/VTBMoreTech5Backend/internal/entity"
)

func OfficeToDTO(office entity.Office) dto.OfficeDTO {
	services := []dto.ServiceDTO{}
	for _, value := range office.Services {
		services = append(services, ServiceToDTO(value))
	}

	charts := ChartsToDTO(office.Charts)

	worksTime := []dto.WorkTimeDTO{}
	for _, value := range office.WorksTime {
		worksTime = append(worksTime, WorksTimeToDTO(value))
	}

	officeDTO := dto.OfficeDTO{
		ID: office.ID, 
    	SalePointName: office.SalePointName,
    	Address: office.Address,
    	Latitude: office.Latitude,
    	Longitude: office.Longitude,
    	Status: office.Status,
    	OfficeType: office.OfficeType,
    	MetroStation: office.MetroStation,
    	Services: services,
    	Charts: charts,
    	WorksTime: worksTime,
	}

	return officeDTO
}

func AtmToDTO(atm entity.Atm) dto.AtmDTO {
	services := []dto.ServiceDTO{}
	for _, value := range atm.Services {
		services = append(services, ServiceToDTO(value))
	}

	atmDTO := dto.AtmDTO{
		ID:atm.ID,
    	Address:atm.Address,
    	Latitude:atm.Latitude,
    	Longitude:atm.Longitude,
    	WorkTimeStart:atm.WorkTimeStart,
    	WorkTimeEnd:atm.WorkTimeEnd,
    	Services:services,
	}

	return atmDTO
}

func ChartsToDTO (charts []entity.Chart) []dto.ChartDTO {
	mapWeekDays := make(map[string][24]int)
	for _, chart := range charts {
		if _, ok := mapWeekDays[chart.WeeksDay]; !ok {
			mapWeekDays[chart.WeeksDay] = [24]int{}
		}

		tmp := mapWeekDays[chart.WeeksDay]

		ind := chart.ChartsTime.Hour()
		tmp[ind] = chart.Chart

		mapWeekDays[chart.WeeksDay] = tmp
	}

	chartsDTO := []dto.ChartDTO{}
	for key, value := range mapWeekDays {
		newChart := dto.ChartDTO{
    		WeeksDay: key,
    		ChartByTime: value[:],
		}
		chartsDTO = append(chartsDTO, newChart)
	}

	return chartsDTO
}

func ServiceToDTO(service entity.Service) dto.ServiceDTO {
	return dto.ServiceDTO{
		ID: service.ID,
		ServiceName: service.ServiceName,
	}
}

func WorksTimeToDTO(worksTime entity.WorksTime) dto.WorkTimeDTO {
	return dto.WorkTimeDTO{
		WeeksDay: worksTime.WeeksDay,
    	WorkTimeStart: worksTime.WorkTimeStart,
    	WorkTimeEnd: worksTime.WorkTimeEnd,
    	IsLegal: worksTime.IsLegal,
	}
}