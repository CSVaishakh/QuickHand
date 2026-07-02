package models

type HandymanType string

const (
	Plumber        HandymanType = "plumber"
	Electrician    HandymanType = "electrician"
	Carpenter      HandymanType = "carpenter"
	Mason          HandymanType = "mason"
	Mechanic       HandymanType = "mechanic"
	HVACTechnician HandymanType = "hvac_technician"
	Landscaper     HandymanType = "landscaper"
	DeepCleaner    HandymanType = "deep_cleaner"
)

type Handyman struct {
	User
	Type HandymanType `gorm:"column:type;type:handyman_type;not null"`
}