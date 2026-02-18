package main

import (
    "context"
    "log"
    "time"

    "github.com/gin-gonic/gin"
    // ... outros imports (swagger, dotEnv)
    
    "arcanum/src/handler"
    "arcanum/src/service"
)

// @title           Arcanum API
// @version         1.0
// @description     API Open Source para múltiplos oráculos (Tarot, Cigano, Osho, etc.)
// @host            localhost:9090
// @BasePath        /v1
func main() {

    r := gin.Default()


    deckService := service.NewDeckService()
    deckHandler := handler.NewDeckHandler(deckService)

    v1 := r.Group("/v1")
    {
        v1.GET("/draw/:deck_type", deckHandler.GetRandomCards)


        v1.GET("/spread/:spread_type", deckHandler.GetComplexSpread)
        
        v1.GET("/tarot", deckHandler.HandleTarot)
        v1.GET("/cigano", deckHandler.HandleCigano)
        v1.GET("/padilha", deckHandler.HandlePadilha)
        v1.GET("/osho", deckHandler.HandleOsho)
        v1.GET("/sibila", deckHandler.HandleSibila)
    }

    log.Println("Arcanum API online em :9090")
    r.Run(":9090")
}

