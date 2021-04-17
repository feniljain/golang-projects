package cmd

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" //Using the psql driver
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "learn"
)

var (
	db *sql.DB
)

//Init is used to initialize the database
func Init() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%v user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Database se connection chalu kar diya hai!")
	//defer db.Close()
}

//AddNumber adds given number into the db
func AddNumber(number string) int {
	sqlStatement := `
	INSERT INTO test (number)
	VALUES($1)
	RETURNING id
	`
	id := 0
	err := db.QueryRow(sqlStatement, number).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record id is:", id)
	return id
}

//UpdateNumber updates the number with given id
func UpdateNumber(id int, number string) {
	sqlStatement := `
	UPDATE test
	SET number = $2
	WHERE id = $1
	`
	res, err := db.Exec(sqlStatement, id, number)
	if err != nil {
		DeleteNumber(id)
		fmt.Println(err)
		return
	}
	count, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Successfully updated", count, "records")
}

//DeleteNumber deleted the entry with given id
func DeleteNumber(id int) {
	sqlStatement := `
	DELETE FROM test
	WHERE id = $1
	`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		fmt.Println(err)
		return
	}
	count, err := res.RowsAffected()
	if err != nil {
		return
	}
	fmt.Println("Deleted", count, "rows")
}

//GetAllNumbers returns the whole table
func GetAllNumbers() []Number {
	sqlStatement := `
	SELECT * FROM test;
	`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var numbers []Number
	for rows.Next() {
		var id int
		var number string
		err := rows.Scan(&id, &number)
		if err != nil {
			panic(err)
		}
		numbers = append(numbers, Number{
			id:     id,
			number: number,
		})
	}
	return numbers
}

//Number struct represents a row in table
type Number struct {
	id     int
	number string
}
