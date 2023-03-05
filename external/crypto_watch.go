package external

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/cryptowatch_challenge/config"
	"github.com/cryptowatch_challenge/internal/models"
	"github.com/cryptowatch_challenge/internal/stores"
	"github.com/gorilla/websocket"
)

type Subscription struct {
	StreamSubscription `json:"streamSubscription"`
}

type StreamSubscription struct {
	Resource string `json:"resource"`
}

type SubscribeRequest struct {
	Subscriptions []Subscription `json:"subscriptions"`
}

type Update struct {
	MarketUpdate struct {
		Market struct {
			MarketId int `json:"marketId,string"`
		} `json:"market"`

		TradesUpdate struct {
			Trades []Trade `json:"trades"`
		} `json:"tradesUpdate"`
	} `json:"marketUpdate"`
}

type Trade struct {
	Timestamp     int `json:"timestamp,string"`
	TimestampNano int `json:"timestampNano,string"`

	Price  string `json:"priceStr"`
	Amount string `json:"amountStr"`
}

type CryptoWatchClient struct {
	config *config.Config
}

func NewCryptoWatchClient(config *config.Config) *CryptoWatchClient {
	return &CryptoWatchClient{
		config: config,
	}
}

func (s *CryptoWatchClient) ListenCryptoWatch(priceStore *stores.PriceStore) {
	APIKEY := s.config.CryptoWatchApiKey
	c, _, err := websocket.DefaultDialer.Dial(s.config.CryptoWatchUrl+APIKEY, nil)
	if err != nil {
		panic(err)
	}
	defer c.Close()

	// Read first message, which should be an authentication response
	_, message, err := c.ReadMessage()
	if err != nil {
		panic(err)
	}

	var authResult struct {
		AuthenticationResult struct {
			Status string `json:"status"`
		} `json:"authenticationResult"`
	}
	err = json.Unmarshal(message, &authResult)
	if err != nil {
		panic(err)
	}

	// Send a JSON payload to subscribe to a list of resources
	// Read more about resources here: https://docs.cryptowat.ch/websocket-api/data-subscriptions#resources
	resources := []string{
		"instruments:125:trades",
	}
	subMessage := struct {
		Subscribe SubscribeRequest `json:"subscribe"`
	}{}
	// No map function in golang :-(
	for _, resource := range resources {
		subMessage.Subscribe.Subscriptions = append(subMessage.Subscribe.Subscriptions, Subscription{StreamSubscription: StreamSubscription{Resource: resource}})
	}
	msg, err := json.Marshal(subMessage)
	if err != nil {
		panic(err)
	}

	err = c.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		panic(err)
	}

	// Process incoming ETH/USD trades
	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Fatal("Error reading from connection", err)
			return
		}

		var update Update
		err = json.Unmarshal(message, &update)
		if err != nil {
			panic(err)
		}

		pricesMap := make(map[int]string)
		prices := make([]*models.Price, 0)

		for _, trade := range update.MarketUpdate.TradesUpdate.Trades {
			pricesMap[update.MarketUpdate.Market.MarketId] = trade.Price
		}

		for marketID, price := range pricesMap {
			if s, err := strconv.ParseFloat(price, 64); err == nil {
				prices = append(prices, &models.Price{
					MarketID: uint32(marketID),
					Price:    s,
				})
			}
		}

		err = priceStore.CreateInBatches(prices, 100).Error
		if err != nil {
			continue
		}

		if len(prices) == 0 {
			continue
		}

		time.Sleep(time.Duration(s.config.TimeIntervalCallCryptoWatchSecond) * time.Second)
	}
}
