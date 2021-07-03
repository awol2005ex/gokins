package models

import (
	"time"
)

type TPipeline struct {
	Id           string    `xorm:"not null pk VARCHAR(64)" json:"id"`
	Uid          string    `xorm:"VARCHAR(64)" json:"uid"`
	Name         string    `xorm:"VARCHAR(255)" json:"name"`
	DisplayName  string    `xorm:"VARCHAR(255)" json:"displayName"`
	PipelineType string    `xorm:"VARCHAR(255)" json:"pipelineType"`
	AccessToken  string    `xorm:"VARCHAR(255)" json:"accessToken"`
	Url          string    `xorm:"VARCHAR(255)" json:"url"`
	Username     string    `xorm:"VARCHAR(255)" json:"username"`
	Deleted      int       `xorm:"default 0 INT(1)" json:"-"`
	DeletedTime  time.Time `xorm:"DATETIME" json:"-"`

	Nick    string `xorm:"-" json:"nick"`
	Avat    string `xorm:"-" json:"avat"`
	Buildln int64  `xorm:"-" json:"buildln"`

	LastId       string    `xorm:"-" json:"lastId"`
	LastStatus   string    `xorm:"-" json:"lastStatus"`
	LastError    string    `xorm:"-" json:"lastError"`
	LastCreated  time.Time `xorm:"-" json:"lastCreated"`
	LastStarted  time.Time `xorm:"-" json:"lastStarted"`
	LastFinished time.Time `xorm:"-" json:"lastFinished"`
}
