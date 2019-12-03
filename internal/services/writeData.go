package services

import (
	"github.com/pkg/errors"
	"myLogComPortToDB/internal/domains"
	"myLogComPortToDB/internal/postgres"
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
