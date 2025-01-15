package db

// import (
// 	"fmt"
// 	// "log"
// 	// "os"
// 	// "time"
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	// "crypto/x509"
// 	// "crypto/tls"
// 	// "gorm.io/gorm/logger"
// 	// "github.com/jackc/pgx/v5/stdlib"
// )

// var sql *gorm.DB

// // InitSQL initializes a connection to the PostgreSQL database.
// func InitSQL(host string, port int, user, password, name string) error {
// 	// Construct the DSN for PostgreSQL


// 	// dsn := url.URL{
//     //     User:     url.UserPassword(conf.User, conf.Password),
//     //     Scheme:   "postgres",
//     //     Host:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
//     //     Path:     conf.DBName,
//     //     RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
//     // }
//     // db, err := gorm.Open("postgres", dsn.String())
// 	// rootCertPool := x509.NewCertPool()
// 	// if ok := rootCertPool.AppendCertsFromPEM([]byte(caCert)); !ok {
// 	// 	return fmt.Errorf("failed to append CA certificate")
// 	// }

// 		// Register connection configuration
// 		// stdlib.RegisterConnConfig(&pgx.ConnConfig{
// 		// 	Config: pgconn.Config{
// 		// 		TLSConfig: &tls.Config{RootCAs: rootCertPool},
// 		// 	},
// 		// })


// 		// _ = postgres.RegisterTLSConfig("custom", &tls.Config{
// 		// 	RootCAs:    rootCertPool,
// 		// })
// 	// host=localhost user=postgres password=postgres dbname=astra port=5432 sslmode=disable TimeZone=UTC
// 	dsn := fmt.Sprintf(
// 		"host=%s user=%s password=%s dbname=%s port=%v sslmode=disable TimeZone=UTC",
// 		host, user, password, name, port,
// 	)
// 	fmt.Println(dsn)

// 	m, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// 	// m, err := gorm.Open(postgres.New(postgres.Config{
// 	// 	DSN: dsn,
// 	// }), &gorm.Config{
// 	// 	PrepareStmt: true,
// 	// })
// 	// m, err := gorm.Open(postgres.Open(dsn), 
// 	// 	&gorm.Config{
// 	// 			Logger: logger.New(
// 	// 				log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
// 	// 				logger.Config{
// 	// 					SlowThreshold:             time.Second, // Slow SQL threshold
// 	// 					LogLevel:                  logger.Info, // Log level
// 	// 					IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
// 	// 					ParameterizedQueries:      false,       // Don't include params in the SQL log
// 	// 					Colorful:                  true,        // Enable color
// 	// 				},
// 	// 			),
// 	// 		},
// 	// )

// 	// Open the database connection
// 	// m, err := gorm.Open(postgres.New(postgres.Config{
// 	// 	DSN: dsn, // Data Source Name
// 	// }), &gorm.Config{
// 	// 	Logger: logger.New(
// 	// 		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
// 	// 		logger.Config{
// 	// 			SlowThreshold:             time.Second, // Slow SQL threshold
// 	// 			LogLevel:                  logger.Info, // Log level
// 	// 			IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
// 	// 			ParameterizedQueries:      false,       // Don't include params in the SQL log
// 	// 			Colorful:                  true,        // Enable color
// 	// 		},
// 	// 	),
// 	// })

// 	if err != nil {
// 		return err
// 	}

// 	sql = m
// 	return nil
// }

// // SQL returns the current database connection instance.
// func SQL() *gorm.DB {
// 	return sql
// }
