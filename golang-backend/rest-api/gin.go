package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

var sGin = gin.New()

func runGin(){
  sGin.GET("/hellogin", func(c *gin.Context){
    c.String(http.StatusOK, "WORLD")
  })

  sGin.GET("/marketgin/:id", func(c *gin.Context){
    c.String(http.StatusOK, "You request this item: %s", c.Param("id"))
  })

  sGin.PUT("/marketgin/:id", func(c *gin.Context){
    c.String(http.StatusAccepted, "You want to edit this item: %s", c.Param("id"))
  })

  sGin.POST("/login", func(c *gin.Context){
    var body Account

    // shouldbindjson only read once while shouldbindbodywithjson can read multiple? why?
    if err:=c.ShouldBindJSON(&body); err!=nil{
      c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
      return
    }

    c.JSON(http.StatusOK, gin.H{"status":"you are logged in"})
  })

  log.Fatal(http.ListenAndServe(":8080", sGin))
}

func HelloHandlerGin(){

}