package boot

import (
	// "context"

	"context"
	"encoding/json"
	"fmt"
	"gold-gym-be/docs"

	// "gold-gym-be/internal/data/auth"

	// "gold-gym-be/pkg/firebaseclient"

	"gold-gym-be/pkg/tracing"
	"log"
	"net/http"

	"gold-gym-be/internal/config"
	jaegerLog "gold-gym-be/pkg/log"

	// Log "gold-gym-be/pkg/logs"

	"gold-gym-be/pkg/firebaseclient"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/db"
	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"google.golang.org/api/option"

	// "golang.org/x/net/trace"
	// "go.opentelemetry.io/otel/trace"
	// "gold-gym-be/pkg/trace"

	goldgymData "gold-gym-be/internal/data/goldgym"
	goldgymServer "gold-gym-be/internal/delivery/http"
	authHandler "gold-gym-be/internal/delivery/http/auth"
	goldgymHandler "gold-gym-be/internal/delivery/http/goldgym"
	goldgymService "gold-gym-be/internal/service/goldgym"

	goldgymStockData "gold-gym-be/internal/data/stock"
	goldgymStockService "gold-gym-be/internal/service/stock"
	// goldgymStockData "gold-gym-be/internal/data/stock"
	// pushNotifData "gold-gym-be/internal/data/pushnotif"
	// pushNotifHandler "gold-gym-be/internal/delivery/http/pushnotif"
	// pushNotifService "gold-gym-be/internal/service/pushnotif"
)

// HTTP will load configuration, do dependency injection and then start the HTTP server
func HTTP() error {
	var (
		// 	ctx = context.Background()
		cred map[string]string
		cfg  *config.Config // Configuration object
	)
	err := config.Init()
	if err != nil {
		log.Fatalf("[CONFIG] Failed to initialize config: %v", err)
	}
	cfg, cred = config.Get()

	rdb := newRedisClient(cfg.Redis)

	// t, err := trace.New(ctx, cfg.Trace.Exporter)
	// if err != nil {
	// 	log.Fatalf("[CONFIG] Failed to initialize tracer: %v", err)
	// }
	// defer t.Shutdown(ctx)

	// Open MySQL DB Connection
	db, err := openDatabases(cfg)
	if err != nil {
		log.Fatalf("[DB] Failed to initialize database connection: %v", err)
	}

	// Open MySQL DB Connection
	f, err := firebaseclient.NewClient(cfg, cred)
	if err != nil {
		log.Fatalf("[FIREBASE] Failed to initialize firebase client: %v", err)
	}
	fs := f.StorageClient

	ctx := context.Background()
	// fmt.Println("cfg.Firebase", cfg.Firebase)
	// fmt.Println("cfg.Database.Master", cfg.Database.Master)
	// fmt.Println("cred", cred)
	firebaseApp, err := openFirebaseClient(ctx, cfg.Firebase, cred)
	if err != nil {
		log.Fatalf("[FIREBASE] Failed to initialize firebase client: %v", err)
	}
	// fmt.Println("firebaseApp", firebaseApp)

	fsdb, err := openFirestoreClient(ctx, firebaseApp)
	if err != nil {
		log.Fatalf("[FIRESTORE] Failed to initialize Firestore client: %v", err)
	}
	defer fsdb.Close()

	// fsdb, err := openFirebaseDatabaseClient(ctx, firebaseApp)
	// if err != nil {
	// 	log.Fatalf("[FIREBASE] Failed to initialize Realtime Database client: %v", err)
	// }

	// Firebase Client Init
	// fcmCredB2BPelapak, err := firebaseclient.NewClient(cfg.Firebase.FcmProjectIDB2BPelapak, cred)
	// if err != nil {
	// 	log.Fatalf("[FIREBASE] Failed to initialize firebase client: %v", err)
	// }
	// fcmB2BPelapak := fcmCredB2BPelapak.MessagingClient

	//
	docs.SwaggerInfo.Host = cfg.Swagger.Host
	docs.SwaggerInfo.Schemes = cfg.Swagger.Schemes

	// Set logger used for jaeger
	logger, _ := zap.NewDevelopment(
		zap.AddStacktrace(zapcore.FatalLevel),
		zap.AddCallerSkip(1),
	)
	zapLogger := logger.With(zap.String("service", "goldgym"))
	zlogger := jaegerLog.NewFactory(zapLogger)
	// loggers := Log.NewLogrusLogger()
	// Set tracer for service
	tracer, closer := tracing.Init("goldgym", zlogger)
	defer closer.Close()

	// httpc := httpclient.NewClient(tracer)
	// ad := auth.New(httpc, cfg.API.Auth)

	sdst := goldgymStockData.New(db, fsdb, fs, rdb, tracer, zlogger)
	ssst := goldgymStockService.New(sdst, tracer, zlogger)

	// Diganti dengan domain yang anda buat
	sd := goldgymData.New(db, tracer, zlogger)
	// ss := goldgymService.New(sd, ad, tracer, zlogger)
	ss := goldgymService.New(sd, tracer, zlogger)
	sh := goldgymHandler.New(ss, ssst, tracer, zlogger)

	sha := authHandler.New(ss, tracer, zlogger)
	// sh := goldgymHandler.New(ss, tracer, zlogger)

	// sdpn := pushNotifData.New(fcmB2BPelapak, loggers)
	// sspn := pushNotifService.New(sdpn, t.Tracer, loggers)
	// spnh := pushNotifHandler.New(sspn, loggers)

	config.PrepareWatchPath()
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		err := config.Init()
		if err != nil {
			log.Printf("[VIPER] Error get config file, %v", err)
		}
		cfg, cred = config.Get()
		masterNew, err := openDatabases(cfg)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize database connection: %v", err)
		} else {
			*db = *masterNew
			sd.InitStmt()
		}

	})
	s := goldgymServer.Server{
		Goldgym: sh,
		Auth:    sha,
		// PushNotification: spnh,
	}

	if err := s.Serve(cfg.Server.Port); err != http.ErrServerClosed {
		return err
	}

	return nil
}

func openDatabases(cfg *config.Config) (master *sqlx.DB, err error) {
	master, err = openConnectionPool("mysql", cfg.Database.Master)
	if err != nil {
		return master, err
	}

	return master, err
}

func openConnectionPool(driver string, connString string) (db *sqlx.DB, err error) {
	// // ----------------------------------- test tunnel -----------------------------------
	// // SSH configuration
	// sshConfig := &ssh.ClientConfig{
	// 	User: "butuhdok",
	// 	Auth: []ssh.AuthMethod{
	// 		ssh.Password("Zgamersz123"),
	// 	},
	// 	HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	// }

	// // Connect to SSH server
	// sshClient, err := ssh.Dial("tcp", "leafeon.rapidplex.com:64000", sshConfig)
	// if err != nil {
	// 	log.Fatalf("Failed to connect to SSH server: %v", err)
	// }
	// // defer sshClient.Close()

	// log.Printf("test %+v", sshClient)

	// // // Create a local forwarding port
	// // localAddr := "localhost:3306"
	// // localListener, err := sshClient.Listen("tcp", localAddr)
	// // if err != nil {
	// // 	fmt.Println("Failed to listen on local port:", err)
	// // 	return
	// // }
	// // defer localListener.Close()

	// // MySQL configuration
	// mysql.RegisterDial("mysql+tcp",
	// 	func(addr string) (net.Conn, error) {
	// 		return sshClient.Dial("tcp", addr)
	// 	})
	// log.Println("test", connString)

	// splitFunc := func(c rune) bool {
	// 	return c == ':' || c == '@' || c == '(' || c == ')'
	// }

	// words := strings.FieldsFunc(connString, splitFunc)
	// // userAndPass := strings.Split(connString, ":")

	// log.Printf("test %+v", words)

	// // if len(words) >= 2 {
	// // user := words[0]
	// // pass := words[1]
	// // tcp := words[2]
	// // ip := words[3]
	// // port := words[4]
	// // database := words[5]
	// // // log.Printf("testText %+v", selectedWord)
	// // // }

	// // // MySQL configuration
	// // mysqlConfig := mysql.Config{
	// // 	User:   user,
	// // 	Passwd: pass,
	// // 	Addr:   ip + "+" + port,
	// // 	Net:    tcp,
	// // 	DBName: database,
	// // }

	// // // Establish a connection to MySQL through SSH tunnel
	// // tunnel, err := sshClient.Dial("tcp", "127.0.0.1:3306")
	// // if err != nil {
	// // 	log.Fatalf("Failed to establish SSH tunnel: %v", err)
	// // }
	// // connString.Conn = tunnel
	// // dsn := connString.FormatDSN()
	// // ----------------------------------- test tunnel -----------------------------------
	db, err = sqlx.Open(driver, connString)
	if err != nil {
		return db, err
	}

	err = db.Ping()
	if err != nil {
		return db, err
	}

	return db, err
}

func newRedisClient(cred config.Redis) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     cred.Host,
		Password: cred.Password,
		DB:       0,
	})
	return client
}

func openFirebaseClient(ctx context.Context, cfg config.FirebaseConfig, cred map[string]string) (*firebase.App, error) {
	credBytes, err := json.Marshal(cred)
	if err != nil {
		return nil, err
	}

	opt := option.WithCredentialsJSON(credBytes)
	app, err := firebase.NewApp(ctx, &firebase.Config{
		ProjectID:     cfg.ProjectID,
		DatabaseURL:   cfg.DatabaseURL,
		StorageBucket: cfg.StorageBucket,
	}, opt)

	if err != nil {
		return nil, err
	}

	return app, nil
}

func openFirestoreClient(ctx context.Context, app *firebase.App) (*firestore.Client, error) {
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func openFirebaseDatabaseClient(ctx context.Context, app *firebase.App) (*db.Client, error) {
	client, err := app.Database(ctx)
	if err != nil {
		fmt.Println("ERR-CLIENT", err)
		return nil, err
	}
	fmt.Println("clientz", client)
	return client, nil
}

// package boot

// import (
// 	// "context"

// 	"gold-gym-be/docs"
// 	"gold-gym-be/internal/data/auth"
// 	"net"

// 	// "gold-gym-be/pkg/firebaseclient"
// 	"gold-gym-be/pkg/httpclient"
// 	"gold-gym-be/pkg/tracing"
// 	"log"
// 	"net/http"
// 	"strings"

// 	"gold-gym-be/internal/config"
// 	jaegerLog "gold-gym-be/pkg/log"

// 	// Log "gold-gym-be/pkg/logs"

// 	"github.com/fsnotify/fsnotify"
// 	"github.com/go-sql-driver/mysql"
// 	"github.com/jmoiron/sqlx"
// 	"github.com/spf13/viper"
// 	"go.uber.org/zap"
// 	"go.uber.org/zap/zapcore"
// 	"golang.org/x/crypto/ssh"

// 	// "golang.org/x/net/trace"
// 	// "go.opentelemetry.io/otel/trace"
// 	// "gold-gym-be/pkg/trace"

// 	goldgymData "gold-gym-be/internal/data/goldgym"
// 	goldgymServer "gold-gym-be/internal/delivery/http"
// 	goldgymHandler "gold-gym-be/internal/delivery/http/goldgym"
// 	goldgymService "gold-gym-be/internal/service/goldgym"
// 	// pushNotifData "gold-gym-be/internal/data/pushnotif"
// 	// pushNotifHandler "gold-gym-be/internal/delivery/http/pushnotif"
// 	// pushNotifService "gold-gym-be/internal/service/pushnotif"
// 	"github.com/casbin/casbin/v2"
// 	"github.com/fsnotify/fsnotify"
// 	"github.com/jmoiron/sqlx"
// 	sqladapter "github.com/Blank-Xu/sqlx-adapter"
// 	"github.com/spf13/viper"
// 	"github.com/uptrace/opentelemetry-go-extra/otelsql"
// 	"github.com/uptrace/opentelemetry-go-extra/otelsqlx"
// 	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"

// 	logger "core-be/pkg/log"
// 	"gold-gym-be/pkg/trace"
// )

// // // HTTP will load configuration, do dependency injection and then start the HTTP server
// // func HTTP() error {
// // 	// var (
// // 	// 	ctx = context.Background()
// // 	// )
// // 	err := config.Init()
// // 	if err != nil {
// // 		log.Fatalf("[CONFIG] Failed to initialize config: %v", err)
// // 	}
// // 	cfg := config.Get()

// // 	// t, err := trace.New(ctx, cfg.Trace.Exporter)
// // 	// if err != nil {
// // 	// 	log.Fatalf("[CONFIG] Failed to initialize tracer: %v", err)
// // 	// }
// // 	// defer t.Shutdown(ctx)

// // 	// Open MySQL DB Connection
// // 	db, err := openDatabases(cfg)
// // 	if err != nil {
// // 		log.Fatalf("[DB] Failed to initialize database connection: %v", err)
// // 	}

// // 	// Firebase Client Init
// // 	// fcmCredB2BPelapak, err := firebaseclient.NewClient(cfg.Firebase.FcmProjectIDB2BPelapak, cred)
// // 	// if err != nil {
// // 	// 	log.Fatalf("[FIREBASE] Failed to initialize firebase client: %v", err)
// // 	// }
// // 	// fcmB2BPelapak := fcmCredB2BPelapak.MessagingClient

// // 	//
// // 	docs.SwaggerInfo.Host = cfg.Swagger.Host
// // 	docs.SwaggerInfo.Schemes = cfg.Swagger.Schemes

// // 	// Set logger used for jaeger
// // 	logger, _ := zap.NewDevelopment(
// // 		zap.AddStacktrace(zapcore.FatalLevel),
// // 		zap.AddCallerSkip(1),
// // 	)
// // 	zapLogger := logger.With(zap.String("service", "goldgym"))
// // 	zlogger := jaegerLog.NewFactory(zapLogger)
// // 	// loggers := Log.NewLogrusLogger()
// // 	// Set tracer for service
// // 	tracer, closer := tracing.Init("goldgym", zlogger)
// // 	defer closer.Close()

// // 	httpc := httpclient.NewClient(tracer)
// // 	ad := auth.New(httpc, cfg.API.Auth)

// // 	// Diganti dengan domain yang anda buat
// // 	sd := goldgymData.New(db, tracer, zlogger)
// // 	ss := goldgymService.New(sd, ad, tracer, zlogger)
// // 	sh := goldgymHandler.New(ss, tracer, zlogger)

// // 	// sdpn := pushNotifData.New(fcmB2BPelapak, loggers)
// // 	// sspn := pushNotifService.New(sdpn, t.Tracer, loggers)
// // 	// spnh := pushNotifHandler.New(sspn, loggers)

// // 	config.PrepareWatchPath()
// // 	viper.WatchConfig()
// // 	viper.OnConfigChange(func(e fsnotify.Event) {
// // 		err := config.Init()
// // 		if err != nil {
// // 			log.Printf("[VIPER] Error get config file, %v", err)
// // 		}
// // 		cfg := config.Get()
// // 		masterNew, err := openDatabases(cfg)
// // 		if err != nil {
// // 			log.Fatalf("[DB] Failed to initialize database connection: %v", err)
// // 		} else {
// // 			*db = *masterNew
// // 			sd.InitStmt()
// // 		}

// // 	})
// // 	s := goldgymServer.Server{
// // 		Goldgym: sh,
// // 		// PushNotification: spnh,
// // 	}

// // 	if err := s.Serve(cfg.Server.Port); err != http.ErrServerClosed {
// // 		return err
// // 	}

// // 	return nil
// // }

// // func openDatabases(cfg *config.Config) (master *sqlx.DB, err error) {
// // 	master, err = openConnectionPool("mysql", cfg.Database.Master)
// // 	if err != nil {
// // 		return master, err
// // 	}

// // 	return master, err
// // }

// // func openConnectionPool(driver string, connString string) (db *sqlx.DB, err error) {
// // 	// ----------------------------------- test tunnel -----------------------------------
// // 	// SSH configuration
// // 	sshConfig := &ssh.ClientConfig{
// // 		User: "butuhdok",
// // 		Auth: []ssh.AuthMethod{
// // 			ssh.Password("Zgamersz123"),
// // 		},
// // 		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
// // 	}

// // 	// Connect to SSH server
// // 	sshClient, err := ssh.Dial("tcp", "leafeon.rapidplex.com:64000", sshConfig)
// // 	if err != nil {
// // 		log.Fatalf("Failed to connect to SSH server: %v", err)
// // 	}
// // 	// defer sshClient.Close()

// // 	log.Printf("test %+v", sshClient)

// // 	// // Create a local forwarding port
// // 	// localAddr := "localhost:3306"
// // 	// localListener, err := sshClient.Listen("tcp", localAddr)
// // 	// if err != nil {
// // 	// 	fmt.Println("Failed to listen on local port:", err)
// // 	// 	return
// // 	// }
// // 	// defer localListener.Close()

// // 	// MySQL configuration
// // 	mysql.RegisterDial("mysql+tcp",
// // 		func(addr string) (net.Conn, error) {
// // 			return sshClient.Dial("tcp", addr)
// // 		})
// // 	log.Println("test", connString)

// // 	splitFunc := func(c rune) bool {
// // 		return c == ':' || c == '@' || c == '(' || c == ')'
// // 	}

// // 	words := strings.FieldsFunc(connString, splitFunc)
// // 	// userAndPass := strings.Split(connString, ":")

// // 	log.Printf("test %+v", words)

// // 	// if len(words) >= 2 {
// // 	// user := words[0]
// // 	// pass := words[1]
// // 	// tcp := words[2]
// // 	// ip := words[3]
// // 	// port := words[4]
// // 	// database := words[5]
// // 	// // log.Printf("testText %+v", selectedWord)
// // 	// // }

// // 	// // MySQL configuration
// // 	// mysqlConfig := mysql.Config{
// // 	// 	User:   user,
// // 	// 	Passwd: pass,
// // 	// 	Addr:   ip + "+" + port,
// // 	// 	Net:    tcp,
// // 	// 	DBName: database,
// // 	// }

// // 	// // Establish a connection to MySQL through SSH tunnel
// // 	// tunnel, err := sshClient.Dial("tcp", "127.0.0.1:3306")
// // 	// if err != nil {
// // 	// 	log.Fatalf("Failed to establish SSH tunnel: %v", err)
// // 	// }
// // 	// connString.Conn = tunnel
// // 	// dsn := connString.FormatDSN()
// // 	// ----------------------------------- test tunnel -----------------------------------
// // 	db, err = sqlx.Open(driver, connString)
// // 	if err != nil {
// // 		return db, err
// // 	}

// // 	err = db.Ping()
// // 	if err != nil {
// // 		return db, err
// // 	}

// // 	return db, err
// // }

// // package boot

// // import (
// 	// "context"
// 	// "core-be/docs"
// 	// log "core-be/pkg/clog"
// 	// "core-be/pkg/httpclient"
// 	// "net/http"

// 	// "core-be/internal/config"

// 	// authDatav2 "core-be/internal/data/auth/v2"
// 	// authHandlerv2 "core-be/internal/delivery/http/auth/v2"
// 	// authServicev2 "core-be/internal/service/auth/v2"

// 	// coreDatav1 "core-be/internal/data/core/v1"
// 	// coreHandlerv1 "core-be/internal/delivery/http/core/v1"
// 	// coreServicev1 "core-be/internal/service/core/v1"

// 	// chatWAData "core-be/internal/data/chat-wa"

// 	// httpServer "core-be/internal/delivery/http"

// 	// "github.com/casbin/casbin/v2"
// 	// "github.com/fsnotify/fsnotify"
// 	// "github.com/jmoiron/sqlx"
// 	// sqladapter "github.com/Blank-Xu/sqlx-adapter"
// 	// "github.com/spf13/viper"
// 	// "github.com/uptrace/opentelemetry-go-extra/otelsql"
// 	// "github.com/uptrace/opentelemetry-go-extra/otelsqlx"
// 	// semconv "go.opentelemetry.io/otel/semconv/v1.4.0"

// 	// logger "core-be/pkg/log"
// 	// "gold-gym-be/pkg/trace"
// // )

// // HTTP will load configuration, do dependency injection and then start the HTTP server
// func HTTP() error {
// 	var (
// 		ctx = context.Background()
// 	)

// 	err := config.Init()
// 	if err != nil {
// 		log.Fatalf("[CONFIG] Failed to initialize config: %v", err)
// 	}
// 	cfg := config.Get()

// 	t, err := trace.New(ctx, cfg.Trace.Exporter)
// 	if err != nil {
// 		log.Fatalf("[CONFIG] Failed to initialize tracer: %v", err)
// 	}
// 	defer t.Shutdown(ctx)

// 	logger := logger.NewLogrusLogger()

// 	httpc := httpclient.NewClient(
// 		httpclient.WithTracer(t.Tracer),
// 	)

// 	coreDB, err := openConnectionPool("mysql", cfg.Database.Master)
// 	if err != nil {
// 		log.Fatalf("[atlasDB] Failed to open sql connection pool: %v", err)
// 	}

// 	casbinAdapter, err := sqladapter.NewAdapter(coreDB, "core_policy")
// 	if err != nil {
// 		log.Fatalf("[CASBIN] NewEnforcer failed to create new adapter: %v", err)
// 	}

// 	authEnforcer, err := casbin.NewEnforcer("auth_model.conf", casbinAdapter)
// 	if err != nil {
// 		log.Fatalf("[CASBIN] NewEnforcer failed to creates an enforcer: %v", err)
// 	}
// 	authEnforcer.AddFunction("coreMatch", authServicev2.KeyMatchFunc)

// 	//
// 	docs.SwaggerInfo.Host = cfg.Swagger.Host
// 	docs.SwaggerInfo.Schemes = cfg.Swagger.Schemes

// 	_authDatav2 := authDatav2.New(coreDB)
// 	_coreDatav1 := coreDatav1.New(coreDB, authEnforcer)
// 	_chatWA := chatWAData.New(httpc, cfg.API.ChatWA)

// 	_coreServicev1 := coreServicev1.New(_coreDatav1, t.Tracer, logger)
// 	_authServicev2 := authServicev2.New(_authDatav2, _coreServicev1, _chatWA, t.Tracer, logger, authEnforcer)

// 	_authHandlerv2 := authHandlerv2.New(_authServicev2, logger)
// 	_coreHandlerv1 := coreHandlerv1.New(_coreServicev1, logger)

// 	config.PrepareWatchPath()
// 	viper.OnConfigChange(func(e fsnotify.Event) {
// 		err := config.Init()
// 		if err != nil {
// 			log.Printf("[VIPER] Error get config file, %v", err)
// 		}
// 		cfg := config.Get()

// 		coreNew, err := openConnectionPool("mysql", cfg.Database.Master)
// 		if err != nil {
// 			log.Printf("[VIPER] Error open db connection, %v", err)
// 		} else {
// 			*coreDB = *coreNew
// 			_coreDatav1.InitStmt()
// 			_authDatav2.InitStmt()
// 		}
// 	})

// 	s := httpServer.Server{
// 		AuthV2:         _authHandlerv2,
// 		CoreV1:         _coreHandlerv1,
// 	}

// 	if err := s.Serve(cfg.Server.Port); err != http.ErrServerClosed {
// 		return err
// 	}

// 	return nil
// }

// func openConnectionPool(driver string, connString string) (db *sqlx.DB, err error) {
// 	db, err = otelsqlx.Open(
// 		driver,
// 		connString,
// 		otelsql.WithDBSystem(semconv.DBSystemMySQL.Value.AsString()),
// 	)
// 	if err != nil {
// 		return db, err
// 	}

// 	err = db.Ping()
// 	if err != nil {
// 		return db, err
// 	}

// 	return db, err
// }
