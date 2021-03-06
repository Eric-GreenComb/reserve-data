package statpruner

import (
	"github.com/KyberNetwork/reserve-data/common/archive"
)

const (
	EXPIRED_ANALYTIC_PATH string = "expired-analytic-data/"
)

type StorageController struct {
	Runner                   ControllerRunner
	Arch                     archive.Archive
	ExpiredPriceAnalyticPath string
}

func NewStorageController(storageControllerRunner ControllerRunner, arch archive.Archive) (StorageController, error) {
	storageController := StorageController{
		storageControllerRunner, arch, EXPIRED_ANALYTIC_PATH,
	}
	return storageController, nil
}
