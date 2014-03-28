package endpoints

import (
	"net/http"

	"github.com/emicklei/go-restful"

	"../database"
)

type SignupEndpoint struct{}

func (r SignupEndpoint) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/signup").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.POST("/").To(r.signup))

	container.Add(ws)
}

func (r SignupEndpoint) signup(request *restful.Request, response *restful.Response) {
	var signup map[string]interface{}
	request.ReadEntity(&signup)

	err := database.Signup(signup)
	if err != nil {
		response.ResponseWriter.WriteHeader(http.StatusBadRequest)
		response.WriteEntity(map[string]string{"message": err.Error()})
	} else {
		response.WriteEntity(map[string]interface{}{"success": true})
	}
}
