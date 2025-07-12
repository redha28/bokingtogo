package services

import (
	"context"

	nationalitydto "github.com/redha28/bookingtogo/internal/dto/nationality"
	"github.com/redha28/bookingtogo/internal/entities"
	"gorm.io/gorm"
)

type NationalityService struct {
	DB *gorm.DB
}

func NewNationalityService(db *gorm.DB) *NationalityService {
	return &NationalityService{DB: db}
}

func (s *NationalityService) CreateNationality(ctx context.Context, req nationalitydto.CreateNationalityRequest) (*nationalitydto.NationalityData, error) {
	nationality := entities.Nationality{
		NationalityName: req.NationalityName,
		NationalityCode: req.NationalityCode,
	}

	if err := s.DB.WithContext(ctx).Create(&nationality).Error; err != nil {
		return nil, err
	}

	dto := nationalityToDto(nationality)
	return &dto, nil
}

func (s *NationalityService) GetNationalities(ctx context.Context) ([]nationalitydto.NationalityData, error) {
	var nationalities []entities.Nationality
	
	if err := s.DB.WithContext(ctx).Find(&nationalities).Error; err != nil {
		return nil, err
	}

	var nationalityDtos []nationalitydto.NationalityData
	for _, nationality := range nationalities {
		nationalityDtos = append(nationalityDtos, nationalityToDto(nationality))
	}

	return nationalityDtos, nil
}

func (s *NationalityService) GetNationalityByID(ctx context.Context, id int) (*nationalitydto.NationalityData, error) {
	var nationality entities.Nationality
	
	if err := s.DB.WithContext(ctx).First(&nationality, id).Error; err != nil {
		return nil, err
	}

	dto := nationalityToDto(nationality)
	return &dto, nil
}

func (s *NationalityService) UpdateNationality(ctx context.Context, id int, req nationalitydto.UpdateNationalityRequest) (*nationalitydto.NationalityData, error) {
	var nationality entities.Nationality
	
	if err := s.DB.WithContext(ctx).First(&nationality, id).Error; err != nil {
		return nil, err
	}

	nationality.NationalityName = req.NationalityName
	nationality.NationalityCode = req.NationalityCode

	if err := s.DB.WithContext(ctx).Save(&nationality).Error; err != nil {
		return nil, err
	}

	dto := nationalityToDto(nationality)
	return &dto, nil
}

func (s *NationalityService) DeleteNationality(ctx context.Context, id int) error {
	if err := s.DB.WithContext(ctx).Delete(&entities.Nationality{}, id).Error; err != nil {
		return err
	}
	return nil
}

func nationalityToDto(nationality entities.Nationality) nationalitydto.NationalityData {
	return nationalitydto.NationalityData{
		NationalityID:   nationality.NationalityID,
		NationalityName: nationality.NationalityName,
		NationalityCode: nationality.NationalityCode,
	}
}
