package repositories

import (
	"database/sql"

	"wthunder/bstack/models"
)

type UserRepositoryStruct struct {
	DBDao *sql.DB
}

type UserWithDetailsStruct struct {
	NanoId string
	Username string
	Details string
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

func (userRepo *UserRepositoryStruct) GetUser(playerNanoId string) (UserWithDetailsStruct, error) {
	dbConn := userRepo.DBDao

	var err error
	var player UserWithDetailsStruct
	
	query := "SELECT p.username,p.nano_id,pd.details FROM player AS p LEFT JOIN player_details AS pd ON p.id=pd.player WHERE p.nano_id=?"
	row := dbConn.QueryRow(query, playerNanoId)
	err = row.Scan(&player.Username, &player.NanoId, &player.Details)
	if err != nil {
		if err == sql.ErrNoRows {
			return player, err
		}
	}

	return player, err
}

func (userRepo *UserRepositoryStruct) GetUsersWithDetails() ([]models.Player, error) {
	dbConn := userRepo.DBDao
	
	query := "SELECT player.username,player.nano_id,player_details.details FROM player LEFT JOIN player_details ON player_details.player=player.id;"
	
	var players []models.Player
	rows, err := dbConn.Query(query)
	if err != nil {
		return players, err
	}
	for rows.Next() {
		var player models.Player
		rows.Scan(&player.Username, &player.NanoId, &player.Details)
		players = append(players, player)
	}

	return players, nil
}
