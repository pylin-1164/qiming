package restful

import "github.com/emicklei/go-restful"



type IRestfulServer interface{
	RegistRoute(server *restful.WebService)
}
var WS = &restful.WebService{}

func  StartRestfulServer() {

	WS.Path("/").
		Consumes(restful.MIME_JSON,restful.MIME_XML,"application/x-www-form-urlencoded").Produces(restful.MIME_JSON,restful.MIME_XML,"application/x-www-form-urlencoded")
	/*ws.
		Path("/users").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML)

	ws.Route(ws.GET("/{user-id}").To(u.findUser).
		Doc("get a user").
		Param(ws.PathParameter("user-id", "identifier of the user").DataType("string")).
		Writes(User{}))

	...

	func (u UserResource) findUser(request *restful.Request, response *restful.Response) {
		id := request.PathParameter("user-id")
		...
	}*/
}
