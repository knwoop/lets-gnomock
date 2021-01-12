package model

import "time"

type User struct {
	UserID    uint64    `json:"user_id"　db:"user_id"`       // user_id
	Username  string    `json:"username"　db:"username"`     // username
	CreatedAt time.Time `json:"created_at"　db:"created_at"` // created_at
}

// Insert inserts the User to the database.
func (u *User) Insert(db XODB) error {
	var err error

	// sql insert query, primary key provided by autoincrement
	const sqlstr = `INSERT INTO users (` +
		`username, created_at` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
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
	return nil
}
