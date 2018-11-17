package faker

// VehicleInfo contains all the major features of a motorized vehicle
type VehicleInfo struct {
	// Vehicle type
	VehicleType string
	// Fuel type
	Fuel string
	// Transmission type
	TransmissionGear string
	// Brand name
	Brand string
	// Vehicle model
	Model string
	// Vehicle model year
	Year int
}

// Vehicle will generate a struct with vehicle information
func (f *Faker) Vehicle() *VehicleInfo {
	return &VehicleInfo{
		VehicleType:      f.VehicleType(),
		Fuel:             f.FuelType(),
		TransmissionGear: f.TransmissionGearType(),
		Brand:            f.CarMaker(),
		Model:            f.CarModel(),
		Year:             f.Year(),
	}

}

// VehicleType will generate a random vehicle type string
func (f *Faker) VehicleType() string {
	return f.getRandValue([]string{"vehicle", "vehicle_type"})
}

// FuelType will generate a random fuel type
func (f *Faker) FuelType() string {
	return f.getRandValue([]string{"vehicle", "fuel_type"})
}

// TransmissionGearType will generate a random transmission type
func (f *Faker) TransmissionGearType() string {
	return f.getRandValue([]string{"vehicle", "transmission_type"})
}

// CarMaker will generate a random car manufacturer
func (f *Faker) CarMaker() string {
	return f.getRandValue([]string{"vehicle", "maker"})
}

// CarModel will generate a random car model
func (f *Faker) CarModel() string {
	return f.getRandValue([]string{"vehicle", "model"})
}
