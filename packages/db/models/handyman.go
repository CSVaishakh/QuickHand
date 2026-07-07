package models

type HandymanType string

const (
	Plumber        	HandymanType = "plumber"
	Electrician    	HandymanType = "electrician"
	Carpenter      	HandymanType = "carpenter"
	Mason          	HandymanType = "mason"
	Mechanic       	HandymanType = "mechanic"
	HVACTechnician 	HandymanType = "hvac_technician"
	Landscaper     	HandymanType = "landscaper"
	DeepCleaner    	HandymanType = "deep_cleaner"
)

func (h HandymanType) MapJobType () JobType {
	switch h {
	   case Plumber:
	      return Plumbing
	   case Electrician:
	      return Electrical
	   case Carpenter:
	      return Carpentry
	   case Mason:
	      return Masonry
	   case Mechanic:
	      return Mechanical
	   case HVACTechnician:
	      return HVAC
	   case Landscaper:
	      return Landscaping
	   case DeepCleaner:
	      return DeepCleaning
	   default:
	      return ""
   }
}

type Handyman struct {
	User
	Type 	HandymanType `gorm:"column:type;type:handyman_type;not null"`
}