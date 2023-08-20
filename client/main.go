package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	_ "github.com/ldmtam/tam-grpc-resolver"

	pb "github.com/ldmtam/calculator_service/proto"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		"tam:///calculator-server.default:8080",
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("dial to calculator server failed: %v", err)
	}
	defer conn.Close()

	client := pb.NewCalculatorServiceClient(conn)

	resp, err := client.Ping(ctx, &pb.PingRequest{})
	if err != nil {
		log.Fatalf("failed to call PING to server: %v", err)
	}

	fmt.Println("Initialized gRPC connection to calculator server successfully. Message:", resp.Message)

	router := setupRouter(client)
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("run http server failed: %v", err)
	}
}

func setupRouter(client pb.CalculatorServiceClient) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	router.GET("/api/v1/add/:a/:b", func(c *gin.Context) {
		a := cast.ToFloat32(c.Param("a"))
		b := cast.ToFloat32(c.Param("b"))

		fmt.Printf("Received ADD request with a = %v and b = %v\n", a, b)

		result, err := client.Add(c.Request.Context(), &pb.AddRequest{A: a, B: b})
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]any{"error": err})
			return
		}

		c.JSON(http.StatusOK, map[string]any{"result": result.Result})
	})

	router.GET("/api/v1/subtract/:a/:b", func(c *gin.Context) {
		a := cast.ToFloat32(c.Param("a"))
		b := cast.ToFloat32(c.Param("b"))

		fmt.Printf("Received SUBTRACT request with a = %v and b = %v\n", a, b)

		result, err := client.Subtract(c.Request.Context(), &pb.SubtractRequest{A: a, B: b})
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]any{"error": err})
			return
		}

		c.JSON(http.StatusOK, map[string]any{"result": result.Result})
	})

	router.GET("/api/v1/multiply/:a/:b", func(c *gin.Context) {
		a := cast.ToFloat32(c.Param("a"))
		b := cast.ToFloat32(c.Param("b"))

		fmt.Printf("Received MULTIPLY request with a = %v and b = %v\n", a, b)

		result, err := client.Multiply(c.Request.Context(), &pb.MultiplyRequest{A: a, B: b})
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]any{"error": err})
			return
		}

		c.JSON(http.StatusOK, map[string]any{"result": result.Result})
	})

	router.GET("/api/v1/divide/:a/:b", func(c *gin.Context) {
		a := cast.ToFloat32(c.Param("a"))
		b := cast.ToFloat32(c.Param("b"))

		fmt.Printf("Received DIVIDE request with a = %v and b = %v\n", a, b)

		result, err := client.Divide(c.Request.Context(), &pb.DivideRequest{A: a, B: b})
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]any{"error": err})
			return
		}

		c.JSON(http.StatusOK, map[string]any{"result": result.Result})
	})

	return router
}
