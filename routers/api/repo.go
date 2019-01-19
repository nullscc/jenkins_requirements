package api

import (
	"github.com/gin-gonic/gin"
	"github.com/nullscc/jenkins_requirements/models"
	"github.com/nullscc/jenkins_requirements/utils"
)

func Add(c *gin.Context) {
	var req models.RepoRequirement
	if err := c.ShouldBind(&req); err != nil {
		utils.JsonRetError(c, "参数错误")
		return
	}

	session := utils.Db.NewSession()
	defer session.Close()

	err := session.Begin()
	if err != nil {
		utils.JsonRetError(c, err.Error())
		return
	}

	repo_requirement := new(models.RepoRequirement)
	has, _ := session.Where("org=? and repo=? and file_name=?", req.Org, req.Repo, req.FileName).Get(repo_requirement)
	if !has {
		repo_requirement.Org = req.Org
		repo_requirement.Repo = req.Repo
		repo_requirement.FileName = req.FileName
		repo_requirement.Sha1 = req.Sha1
		_, err = session.Insert(repo_requirement)
	} else {
		repo_requirement.Sha1 = req.Sha1
		_, err = session.Where("id = ?", repo_requirement.Id).Cols("sha1").Update(repo_requirement)
	}

	if err != nil {
		session.Rollback()
		utils.JsonRetError(c, err.Error())
		return
	}
	err = session.Commit()
	if err != nil {
		session.Rollback()
		utils.JsonRetError(c, err.Error())
		return
	}

	utils.JsonRetEmpty(c)
}

func Get(c *gin.Context) {
	var req models.RepoRequirement
	if err := c.ShouldBindQuery(&req); err != nil {
		utils.JsonRetError(c, "参数错误")
		return
	}

	repo_requirement := new(models.RepoRequirement)
	has, err := utils.Db.Where("org=? and repo=? and file_name=?", req.Org, req.Repo, req.FileName).Get(repo_requirement)
	if err != nil {
		utils.JsonRetError(c, err.Error())
		return
	}
	if !has {
		utils.JsonRet(c, &gin.H{"sha1": ""})
		return
	}

	utils.JsonRet(c, &gin.H{"sha1": repo_requirement.Sha1})
}
