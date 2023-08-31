type EaterService interface {
	SignupEater(ctx context.Context, phoneNumber string) (string, error)
	ConfirmSMSCode(ctx context.Context, eaterID, smsCode string) (*models.EaterProfile, error)
	GetEaterProfile(ctx context.Context, eaterID string) (*models.EaterProfile, error)
	UpdateEaterProfile(ctx context.Context, eaterID, name, imageUrl string) (*models.EaterProfile, error)
	}


func (s * eaterSvcImpl) handleNewEater(ctx context,Context, phoneNumber) {
	var (
		eaterID = rand.UUID()
		eaterName = fmt.Sprintf("eater-%", rand.NumericString(5))
		salt = crypto.GenerateSalt()
		saltedPass = crypto.Combine(salt, phoneNumber)
		passHash = crypto.HashPassword(saltedPass)
		now = time.Now().UTC()
	)

	eater := models.Eater{
		ID: eaterID
		PhoneNumber: phoneNumber,
		PasswordHash: passHash,
		PasswordSalt: salt,
		CreatedAt: now,
		UpdatedAt: now,
	}

	EaterProfile := models.EaterProfile{
		EaterID: eaterID,
		PhoneNumber: phoneNumber,
		Name: eaterName,
		ImageUrl: "",
		CreatedAt: now,
		UpdatedAt: now,
	}

	smsCode := models.EaterSmsCode{
		EaterID: eaterID,
		Code: rand.NumericString(5),
		ExpiresIn: 300,
		CreatedAt: now,
	}

	err := s.eaterRepo.WithTx(ctx, func(r repositories.EaterRepository) error {
		if err := s.eaterRepo.SaveEater(ctx, &eater): err != nil {
			return err
		}

		if err := r.SaveEaterProfile(ctx, &eaterProfile); err != nil {
			return err
		}

		if err := r.SaveEaterSmsCode(ctx, $smsCode); err != nil {
			return err
		}
		return nil
	}) 	
	if err != nil {
		return "", err
	}	

	smsMsg := fmt.Sprintf("Food.uz code: %s", smsCode.Code)
	if err := s.smsClient.SendMessage(ctx, phoneNumber, smsMsg); err != nil {
		return "", err
	}

	return eaterID, nil
}

func (s *eaterSvcImpl) handleExistingEater(ctx, context.Context, eaterID string) {
	eater, err := s.eaterRepo.GetEater(ctx, eaterID)

	if err != nil {
		return "", err
	}	

	smsCode := models.EaterSmsCode{
		EaterID: eaterID,
		Code: rand.NumericString(5),
		ExpiresIn: 300,
		CreatedAt: now,
	}
	if err := s.eaterRepo.SaveEaterSmsCode(ctx, &smsCode); if err != nil {
		return "", err
	}	

	smsMsg := fmt.Sprintf("Food.uz code: %s", smsCode.Code)
	
	if err := s.smsClient.SendMessage(ctx, phoneNumber, smsMsg); err != nil {
		return "", err
	}

	return eaterID, nil
}

