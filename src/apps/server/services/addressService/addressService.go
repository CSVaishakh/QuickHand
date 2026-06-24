package addressService

import (
	"github.com/CSVaishakh/QuickHand/src/packages/db/models"
	repo "github.com/CSVaishakh/QuickHand/src/packages/db/repositories"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AddressService struct {
	addressRepo		*repo.AddressRepository
	db				*gorm.DB
}

func NewAddressService (
	addressRepo 	*repo.AddressRepository,
	db				*gorm.DB,
) *AddressService{
	return &AddressService{
		addressRepo: addressRepo,
		db: db,
	}
}

func (s *AddressService) AddNewAddress(req AddAddressReq)(AddAddressRes, error){
	id, err := uuid.Parse(req.UserId)
	if err != nil {
		return AddAddressRes{}, err
	}
	
	address := models.Address{
		UserID: id,
		HouseNo: req.HouseNo,
		Street: req.Street,
		City: req.City,
		State: req.State,
		Country: req.Country,
		Pincode: req.Pincode,
	}
	
	err = s.addressRepo.AddAddress(&address,s.db)
	if err != nil {
		return AddAddressRes{}, err
	}

	return AddAddressRes{
		Address: address,
	}, nil
}

func (s *AddressService) UpdateAddress (req UpdateAddressReq) (UpdateAddressRes, error){
	address_id,err := uuid.Parse(req.AddressID)
	if err != nil {
		return UpdateAddressRes{}, err
	}
	user_id, err := uuid.Parse(req.UserId)
	if err != nil {
		return UpdateAddressRes{}, err
	}
	
	address := models.Address{
		AddressID: address_id,
		UserID: user_id,
		HouseNo: req.HouseNo,
		Street: req.Street,
		City: req.City,
		State: req.State,
		Country: req.Country,
		Pincode: req.Pincode,
	}
	err = s.addressRepo.UpdateAddress(&address,s.db)
	if err != nil {
		return UpdateAddressRes{}, err
	}

	address_retrived, err := s.addressRepo.GetByAddressID(req.AddressID, s.db)
	if err != nil {
		return UpdateAddressRes{}, err
	}
	
	if address != address_retrived {
		return UpdateAddressRes{}, ErrAddressUpdateFailed
	}
	
	return UpdateAddressRes{
		Address:	address_retrived,
	}, nil
}

func (s *AddressService) GetAddresses(req GetAddressesReq)(GetAddressesRes, error){
	addresses, err := s.addressRepo.GetAddresses(req.UserId, s.db)
	if err != nil {
		return GetAddressesRes{}, err
	}

	return GetAddressesRes{
		Addresses: addresses,
	}, nil
}