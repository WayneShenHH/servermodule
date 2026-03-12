package cmd

import (
	"errors"
	"fmt"
	"math/rand"
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
		if len(args) < 1 {
			return
		}

		pwdLen, err := strconv.Atoi(args[0])
		if err != nil {
			panic(err)
		}

		if len(args) < 2 {
			args = append(args, "lud")
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
	if option == "" {
		option = "lud"
	}

	basic := []byte{}

	charset := ""
	if strings.ContainsRune(option, 'l') {
		charset += charset_lower
		basic = append(basic, charset_lower[rand.Intn(len(charset_lower))])
	}
	if strings.ContainsRune(option, 'u') {
		charset += charset_upper
		basic = append(basic, charset_upper[rand.Intn(len(charset_upper))])
	}
	if strings.ContainsRune(option, 'd') {
		charset += charset_number
		basic = append(basic, charset_number[rand.Intn(len(charset_number))])
	}
	if strings.ContainsRune(option, 's') {
		charset += charset_symbol
		basic = append(basic, charset_symbol[rand.Intn(len(charset_symbol))])
	}

	if len(charset) == 0 {
		return "", errors.New("options wrong:" + option)
	}

	if len(basic) > length {
		return "", fmt.Errorf("length wrong: %v", length)
	}

	password := make([]byte, 0, length-len(basic))

	for i := 0; i < length-len(basic); i++ {
		num := rand.Intn(len(charset))
		password = append(password, charset[num])
	}

	password = append(password, basic...)

	// shuffle password
	rand.Shuffle(len(password), func(i, j int) {
		password[i], password[j] = password[j], password[i]
	})

	return string(password), nil
}

func init() {
	RootCmd.AddCommand(genPasswordCmd)
}
