
package db
"database/sql"
_"github.com/go-sql-driver/mysql"
/*Create mysql connection*/
func CreateCon() *sql.DB {
db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test")
if err != nil {
fmt.Println(err.Error())
fmt.Println("db is connected")
//defer db.Close()
// make sure connection is available
err = db.Ping()
fmt.Println(err)
if err != nil {
fmt.Println("MySQL db is not connected")
fmt.Println(err.Error())
