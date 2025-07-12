package services

import (
	"context"

	customerdto "github.com/redha28/bookingtogo/backend/internal/dto/customer"
	"github.com/redha28/bookingtogo/backend/internal/entities"

	"gorm.io/gorm"
)

type CustomerService struct {
	DB *gorm.DB
}

func NewCustomerService(db *gorm.DB) *CustomerService {
	return &CustomerService{DB: db}
}

func (s *CustomerService) CreateCustomer(ctx context.Context, req customerdto.CreateCustomerRequest) (*customerdto.CustomerData, error) {
	customer := entities.Customer{
		CstName:       req.CstName,
		CstEmail:      req.CstEmail,
		CstDob:        req.CstDob,
		CstPhoneNum:   req.CstPhoneNum,
		NationalityID: req.NationalityID,
	}

	for _, f := range req.FamilyList {
		customer.FamilyList = append(customer.FamilyList, entities.Family{
			FlName:     f.FlName,
			FlDob:      f.FlDob,
			FlRelation: f.FlRelation,
		})
	}

	if err := s.DB.WithContext(ctx).Create(&customer).Error; err != nil {
		return nil, err
	}

	dto := customerToCustomerDto(customer)
	return &dto, nil
}

func (s *CustomerService) GetCustomers(ctx context.Context) ([]customerdto.CustomerData, error) {
	var customers []entities.Customer
	
	if err := s.DB.WithContext(ctx).Preload("FamilyList").Find(&customers).Error; err != nil {
		return nil, err
	}

	var customerDtos []customerdto.CustomerData
	for _, customer := range customers {
		customerDtos = append(customerDtos, customerToCustomerDto(customer))
	}

	return customerDtos, nil
}

func (s *CustomerService) GetCustomerByID(ctx context.Context, id int) (*customerdto.CustomerData, error) {
	var customer entities.Customer
	
	if err := s.DB.WithContext(ctx).Preload("FamilyList").First(&customer, id).Error; err != nil {
		return nil, err
	}

	dto := customerToCustomerDto(customer)
	return &dto, nil
}

func (s *CustomerService) UpdateCustomer(ctx context.Context, id int, req customerdto.CreateCustomerRequest) (*customerdto.CustomerData, error) {
	var customer entities.Customer
	
	if err := s.DB.WithContext(ctx).First(&customer, id).Error; err != nil {
		return nil, err
	}

	// Update customer fields
	customer.CstName = req.CstName
	customer.CstEmail = req.CstEmail
	customer.CstDob = req.CstDob
	customer.CstPhoneNum = req.CstPhoneNum
	customer.NationalityID = req.NationalityID

	// Delete existing family members
	s.DB.WithContext(ctx).Where("cst_id = ?", id).Delete(&entities.Family{})

	// Add new family members
	customer.FamilyList = nil
	for _, f := range req.FamilyList {
		customer.FamilyList = append(customer.FamilyList, entities.Family{
			FlName:     f.FlName,
			FlDob:      f.FlDob,
			FlRelation: f.FlRelation,
		})
	}

	if err := s.DB.WithContext(ctx).Save(&customer).Error; err != nil {
		return nil, err
	}

	dto := customerToCustomerDto(customer)
	return &dto, nil
}

func (s *CustomerService) DeleteCustomer(ctx context.Context, id int) error {
	// Delete family members first
	if err := s.DB.WithContext(ctx).Where("cst_id = ?", id).Delete(&entities.Family{}).Error; err != nil {
		return err
	}

	// Delete customer
	if err := s.DB.WithContext(ctx).Delete(&entities.Customer{}, id).Error; err != nil {
		return err
	}

	return nil
}

func customerToCustomerDto(customer entities.Customer) customerdto.CustomerData {
	dto := customerdto.CustomerData{
		CstID:         customer.CstID,
		CstName:       customer.CstName,
		CstEmail:      customer.CstEmail,
		CstDob:        customer.CstDob,
		CstPhoneNum:   customer.CstPhoneNum,
		NationalityID: customer.NationalityID,
	}

	for _, fam := range customer.FamilyList {
		dto.FamilyList = append(dto.FamilyList, customerdto.FamilyResponse{
			FlName:     fam.FlName,
			FlDob:      fam.FlDob,
			FlRelation: fam.FlRelation,
		})
	}

	return dto
}
