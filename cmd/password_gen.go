package cmd

import (
	"crypto/rand"
	"math/big"
	"strconv"

	"github.com/spf13/cobra"

	"github.com/WayneShenHH/servermodule/logger"
)

var genPasswordCmd = &cobra.Command{
	Short: "generate password",
	Long:  `generate password`,
	Use:   "gen:pwd",

	Run: func(cmd *cobra.Command, args []string) {
		logger.Debug(cmd.Short)
		logger.Info("args:", args)
		if len(args) < 1 {
			return
		}

		pwdLen, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		pwd, err := generatePassword(pwdLen)
		if err != nil {
			panic(err)
		}

		logger.Infof("password: %v", pwd)
	},
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+"
const charset2 = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generatePassword(length int) (string, error) {
	password := make([]byte, length)

	for i := range password {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset2))))
		if err != nil {
			return "", err
		}
		password[i] = charset2[num.Int64()]
	}

	return string(password), nil
}

func init() {
	RootCmd.AddCommand(genPasswordCmd)
}
