package main

import (
	"log"
    // "os"
	"time"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
    pb "../proto"
    "net/http"
    // "github.com/gorilla/mux"
    "github.com/gin-gonic/gin"
)

const (
	address     = "ec2-18-218-144-18.us-east-2.compute.amazonaws.com:50052"
	defaultName = "world"
)

func GetEventos(c *gin.Context) {
    conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
        
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewMicroClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 60)
	defer cancel()

    start := time.Now()
    name := defaultName
	message, err := client.Ping(ctx,  &pb.PingRequest{Message: name})
	end := time.Now()
    fmt.Println(end.Sub(start))
    log.Println(message)
    if message == nil {
        c.JSON(http.StatusInternalServerError, gin.H{"not connected": "no conectado"})
    } else {
        c.JSON(http.StatusOK, gin.H{"eventos": "Los eventos"})
    }
}

func ComprarBoletos(c *gin.Context)  {
    usuario_id := c.PostForm("usuario_id")
    asiento_id := c.PostForm("asiento_id")
    fmt.Printf("El usuario id es: %s, el asiento id es: %s", usuario_id, asiento_id)
    c.String(http.StatusOK, "Comprar boletos")
}

func main() {
    // Set up a connection to the server.
    // r := mux.NewRouter()
    // r.HandleFunc("/", Ping)
    router := gin.Default()
    router.GET("/api/eventos", GetEventos)
    router.POST("/api/comprarBoletos", ComprarBoletos)
	router.Run(":3001")
}