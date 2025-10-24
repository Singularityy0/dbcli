package main
import (
	"database/sql"
	"fmt"
	"log"
	"os"
	_ "github.com/go-sql-driver/mysql"
)
const (
	dbHost     = "localhost"
	dbPort     = "3306"
	dbUser     = "root"
	dbPassword = "apna_password_daaldo_yaar_mera_leak_nhi_krrha"
	dbName     = "college"
)

func tabu(db *sql.DB) error{
	fmt.Println("Tables")
	
	q1 := "SHOW TABLES"
	rows, err := db.Query(q1)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	defer rows.Close()

	var tflag int
	for rows.Next() {
		var t_i string
		if err := rows.Scan(&t_i); err != nil {
			return fmt.Errorf("error: %v", err)
		}
		tflag++
		fmt.Printf("%d. %s\n", tflag, t_i)
	}

	if tflag == 0 {
		fmt.Println("No tablles")
	}

	return nil
}

func main() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error connecting :", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		//log.Fatal("Error:", err)
	}

	fmt.Println("Connected nig")

	command := os.Args[1]

	switch command {
	case "list":
		if err := tabu(db); err != nil {
			log.Fatal("Error:", err)
		}

	
	}
}