package main

import (
	"Arcanum/internal/adapters/handler"
	"Arcanum/internal/adapters/repositories"
	"Arcanum/internal/core/domain"
	"Arcanum/internal/core/ports"
	"Arcanum/internal/core/service"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	// ... outros imports (swagger, dotEnv)
)

// @title           Arcanum API
// @version         1.0
// @description     API Open Source para múltiplos oráculos (Tarot, Cigano, Osho, etc.)
// @host            localhost:9090
// @BasePath        /v1
func main() {

	repo := repositories.NewMemoryDeckRepository()

	err := loadJSONToMemory(repo, "../../assets/decks")
	if err != nil {
		log.Fatalf("Erro ao carregar baralhos: %v", err)
	}

	r := gin.Default()

	deckService := service.NewDeckService(repo)
	deckHandler := handler.NewDeckHandler(deckService)

	v1 := r.Group("/v1")
	{
		v1.GET("/draw/:deck_type", deckHandler.GetRandomCards)

		v1.GET("/spread/:spread_type", deckHandler.GetComplexSpread)

		/*v1.GET("/tarot", deckHandler.HandleTarot)
		v1.GET("/cigano", deckHandler.HandleCigano)
		v1.GET("/padilha", deckHandler.HandlePadilha)
		v1.GET("/osho", deckHandler.HandleOsho)
		v1.GET("/sibila", deckHandler.HandleSibila)*/
	}

	log.Println("Arcanum API online em :9090")
	r.Run(":9090")
}

func loadJSONToMemory(repo ports.DeckRepository, folderPath string) error {
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			filePath := filepath.Join(folderPath, file.Name())
			content, err := os.ReadFile(filePath)
			if err != nil {
				log.Printf("Erro ao ler arquivo %s: %v", file.Name(), err)
				continue
			}

			var deck domain.Deck
			if err := json.Unmarshal(content, &deck); err != nil {
				log.Printf("Erro no parse do JSON %s: %v", file.Name(), err)
				continue
			}

			slug := strings.TrimSuffix(file.Name(), ".json")
			deck.Slug = slug

			repo.Save(deck)
			log.Printf("Baralho carregado: %s (%d cartas)", deck.Name, len(deck.Cards))
		}
	}
	return nil
}
