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
		return s.handleExistingPhoneNumber(ctx, eater.ID)
	}

	return s.handleNewPhoneNumber(ctx, phoneNumber)
}

func (s *eaterSvcImpl) ConfirmSMSCode(ctx context.Context, eaterID, smsCode string) (*models.EaterProfile, error) {
	return nil, nil
}

func (s *eaterSvcImpl) GetEaterProfile(ctx context.Context, eaterID string) (*models.EaterProfile, error) {
	return nil, nil
}

func (s *eaterSvcImpl) UpdateEaterProfile(ctx context.Context, eaterID, name, imageUrl string) (*models.EaterProfile, error) {
	return nil, nil
}
