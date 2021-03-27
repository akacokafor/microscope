package api

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/akacokafor/microscope/internal"
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
	indexHandler := func(w http.ResponseWriter, r *http.Request) {
		indexTemplate.ExecuteTemplate(w, "index.html", MicroScopeViewModel{
			CssFilePath: fmt.Sprintf("%sapp.css", staticDir),
			JsFilePath:  fmt.Sprintf("%sapp.js", staticDir),
			AppName:     "Microscope",
			Data: map[string]interface{}{
				"timezone":  "Africa/Lagos",
				"path":      basePath,
				"recording": false,
				"apiPath":   fmt.Sprintf("%s/api/", basePath),
			},
		})
	}

	router := httprouter.New()

	router.Handler("GET", fmt.Sprintf("%s*path", staticDir), http.StripPrefix(staticDir, http.FileServer(http.FS(staticFs))))
	router.GET(basePath, func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		indexHandler(rw, r)
	})
	router.GET(fmt.Sprintf("%s/#/*paths", basePath), func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		indexHandler(rw, r)
	})

	router.GET(fmt.Sprintf("%s/api/requests", basePath), ContentTypeJson(func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		rw.Write([]byte(`{"entries": []}`))
	}))
	router.POST(fmt.Sprintf("%s/api/requests", basePath), ContentTypeJson(func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		rw.Write([]byte(`{"entries": []}`))
	}))

	gocraftApi := NewGoCraftApi(options)
	router.GET(fmt.Sprintf("%s/api/queues", basePath), gocraftApi.queues)
	router.GET(fmt.Sprintf("%s/api/worker_pools", basePath), gocraftApi.workerPools)
	router.GET(fmt.Sprintf("%s/api/busy_workers", basePath), gocraftApi.busyWorkers)
	router.GET(fmt.Sprintf("%s/api/retry_jobs", basePath), gocraftApi.retryJobs)
	router.GET(fmt.Sprintf("%s/api/scheduled_jobs", basePath), gocraftApi.scheduledJobs)
	router.GET(fmt.Sprintf("%s/api/dead_jobs", basePath), gocraftApi.deadJobs)
	router.POST(fmt.Sprintf("%s/api/delete_dead_job/:died_at/:job_id", basePath), gocraftApi.deleteDeadJob)
	router.POST(fmt.Sprintf("%s/api/retry_dead_job/:died_at/:job_id", basePath), gocraftApi.retryDeadJob)
	router.POST(fmt.Sprintf("%s/api/delete_all_dead_jobs", basePath), gocraftApi.deleteAllDeadJobs)
	router.POST(fmt.Sprintf("%s/api/retry_all_dead_jobs", basePath), gocraftApi.retryAllDeadJobs)
	return router
}
