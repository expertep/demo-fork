package model

import (
	"mime/multipart"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	InvoiceNo string             `bson:"invoice_no,omitempty" json:"invoice_no"`
	RefNo     string             `bson:"ref_no,omitempty" json:"ref_no"`
	Type      string             `bson:"type,omitempty" json:"type"  binding:"required"`
	Desc      string             `bson:"desc,omitempty" json:"desc"`
	Invoices  []Invoice          `bson:"invoices,omitempty" json:"invoices"`
	Payment   TransactionPayment `bson:"payment,omitempty" json:"payment"`
	Amount    float64            `bson:"amount" json:"amount" binding:"required"`
	After     float64            `bson:"after" json:"after"  binding:"required"`
	Before    float64            `bson:"before" json:"before"  binding:"required"`
	Status    string             `bson:"status,omitempty" json:"status"`
	UserId    string             `bson:"user_id,omitempty" json:"user_id"`
	Username  string             `bson:"username,omitempty" json:"username"`
	CreateAt  time.Time          `bson:"create_at,omitempty" json:"create_at"`
	UpdateAt  time.Time          `bson:"update_at,omitempty" json:"update_at"`

	LoginName string `bson:"login_name,omitempty" json:"login_name"`
	LoginID   string `bson:"login_id,omitempty" json:"login_id"`
}

type RequestChangeCreditTransaction struct {
	Desc     string  `bson:"desc,omitempty" json:"desc"`
	Amount   float64 `bson:"amount" json:"amount" binding:"required"`
	Username string  `bson:"username,omitempty" json:"username"  binding:"required"`
}

type Balance struct {
	Balance int64 `bson:"balance" json:"balance"`
}
type TransactionPayment struct {
	ID       primitive.ObjectID `bson:"id,omitempty" json:"id"`
	Type     string             `bson:"type,omitempty" json:"type"`
	Name     string             `bson:"name,omitempty" json:"name"`
	Desc     string             `bson:"desc,omitempty" json:"desc"`
	Time     time.Time          `bson:"time,omitempty" json:"time"`
	Amount   float64            `bson:"amount,omitempty" json:"amount"`
	Image    string             `bson:"image,omitempty" json:"image"`
	Receiver Receiver           `bson:"receiver,omitempty" json:"receiver"`
	Sender   Receiver           `bson:"sender,omitempty" json:"sender"`
	Hash     string             `bson:"hash,omitempty" json:"hash"`
}

type Invoice struct {
	Name     string    `bson:"name,omitempty" json:"name"`
	Desc     string    `json:"desc" form:"desc"  binding:"required"`
	Credit   int32     `json:"credit" form:"credit"  binding:"required"`
	Amount   float64   `bson:"amount" json:"amount"`
	CreateAt time.Time `bson:"create_at,omitempty" json:"create_at"`
	UpdateAt time.Time `bson:"update_at,omitempty" json:"update_at"`
}

type Order struct {
	Type        string                `bson:"type,omitempty" json:"type"`
	Desc        string                `bson:"desc,omitempty" json:"desc"`
	PaymentType string                `bson:"payment_type,omitempty" json:"payment_type"`
	Amount      int                   `bson:"amount,omitempty" json:"amount"`
	Image       *multipart.FileHeader `bson:"image,omitempty" json:"image"`
	UserId      string                `bson:"user_id,omitempty" json:"user_id"`
	Username    string                `bson:"username,omitempty" json:"username"`
	CreateAt    time.Time             `bson:"create_at,omitempty" json:"create_at"`
	UpdateAt    time.Time             `bson:"update_at,omitempty" json:"update_at"`
}

type Receiver struct {
	BankCode   string `json:"bank_code"`
	BankName   string `json:"bank_name"`
	BankNumber string `json:"bank_number"`
}
type VerifyRes struct {
	Amount    float64   `json:"amount"`
	Datetime  time.Time `json:"datetime"`
	ImageSlip string    `json:"image_slip"`
	Receiver  Receiver  `json:"receiver"`
	ScanToken string    `json:"scan_token"`
	Sender    Receiver  `json:"sender"`
	TransRef  string    `json:"trans_ref"`
}
