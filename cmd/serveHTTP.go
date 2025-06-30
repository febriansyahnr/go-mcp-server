package cmd

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"time"

	"github.com/paper-indonesia/pdk/go/monitoring"
	pdkLogger "github.com/paper-indonesia/pdk/v2/logger"
	pdkMySql "github.com/paper-indonesia/pdk/v2/mySqlExt"
	pdkNewRelic "github.com/paper-indonesia/pdk/v2/newRelicExt"
	"github.com/paper-indonesia/pdk/v2/otelExt"
	pdkRedis "github.com/paper-indonesia/pdk/v2/redisExt"
	"github.com/paper-indonesia/pg-mcp-server/config"
	"github.com/paper-indonesia/pg-mcp-server/constant"
	disbursementHandler "github.com/paper-indonesia/pg-mcp-server/internal/handlers/disbursement"
	extraHandler "github.com/paper-indonesia/pg-mcp-server/internal/handlers/extra"
	backendportalRepository "github.com/paper-indonesia/pg-mcp-server/internal/repository/backendPortal"
	snapcoreRepository "github.com/paper-indonesia/pg-mcp-server/internal/repository/snapcore"
	disbursementService "github.com/paper-indonesia/pg-mcp-server/internal/service/disbursement"
	pkgMonitor "github.com/paper-indonesia/pg-mcp-server/pkg/monitor"
	"github.com/paper-indonesia/pg-mcp-server/pkg/mySqlExt"
	"github.com/paper-indonesia/pg-mcp-server/pkg/rabbitMqExt"
	"github.com/paper-indonesia/pg-mcp-server/pkg/redisExt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(serveHTTPCmd)
}

var serveHTTPCmd = &cobra.Command{
	Use:   "serveHTTP",
	Short: "Start HTTP server",
	Long:  `Start MCP Server with HTTP`,
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		// Init config
		conf, secret, err := config.LoadConfig(cfgFile, scrtFile)
		if err != nil {
			log.Fatalf("Unable to load configuration and secret: %v", err)
		}

		isDevelopment := true
		if conf.Environment == constant.EnvironmentProduction {
			isDevelopment = false
		}

		pdkLog, err := pdkLogger.NewZapLogger(
			pdkLogger.Config{
				IsDevelopment: isDevelopment,
				Environment:   conf.Environment,
				ServiceName:   conf.ServiceName + "-http",
			},
		)
		if err != nil {
			log.Fatalf("Unable to initialize logger: %v", err)
			panic(err)
		}
		defer pdkLog.Sync()
		defer pdkLog.Info(context.Background(), "Service successfully stopped")

		otelOpts := []otelExt.OptionFunc{}
		if conf.OTLPConfig.Insecure {
			otelOpts = append(otelOpts, otelExt.WithInsecure())
		}
		if conf.OTLPConfig.TLSClientConfig != nil {
			otelOpts = append(otelOpts, otelExt.WithTLSClientConfig(&tls.Config{
				InsecureSkipVerify: conf.OTLPConfig.TLSClientConfig.InsecureSkipVerify,
			}))
		}
		otelExt, err := otelExt.New(
			otelExt.Config{
				ServiceName:  conf.ServiceName + "-http",
				Environment:  conf.Environment,
				OTLPEndpoint: conf.OTLPConfig.Host,
				LicenseKey:   secret.NewRelicLicenseKey,
			},
			otelOpts...,
		)
		if err != nil {
			fmt.Printf("Unable to init opentelemetry, %v", err)
			panic(err)
		}
		defer otelExt.Shutdown(ctx)

		// Init New Relic
		newRelicExt, err := pdkNewRelic.New(
			pdkNewRelic.Config{
				ServiceName: conf.ServiceName + "-http-" + conf.Environment,
				Environment: conf.Environment,
				LicenseKey:  secret.NewRelicLicenseKey,
			},
		)
		if err != nil {
			fmt.Printf("Unable to init new relic, %v", err)
			panic(err)
		}
		defer newRelicExt.GetApp().Shutdown(10 * time.Second)

		// Statsd Monitoring
		monitor, err := monitoring.New(
			conf.ServiceName+"-http-"+conf.Environment,
			secret.StatsdHost,
			secret.StatsdPort,
		)
		if err != nil {
			fmt.Printf("Unable to init monitoring, %v", err)
			panic(err)
		}
		monitor = monitor.WithOtelMeterProvider(otelExt.MeterProvider())
		pkgMonitor.SetGlobalMonitoring(monitor)

		// snap core db client
		snapCoreDBClient, err := mySqlExt.New(
			pdkMySql.Config{
				Host:         conf.MySQLConfig.Host,
				Port:         conf.MySQLConfig.Port,
				Username:     secret.MySQLSecret.SnapCore.Username,
				Password:     secret.MySQLSecret.SnapCore.Password,
				DBName:       secret.MySQLSecret.SnapCore.Database,
				MaxIdleConns: conf.MySQLConfig.MaxIdleConns,
				MaxIdleTime:  conf.MySQLConfig.MaxOpenConns,
				MaxLifeTime:  conf.MySQLConfig.MaxLifeTime,
				MaxOpenConns: conf.MySQLConfig.MaxOpenConns,
				SlaveHost:    conf.MySQLConfig.SlaveHost,
				SlavePort:    conf.MySQLConfig.SlavePort,
			},
			pdkMySql.WithLogger(pdkLog),
			pdkMySql.WithTracerProvider(otelExt.TracerProvider()),
			pdkMySql.WithMetricProvider(otelExt.MeterProvider()),
		)
		if err != nil {
			fmt.Printf("Unable to init mysql, %v", err)
			panic(err)
		}
		defer snapCoreDBClient.Close()

		// backend portal db client
		backendPortalDBClient, err := mySqlExt.New(
			pdkMySql.Config{
				Host:         conf.MySQLConfig.Host,
				Port:         conf.MySQLConfig.Port,
				Username:     secret.MySQLSecret.BackendPortal.Username,
				Password:     secret.MySQLSecret.BackendPortal.Password,
				DBName:       secret.MySQLSecret.BackendPortal.Database,
				MaxIdleConns: conf.MySQLConfig.MaxIdleConns,
				MaxIdleTime:  conf.MySQLConfig.MaxOpenConns,
				MaxLifeTime:  conf.MySQLConfig.MaxLifeTime,
				MaxOpenConns: conf.MySQLConfig.MaxOpenConns,
				SlaveHost:    conf.MySQLConfig.SlaveHost,
				SlavePort:    conf.MySQLConfig.SlavePort,
			},
			pdkMySql.WithLogger(pdkLog),
			pdkMySql.WithTracerProvider(otelExt.TracerProvider()),
			pdkMySql.WithMetricProvider(otelExt.MeterProvider()),
		)
		if err != nil {
			fmt.Printf("Unable to init mysql, %v", err)
			panic(err)
		}
		defer backendPortalDBClient.Close()

		// Redis
		cacheClient, err := redisExt.New(
			pdkRedis.Config{
				Addr:             conf.RedisConfig.Host + ":" + conf.RedisConfig.Port,
				Password:         secret.RedisSecret.Password,
				DB:               conf.RedisConfig.CacheDB,
				IsRedsyncEnabled: true,
				IsLimiterEnabled: true,
			},
			pdkRedis.WithTracerProvider(otelExt.TracerProvider()),
			pdkRedis.WithMetricProvider(otelExt.MeterProvider()),
		)

		pdkLog.Info(context.Background(), "Redis cache client created", pdkLogger.Field{
			Type:   pdkLogger.StringType,
			String: cacheClient.Ping(context.Background()).String(),
		})
		if err != nil {
			fmt.Printf("Unable to init redis cache, %v", err)
			panic(err)
		}
		defer cacheClient.Close()

		// Rabbit mq
		rabbitMqExt, err := rabbitMqExt.New(
			conf.RabbitMQConfig,
			secret.RabbitMQSecret,
			pdkLog,
			newRelicExt,
			rabbitMqExt.WithContext(ctx),
		)
		if err != nil {
			fmt.Printf("Unable to init rabbitmq, %v", err)
			panic(err)
		}

		defer rabbitMqExt.Close()

		// setup repositories
		backendPortalRepo := backendportalRepository.New(
			conf,
			backendportalRepository.WithDBClient(backendPortalDBClient),
			backendportalRepository.WithLogger(pdkLog),
		)

		snapcoreRepo := snapcoreRepository.New(
			conf,
			snapcoreRepository.WithDBClient(snapCoreDBClient),
			snapcoreRepository.WithLogger(pdkLog),
		)

		// setup services
		disbursementService := disbursementService.New(
			conf,
			disbursementService.WithBackendPortalRepository(backendPortalRepo),
			disbursementService.WithSnapCoreRepository(snapcoreRepo),
		)

		// setup handlers
		extrasHandler := extraHandler.New()
		disbursementHandler := disbursementHandler.New(
			conf,
			disbursementHandler.WithDisbursementService(disbursementService),
		)

		// setup mcp server
		mcpServer := NewMCPServer(
			conf,
			WithExtraHandler(extrasHandler),
			WithDisbursementHandler(disbursementHandler),
		)
		mcpServer.StartSSE()
	},
}
