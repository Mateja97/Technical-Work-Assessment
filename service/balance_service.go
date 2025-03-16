package service

import (
	"alluvial-task/client"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"sync"
)

type BalanceService struct {
	clientPool *client.EthClientPool
	cache      *sync.Map
}

type GetBalanceResponse struct {
	JsonRpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  string `json:"result"`
}

func NewBalanceService(clientPool *client.EthClientPool, cache *sync.Map) *BalanceService {
	return &BalanceService{clientPool: clientPool, cache: cache}
}

func (s *BalanceService) GetBalance(address string) (string, error) {
	if balance, found := s.cache.Load(address); found {
		return balance.(string), nil
	}

	numberOfClients := s.clientPool.Len()
	clients := s.clientPool.GetClients(numberOfClients)
	var wg sync.WaitGroup
	results := make(chan string, numberOfClients)
	errors := make(chan error, numberOfClients)

	for _, clientURL := range clients {
		wg.Add(1)
		go func(_url string) {
			defer wg.Done()
			balance, err := s.queryClient(_url, address)
			if err != nil {
				errors <- err
				return
			}
			results <- balance
		}(clientURL)
	}

	wg.Wait()
	close(results)
	close(errors)

	var balances []string
	for balance := range results {
		balances = append(balances, balance)
	}
	for err := range errors {
		log.Println("Client error:", err)
	}

	balance, err := s.determineConsistentBalance(balances)
	if err != nil {
		return "", err
	}

	s.cache.Store(address, balance)

	return balance, nil
}

func (s *BalanceService) queryClient(clientURL, address string) (string, error) {

	clientURL, clientId := s.clientPool.GetClient()
	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_getBalance",
		"params":  []interface{}{address, "latest"},
		"id":      clientId,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to encode JSON payload: %w", err)
	}

	resp, err := http.Post(clientURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to make request to Ethereum client: %w", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var response GetBalanceResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return "", fmt.Errorf("failed to parse JSON response: %w", err)

	}

	hexValue := response.Result
	if len(hexValue) < 2 || hexValue[:2] != "0x" {
		return "", fmt.Errorf("invalid hexadecimal value: %w", err)

	}

	balanceWei := new(big.Int)
	_, success := balanceWei.SetString(hexValue[2:], 16)
	if !success {
		return "", fmt.Errorf("failed to parse hexadecimal value: %s", hexValue)
	}

	balance := weiToEther(balanceWei)
	return balance.String(), nil
}

func weiToEther(wei *big.Int) *big.Float {
	if wei.Int64() == 0 {
		return big.NewFloat(0)
	}
	ether := new(big.Float).SetInt(wei)
	ether.Quo(ether, big.NewFloat(1e18))
	return ether
}

func (s *BalanceService) determineConsistentBalance(balances []string) (string, error) {
	counts := make(map[string]int)
	for _, balance := range balances {
		counts[balance]++
	}

	var consistentBalance string
	maxCount := 0
	for balance, count := range counts {
		if count > maxCount {
			consistentBalance = balance
			maxCount = count
		}
	}

	if maxCount < s.clientPool.Len()/2+1 {
		return "", fmt.Errorf("no quorum reached for balance")
	}

	return consistentBalance, nil
}
