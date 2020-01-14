package api

import controllers "github.com/eai-rts/gate5/internal/api/conrtollers"

var server = controllers.Server{}

// Run is API's entry point
func Run() {
	controllers.NewServer().Run(":8080")
}
