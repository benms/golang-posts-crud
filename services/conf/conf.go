package conf

import (
	"fmt"
	"os"
	"strconv"

	"github.com/rs/zerolog/log"
)

const (
	hostKey       = "PHAROS_HOST"
	portKey       = "PHAROS_PORT"
	dbHostKey     = "PHAROS_DB_HOST"
	dbPortKey     = "PHAROS_DB_PORT"
	dbNameKey     = "PHAROS_DB_NAME"
	dbUserKey     = "PHAROS_DB_USER"
	dbPasswordKey = "PHAROS_DB_PASSWORD"
	jwtSecretKey  = "PHAROS_JWT_SECRET"
)

type Config struct {
	Host       string
	Port       string
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
	JwtSecret  string
	Env        string
}

func NewConfig(env string) Config {
	host, ok := os.LookupEnv(hostKey)
	if !ok || host == "" {
		logPanic(hostKey)
	}
	port, ok := os.LookupEnv(portKey)
	if !ok || port == "" {
		if _, err := strconv.Atoi(port); err != nil {
			logPanic(portKey)
		}
	}
	dbHost, ok := os.LookupEnv(dbHostKey)
	if !ok || dbHost == "" {
		logPanic(dbHostKey)
	}
	dbPort, ok := os.LookupEnv(dbPortKey)
	if !ok || dbPort == "" {
		if _, err := strconv.Atoi(dbPort); err != nil {
			logPanic(dbPortKey)
		}
	}
	dbName, ok := os.LookupEnv(dbNameKey)
	if !ok || dbName == "" {
		logPanic(dbNameKey)
	}
	dbUser, ok := os.LookupEnv(dbUserKey)
	if !ok || dbUser == "" {
		logPanic(dbUserKey)
	}
	dbPassword, ok := os.LookupEnv(dbPasswordKey)
	if !ok || dbPassword == "" {
		logPanic(dbPasswordKey)
	}
	jwtSecret, ok := os.LookupEnv(jwtSecretKey)
	if !ok || jwtSecret == "" {
		logPanic(jwtSecretKey)
	}

	return Config{
		Host:       host,
		Port:       port,
		DbHost:     dbHost,
		DbPort:     dbPort,
		DbName:     dbName,
		DbUser:     dbUser,
		DbPassword: dbPassword,
		JwtSecret:  jwtSecret,
		Env:        env,
	}
}

// not used: just set test env
func NewTestConfig() Config {
	// godotenv.Load("../../.env")
	testConfig := NewConfig("dev")
	testConfig.DbName = fmt.Sprintf("%s_test", testConfig.DbName)
	return testConfig
}

func logPanic(envVar string) {
	log.Panic().Str("envVar", envVar).Msg("ENV variable not set or value not valid")
}
