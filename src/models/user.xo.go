// Package models contains the types for schema 'test_database'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"errors"
	"time"
)

// User represents a row from 'test_database.users'.
type User struct {
	UserID    uint64    `json:"user_id"`    // user_id
	Username  string    `json:"username"`   // username
	CreatedAt time.Time `json:"created_at"` // created_at

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the User exists in the database.
func (u *User) Exists() bool {
	return u._exists
}

// Deleted provides information if the User has been deleted from the database.
func (u *User) Deleted() bool {
	return u._deleted
}

// Insert inserts the User to the database.
func (u *User) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if u._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO test_database.users (` +
		`username, created_at` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
	XOLog(sqlstr, u.Username, u.CreatedAt)
	res, err := db.Exec(sqlstr, u.Username, u.CreatedAt)
	if err != nil {
		return err
	}

	// retrieve id
	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	// set primary key and existence
	u.UserID = uint64(id)
	u._exists = true

	return nil
}

// Update updates the User in the database.
func (u *User) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !u._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if u._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE test_database.users SET ` +
		`username = ?, created_at = ?` +
		` WHERE user_id = ?`

	// run query
	XOLog(sqlstr, u.Username, u.CreatedAt, u.UserID)
	_, err = db.Exec(sqlstr, u.Username, u.CreatedAt, u.UserID)
	return err
}

// Save saves the User to the database.
func (u *User) Save(db XODB) error {
	if u.Exists() {
		return u.Update(db)
	}

	return u.Insert(db)
}

// Delete deletes the User from the database.
func (u *User) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !u._exists {
		return nil
	}

	// if deleted, bail
	if u._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM test_database.users WHERE user_id = ?`

	// run query
	XOLog(sqlstr, u.UserID)
	_, err = db.Exec(sqlstr, u.UserID)
	if err != nil {
		return err
	}

	// set deleted
	u._deleted = true

	return nil
}

// UserByUsername retrieves a row from 'test_database.users' as a User.
//
// Generated from index 'ui_username'.
func UserByUsername(db XODB, username string) (*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`user_id, username, created_at ` +
		`FROM test_database.users ` +
		`WHERE username = ?`

	// run query
	XOLog(sqlstr, username)
	u := User{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, username).Scan(&u.UserID, &u.Username, &u.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

// UserByUserID retrieves a row from 'test_database.users' as a User.
//
// Generated from index 'users_user_id_pkey'.
func UserByUserID(db XODB, userID uint64) (*User, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`user_id, username, created_at ` +
		`FROM test_database.users ` +
		`WHERE user_id = ?`

	// run query
	XOLog(sqlstr, userID)
	u := User{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, userID).Scan(&u.UserID, &u.Username, &u.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
