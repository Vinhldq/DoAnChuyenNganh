package vo

type Beds struct {
	SingleBed           uint `json:"single_bed"`
	DoubleBed           uint `json:"double_bed"`
	LargeDoubleBed      uint `json:"large_double_bed"`
	ExtraLargeDoubleBed uint `json:"extra_large_double_bed"`
}

type FacilitiesAccommodationDetail struct {
	WiFi         bool `json:"wifi"`
	AirCondition bool `json:"air_condition"`
	TV           bool `json:"tv"`
}

type GetAccommodationDetailsInput struct {
	AccommodationId string `json:"accommodation_id"`
}

type CreateAccommodationDetailInput struct {
	AccommodationId string                        `json:"accommodation_id" validate:"required"`
	Name            string                        `json:"name" validate:"required"`
	Guests          uint8                         `json:"guests" validate:"required"`
	Beds            Beds                          `json:"beds" validate:"required"`
	Facilities      FacilitiesAccommodationDetail `json:"facilities" validate:"required"`
	AvailableRooms  uint8                         `json:"available_rooms" validate:"required"`
	Price           float64                       `json:"price" validate:"required"`
	DiscountId      string                        `json:"discount_id"`
}

type CreateAccommodationDetailOutput struct {
	Id              string                        `json:"id"`
	AccommodationId string                        `json:"accommodation_id"`
	Name            string                        `json:"name"`
	Guests          uint8                         `json:"guests"`
	Beds            Beds                          `json:"beds"`
	Facilities      FacilitiesAccommodationDetail `json:"facilities"`
	AvailableRooms  uint8                         `json:"available_rooms"`
	Price           string                        `json:"price"`
	DiscountId      string                        `json:"discount_id"`
	Images          []string                      `json:"images"`
}

type GetAccommodationDetailsOutput struct {
	Id              string                        `json:"id"`
	AccommodationId string                        `json:"accommodation_id"`
	Name            string                        `json:"name"`
	Guests          uint8                         `json:"guests"`
	Beds            Beds                          `json:"beds"`
	Facilities      FacilitiesAccommodationDetail `json:"facilities"`
	AvailableRooms  uint8                         `json:"available_rooms"`
	Price           string                        `json:"price"`
	DiscountId      string                        `json:"discount_id"`
	Images          []string                      `json:"images"`
}

type UpdateAccommodationDetailInput struct {
	Id              string                        `json:"id" validate:"required"`
	AccommodationId string                        `json:"accommodation_id"`
	Name            string                        `json:"name"`
	Guests          uint8                         `json:"guests"`
	Beds            Beds                          `json:"beds"`
	Facilities      FacilitiesAccommodationDetail `json:"facilities"`
	AvailableRooms  uint8                         `json:"available_rooms"`
	Price           float64                       `json:"price"`
	DiscountId      string                        `json:"discount_id"`
}

type UpdateAccommodationDetailOutput struct {
	Id              string                        `json:"id"`
	AccommodationId string                        `json:"accommodation_id"`
	Name            string                        `json:"name"`
	Guests          uint8                         `json:"guests"`
	Beds            Beds                          `json:"beds"`
	Facilities      FacilitiesAccommodationDetail `json:"facilities"`
	AvailableRooms  uint8                         `json:"available_rooms"`
	Price           string                        `json:"price"`
	DiscountId      string                        `json:"discount_id"`
	Images          []string                      `json:"images"`
}

type DeleteAccommodationDetailInput struct {
	Id string `uri:"id" validate:"required"`
}
