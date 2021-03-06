package fetcher

import (
	"github.com/KyberNetwork/reserve-data/common"
)

// Exchange is the common interface of centralized exchanges.
type Exchange interface {
	ID() common.ExchangeID
	Name() string
	FetchPriceData(timepoint uint64) (map[common.TokenPairID]common.ExchangePrice, error)
	FetchEBalanceData(timepoint uint64) (common.EBalanceEntry, error)
	OrderStatus(id string, base, quote string) (string, error)
	DepositStatus(id common.ActivityID, txHash, currency string, amount float64, timepoint uint64) (string, error)
	WithdrawStatus(id, currency string, amount float64, timepoint uint64) (string, string, error)
}
