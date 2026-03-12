package cmd

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strconv"
	"strings"

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
		if len(args) < 2 {
			return
		}

		pwdLen, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		pwd, err := generatePassword(pwdLen, args[1])
		if err != nil {
			panic(err)
		}

		logger.Infof("password: %v", pwd)
	},
}

const charset_symbol = "!@#$%^&*()-_=+"
const charset_number = "0123456789"
const charset_upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const charset_lower = "abcdefghijklmnopqrstuvwxyz"

func generatePassword(length int, option string) (string, error) {
	password := make([]byte, length)

	charset := ""
	if strings.ContainsRune(option, 'l') {
		charset += charset_lower
	}
	if strings.ContainsRune(option, 'u') {
		charset += charset_upper
	}
	if strings.ContainsRune(option, 'd') {
		charset += charset_number
	}
	if strings.ContainsRune(option, 's') {
		charset += charset_symbol
	}

	if len(charset) == 0 {
		return "", errors.New("options wrong:" + option)
	}

	for i := range password {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[num.Int64()]
	}

	return string(password), nil
}

func init() {
	RootCmd.AddCommand(genPasswordCmd)
}
