package repositories

import (
	"database/sql"
)

type TeamRepositoryStruct struct {
	DBBao *sql.DB
}

func (teamRepo *TeamRepositoryStruct) CreateTeam() error {
	var err error

	return err
}
