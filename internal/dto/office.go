package dto

type OfficeDTO struct {
	ID            int64
	SalePointName string
	Address       string
	Latitude      float32
	Longitude     float32
	Status        bool
	OfficeType    string
	MetroStation  string
	Services      []ServiceDTO
	Charts        []ChartDTO
	WorksTime     []WorkTimeDTO
}
