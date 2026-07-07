package addressService

import (
	"github.com/CSVaishakh/QuickHand/packages/db/models"
	"github.com/google/uuid"
)

type AddAddressReq struct{
	UserId 		uuid.UUID
	HouseNo 		string
	Street  		string 
	City    		string 
	State  		string 
	Country 		string 
	Pincode 		string 
}

type AddAddressRes struct{
	Address models.Address
}

type GetAddressesReq struct{
	UserId 	uuid.UUID
}

type GetAddressesRes struct{
	Addresses []models.Address
}

type UpdateAddressReq struct {
	AddressID 	uuid.UUID
	UserId 		uuid.UUID
	HouseNo 		string
	Street  		string 
	City    		string 
	State  		string 
	Country 		string 
	Pincode 		string 
}

type UpdateAddressRes struct {
	Address models.Address
}