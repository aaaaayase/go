package main

type XiaomiCredential struct {
	Uid     string `json:"uid" binding:"required"`
	Session string `json:"session" binding:"required"`
}

type SignInRequest struct {
	Idp              string           `json:"idp" binding:"required"`
	XiaomiCredential XiaomiCredential `json:"credential" binding:"required"`
	UpushDeviceToken string           `json:"upush_device_token" binding:"required"`
}
