package pg

import (
	"database/sql"
	"fmt"
	"log"
)

//
// Connected
//
func Connected(src string) (db *sql.DB) {
	var err error
	db, err = sql.Open("postgres", src)
	if err != nil {
		log.Fatalf("DB Open err: %v", err)
	}
	return
}

//
// ResetSeq
//
const alterSeq = `
	ALTER SEQUENCE %v RESTART WITH 1
`

func ResetSeq(seq string) {
	if _, err := DB.Exec(fmt.Sprintf(alterSeq, seq)); err != nil {
		log.Fatalf("DB Reset Sequence err: %v", err)
	}
}

const deleteAllQuery = `
	DELETE FROM %v
`

func DeleteAll(table string) {
	if _, err := DB.Exec(fmt.Sprintf(deleteAllQuery, table)); err != nil {
		log.Fatalf("DB Delete All Tickets err: %v", err)
	}
}
