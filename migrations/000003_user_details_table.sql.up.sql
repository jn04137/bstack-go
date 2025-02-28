CREATE TABLE player_details (
	id INT NOT NULL AUTO_INCREMENT,
	player INT NOT NULL,
	details TEXT(500),
	PRIMARY KEY (id),
	UNIQUE (player),
	FOREIGN KEY (player) REFERENCES player(id)
);
