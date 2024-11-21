package helper

import (
	"github.com/sirupsen/logrus"
	"github.com/witnsby/web3_offline_coding_test/src/pkg/model"
	"os"
)

func GetEnvVariable() *model.EnvironmentVariables {
	privateKeyHex, ok := os.LookupEnv("PRIVATEKEYHEX")
	if !ok {
		logrus.Warning("Private Key Hex should be added to Environment as PRIVATEKEYHEX")
	}

	return &model.EnvironmentVariables{
		PrivateKeyHex: privateKeyHex,
	}
}
