CREATE TABLE IF NOT EXISTS usersprofile(
   username VARCHAR (50) UNIQUE NOT NULL,
   name VARCHAR (50),
   address VARCHAR (500),
   bod DATE,
   email VARCHAR (50),
   PRIMARY KEY (username),
   FOREIGN KEY (username) REFERENCES users (username)
);