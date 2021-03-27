package api

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/akacokafor/microscope/internal"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

type MicroScopeViewModel struct {
	CssFilePath string
	JsFilePath  string
	AppName     string
	Data        map[string]interface{}
}

func ContentTypeJson(next httprouter.Handle) httprouter.Handle {
	return func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
		next(rw, r, p)
	}
}

func NewHTTPRouter(prefix string, isProduction bool, options GoCraftOptions) http.Handler {
	router := httprouter.New()
	RegisterRoutes(router, prefix, isProduction, options)
	return router
}

func RegisterRoutes(router *httprouter.Router, prefix string, isProduction bool, options GoCraftOptions) {
	staticFs, err := internal.GetStaticFileSystem(isProduction)
	if err != nil {
		logrus.WithError(err).Fatalf("error creating file system")
	}

	indexTemplate, err := template.New("index_template").ParseFS(staticFs, "*.html")
	if err != nil {
		logrus.WithError(err).Fatalf("error creating file system")
	}

	staticDir := "/static/"
	basePath := fmt.Sprintf("/%s", prefix)
	baseRouterPath := fmt.Sprintf("/%s", prefix)
	indexHandler := func(w http.ResponseWriter, r *http.Request) {
		indexTemplate.ExecuteTemplate(w, "index.html", MicroScopeViewModel{
			CssFilePath: fmt.Sprintf("%s%s", staticDir, mix(staticFs, "/app.css", isProduction)),
			JsFilePath:  fmt.Sprintf("%s%s", staticDir, mix(staticFs, "/app.js", isProduction)),
			AppName:     "Microscope",
			Data: map[string]interface{}{
				"timezone":  "Africa/Lagos",
				"path":      basePath,
				"recording": false,
				"apiPath":   fmt.Sprintf("%s/api/", basePath),
			},
		})
	}

	router.Handler("GET", fmt.Sprintf("%s*path", staticDir), http.StripPrefix(staticDir, http.FileServer(http.FS(staticFs))))
	router.GET(basePath, func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		indexHandler(rw, r)
	})
	router.GET(fmt.Sprintf("%s/#/*paths", baseRouterPath), func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		indexHandler(rw, r)
	})

	router.GET(fmt.Sprintf("%s/api/requests", baseRouterPath), ContentTypeJson(func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		rw.Write([]byte(`{"entries": []}`))
	}))
	router.POST(fmt.Sprintf("%s/api/requests", baseRouterPath), ContentTypeJson(func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		rw.Write([]byte(`{"entries": []}`))
	}))

	gocraftApi := NewGoCraftApi(options)
	router.GET(fmt.Sprintf("%s/api/queues", baseRouterPath), gocraftApi.queues)
	router.GET(fmt.Sprintf("%s/api/worker_pools", baseRouterPath), gocraftApi.workerPools)
	router.GET(fmt.Sprintf("%s/api/busy_workers", baseRouterPath), gocraftApi.busyWorkers)
	router.GET(fmt.Sprintf("%s/api/retry_jobs", baseRouterPath), gocraftApi.retryJobs)
	router.GET(fmt.Sprintf("%s/api/scheduled_jobs", baseRouterPath), gocraftApi.scheduledJobs)
	router.GET(fmt.Sprintf("%s/api/dead_jobs", baseRouterPath), gocraftApi.deadJobs)
	router.POST(fmt.Sprintf("%s/api/delete_dead_job/:died_at/:job_id", baseRouterPath), gocraftApi.deleteDeadJob)
	router.POST(fmt.Sprintf("%s/api/retry_dead_job/:died_at/:job_id", baseRouterPath), gocraftApi.retryDeadJob)
	router.POST(fmt.Sprintf("%s/api/delete_all_dead_jobs", baseRouterPath), gocraftApi.deleteAllDeadJobs)
	router.POST(fmt.Sprintf("%s/api/retry_all_dead_jobs", baseRouterPath), gocraftApi.retryAllDeadJobs)
}

func RegisterRoutesOnGin(router gin.IRoutes, prefix string, isProduction bool, options GoCraftOptions) {
	staticFs, err := internal.GetStaticFileSystem(isProduction)
	if err != nil {
		logrus.WithError(err).Fatalf("error creating file system")
	}

	indexTemplate, err := template.New("index_template").ParseFS(staticFs, "*.html")
	if err != nil {
		logrus.WithError(err).Fatalf("error creating file system")
	}

	staticDir := "/static/"
	basePath := fmt.Sprintf("/%s", prefix)
	baseRouterPath := fmt.Sprintf("/%s", prefix)
	indexHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		indexTemplate.ExecuteTemplate(w, "index.html", MicroScopeViewModel{
			CssFilePath: fmt.Sprintf("%s%s", staticDir, mix(staticFs, "/app.css", isProduction)),
			JsFilePath:  fmt.Sprintf("%s%s", staticDir, mix(staticFs, "/app.js", isProduction)),
			AppName:     "Microscope",
			Data: map[string]interface{}{
				"timezone":  "Africa/Lagos",
				"path":      basePath,
				"recording": false,
				"apiPath":   fmt.Sprintf("%s/api/", basePath),
			},
		})
	}

	router.GET(fmt.Sprintf("%s*path", staticDir), func(c *gin.Context) {
		http.StripPrefix(staticDir, http.FileServer(http.FS(staticFs))).ServeHTTP(c.Writer, c.Request)
	})
	router.GET(basePath, func(gtx *gin.Context) {
		indexHandler(gtx.Writer, gtx.Request)
	})
	router.GET(fmt.Sprintf("%s/#/*paths", baseRouterPath), func(gtx *gin.Context) {
		indexHandler(gtx.Writer, gtx.Request)
	})

	adapter := func(v httprouter.Handle) gin.HandlerFunc {
		return func(gtx *gin.Context) {
			var px httprouter.Params
			copier.Copy(&px, &gtx.Params)
			v(gtx.Writer, gtx.Request, px)
		}
	}

	gocraftApi := NewGoCraftApi(options)
	router.GET(fmt.Sprintf("%s/api/queues", baseRouterPath), adapter(gocraftApi.queues))
	router.GET(fmt.Sprintf("%s/api/worker_pools", baseRouterPath), adapter(gocraftApi.workerPools))
	router.GET(fmt.Sprintf("%s/api/busy_workers", baseRouterPath), adapter(gocraftApi.busyWorkers))
	router.GET(fmt.Sprintf("%s/api/retry_jobs", baseRouterPath), adapter(gocraftApi.retryJobs))
	router.GET(fmt.Sprintf("%s/api/scheduled_jobs", baseRouterPath), adapter(gocraftApi.scheduledJobs))
	router.GET(fmt.Sprintf("%s/api/dead_jobs", baseRouterPath), adapter(gocraftApi.deadJobs))
	router.POST(fmt.Sprintf("%s/api/delete_dead_job/:died_at/:job_id", baseRouterPath), adapter(gocraftApi.deleteDeadJob))
	router.POST(fmt.Sprintf("%s/api/retry_dead_job/:died_at/:job_id", baseRouterPath), adapter(gocraftApi.retryDeadJob))
	router.POST(fmt.Sprintf("%s/api/delete_all_dead_jobs", baseRouterPath), adapter(gocraftApi.deleteAllDeadJobs))
	router.POST(fmt.Sprintf("%s/api/retry_all_dead_jobs", baseRouterPath), adapter(gocraftApi.retryAllDeadJobs))
}
