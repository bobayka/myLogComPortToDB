package postgres

import (
	"database/sql"
	"github.com/pkg/errors"
	"myLogComPortToDB/internal/domains"
)

const (
	insertAccelDataQuery = `INSERT INTO Accel (x, y, z) VALUES ($1, $2, $3) RETURNING ID`
	insertGyroDataQuery  = `INSERT INTO Gyro (x, y, z, accel_id) VALUES ($1, $2, $3, $4)`
)

type Storage struct {
	statementStorage

	insertAccelDataStmt *sql.Stmt
	insertGyroDataStmt  *sql.Stmt
}

func NewStorage(db *sql.DB) (*Storage, error) {
	storage := &Storage{statementStorage: newStatementsStorage(db)}
	statements := []stmt{
		{Query: insertAccelDataQuery, Dst: &storage.insertAccelDataStmt},
		{Query: insertGyroDataQuery, Dst: &storage.insertGyroDataStmt},
	}
	if err := storage.initStatements(statements); err != nil {
		return nil, errors.Wrap(err, "can't create statements")
	}
	return storage, nil
}
func (u *Storage) InsertGyroAccelData(accelData domains.AccelData, gyroData domains.GyroData) error {
	var accID int64
	err := u.insertAccelDataStmt.QueryRow(accelData.X, accelData.Y, accelData.Z).Scan(&accID)
	if err != nil {
		return errors.Wrap(err, "cant insert accel data")
	}
	_, err = u.insertGyroDataStmt.Exec(gyroData.X, gyroData.Y, gyroData.Z, accID)
	if err != nil {
		return errors.Wrap(err, "cant insert gyro data")
	}
	return nil
}
