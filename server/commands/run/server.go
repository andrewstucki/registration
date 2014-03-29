package run

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"time"
	"compress/gzip"
	"io"
	"strings"

	"github.com/GeertJohan/go.rice"
	"github.com/emicklei/go-restful"

	"../../commands"
	"../../database"
	"../../endpoints"
	"../../globals"
	"../../network"
	"../../utils"
)

var ServerCmd = &commands.Command{
	Name: "run:server",
	Usage: `
Run:server starts the registration server.

The following flags are available:

  -backgrounds
    The folder with 7 alternate backgrounds to use. Images must be named bg1.jpg ... bg7.jpg (defaults to none).
  -config
    The path to the web page configuration file (defaults to "config.json").
  -schema
	The path to the schema file (defaults to "schema.json").
  -database
    The path to the registration database (defaults to "registration.db").
  -interface
    Interface for ad-hoc network (default auto-detects interface).
  -network
    Name of the ad-hoc network (defaults to "registrationNetwork").
  -security
    Security type for the network, options are WEP40 or WEP104. (defaults to none).
  -password
    The password for the ad-hoc network (must be specified if -security is used is used).
  -private
    Option to create a private ad-hoc network. Use in conjunction with interface, network, password, security, and channel flags (defaults to false).
  -channel
    Channel for the ad-hoc network (defaults to 6).
  -port
    Port for serving registration page (defaults to 8080).
  `,
	Summary: "Run the registration server",
	Run:     runServer,
}

var backgroundsPath string
var configPath string
var schemaPath string
var databasePath string
var interfaceName string
var privateNetwork bool
var networkName string
var networkChannel uint
var security string
var password string
var port uint

var config []byte

var quit = make(chan os.Signal, 1)

func runServer(cmd *commands.Command, args ...string) {
	var ip string
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	config = file

	rand.Seed(time.Now().UTC().UnixNano())
	if !globals.Debug {
		log.SetOutput(ioutil.Discard)
	}
	signal.Notify(quit, os.Interrupt)
	fmt.Printf("Using database at %s...\n", databasePath)
	err = database.Initialize(databasePath)
	defer database.Cleanup()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Using schema at %s...\n", schemaPath)
	err = database.LoadSchemaFromFile(schemaPath)
	if err != nil {
		panic(err)
	}

	if privateNetwork {
		fmt.Printf("Configuring network.\n")
		ip = network.Start(interfaceName, security, networkName, networkChannel, password)
		defer network.Cleanup()
	} else {
		ip = "localhost"
	}

	serverLoop(port)
	site := fmt.Sprintf("http://%s:%d", ip, port)
	fmt.Printf("\nServer listening at %s!\n", site)
	utils.OpenSite(site)

	<-quit
	fmt.Println("\nShutting down server")
}

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}
 
func (w gzipResponseWriter) Write(b []byte) (int, error) {
  if "" == w.Header().Get("Content-Type") {
      // If no content type, apply sniffing algorithm to un-gzipped body.
      w.Header().Set("Content-Type", http.DetectContentType(b))
  }
  return w.Writer.Write(b)
}

func gzipHandler(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			h.ServeHTTP(w, r)
			return
		}
		w.Header().Set("Content-Encoding", "gzip")
		gz := gzip.NewWriter(w)
		defer gz.Close()
		gzr := gzipResponseWriter{Writer: gz, ResponseWriter: w}
		h.ServeHTTP(gzr, r)
	}
}

func serverLoop(port uint) {
	wsContainer := restful.NewContainer()
	signup := endpoints.SignupEndpoint{}
	signup.Register(wsContainer)
	addr := fmt.Sprintf(":%d", port)

	if backgroundsPath == "" {
		backgroundBox := rice.MustFindBox("../../backgrounds")
		wsContainer.Handle("/backgrounds/", http.StripPrefix("/backgrounds/", gzipHandler(http.FileServer(backgroundBox.HTTPBox()))))
		log.Println("Using built-in backgrounds")
	} else {
		wsContainer.Handle("/backgrounds/", http.StripPrefix("/backgrounds/", gzipHandler(http.FileServer(http.Dir(backgroundsPath)))))
		log.Printf("Serving backgrounds from %s", backgroundsPath)
	}

	assetBox := rice.MustFindBox("../../assets")

	// register on different handler so that the path prefixes don't conflict
	wsContainer.Handle("/", gzipHandler(http.FileServer(assetBox.HTTPBox())))

	schemaService := new(restful.WebService)
	schemaService.Path("/schema.json").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	schemaService.Route(schemaService.GET("/").To(serveSchema))
	wsContainer.Add(schemaService)

	configService := new(restful.WebService)
	configService.Path("/config.json").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	configService.Route(configService.GET("/").To(serveConfig))
	wsContainer.Add(configService)

	go func() {
		// wrap so that default mux is still useable
		log.Fatal(http.ListenAndServe(addr, wsContainer))
	}()
}

func serveSchema(request *restful.Request, response *restful.Response) {
	response.Header().Set(restful.HEADER_ContentType, restful.MIME_JSON)
	response.Write([]byte(database.Schema()))
}

func serveConfig(request *restful.Request, response *restful.Response) {
	response.Header().Set(restful.HEADER_ContentType, restful.MIME_JSON)
	response.Write(config)
}

func init() {
	ServerCmd.Flag.StringVar(&backgroundsPath, "backgrounds", "", "")
	ServerCmd.Flag.StringVar(&configPath, "config", "config.json", "")
	ServerCmd.Flag.StringVar(&schemaPath, "schema", "schema.json", "")
	ServerCmd.Flag.StringVar(&databasePath, "database", "registration.db", "")
	ServerCmd.Flag.StringVar(&interfaceName, "interface", "", "")
	ServerCmd.Flag.StringVar(&networkName, "network", "registrationNetwork", "")
	ServerCmd.Flag.StringVar(&security, "security", "", "")
	ServerCmd.Flag.StringVar(&password, "password", "", "")
	ServerCmd.Flag.BoolVar(&privateNetwork, "private", false, "")
	ServerCmd.Flag.UintVar(&networkChannel, "channel", 6, "")
	ServerCmd.Flag.UintVar(&port, "port", 8080, "")
}
