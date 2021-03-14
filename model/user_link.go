package model

import (
	"onbio/logger"
	"time"

	"go.uber.org/zap"
)

const (
	LinkTableName = "t_user_link"
)

/***

CREATE TABLE `t_user_link` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户id',
  `link_url` varchar(255) NOT NULL DEFAULT '' COMMENT '用户链接',
  `link_desc` varchar(2048) NOT NULL DEFAULT '' COMMENT '内容简述',
  `link_img` varchar(255) NOT NULL DEFAULT '' COMMENT '链接首图',
  `operator` varchar(255) NOT NULL DEFAULT '' COMMENT '操作人',
  `use_flag` tinyint(2) NOT NULL DEFAULT '0' COMMENT '是否有效',
  `create_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `last_updated_time` bigint(20) NOT NULL DEFAULT '0' COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  KEY `idx_user` (`user_id`)
) ENGINE=INNODB AUTO_INCREMENT=1  DEFAULT CHARSET=utf8 COMMENT='用户链接表';
**/
type Link struct {
	ID              uint64 `gorm:"primaryKey"  json:"id"`
	UserID          uint64 `gorm:"column:user_id" json:"user_id"`
	LinkTitle       string `gorm:"column:link_title" json:"link_title"`
	LinkUrl         string `gorm:"column:link_url" json:"link_url"`
	LinkDesc        string `gorm:"column:link_desc" json:"link_desc"`
	LinkImg         string `gorm:"column:link_img" json:"link_img"`
	Position        uint64 `gorm:"column:position" json:"position"`
	IsSpecial       int    `gorm:"column:is_special" json:"is_special"`
	Operator        string `gorm:"column:operator" json:"operator"`
	UseFlag         int    `gorm:"column:use_flag" json:"use_flag"`
	CreateTime      uint64 `gorm:"column:create_time" json:"create_time"`
	LastUpdatedTime uint64 `gorm:"column:last_updated_time" json:"last_updated_time"`
}

//删除用户链接,已经有有效开关了，这是直接删
func DeleteUserLink(userID, linkID uint64) (err error) {
	db := getMysqlConn().Table(LinkTableName)
	err = db.Where("id = ? and user_id = ? ", linkID, userID).Delete(Link{}).Error
	if err != nil {
		logger.Error("delete link info ", zap.Uint64("userId", userID), zap.Uint64("link", linkID))
		return
	}
	return
}

func CreateLink(userID, position uint64, linkUrl, linkDesc, linkImg, linkTitle string) (ID uint64, err error) {

	newLink := Link{
		UserID:          userID,
		LinkUrl:         linkUrl,
		LinkTitle:       linkTitle,
		LinkDesc:        linkDesc,
		LinkImg:         linkImg,
		IsSpecial:       0,
		UseFlag:         1,
		Position:        position,
		CreateTime:      uint64(time.Now().Unix()),
		LastUpdatedTime: uint64(time.Now().Unix()),
	}

	if linkTitle == "" {
		newLink.UseFlag = 0
	}

	db := getMysqlConn().Table(LinkTableName)
	db = db.Create(&newLink)
	if db.Error != nil {
		logger.Error("CreateLink::Find error: %s", zap.Error(db.Error))
		return 0, db.Error
	}
	ID = newLink.ID
	return
}

//获取单个链接详情
func GetUserLinkByID(id uint64) (linkItem Link, err error) {
	if id == 0 {
		return
	}
	db := getMysqlConn().Table(LinkTableName)

	db = db.Where("id = ?", id)

	//db = db.Where("use_flag = 1")

	err = db.First(&linkItem).Error

	if err != nil {
		logger.Error("get user link by id failed ", zap.Uint64("id", id))
		return
	}
	return
}

//更新链接记录
func UpdateLinkByID(linkID, userID uint64, info Link) (err error) {
	link := Link{
		ID:     linkID,
		UserID: userID,
	}

	updates := map[string]interface{}{}

	if info.LinkUrl != "" {
		updates["link_url"] = info.LinkUrl
	}
	if info.LinkImg != "default" {
		updates["link_img"] = info.LinkImg
	}

	if info.LinkDesc != "default" {
		updates["link_desc"] = info.LinkDesc
	}

	if info.LinkTitle != "default" {
		updates["link_title"] = info.LinkTitle
	}

	if info.Position != 0 {
		updates["position"] = info.Position
	}

	if info.IsSpecial != -1 {
		updates["is_special"] = info.IsSpecial
	}

	if info.UseFlag != -1 {
		updates["use_flag"] = info.UseFlag
	}

	updates["last_updated_time"] = uint64(time.Now().Unix())

	db := getMysqlConn().Table(LinkTableName)
	err = db.Where("id = ? and user_id = ? ", link.ID, link.UserID).Updates(updates).Error
	if err != nil {
		logger.Error("update link info ", zap.Any("model", link), zap.Any("updates", updates))
		return
	}
	return
}

//获取用户有效链接 + 分页
func GetUserLinkListWithPage(userID uint64, page, pageSize int) (linkList []*Link, count int, err error) {
	db := getMysqlConn().Table(LinkTableName)
	//db.LogMode(true)
	if userID != 0 {
		db = db.Where("user_id = ?", userID)
	}
	db = db.Where("use_flag = 1")

	err = db.Count(&count).Error
	if err != nil {
		logger.Error("get user link count from db failed ")
		return
	}

	limit := pageSize
	offset := (page - 1) * pageSize

	db = db.Limit(limit)
	db = db.Offset(offset)

	db = db.Order("position desc")

	err = db.Find(&linkList).Error

	if err != nil {
		logger.Error("get linkList from db failed ", zap.Uint64("userId", userID))
		return
	}

	return
}

//获取用户链接列表
func GetUserLinkList(userID uint64, page, pageSize int) (linkList []*Link, count int, err error) {

	db := getMysqlConn().Table(LinkTableName)

	err = db.Raw("select * from t_user_link where user_id = ? order by  position desc ", userID).Scan(&linkList).Error
	if err != nil {
		logger.Error("get user link from db failed ")
		return
	}

	if userID != 0 {
		db = db.Where("user_id = ?", userID)
	}
	err = db.Count(&count).Error
	if err != nil {
		logger.Error("get user link count from db failed ")
		return
	}
	return
}
