package response

import "github.com/MuhammadSuryono/module-golang-server/server"

func BuildResponseJson(code int, value interface{}) {
	server.Context.JSON(code, value)
}
