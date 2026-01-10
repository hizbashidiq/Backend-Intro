package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pelletier/go-toml/v2"
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

  sGin.POST("/test", handler)

  sGin.GET("/test/:details", handler)

  sGin.POST("/price", func(c *gin.Context) {
    body := Body{}
    if err:=c.ShouldBindJSON(&body); err!=nil{
      var ve validator.ValidationErrors
      if errors.As(err, &ve){
        out := make([]ErrMsg, len(ve))
        for i, fe:= range ve{
          out[i] = ErrMsg{fe.Field(), getErrMsg(fe)}
        }
        c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
      }
      // c.AbortWithStatusJSON(http.StatusBadRequest,
      //   gin.H{
      //     "error": "VALIDATEERR-1",
      //     // "message": "invalid inputs. Please check your inputs",
      //     "message":err.Error(),
      //   })
      return
    }
    c.JSON(http.StatusOK, &body)

  })

  sGin.POST("/toml", func(c *gin.Context){
    uri := URI{}
    // shouldbind is i/o stream while shouldbindbody is memory buffer
    // mostly you will only need shouldbind unless advance and specific case of middleware
    if err:=c.ShouldBindWith(&uri, Toml{});err!=nil{
      c.AbortWithError(http.StatusBadRequest, err)
      return
    }
    c.JSON(http.StatusOK, uri)
  })

  log.Fatal(http.ListenAndServe(":8080", sGin))
}

func handler(c *gin.Context){
  details := c.Param("details")

  if details == ""{
    body := Body{}
    if err:=c.BindJSON(&body);err!=nil{
      // bindjson is stream I/O, while shouldbindbodywith copy the body buffer and add it to context
      c.AbortWithError(http.StatusBadRequest, err)
      return
    }
    fmt.Println(body)
    c.JSON(http.StatusAccepted, &body)
  }else{
    uri := URI{}
    if err:=c.BindUri(&uri);err!=nil{
      c.AbortWithError(http.StatusBadRequest, err)
      return
    }
    fmt.Println(uri)
    c.JSON(http.StatusAccepted, &uri)
  }
}

type ErrMsg struct{
  Field string `json:"field"`
  Message string `jsong:"message"`
}

func getErrMsg(fe validator.FieldError) string{
  switch fe.Tag(){
  case "required":
    return "This field is required"
  case "lte":
    return "Should be less than "+fe.Param()
  case "gte":
    return "Should be more than "+fe.Param()
  }
  return "Unknown Error"
}

type Toml struct{
}

// customize binding need Name() and Bind(r *http.Request, i interface{}) method
func (t Toml)Name() string{
  return "toml"
}

func (t Toml)Bind(r *http.Request, i interface{}) error{
  tD := toml.NewDecoder(r.Body)
  return tD.Decode(i)
}

// to be able to use ShouldBindBodyWith, need to implement BindBody method
func (t Toml)BindBody(bytes []byte, i interface{})error{
  return toml.Unmarshal(bytes, i)
}

type URI struct{
  Details string `json:"name" uri:"details" binding:"required"`
  // uri:details is a gin tag. it tells to bind path parameter named :details
}

type Body struct{
  Name string `json:"name"`
  Price uint `json:"price" binding:"required,gte=10,lte=1000"`
}