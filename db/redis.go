package db

import (
	"encoding/json"
	"github.com/ptflp/adrenalinBank/models"
)
import "github.com/go-redis/redis"

var client *redis.Client

func InitRedis() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
	})

	client.FlushAll()
}

func getMoveKey(game string) string {
	return "game_" + game
}

func SaveOperation(operation models.Operation) error {
	b, _ := json.Marshal(operation)

	err := client.Set(operation.Id, b, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

func SaveOperations(o []models.Operation) error {
	b, _ := json.Marshal(o)

	err := client.Set("operations", b,0).Err()
	if err != nil {
		return err
	}
	return nil
}

func GetData(key string) (row string, err error){
	row, err = client.Get(key).Result()
	if err != nil {
		return "", err
	}

	return row, err
}

func GetMoves(game string) (moves models.MoveObjects, err error) {
	rows, err := client.LRange(getMoveKey(game), 0, -1).Result()
	if err != nil {
		return moves, err
	}

	ms := []models.WebMoveObject{}
	for _, val := range rows {
		var m = models.WebMoveObject{}
		err = json.Unmarshal([]byte(val), &m)
		ms = append(ms, m)
	}

	return ms, err
}

func DeleteTestMoves() {
	client.FlushAll()
}
