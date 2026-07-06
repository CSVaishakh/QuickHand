package addressService

import (
	"github.com/CSVaishakh/QuickHand/src/packages/db/models"
	repo "github.com/CSVaishakh/QuickHand/src/packages/db/repositories"
)

type AddressService struct {
	addressRepo		*repo.AddressRepository
}

func NewAddressService (
	addressRepo 	*repo.AddressRepository,
) *AddressService{
	return &AddressService{
		addressRepo: 	addressRepo,
	}
}

func (s *AddressService) AddNewAddress(req AddAddressReq)(AddAddressRes, error){
	
	address := models.Address{
		UserID: 		req.UserId,
		HouseNo: 	req.HouseNo,
		Street: 		req.Street,
		City: 		req.City,
		State: 		req.State,
		Country: 	req.Country,
		Pincode: 	req.Pincode,
	}
	
	err := s.addressRepo.AddAddress(&address)
	if err != nil {
		return AddAddressRes{}, err
	}

	return AddAddressRes{
		Address: address,
	}, nil
}

func (s *AddressService) UpdateAddress (req UpdateAddressReq) (UpdateAddressRes, error){
	
	address := models.Address{
		AddressID: 	req.AddressID,
		UserID: 		req.UserId,
		HouseNo: 	req.HouseNo,
		Street: 		req.Street,
		City: 		req.City,
		State: 		req.State,
		Country: 	req.Country,
		Pincode: 	req.Pincode,
	}

	address_retrived, err := s.addressRepo.GetByAddressID(req.AddressID)
	if err != nil {
		return UpdateAddressRes{}, err
	}

	if address_retrived.UserID != address.UserID {
		return UpdateAddressRes{}, ErrAddressNotFoundForUser
	}
	
	err = s.addressRepo.UpdateAddress(&address)
	if err != nil {
		return UpdateAddressRes{}, err
	}
	
	return UpdateAddressRes{
		Address:	address,
	}, nil
}

func (s *AddressService) GetAddresses(req GetAddressesReq)(GetAddressesRes, error){
	addresses, err := s.addressRepo.GetAddresses(req.UserId)
	if err != nil {
		return GetAddressesRes{}, err
	}

	return GetAddressesRes{
		Addresses: addresses,
	}, nil
}