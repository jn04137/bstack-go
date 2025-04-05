package repositories

import (
	"log"
	"database/sql"

	"wthunder/bstack/models"
)

type TeamRepositoryStruct struct {
	DBBao *sql.DB
}

type TeamAndOwner struct {
	TeamId int
	TeamNanoId string
	TeamName string
	TeamDesc string
	TeamCreatedAt string
	TeamUpdatedAt string
	OwnerUsername string
}

func (teamRepo *TeamRepositoryStruct) CreateTeam() error {
	var err error

	return err
}

func (teamRepo *TeamRepositoryStruct) GetTeams() ([]TeamAndOwner, error) {
	dbConn := teamRepo.DBBao
	teams := []TeamAndOwner{}
	var err error = nil

	query := `SELECT t.id, t.nano_id, t.team_name, t.team_desc, t.created_at, t.updated_at, p.username 
		FROM team as t INNER JOIN player as p ON p.id=t.team_owner`

		rows, err := dbConn.Query(query)
		for rows.Next() {
			var tO TeamAndOwner
			rows.Scan(&tO.TeamId, &tO.TeamNanoId, &tO.TeamName, &tO.TeamDesc, &tO.TeamCreatedAt, &tO.TeamUpdatedAt, &tO.OwnerUsername)
			teams = append(teams, tO)
		}

	return teams, err
}

func (teamRepo *TeamRepositoryStruct) GetTeam(teamNanoId string) (TeamAndOwner, error) {
	dbConn := teamRepo.DBBao
	var err error
	var team TeamAndOwner

	query := `SELECT team_name, team_desc FROM team WHERE team.nano_id=?`
	row := dbConn.QueryRow(query, teamNanoId)
	err = row.Scan(&team.TeamName, &team.TeamDesc)

	return team, err
}

func (teamRepo *TeamRepositoryStruct) GetPlayersOnTeam(teamNanoId string) ([]models.Player, error) {
	dbConn := teamRepo.DBBao
	var players []models.Player

	// Get users from the database given the teamNanoId
	query := `SELECT p.username FROM user_on_team as uot 
		INNER JOIN team as t ON uot.team=t.id 
		INNER JOIN player as p on uot.player=p.id 
		WHERE t.nano_id=(?);`

	rows, err := dbConn.Query(query, teamNanoId)
	if err != nil {
		return nil, err
	}

	log.Printf("These are the rows: %v", rows)

	for rows.Next() {
		var p models.Player
		err = rows.Scan(&p.Username)
		if err != nil {
			return nil, err
		}
		players = append(players, p)
	}

	return players, nil
}


