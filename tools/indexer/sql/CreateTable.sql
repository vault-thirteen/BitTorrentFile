--// This is a template command (query) used to create a table for storing //--
--// the results of the indexer tool. It is written for MySQL SQL dialect. //--

CREATE TABLE test.BTFiles (
	Id BIGINT auto_increment NOT NULL,
	TorrentFile VARCHAR(256) NOT NULL,
	TorrentPath VARCHAR(1024) NOT NULL,
	StoredFile VARCHAR(1024) NOT NULL,
	StoredPath VARCHAR(1024) NOT NULL,
	StoredSize BIGINT unsigned NOT NULL,
	CONSTRAINT BTFiles_PK PRIMARY KEY (Id)
)
ENGINE=InnoDB 
DEFAULT CHARSET=utf8mb4 
COLLATE=utf8mb4_0900_ai_ci;
