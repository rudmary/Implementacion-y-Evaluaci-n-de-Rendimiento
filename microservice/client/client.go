// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"database/sql"
// 	_  "github.com/go-sql-driver/mysql"
// 	"github.com/joho/godotenv"
// 	"github.com/go-redis/redis"
// )

// func checkErr(err error) {
//     if err != nil {
// 		log.Fatal("Error loading .env file")
//         panic(err)
//     }
// }

// func ExampleClient() {
// 	client := redis.NewClient(&redis.Options{
// 		Addr: "localhost:6379",
// 		Password: "",
// 		DB: 0,
// 	})
// 	err := client.Set("key", "value", 0).Err()
// 	if err != nil {
// 		panic(err)
// 	}

// 	val, err := client.Get("key").Result()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("key", val)

// 	val2, err := client.Get("key2").Result()
// 	if err == redis.Nil {
// 		fmt.Println("key2 does not exist")
// 	} else if err != nil {
// 		panic(err)
// 	} else {
// 		fmt.Println("key2", val2)
// 	}
// 	// Output: key value
// 	// key2 does not exist
// }

// func main() {
// 	err := godotenv.Load()
// 	checkErr(err)
	
// 	MYSQL_DATABASE := os.Getenv("MYSQL_DATABASE")
// 	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
// 	MYSQL_USER := os.Getenv("MYSQL_USER")
// 	dbConfig := fmt.Sprintf("%s:%s@/%s", MYSQL_USER, MYSQL_PASSWORD, MYSQL_DATABASE) // username:password@protocol(address)/dbname?param=value
	
// 	db, err := sql.Open("mysql", dbConfig)
// 	err = db.Ping()
// 	checkErr(err)
// 	ExampleClient()

// 	// db, err := sql.Open("mysql", dbConfig)
// 	// if err != nil {
// 	// 	panic(err.Error())
// 	// }

// 	// rows, err := db.Query("SELECT * FROM bots")
// 	// columns, err := rows.Columns()
// 	// values := make([]sql.RawBytes, len(columns))
// 	// scanArgs := make([]interface{}, len(values))
	
// 	// for i := range values {
// 	// 	scanArgs[i] = &values[i]
// 	// }
	
// 	// for rows.Next() {
// 	// 	err = rows.Scan(scanArgs...)
// 	// 	if err != nil {
// 	// 		panic(err.Error())
// 	// 	}
// 	// 	var value string
// 	// 	for i, col := range values {
// 	// 		if col == nil {
// 	// 			value = "NULL"
// 	// 		} else {
// 	// 			value = string(col)
// 	// 		}
// 	// 		fmt.Println(columns[i], ": ", value)
// 	// 	}
// 	// 	fmt.Println("-----------------------------------")
// 	// }
// 	// defer db.Close()
// }


package main
import (
    "fmt"
    "log"
    "net"
    // "errors"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "golang.org/x/net/context"
    pb "../proto"
)

const (
	port = ":50052"
)

type server struct{}


func checkErr(err error) {
    if err != nil {
		log.Fatal("Error loading .env file")
        panic(err)
    }
}

func (s *server) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingReply, error) {
    log.Printf("Received: %v", in.Message)
	return &pb.PingReply{Message: "Hello aaa" + in.Message}, nil
}

func main() {
    fmt.Println("Servidor iniciado")
    lis, err := net.Listen("tcp", port)
    checkErr(err)
    s := grpc.NewServer()
    pb.RegisterMicroServer(s, &server{})
    reflection.Register(s)
    if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}