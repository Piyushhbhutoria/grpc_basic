package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"grpc_tut/proto"
	"log"
	"net/http"
	"strconv"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	client := proto.NewAddServiceClient(conn)

	g := gin.Default()
	g.GET("/add/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Param"})
			return
		}
		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Param"})
			return
		}
		req := &proto.Request{A: int64(a), B: int64(b)}
		if response, err := client.Add(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{"result": fmt.Sprint(response.Result)})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Param"})
		}
	})
	g.GET("/mull/:a/:b", func(ctx *gin.Context) {
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Param"})
			return
		}
		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Param"})
			return
		}
		req := &proto.Request{A: int64(a), B: int64(b)}
		if response, err := client.Multiply(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{"result": fmt.Sprint(response.Result)})
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Param"})
		}
	})

	if err := g.Run(":8081"); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
