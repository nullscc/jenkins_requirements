package models

type RepoRequirement struct {
	Id       int64  `json:"_" xorm:"'id' pk autoincr"`
	Org      string `form:"org" json:"org" xorm:"'org' varchar(50) notnull" binding:"required"`
	Repo     string `form:"repo" json:"repo" xorm:"'repo' varchar(50) notnull" binding:"required"`
	FileName string `form:"file_name" json:"file_name" xorm:"'file_name' varchar(50) notnull" binding:"required"`
	Sha1     string `form:"sha1" json:"sha1" xorm:"'sha1' varchar(64) notnull" binding:"-"`
}
