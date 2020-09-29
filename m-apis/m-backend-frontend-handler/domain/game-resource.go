package domain

import (
	v1gameengine "github.com/fvukojevic/grpc_test/m-apis/m-game-engine/v1"
	v1highscore "github.com/fvukojevic/grpc_test/m-apis/m-highscore/v1"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"strconv"
)

type gameResource struct {
	gameEngineClient v1gameengine.GameEngineClient
	highScoreClient  v1highscore.GameClient
}

func NewGameResource(gameClient v1gameengine.GameEngineClient, highScoreClient v1highscore.GameClient) *gameResource {
	return &gameResource{
		gameEngineClient: gameClient,
		highScoreClient:  highScoreClient,
	}
}

func (gr *gameResource) SetHighScore(c *gin.Context) {
	highScoreString := c.Param("hs")
	highScoreFloat64, err := strconv.ParseFloat(highScoreString, 64)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert to float")
	}

	_, err = gr.highScoreClient.SetHighScore(c, &v1highscore.SetHighScoreRequest{
		HighScore: highScoreFloat64,
	})
	if err != nil {
		log.Error().Err(err).Msg("error while setting highscore")
	}
}

func (gr *gameResource) GetHighScore(c *gin.Context) {
	highscoreResponse, err := gr.highScoreClient.GetHighScore(c, &v1highscore.GetHighScoreRequest{})

	if err != nil {
		log.Error().Err(err).Msg("error while getting highscore")
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return
	}

	hsString := strconv.FormatFloat(highscoreResponse.HighScore, 'e', -1, 64)
	c.JSON(200, gin.H{
		"hs": hsString,
	})
}

func (gr *gameResource) GetSize(c *gin.Context) {
	getSizeResponse, err := gr.gameEngineClient.GetSize(c, &v1gameengine.GetSizeRequest{})

	if err != nil {
		log.Error().Err(err).Msg("error while getting highscore")
		c.JSON(500, gin.H{
			"error": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"size": getSizeResponse.GetSize(),
	})
}

func (gr *gameResource) SetScore(c *gin.Context) {
	scoreString := c.Param("score")
	scoreFloat64, err := strconv.ParseFloat(scoreString, 64)
	if err != nil {
		log.Error().Err(err).Msg("Failed to convert to float")
	}

	_, err = gr.gameEngineClient.SetScore(c, &v1gameengine.SetScoreRequest{
		Score: scoreFloat64,
	})

	if err != nil {
		log.Error().Err(err).Msg("error while setting score")
	}
}
