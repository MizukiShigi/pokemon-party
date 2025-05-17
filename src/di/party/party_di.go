package party

import (
	"database/sql"

	_partyHandler "github.com/MizukiShigi/go_pokemon/handler/party"
	"github.com/MizukiShigi/go_pokemon/repository"
	_partyUsecase "github.com/MizukiShigi/go_pokemon/usecase/party"
)

func InitParty(db *sql.DB) _partyHandler.IPartyHandler {
	partyRepository := repository.NewPartyRepository(db)
	partyUsecase := _partyUsecase.NewPartyUsecase(partyRepository)
	return _partyHandler.NewPartyHandler(partyUsecase)
}
