CREATE TABLE user_on_team (
	id INT NOT NULL AUTO_INCREMENT,
	player INT NOT NULL,
	team INT NOT NULL,
	PRIMARY KEY(id),
	FOREIGN KEY (player) REFERENCES player(id),
	FOREIGN KEY (team) REFERENCES team(id),
	UNIQUE (player, team)
);
