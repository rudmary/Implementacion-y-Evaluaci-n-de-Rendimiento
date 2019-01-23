// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
	// "database/sql"
	// _  "github.com/go-sql-driver/mysql"
	// "github.com/joho/godotenv"
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
    // "strconv"
    "os"
    // "errors"
    "database/sql"
	_  "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    // "golang.org/x/net/context"
    pb "../proto"
)

const (
	port = ":50051"
)

type server struct{}


func checkErr(err error) {
    if err != nil {
		log.Fatal("Error loading .env file")
        panic(err)
    }
}

// MYSQL_HOST=127.0.0.1
// MYSQL_DATABASE=tickets
// MYSQL_PASSWORD=mysqldb
// MYSQL_USER=root

// MYSQL_HOST=ec2-18-222-178-65.us-east-2.compute.amazonaws.com
// MYSQL_DATABASE=tickets
// MYSQL_PASSWORD=ruddy
// MYSQL_USER=ruddy

func GetEventosFromDB() ([50]pb.Evento, error) {
	err := godotenv.Load()
	checkErr(err)
    var eventos [50]pb.Evento
	MYSQL_DATABASE := os.Getenv("MYSQL_DATABASE")
	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
    MYSQL_USER := os.Getenv("MYSQL_USER")
    MYSQL_HOST := os.Getenv("MYSQL_HOST")
    // user:password@tcp(localhost:5555)/dbname
	dbConfig := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_DATABASE) // username:password@protocol(address)/dbname?param=value
    // dbConfig := fmt.Sprintf("%s:%s@/%s", MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_DATABASE) 
    db, err := sql.Open("mysql", dbConfig)
    rows, err := db.Query("select * from eventos inner join localidades where eventos.localidad_id = localidades.id")
    checkErr(err)
    defer db.Close()
    i := 0

    for rows.Next() {
        
            var id int64
            var localidad_id int64
            var nombre string
            var fechaCreacion string
            var id2 int64
            var tipo string
            var descripcion string
            err = rows.Scan(&id, &localidad_id, &nombre, &fechaCreacion, &id2, &tipo, &descripcion)
            checkErr(err)
            evento :=  pb.Evento{}
            evento.Id = id
            evento.Nombre = nombre
            evento.FechaCreacion = fechaCreacion
            evento.TipoLocalidad = tipo
            evento.LocalidadId = localidad_id
            evento.Descripcion = descripcion
            eventos[i] = evento
            i++
        } 
    checkErr(err)

    return eventos, nil
}

func GetAsientosFromDB(localidad_id int64) ([400]pb.Asiento, error) {
	err := godotenv.Load()
	checkErr(err)
    var asientos [400]pb.Asiento
	MYSQL_DATABASE := os.Getenv("MYSQL_DATABASE")
	MYSQL_PASSWORD := os.Getenv("MYSQL_PASSWORD")
    MYSQL_USER := os.Getenv("MYSQL_USER")
    MYSQL_HOST := os.Getenv("MYSQL_HOST")
    // user:password@tcp(localhost:5555)/dbname
	dbConfig := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_DATABASE)// username:password@protocol(address)/dbname?param=value

    db, err := sql.Open("mysql", dbConfig)
    rows, err := db.Query("select * from asientos inner join localidades_asientos where localidades_asientos.localidad_id = 1") // ,strconv.FormatInt(localidad_id, 10)
    checkErr(err)
    defer db.Close()
    i := 0

    for rows.Next() {
            var id int64
            var localidad_asiento_id int64
            var localidad_id int64
            var categoria string
            // var fechaCreacion string
            var asientos_id int64
            var descripcion string
            err = rows.Scan(&id, &categoria, &descripcion, &localidad_asiento_id, &localidad_id, &asientos_id)
            checkErr(err)
            asiento :=  pb.Asiento{}
            asiento.Id = id
            asiento.Categoria = categoria
            asiento.Descripcion = descripcion
            asientos[i] = asiento
            i++
        }
    checkErr(err)

    return asientos, nil
}

// func (s *server) GetEventos(in *pb.RequestEvento, stream pb.Micro_GetEventosServer) error {
//     var eventos [50]pb.Evento
//     eventos, err := GetEventosFromDB()
//     checkErr(err)
//     for _, evento := range eventos {
//         if err := stream.Send(&evento); err != nil {
//             return err
//         }
//     }
//     return nil
// }

// Proto definitions

func (s *server) GetEventos(in *pb.RequestEvento, stream pb.Micro_GetEventosServer) error {
    var eventos [50]pb.Evento
    eventos, err := GetEventosFromDB()
    checkErr(err)
    for _, evento := range eventos {
        if err := stream.Send(&evento); err != nil {
            return err
        }
    }
    return nil
}

func (s *server) GetAsientos(in *pb.RequestAsiento, stream pb.Micro_GetAsientosServer) error {
    var asientos [400]pb.Asiento
    asientos, err := GetAsientosFromDB(in.LocalidadId)
    checkErr(err)
    for _, asiento := range asientos {
        if err := stream.Send(&asiento); err != nil {
            return err
        }
    }
    return nil
}

// func (s *server) ComprarBoleto(ctx context.Context, in *pb.RequestComprarBoleto) (*pb.ReplyBoleto, error) {
//     var eventos [50]pb.Evento
//     eventos, err := GetEventosFromDB()
//     checkErr(err)
//     for _, evento := range eventos {
//         if err := stream.Send(&evento); err != nil {
//             return err
//         }
//     }
//     return nil
// }



func main() {
    fmt.Println("Servidor corriento")
    lis, err := net.Listen("tcp", port)
    checkErr(err)
    s := grpc.NewServer()
    pb.RegisterMicroServer(s, &server{})
    reflection.Register(s)
    if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}