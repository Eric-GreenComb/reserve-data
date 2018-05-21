package http_runner

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"math"

	"github.com/KyberNetwork/reserve-data/common"
	raven "github.com/getsentry/raven-go"
	"github.com/gin-contrib/sentry"
	"github.com/gin-gonic/gin"
)

// MAX_TIMESPOT is the default time point to return in case the
// timestamp parameter in request is omit or malformed.
const MAX_TIMESPOT uint64 = math.MaxUint64

// HttpRunnerServer is the HTTP ticker server.
type HttpRunnerServer struct {
	runner *HttpRunner
	host   string
	r      *gin.Engine
	http   *http.Server
}

// getTimePoint returns the timepoint from query parameter.
// If no timestamp parameter is supplied, or it is invalid, returns the default one.
func getTimePoint(c *gin.Context) uint64 {
	timestamp := c.DefaultQuery("timestamp", "")
	if timestamp == "" {
		log.Printf("Interpreted timestamp(%s) to default - %d\n", timestamp, MAX_TIMESPOT)
		return MAX_TIMESPOT
	} else {
		timepoint, err := strconv.ParseUint(timestamp, 10, 64)
		if err != nil {
			log.Printf("Interpreted timestamp(%s) to default - %d\n", timestamp, MAX_TIMESPOT)
			return MAX_TIMESPOT
		} else {
			log.Printf("Interpreted timestamp(%s) to %d\n", timestamp, timepoint)
			return timepoint
		}
	}
}

func (self *HttpRunnerServer) otick(c *gin.Context) {
	timepoint := getTimePoint(c)
	self.runner.oticker <- common.TimepointToTime(timepoint)
	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
		},
	)
}

func (self *HttpRunnerServer) atick(c *gin.Context) {
	timepoint := getTimePoint(c)
	self.runner.aticker <- common.TimepointToTime(timepoint)
	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
		},
	)
}

func (self *HttpRunnerServer) rtick(c *gin.Context) {
	timepoint := getTimePoint(c)
	self.runner.rticker <- common.TimepointToTime(timepoint)
	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
		},
	)
}

func (self *HttpRunnerServer) btick(c *gin.Context) {
	timepoint := getTimePoint(c)
	self.runner.bticker <- common.TimepointToTime(timepoint)
	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
		},
	)
}

func (self *HttpRunnerServer) ttick(c *gin.Context) {
	timepoint := getTimePoint(c)
	self.runner.tticker <- common.TimepointToTime(timepoint)
	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
		},
	)
}

func (self *HttpRunnerServer) gtick(c *gin.Context) {
	timepoint := getTimePoint(c)
	self.runner.globalDataTicker <- common.TimepointToTime(timepoint)
	c.JSON(
		http.StatusOK,
		gin.H{
			"success": true,
		},
	)
}

func (self *HttpRunnerServer) init() {
	self.r.GET("/otick", self.otick)
	self.r.GET("/atick", self.atick)
	self.r.GET("/rtick", self.rtick)
	self.r.GET("/btick", self.btick)
	self.r.GET("/ttick", self.ttick)
	self.r.GET("/gtick", self.gtick)
}

func (self *HttpRunnerServer) Start() error {
	if self.http == nil {
		self.http = &http.Server{
			Addr:    self.host,
			Handler: self.r,
		}
		return self.http.ListenAndServe()
	} else {
		return errors.New("server start already")
	}
}

func (self *HttpRunnerServer) Stop() error {
	if self.http != nil {
		err := self.http.Shutdown(nil)
		self.http = nil
		return err
	} else {
		return errors.New("server stop already")
	}
}

// NewHttpRunnerServer creates a new instance of HttpRunnerServer.
func NewHttpRunnerServer(runner *HttpRunner, host string) *HttpRunnerServer {
	r := gin.Default()
	r.Use(sentry.Recovery(raven.DefaultClient, false))
	server := HttpRunnerServer{
		runner,
		host,
		r,
		nil,
	}
	server.init()
	return &server
}
