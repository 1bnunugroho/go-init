CREATE TABLE IF NOT EXISTS userskills(
   userSkillID serial primary key,
   username VARCHAR (50),
   skillID INT,
   skillLevelID INT,
   FOREIGN KEY (username) REFERENCES users (username),
   FOREIGN KEY (skillID) REFERENCES skill (skillID),
   FOREIGN KEY (skillLevelID) REFERENCES skilllevel (skillLevelID)
);