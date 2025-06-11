package main

import (
	"time"
	pb "week1/auth-server/proto"
)

type Server struct {
	pb.UnimplementedAuthServer
}

type XiaomiRequest struct {
	AppID     string `json:"appId" binding:"required"`
	Session   string `json:"session" binding:"required"`
	Uid       string `json:"uid" binding:"required"`
	Signature string `json:"signature" binding:"required"`
}

type XiaomiResponse struct {
	Errcode int32  `json:"errcode" binding:"required"`
	ErrMsg  string `json:"errMsg"`
	Adult   string `json:"adult"`
	Age     int32  `json:"age"`
}

type RealNameCache struct {
	ExternalID  string    `json:"external_id" gorm:"column:external_id;type:varchar(255);not null"`
	Idp         uint8     `json:"idp" gorm:"column:idp;type:tinyint unsigned;not null"`
	PersonID    uint32    `json:"person_id" gorm:"column:person_id;type:int unsigned;not null"`
	DateOfBirth time.Time `json:"date_of_birth" gorm:"column:date_of_birth;type:date;not null"`
}

type ComboUser struct {
	ComboID      uint64            `json:"combo_id" gorm:"primaryKey;column:combo_id;type:bigint unsigned;not null;"`
	Idp          uint8             `json:"idp" gorm:"primaryKey;column:idp;type:tinyint unsigned;not null;"`
	ExternalID   string            `json:"external_id" gorm:"column:external_id;type:varchar(255);not null;"`
	ExternalName string            `json:"external_name" gorm:"column:external_name;type:varchar(255);default:'';"`
	GameID       string            `json:"game_id" gorm:"column:game_id;type:varchar(32);not null;"`
	Meta         map[string]string `json:"meta" gorm:"column:meta;type:json;"`
	ClientMeta   map[string]string `json:"client_meta" gorm:"column:client_meta;type:json;"`
	PersonID     *uint32           `json:"person_id" gorm:"column:person_id;type:int unsigned;"`
	DateOfBirth  *time.Time        `json:"date_of_birth" gorm:"column:date_of_birth;type:date;"`
	CreatedAt    time.Time         `json:"created_at" gorm:"column:created_at;type:datetime;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt    time.Time         `json:"updated_at" gorm:"column:updated_at;type:datetime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}
