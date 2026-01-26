package Store

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/ssh"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var _ = godotenv.Load()

func proxyConnection(local, remote net.Conn) {
	go func() {
		defer local.Close()
		defer remote.Close()
		_, _ = io.Copy(local, remote)
	}()
	go func() {
		defer local.Close()
		defer remote.Close()
		_, _ = io.Copy(remote, local)
	}()
}

var DB *gorm.DB
var DBMutex sync.Mutex

func ConnectPsql() (*gorm.DB, error) {
	DBMutex.Lock()
	defer DBMutex.Unlock()
	if DB != nil {
		return DB, nil
	}
	deployMode := os.Getenv("DEPLOYMODE") // Verifica o modo de deploy

	// Variáveis comuns para ambos os modos
	dbUser := os.Getenv("USER_DB")
	dbPassword := os.Getenv("PASSWORD_DB")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("HOST_DB") // Host do banco (público ou interno)
	dbPort := os.Getenv("DB_PORT") // Porta do banco de dados
	if dbPort == "" {
		dbPort = "5432" // Porta padrão se não definida
	}
	fmt.Println("Teste", dbUser, dbPassword, dbName)
	var psqlInfo string
	if deployMode == "production" {
		dbUser := os.Getenv("USER_DB")
		dbPassword := os.Getenv("PASSWORD_DB")
		dbName := os.Getenv("DB_NAME")
		dbHost := os.Getenv("HOST_DB") // Host do banco (público ou interno)
		dbPort := os.Getenv("DB_PORT") // Porta do banco de dados
		psqlInfo = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	} else {
		fmt.Println("🚀🔗->> CONEXÃO USANDO SSH")
		// Modo com SSH
		sshHost := os.Getenv("SSH_DB_HOST")
		sshPort := os.Getenv("SSH_DB_PORT")
		sshUser := os.Getenv("SSH_DB_USER")
		sshPassword := os.Getenv("SSH_DB_PASSWORD")
		localPort := os.Getenv("LOCAL_PORT")
		if localPort == "" {
			localPort = "15439" // Porta local padrão para o túnel
		}

		// Configuração do cliente SSH
		sshConfig := &ssh.ClientConfig{
			User: sshUser,
			Auth: []ssh.AuthMethod{
				ssh.Password(sshPassword),
			},
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
			Timeout:         10 * time.Second,
		}

		// Conectar via SSH
		sshConn, err := ssh.Dial("tcp", sshHost+":"+sshPort, sshConfig)
		if err != nil {
			return nil, fmt.Errorf("falha na conexão SSH: %v", err)
		}

		// Criar túnel
		listener, err := net.Listen("tcp", "localhost:"+localPort)
		if err != nil {
			return nil, fmt.Errorf("erro ao criar listener local: %v", err)
		}

		go func() {
			for {
				localConn, err := listener.Accept()
				if err != nil {
					continue
				}

				// Conectar ao banco via SSH (host interno)
				remoteConn, err := sshConn.Dial("tcp", dbHost+":"+dbPort)
				if err != nil {
					localConn.Close()
					continue
				}

				go proxyConnection(localConn, remoteConn)
			}
		}()

		time.Sleep(1 * time.Second) // Espera estabilização do túnel

		// String de conexão via túnel
		psqlInfo = fmt.Sprintf("host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable",
			localPort, dbUser, dbPassword, dbName)
		dbHost = "localhost"
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, localPort, dbName)
	}

	// Conexão com o banco (comum para ambos os modos)
	sqlDB, err := sql.Open("pgx", psqlInfo)

	if err != nil {
		return nil, fmt.Errorf("erro ao abrir conexão com PostgreSQL: %v", err)
	}

	// Testar conexão
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := sqlDB.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("ping ao banco de dados falhou: %v", err)
	}
	// Configurações do pool de conexões
	sqlDB.SetMaxOpenConns(5)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)
	sqlDB.SetConnMaxIdleTime(2 * time.Minute)
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err == nil {
		DB = gormDB
		return gormDB, nil
	}
	return nil, err
}
