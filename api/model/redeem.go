package model

import "time"

type Redeem struct {
	RedeemID      uint      `gorm:"primaryKey" json:"redeem_id"`
	UserID        uint      `json:"user_id" binding:"required"`
	KodeVoucher   string    `json:"kode_voucher" binding:"required"`
	TanggalRedeem time.Time `json:"tanggal_redeem"`
}
