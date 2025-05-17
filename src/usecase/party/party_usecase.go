package party

import "github.com/MizukiShigi/go_pokemon/domain"

type PartyUsecase struct {
	pr domain.IPartyRepository
}

func NewPartyUsecase(pr domain.IPartyRepository) domain.IPartyUsecase {
	return &PartyUsecase{pr}
}

func (pu *PartyUsecase) AddParty(partyList []domain.Party) error {
	err := pu.pr.AddParty(partyList)
	if err != nil {
		return err
	}
	return nil
}