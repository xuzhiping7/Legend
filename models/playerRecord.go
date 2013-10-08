package models

import (
	"time"
)

type PlayerRecord struct {
	Flag   int
	Record [10]string
}

func NewPlayerRecord() *PlayerRecord {
	temp := &PlayerRecord{}
	temp.InitRecord()
	return temp
}

//初始化所有记录
func (this *PlayerRecord) InitRecord() {
	this.Flag = 0
	for i := 0; i < len(this.Record); i++ {
		this.Record[i] = ""
	}
}

//增加一个玩家事件记录
func (this *PlayerRecord) AddRecord(s string) {
	this.Record[this.Flag] = time.Now().Format("[01.02 15:04]") + s
	this.Flag++
	if this.Flag >= len(this.Record) {
		this.Flag = 0
	}
}

//遍历输出当前所有事件记录
func (this *PlayerRecord) GetAllRecord() (s string) {

	s = ""
	for i := this.Flag - 1; i >= 0; i-- {
		if this.Record[i] != "" {
			s += this.Record[i] + "\n"
		}
	}

	for i := this.Flag; i < len(this.Record); i++ {
		if this.Record[i] != "" {
			s += this.Record[i] + "\n"
		}
	}

	return s
}
