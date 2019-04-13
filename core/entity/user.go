package entity

import "time"

// User_Profile (PO)
type User_Profile struct {
	ID             uint
	Name           string
	Phone          string
	Gender         int
	Avatar         string
	Register_Mode  int
	Third_Party_ID string
	Create_Time    time.Time
	Update_Time    time.Time
}

// User_Auth (PO)
type User_Auth struct {
	ID          uint
	Encrypt_Pwd string
	User_ID     uint
	Create_Time time.Time
	Update_Time time.Time
}
