package config

import (
	"fmt"

	"github.com/sandrolain/go-utilities/pkg/envutils"
)

const (
	DEF_MONGOBD_DATABASE      = "identity"
	DEF_MONGODB_TIMEOUT       = 5
	DEF_GRPC_ADMIN_PORT       = 1984
	DEF_GRPC_CLIENT_PORT      = 1985
	DEF_OTP_RECOVERY_LENGTH   = 4
	DEF_OTP_RECOVERY_SIZE     = 8
	DEF_OTP_ISSUER            = "identity"
	DEF_KEY_LENGTH            = 32
	DEF_JWT_ISSUER            = "identity"
	DEF_OTP_REQUEST_MINUTES   = 10
	DEF_LOGIN_SESSION_MINUTES = 30
	DEF_MACHINE_KEY_MINUTES   = 525_600 // 1 year
	DEF_LOGIN_MAX_FAILS       = 3
	DEF_LOGIN_LOCKOUT_MINUTES = 30
)

const (
	ENV_MASTER_KEY_B64 = "ID_MASTER_KEY_B64"
	ENV_MONGODB_URI    = "ID_MONGODB_URI"
	ENV_OTP_ISSUER     = "ID_OTP_ISSUER"
	ENV_JWT_ISSUER     = "ID_JWT_ISSUER"
)

type RecoveryTokensConfig struct {
	Length int
	Size   int
}

type MongoDBConfig struct {
	URI      string
	Database string
	Timeout  int
}

type GRPCConfig struct {
	Port int
}

type OTPConfig struct {
	Issuer         string
	RecoveryTokens RecoveryTokensConfig
}

type JWTConfig struct {
	Issuer string
}
type SecureKeyConfig struct {
	Length    int
	MasterKey []byte
}

type SessionConfig struct {
	OTPRequestMinutes   int
	LoginSessionMinutes int
	MachineKeyMinutes   int
}

type LoginConfig struct {
	MaxFails       int
	LockoutMinutes int
}

type Config struct {
	MongoDB    MongoDBConfig
	AdminGRPC  GRPCConfig
	ClientGRPC GRPCConfig
	OTP        OTPConfig
	JWT        JWTConfig
	Session    SessionConfig
	Login      LoginConfig
	SecureKey  SecureKeyConfig
}

func GetConfiguration() (Config, error) {
	mongoDbURI, err := envutils.RequireEnvString(ENV_MONGODB_URI)
	if err != nil {
		return Config{}, err
	}

	mk, err := envutils.RequireEnvBase64(ENV_MASTER_KEY_B64)
	if err != nil {
		return Config{}, err
	}
	mkl := len(mk)
	if mkl != 32 {
		return Config{}, fmt.Errorf("invalid Master Key length: %v", mkl)
	}

	cfg := Config{
		AdminGRPC: GRPCConfig{
			Port: DEF_GRPC_ADMIN_PORT,
		},
		ClientGRPC: GRPCConfig{
			Port: DEF_GRPC_CLIENT_PORT,
		},
		SecureKey: SecureKeyConfig{
			MasterKey: mk,
			Length:    DEF_KEY_LENGTH,
		},
		OTP: OTPConfig{
			RecoveryTokens: RecoveryTokensConfig{
				Length: DEF_OTP_RECOVERY_LENGTH,
				Size:   DEF_OTP_RECOVERY_SIZE,
			},
			Issuer: envutils.GetEnvString(ENV_OTP_ISSUER, DEF_OTP_ISSUER),
		},
		JWT: JWTConfig{
			Issuer: envutils.GetEnvString(ENV_JWT_ISSUER, DEF_JWT_ISSUER),
		},
		Session: SessionConfig{
			OTPRequestMinutes:   DEF_OTP_REQUEST_MINUTES,
			LoginSessionMinutes: DEF_LOGIN_SESSION_MINUTES,
			MachineKeyMinutes:   DEF_MACHINE_KEY_MINUTES,
		},
		Login: LoginConfig{
			MaxFails:       DEF_LOGIN_MAX_FAILS,
			LockoutMinutes: DEF_LOGIN_LOCKOUT_MINUTES,
		},
		MongoDB: MongoDBConfig{
			URI:      mongoDbURI,
			Database: DEF_MONGOBD_DATABASE,
			Timeout:  DEF_MONGODB_TIMEOUT,
		},
	}

	return cfg, nil
}
