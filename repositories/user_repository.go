package repositories

import (
	"database/sql"

	"wthunder/bstack/models"
)

type UserRepositoryStruct struct {
	DBDao *sql.DB
}

func (userRepo *UserRepositoryStruct) CreateUser(player models.Player) error {
	dbConn := userRepo.DBDao

	query := "INSERT INTO player (nano_id,username,password) VALUES (?,?,?);"
	_, dbErr := dbConn.Exec(query, player.NanoId, player.Username, player.Password)
	return dbErr
}

func (userRepo *UserRepositoryStruct) GetUserWithPassHash(username string) (models.Player, error) {
	var err error
	dbConn := userRepo.DBDao

	player := models.Player{}
	query := "SELECT username,nano_id,password FROM player WHERE username=?;"
	row := dbConn.QueryRow(query, username)
	err = row.Scan(&player.Username, &player.NanoId,&player.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return player, err
		}
	}
	return player, err
}

func (userRepo *UserRepositoryStruct) GetUsers() ([]models.Player, error) {
	dbConn := userRepo.DBDao
	
	query := "SELECT username,nano_id FROM player;"
	
	var players []models.Player
	rows, err := dbConn.Query(query)
	if err != nil {
		return players, err
	}
	for rows.Next() {
		var player models.Player
		rows.Scan(&player.Username, &player.NanoId)
		players = append(players, player)
	}

	return players, nil
}
