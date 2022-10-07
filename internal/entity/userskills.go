package entity

// User represents a user.
type UsersSkills struct {
	userSkillID  int    `json:"userSkillID"`
	username     string `json:"username"`
	skillID      int    `json:"skillID"`
	skillLevelID int    `json:"skillLevelID"`
}
