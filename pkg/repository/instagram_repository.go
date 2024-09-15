package repository

type InstagramRepository struct {
}

func NewInstagramRepository() *InstagramRepository {
	return &InstagramRepository{}
}

func (i *InstagramRepository) SendDirectMessage(message string, toUsername string) error {
	return nil
}
