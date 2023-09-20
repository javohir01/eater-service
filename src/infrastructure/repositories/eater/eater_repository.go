package eater

import (
	"context"

	"github.com/javohir01/eater-service/src/domain/eater/models"
	"github.com/javohir01/eater-service/src/domain/eater/repositories"
	"gorm.io/gorm"
)

const (
	tableEaters        = "eater.eaters"
	tableEaterSmsCodes = "eater.eater_sms_codes"
	tableEaterProfiles = "eater.eater_profiles"
)

type eaterRepoImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repositories.EaterRepository {
	return &eaterRepoImpl{
		db: db,
	}
}
func (r *eaterRepoImpl) WithTx(ctx context.Context, f func(r repositories.EaterRepository) error) error {
	if err := r.db.Transaction(func(tx *gorm.DB) error {
		r := eaterRepoImpl{
			db: tx,
		}
		if err := f(&r); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
func (r *eaterRepoImpl) SaveEater(ctx context.Context, eater *models.Eater) error {
	result := r.db.WithContext(ctx).Table(tableEaters).Create(eater)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *eaterRepoImpl) UpdateEater(ctx context.Context, eater *models.Eater) error {
	result := r.db.WithContext(ctx).Table(tableEaters).Save(eater)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *eaterRepoImpl) DeleteEater(ctx context.Context, eaterID string) error {
	var eater models.Eater
	result := r.db.WithContext(ctx).Table(tableEaters).Delete(&eater, "eater_id = ?", eaterID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *eaterRepoImpl) GetEater(ctx context.Context, eaterID string) (*models.Eater, error) {
	var eater models.Eater
	result := r.db.WithContext(ctx).Table(tableEaters).First(&eater, "eater_id = ?", eaterID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &eater, nil
}
func (r *eaterRepoImpl) GetEaterByPhoneNumber(ctx context.Context, phoneNumber string) (*models.Eater, error) {
	var eater models.Eater
	result := r.db.WithContext(ctx).Table(tableEaters).First(&eater, "phone_number = ?", phoneNumber)
	if result.Error != nil {
		return nil, result.Error
	}
	return &eater, nil
}
func (r *eaterRepoImpl) SaveEaterSmsCode(ctx context.Context, smsCode *models.EaterSmsCode) error {
	result := r.db.WithContext(ctx).Table(tableEaterSmsCodes).Create(smsCode)
	if result.Error != nil {
		return result.Eater
	}
	return nil

}
func (r *eaterRepoImpl) GetEaterSmsCode(ctx context.Context, eaterID, code string) (*models.EaterSmsCode, error) {
	var smsCode models.EaterSmsCode
	result := r.db.WithContext(ctx).Table(tableEaterSmsCodes).Where("eater_id = ? and code = ?", eaterID, code).First(&smsCode)
	if result.Error != nil {
		return nil, result.Error
	}
	return &smsCode, nil
}
func (r *eaterRepoImpl) SaveEaterProfile(ctx context.Context, profile *models.EaterProfile) error {
	result := r.db.WithContext(ctx).Table(tableEaterProfiles).Create(profile)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *eaterRepoImpl) UpdateEaterProfile(ctx context.Context, profile *models.EaterProfile) error {
	result := r.db.WithContext(ctx).Table(tableEaterProfiles).Save(profile)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *eaterRepoImpl) UpdateEaterProfilePhoneNumberConfirmed(ctx context.Context, eaterID string, confirmed bool) error {
	result := r.db.WithContext(ctx).Table(tableEaterProfiles).Where("eater_id = ?", eaterID).Update("is_phone_number_confirmed = ", confirmed)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *eaterRepoImpl) DeleteEaterProfile(ctx context.Context, eaterID string) error {
	var profile models.EaterProfile
	result := r.db.WithContext(ctx).Table(tableEaterProfiles).Delete(&profile, "eater_id = ?", eaterID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *eaterRepoImpl) GetEaterProfile(ctx context.Context, eaterID string) (*models.EaterProfile, error) {
	var profile models.EaterProfile
	result := r.db.WithContext(ctx).Table(tableEaterProfiles).First(&profile, "eater_id = ?", eaterID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &profile, nil
}
