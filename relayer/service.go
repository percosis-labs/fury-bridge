package relayer

type Service struct {
}

func NewService(ethRpc string, ethBridgeAddr string, furyGrpc string, relayerMnemonic string) (*Service, error) {
	return &Service{}, nil
}

func (s *Service) Run() error {
	return nil
}
