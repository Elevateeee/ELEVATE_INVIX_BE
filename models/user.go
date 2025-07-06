package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type User struct {
	ID 				primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Username 		string `bson:"username" json:"username"`
	Email 			string `bson:"email" json:"email"`
	Password 		string `bson:"password" json:"-"`
	Phone 			string `bson:"phone" json:"phone"`
	IsVerified   	bool   `bson:"is_verified" json:"is_verified"`
	CreatedAt 		primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt 		primitive.DateTime `bson:"updated_at" json:"updated_at"`
}


// cara biar nentuin field harus di isi adalah dengan tag `bson:"nama_field" json:"nama_field"`
// contoh:
// `bson:"username" json:"username"`
// ini akan membuat field username harus diisi ketika melakukan operasi CRUD pada MongoDB


// untuk membuat field tidak di tampilkan pada response, bisa menggunakan tag `json:"-"`
// contoh:
// `Password string `bson:"password" json:"-"``
// ini akan membuat field Password tidak ditampilkan pada response JSON

// untuk membuat field tidak perlu diisi tapi tetap di kirim ke database, bisa menggunakan tag `bson:"nama_field,omitempty"`
// contoh:
// `Token string `bson:"token,omitempty" json:"token"`
// ini akan membuat field Token tidak perlu diisi ketika melakukan operasi CRUD pada MongoDB, tapi tetap dikirim ke database jika ada nilainya
