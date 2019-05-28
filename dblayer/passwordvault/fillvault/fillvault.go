package main

import (
	"crypto/md5"
	"mygoedu/dblayer/passwordvault"
)

func main() {
	db, err := passwordvault.ConnectPasswordVault()
	if err != nil {
		return
	}
	pass1 := md5.Sum([]byte("test1"))
	pass2 := md5.Sum([]byte("test2"))
	pass3 := md5.Sum([]byte("test3"))
	passwordvault.AddBytesToVault(db, "test1", pass1[:])
	passwordvault.AddBytesToVault(db, "test2", pass2[:])
	passwordvault.AddBytesToVault(db, "test3", pass3[:])
	db.Close()
}
