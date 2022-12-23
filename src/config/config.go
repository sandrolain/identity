package config

import (
	"fmt"

	"github.com/sandrolain/go-utilities/pkg/envutils"
	"github.com/sandrolain/identity/src/keys"
)

const (
	DEF_MONGOBD_DATABASE      = "identity"
	DEF_MONGODB_TIMEOUT       = 5
	DEF_REDIS_TIMEOUT         = 5
	DEF_GRPC_ADMIN_PORT       = 1984
	DEF_GRPC_CLIENT_PORT      = 1985
	DEF_TOTP_RECOVERY_LENGTH  = 4
	DEF_TOTP_RECOVERY_SIZE    = 8
	DEF_TOTP_ISSUER           = "identity"
	DEF_KEY_LENGTH            = 32
	DEF_JWT_ISSUER            = "identity"
	DEF_TOTP_REQUEST_MINUTES  = 10
	DEF_LOGIN_SESSION_MINUTES = 30
	DEF_MACHINE_KEY_MINUTES   = 525_600 // 1 year
	DEF_LOGIN_MAX_FAILS       = 3
	DEF_LOGIN_LOCKOUT_MINUTES = 30
)

const (
	ENV_MASTER_KEY_B64   = "ID_MASTER_KEY_B64"
	ENV_MONGODB_URI      = "ID_MONGODB_URI"
	ENV_REDIS_HOST       = "ID_REDIS_HOST"
	ENV_REDIS_PASSWORD   = "ID_REDIS_PASSWORD"
	ENV_TOTP_ISSUER      = "ID_TOTP_ISSUER"
	ENV_JWT_ISSUER       = "ID_JWT_ISSUER"
	ENV_ADMIN_CERT_FILE  = "ID_ADMIN_CERT_FILE"
	ENV_ADMIN_KEY_FILE   = "ID_ADMIN_KEY_FILE"
	ENV_CLIENT_CERT_FILE = "ID_CLIENT_CERT_FILE"
	ENV_CLIENT_KEY_FILE  = "ID_CLIENT_KEY_FILE"
)

type RecoveryTokensConfig struct {
	Length int
	Size   int
}

type MongoDbConfig struct {
	Uri      string
	Database string
	Timeout  int
}

type RedisConfig struct {
	Host     string
	Password string
	Timeout  int
}

type GrpcConfig struct {
	Port     int
	CertFile string
	KeyFile  string
}

type TotpConfig struct {
	Issuer         string
	RecoveryTokens RecoveryTokensConfig
}

type JwtConfig struct {
	Issuer string
}

type SecureKeyConfig struct {
	Length    int
	MasterKey keys.MasterKey
}

type SessionConfig struct {
	TotpRequestMinutes  int
	LoginSessionMinutes int
	MachineKeyMinutes   int
}

type LoginConfig struct {
	MaxFails       int
	LockoutMinutes int
}

type Config struct {
	MongoDb    MongoDbConfig
	Redis      RedisConfig
	AdminGrpc  GrpcConfig
	ClientGrpc GrpcConfig
	Totp       TotpConfig
	Jwt        JwtConfig
	Session    SessionConfig
	Login      LoginConfig
	SecureKey  SecureKeyConfig
}

func GetDefaultConfiguration() Config {
	return Config{
		AdminGrpc: GrpcConfig{
			Port:     DEF_GRPC_ADMIN_PORT,
			CertFile: "",
			KeyFile:  "",
		},
		ClientGrpc: GrpcConfig{
			Port:     DEF_GRPC_CLIENT_PORT,
			CertFile: "",
			KeyFile:  "",
		},
		SecureKey: SecureKeyConfig{
			MasterKey: [32]byte{},
			Length:    DEF_KEY_LENGTH,
		},
		Totp: TotpConfig{
			RecoveryTokens: RecoveryTokensConfig{
				Length: DEF_TOTP_RECOVERY_LENGTH,
				Size:   DEF_TOTP_RECOVERY_SIZE,
			},
			Issuer: DEF_TOTP_ISSUER,
		},
		Jwt: JwtConfig{
			Issuer: DEF_JWT_ISSUER,
		},
		Session: SessionConfig{
			TotpRequestMinutes:  DEF_TOTP_REQUEST_MINUTES,
			LoginSessionMinutes: DEF_LOGIN_SESSION_MINUTES,
			MachineKeyMinutes:   DEF_MACHINE_KEY_MINUTES,
		},
		Login: LoginConfig{
			MaxFails:       DEF_LOGIN_MAX_FAILS,
			LockoutMinutes: DEF_LOGIN_LOCKOUT_MINUTES,
		},
		MongoDb: MongoDbConfig{
			Uri:      "",
			Database: DEF_MONGOBD_DATABASE,
			Timeout:  DEF_MONGODB_TIMEOUT,
		},
		Redis: RedisConfig{
			Host:     "",
			Password: "",
			Timeout:  DEF_REDIS_TIMEOUT,
		},
	}
}

func GetConfiguration() (cfg Config, err error) {
	mongoDbURI, err := envutils.RequireEnvString(ENV_MONGODB_URI)
	if err != nil {
		return
	}

	redisHost, err := envutils.RequireEnvString(ENV_REDIS_HOST)
	if err != nil {
		return
	}

	redisPassword, err := envutils.RequireEnvString(ENV_REDIS_PASSWORD)
	if err != nil {
		return
	}

	mk, err := envutils.RequireEnvBase64(ENV_MASTER_KEY_B64)
	if err != nil {
		return
	}
	mkl := len(mk)
	if mkl != 32 {
		err = fmt.Errorf("invalid Master Key length: %v", mkl)
		return
	}

	adminCertFile, err := envutils.RequireEnvPath(ENV_ADMIN_CERT_FILE)
	if err != nil {
		return
	}
	adminKeyFile, err := envutils.RequireEnvPath(ENV_ADMIN_KEY_FILE)
	if err != nil {
		return
	}

	clientCertFile, err := envutils.RequireEnvPath(ENV_CLIENT_CERT_FILE)
	if err != nil {
		return
	}
	clientKeyFile, err := envutils.RequireEnvPath(ENV_CLIENT_KEY_FILE)
	if err != nil {
		return
	}

	cfg = GetDefaultConfiguration()
	cfg.AdminGrpc.CertFile = adminCertFile
	cfg.AdminGrpc.KeyFile = adminKeyFile
	cfg.ClientGrpc.CertFile = clientCertFile
	cfg.ClientGrpc.KeyFile = clientKeyFile
	cfg.SecureKey.MasterKey = keys.MasterKeyFromBytes(mk)
	cfg.Totp.Issuer = envutils.GetEnvString(ENV_TOTP_ISSUER, DEF_TOTP_ISSUER)
	cfg.Jwt.Issuer = envutils.GetEnvString(ENV_JWT_ISSUER, DEF_JWT_ISSUER)
	cfg.MongoDb.Uri = mongoDbURI
	cfg.Redis.Host = redisHost
	cfg.Redis.Password = redisPassword

	return
}
