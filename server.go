package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func test(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hi hello\n")
}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "home!\n")
}

func execSql(db *sql.DB, sql string) {
	fmt.Printf("executing statment:-> %s\n", sql)
	_, err := db.Exec(sql)
	if err != nil {
		log.Printf("%q: %s\n", err, sql)
		return
	}
}

func seedDB(db *sql.DB) {
	fmt.Println("Seeding DB")

	userTableSql := "create table user (id integer primary key, name text not null, email text)"
	execSql(db, userTableSql)

	//insert dummy user data
	userInsertsql := "insert into user values (1, 'vp', 'vp@email.com')"
	execSql(db, userInsertsql)
	userInsertsql2 := "insert into user values (2, 'ka', 'ka@email.com')"
	execSql(db, userInsertsql2)

	coffeeTableSql := "create table coffee (id integer primary key, name text, origin text, profile text)"
	execSql(db, coffeeTableSql)

	coffeeTableInsertSql := "insert into coffee values (1, 'guji', 'columbia', 'chocolate')"
	execSql(db, coffeeTableInsertSql)

	coffeeTableSql2 := "insert into coffee values (2, 'geisha', 'ethiopia', 'milk')"
	execSql(db, coffeeTableSql2)

	join_users_coffee_table := `
	create table join_user_coffee
		(
		user_id integer not null,
	 	coffee_id integer not null,
		primary key (user_id, coffee_id),
		foreign key (user_id) references user (id)
		foreign key (coffee_id) references coffee (id)
		)
	`
	execSql(db, join_users_coffee_table)

	user_coffee_sql := "insert into join_user_coffee values (1, 2)"
	execSql(db, user_coffee_sql)

	user_coffee_sql2 := "insert into join_user_coffee values (1, 1)"
	execSql(db, user_coffee_sql2)

	user_coffee_sql3 := "insert into join_user_coffee values (2, 1)"
	execSql(db, user_coffee_sql3)

	/*
		add index
		why:
			- much faster data retrieval (read) without scanning the full table
			- especially when it comes to queries involving ordering, where, aggregation

		cons:
			- disk space -> they take disk space which can be fine but for huge tables can be something to think about
			- inserting/updating/deleting record requires updating all indexes, which can be constly for many write operations
	*/
	idx_sql := "create unique index user_index on user(id)"
	execSql(db, idx_sql)
}

func initDB() {
	fmt.Println("Initializing/confirming database...")

	//Clean slate for now, nuke the db file
	// os.Remove("./db1.db")

	//check if db exists
	db, err := sql.Open("sqlite3", "./db1.db")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	seedDB(db)

	//query some stuff here

	q1 := "select * from user"
	rows, err := db.Query(q1)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var email string

		err = rows.Scan(&id, &name, &email)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(id, name, email)
	}

	defer db.Close()
}

func main() {
	fmt.Println("server init")

	initDB()

	// register handlers
	http.HandleFunc("/test", test)
	http.HandleFunc("/", home)

	http.ListenAndServe(":8080", nil)

}
