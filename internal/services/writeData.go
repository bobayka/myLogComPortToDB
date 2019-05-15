package services

import (
	"github.com/bobayka/myLogComPortToDB/internal/domains"
	"github.com/bobayka/myLogComPortToDB/internal/postgres"
	"github.com/pkg/errors"
)

type WriteDataService struct {
	StmtsStorage *postgres.Storage
}

func NewWriteDataService(StmtsStorage *postgres.Storage) *WriteDataService {
	return &WriteDataService{StmtsStorage: StmtsStorage}
}
func (wds *WriteDataService) WriteToDB(data *domains.GyroAccelData) error {
	err := wds.StmtsStorage.InsertGyroAccelData(data.AccelData, data.GyroData)
	if err != nil {
		return errors.Wrap(err, "cant insert gyro accel data")
	}
	return nil
}
