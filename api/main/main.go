package main

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	fileutil "github.com/jhegarty14/comhrac-api/pkg/fileutil"
	vparse "github.com/jhegarty14/comhrac-api/pkg/version"

	log "github.com/sirupsen/logrus"
)

func init() {
	if "DEV" == strings.ToUpper(os.Getenv("ENV")) {
		log.SetFormatter(&log.TextFormatter{})
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetFormatter(&log.JSONFormatter{})
		log.SetLevel(log.InfoLevel)
	}
}

func addMainRouteHandlers(r *mux.Router) {
	r.HandleFunc("/", HomepageHandler).Methods("GET")
	r.HandleFunc("/", HomepageHandler).Methods("POST")
}

// CreateRouter => init router and pass it to subroute trees
func CreateRouter() *mux.Router {
	r := mux.NewRouter()
	addMainRouteHandlers(r)
	addUsersRouterHandlers(r)

	return r
}

func main() {
	// ===========================================================================
	// Load environment variables
	// ===========================================================================
	var (
		env     = strings.ToUpper(os.Getenv("ENV")) // LOCAL, DEV, STG, PRD
		port    = os.Getenv("PORT")                 // server traffic on this port
		version = os.Getenv("VERSION")              // path to VERSION file
	)
	// ===========================================================================
	// Read environment variables
	// ===========================================================================
	config, configErr := fileutil.ReadPropertiesFile(strings.ToLower(env))
	if configErr != nil {
		log.WithFields(log.Fields{
			"env":  env,
			"err":  configErr,
			"path": os.Getenv("VERSION"),
		}).Fatal("Can't find a config file")
		os.Exit(1)
	}
	// ===========================================================================
	// Read version information
	// ===========================================================================
	version, err := vparse.ParseVersionFile(version)
	if err != nil {
		log.WithFields(log.Fields{
			"env":  env,
			"err":  err,
			"path": os.Getenv("VERSION"),
		}).Fatal("Can't find a VERSION file")
		return
	}
	log.WithFields(log.Fields{
		"env":     env,
		"path":    os.Getenv("VERSION"),
		"version": version,
	}).Info("Loaded VERSION file")
	// ===========================================================================
	// Create database connection
	// ===========================================================================
	dbUser, dbUserExists := config["GO_DB_USER"]
	dbPwd, dbPwdExists := config["GO_DB_PASSWORD"]
	dbHost, dbHostExists := config["GO_DB_HOST"]
	dbPort, dbPortExists := config["GO_DB_PORT"]
	dbName, dbNameExists := config["GO_DB_NAME"]
	if !dbUserExists || !dbPwdExists || !dbHostExists || !dbPortExists || !dbNameExists {
		log.Fatal(env + "config file does not contain complete Postgres URL")
		os.Exit(1)
	}
	url := "postgres://" + dbUser + ":" + dbPwd + "@" + dbHost + ":" + dbPort + "/" + dbName
	dbConn, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	// Get generic database object sql.DB to use its functions
	dB, err := dbConn.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	dB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	dB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	dB.SetConnMaxLifetime(time.Hour / 4)

	if err != nil {
		log.Fatal(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	// set gorm sqlDB to gloabl DB variable
	sqldb.globalDB = dB

	defer sqldb.globalDB.Close()
	// ===========================================================================
	// Start application
	// ===========================================================================
	router := CreateRouter()

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}
