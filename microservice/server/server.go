package main

import (
	"log"
    "io"
	"time"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
    pb "../proto"
    "net/http"
    // "github.com/gorilla/mux"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

const (
    address     = "localhost:50051" 
    // localhost
    // 1. ec2-13-58-52-105.us-east-2.compute.amazonaws.com  
    // 2. ec2-18-218-144-18.us-east-2.compute.amazonaws.com
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

    stream, err := client.GetEventos(ctx, &pb.RequestEvento{})
    var names [50]gin.H
    i := 0
    for {
		evento, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", c, err)
        }
        
        if evento != nil {
            names[i] = gin.H{"nombre": evento.Nombre,"id": evento.Id, "fechaCreacion": evento.FechaCreacion, "tipoLocalidad": evento.TipoLocalidad, "localidad_id": evento.LocalidadId, "descripcion": evento.Descripcion }
        }
        i++
    }
    c.JSON(http.StatusOK, names)
}

func GetAsientos(c *gin.Context) {
    conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
        
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewMicroClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 60)
	defer cancel()
    // LocalidadId := c.Param("name")
    LocalidadId := c.Param("localidaId")
    stream, err := client.GetAsientos(ctx, &pb.RequestAsiento{ LocalidadId:LocalidadId })
    var names [400]gin.H
    i := 0
    for {
		evento, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.ListFeatures(_) = _, %v", c, err)
        }
        
        if evento.Categoria != "" {
            names[i] = gin.H{"id": evento.Id, "categoria": evento.Categoria, "descripcion": evento.Descripcion }
        }
        i++
    }
    c.JSON(http.StatusOK, names)
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
    router.Use(cors.Default())
    router.GET("/api/eventos", GetEventos)
    router.POST("/api/comprarBoletos", ComprarBoletos)
    router.GET("/api/asientos/:localidaId", GetAsientos)
	router.Run(":3000")
}