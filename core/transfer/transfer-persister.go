package transfer

import (
	"context"

	"github.com/NicolasMartino/simplebank/core/accounts"
	db "github.com/NicolasMartino/simplebank/db/sqlc"
)

type TransferPersister struct {
	store            db.Store
	accountRetriever *accounts.AccountRetriever
}

type TransferDTO struct {
	FromAccountID int64  `json:"from_account_id" binding:"required,min=1"`
	ToAccountID   int64  `json:"to_account_id" binding:"required,min=1"`
	Amount        int64  `json:"amount" binding:"required,gt=0"`
	Currency      string `json:"currency" binding:"required,oneof=USD EUR CAD"`
}

func NewTransferPersister(store db.Store, accountRetriever *accounts.AccountRetriever) *TransferPersister {
	transferPersister := &TransferPersister{
		store:            store,
		accountRetriever: accountRetriever,
	}

	return transferPersister
}

func (transferPersister *TransferPersister) CreateTransfer(ctx context.Context, dto TransferDTO) (result db.TransferTxResult, err error) {
	// validate accounts
	if err = transferPersister.accountRetriever.ValidateAccount(ctx, dto.FromAccountID, dto.Currency); err != nil {
		return
	}
	if err = transferPersister.accountRetriever.ValidateAccount(ctx, dto.ToAccountID, dto.Currency); err != nil {
		return
	}

	args := db.TransferTxParams{
		FromAccountID: dto.FromAccountID,
		ToAccountID:   dto.ToAccountID,
		Amount:        float64(dto.Amount),
	}
	return transferPersister.store.TransferTx(ctx, args)
}
