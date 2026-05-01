type EaterService interface {
	SignupEater(ctx context.Context, phoneNumber string) (string, error)
	ConfirmSMSCode(ctx context.Context, eaterID, smsCode string) (*models.EaterProfile, error)
	GetEaterProfile(ctx context.Context, eaterID string) (*models.EaterProfile, error)
	UpdateEaterProfile(ctx context.Context, eaterID, name, imageUrl string) (*models.EaterProfile, error)
}

type eaterSvcImpl struct {
	eaterRepo repositories.EaterRepository
	smsClient sms.Client
	logger    *zap.Logger
}

func NewEaterService(eaterRepo repositories.EaterRepository, smsClient sms.Client, logger *zap.Logger) EaterService {
	return &eaterSvcImpl{
		eaterRepo: eaterRepo,
		smsClient: smsClient,
		logger:    logger,
	}
}

func (s *eaterSvcImpl) SignupEater(ctx context.Context, phoneNumber string) (string, error) {
	phoneNumberExist := true

	eater, err := s.eaterRepo.GetEaterByPhoneNumber(ctx, phoneNumber)
	if err != nil {
		phoneNumberExist = false
	}
	if phoneNumberExist {
		return s.handleExistingEater(ctx, eater.ID)
	}

	return s.handleNewEater(ctx, phoneNumber)
}

func (s *eaterSvcImpl) handleNewEater(ctx context.Context, phoneNumber string) (string, error) {
	var (
		eaterID = rand. UUID()
		eaterName = fmt.Sprintf("eater-%s", rand.NumbericString(5))
		salt = crypto.GenerateSalt()
		saltedPass = crypto.Combine(salt, phoneNumber)
		passHash = crypto.HashPassword(saltedPass)
		now = time.Now().UTC()
	)

	eater := models.Eater{
		ID:        eaterID,
		PhoneNumber: phoneNumber,
		PasswordHash: passHash,
		PasswordSalt: salt,
		CreatedAt: now,
		UpdatedAt: now,
	}

	eaterProfile := models.EaterProfile{
		EaterID: eaterID,
		Name:      eaterName,
		PhoneNumber: phoneNumber,
		ImageURL: "",
		CreatedAt: now,
		UpdatedAt: now,
	}

	smsCode := models.EaterSmsCode{
		EaterID: eaterID,
		Code: = rand.NumericString(5),
		ExpiresIn: 5 * 60, // 5 minutes
		CreatedAt: now,
	}

	err := s.eaterRepo.WithTx(ctx, func(r repositories.EaterRepository) error {
		if err := r.SaveEater(ctx, &eater); err != nil {
			return err
		}
		
		if err := r.SaveEaterProfile(ctx, &eaterProfile); err != nil {
			return err
		}

		if err := r.SaveEaterSmsCode(ctx, &smsCode); err != nil {
			return err
		}

		smsMsg := fmt.Sprintf("Food.com code is: %s", smsCode.Code)
		if err := s.smsClient.SendMessage(ctx, phoneNumber, smsMsg); err != nil {
			return "", err
		}
		return eaterID, nil
	})

	if err != nil {
		s.logger.Error("failed to create new eater", zap.Error(err))
		return "", err
	}

	return eater.ID, nil

}

func (s *eaterSvcImpl) handleExistingEater(ctx context.Context, eaterID string) (string, error) {
	eater, err := s.eaterRepo.GetEater(ctx, eaterID)
	if err != nil {
		s.logger.Error("failed to get eater by id", zap.String("eater_id", eaterID), zap.Error(err))
		return "", err
	}

	smsCode := models.EaterSmsCode{
		EaterID: eaterID,
		Code: = rand.NumericString(5),
		ExpiresIn: 5 * 60, // 5 minutes
		CreatedAt: time.Now().UTC(),
	}

	if err := s.eaterRepo.SaveEaterSmsCode(ctx, &smsCode); err != nil {
		s.logger.Error("failed to save sms code for existing eater", zap.String("eater_id", eaterID), zap.Error(err))
		return "", err
	}

	smsMsg := fmt.Sprintf("Food.com code is: %s", smsCode.Code)
	if err := s.smsClient.SendMessage(ctx, eater.PhoneNumber, smsMsg); err != nil {
		s.logger.Error("failed to send sms code for existing eater", zap.String("eater_id", eaterID), zap.Error(err))
		return "", err
	}
	return eaterID, nil
}

func (s *eaterSvcImpl) ConfirmSMSCode(ctx context.Context, eaterID, smsCode string) (*models.EaterProfile, error) {
	smsCode, err := s.eaterRepo.GetEaterSmsCode(ctx, eaterID, code);
	
	if err != nil {
		s.logger.Error("failed to get sms code for eater", zap.String("eater_id", eaterID), zap.Error(err))
		return nil, err
	}

	return e, nil
}

func (s *eaterSvcImpl) GetEaterProfile(ctx context.Context, eaterID string) (*models.EaterProfile, error) {
	return nil, nil
}

func (s *eaterSvcImpl) UpdateEaterProfile(ctx context.Context, eaterID, name, imageUrl string) (*models.EaterProfile, error) {
	return nil, nil
}
