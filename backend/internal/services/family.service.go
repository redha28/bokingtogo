package services

import (
	"context"

	familydto "github.com/redha28/bookingtogo/internal/dto/family"
	"github.com/redha28/bookingtogo/internal/entities"
	"gorm.io/gorm"
)

type FamilyService struct {
	DB *gorm.DB
}

func NewFamilyService(db *gorm.DB) *FamilyService {
	return &FamilyService{DB: db}
}

func (s *FamilyService) CreateFamily(ctx context.Context, req familydto.CreateFamilyRequest) (*familydto.FamilyData, error) {
	family := entities.Family{
		CstID:      req.CstID,
		FlName:     req.FlName,
		FlDob:      req.FlDob,
		FlRelation: req.FlRelation,
	}

	if err := s.DB.WithContext(ctx).Create(&family).Error; err != nil {
		return nil, err
	}

	dto := familyToDto(family)
	return &dto, nil
}

func (s *FamilyService) GetFamiliesByCustomerID(ctx context.Context, customerID int) ([]familydto.FamilyData, error) {
	var families []entities.Family
	
	if err := s.DB.WithContext(ctx).Where("cst_id = ?", customerID).Find(&families).Error; err != nil {
		return nil, err
	}

	var familyDtos []familydto.FamilyData
	for _, family := range families {
		familyDtos = append(familyDtos, familyToDto(family))
	}

	return familyDtos, nil
}

func (s *FamilyService) GetFamilyByID(ctx context.Context, id int) (*familydto.FamilyData, error) {
	var family entities.Family
	
	if err := s.DB.WithContext(ctx).First(&family, id).Error; err != nil {
		return nil, err
	}

	dto := familyToDto(family)
	return &dto, nil
}

func (s *FamilyService) UpdateFamily(ctx context.Context, id int, req familydto.UpdateFamilyRequest) (*familydto.FamilyData, error) {
	var family entities.Family
	
	if err := s.DB.WithContext(ctx).First(&family, id).Error; err != nil {
		return nil, err
	}

	family.FlName = req.FlName
	family.FlDob = req.FlDob
	family.FlRelation = req.FlRelation

	if err := s.DB.WithContext(ctx).Save(&family).Error; err != nil {
		return nil, err
	}

	dto := familyToDto(family)
	return &dto, nil
}

func (s *FamilyService) DeleteFamily(ctx context.Context, id int) error {
	if err := s.DB.WithContext(ctx).Delete(&entities.Family{}, id).Error; err != nil {
		return err
	}
	return nil
}

func familyToDto(family entities.Family) familydto.FamilyData {
	return familydto.FamilyData{
		FlID:       family.FlID,
		CstID:      family.CstID,
		FlName:     family.FlName,
		FlDob:      family.FlDob,
		FlRelation: family.FlRelation,
	}
}
