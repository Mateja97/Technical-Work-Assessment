package config

import (
	"errors"
	"os"
	"strings"
)

type Config struct {
	ethClients    []string
	serverAddress string
}

var c *Config

func LoadConfig() error {

	ethClients := os.Getenv("ETH_CLIENTS")
	serverAddress := os.Getenv("SERVER_ADDRESS")

	if ethClients == "" {
		return errors.New("missing eth clients")
	}
	ethClientList := make([]string, 0)
	for _, part := range strings.Split(ethClients, ",") {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			ethClientList = append(ethClientList, trimmed)
		}
	}

	c = &Config{
		ethClients:    ethClientList,
		serverAddress: serverAddress,
		//EthClients:         []string{"https://eth-mainnet.g.alchemy.com/v2/SYpDdGSITBpoS7VC6Duq02FmvWmbaS2i", "https://mainnet.infura.io/v3/03165944e0a24e349d21b977cad5e8a2", "https://virtual.mainnet.rpc.tenderly.co/b6046e1c-b375-407f-a681-0048d52f6630"},
	}
	return nil
}

func EthClients() []string {
	return c.ethClients
}

func ServerAddress() string {
	return c.serverAddress
}
