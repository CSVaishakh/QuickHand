package addressService

import "github.com/CSVaishakh/QuickHand/src/packages/db/models"

type AddAddressReq struct{
	UserId 		string
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
	UserId 	string
}

type GetAddressesRes struct{
	Addresses []models.Address
}

type UpdateAddressReq struct {
	AddressID 	string
	UserId 		string
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