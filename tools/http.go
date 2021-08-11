package tools

import (
	"crypto/rand"
	"go.uber.org/zap"
	"log"
	"math/big"
	"net"
)

func GetRandomDynamicPort(maxPort int64, minPort int64) int {
	port, err := rand.Int(rand.Reader, big.NewInt(maxPort-minPort))
	if err != nil {
		log.Fatalf("could not generate a random port: %v", err)
	}
	return int(port.Int64() + minPort)
}

func ListAddresses(port int, logger *zap.SugaredLogger) {
	interfaces, err := net.Interfaces()
	if err != nil {
		logger.Fatal("could not get network interfaces", err)
	}
	var ip net.IP
	for _, iface := range interfaces {
		addresses, err := iface.Addrs()
		if err != nil {
			logger.Error("could not get addresses for interface", iface.Name, err)
		}
		for _, addr := range addresses {
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP.To4()
			case *net.IPAddr:
				ip = v.IP.To4()
			}
			if ip != nil {
				logger.Infof("listening on %s:%d", ip, port)
			}
		}
	}
}
