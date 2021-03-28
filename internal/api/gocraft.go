package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gocraft/work"
	"github.com/gomodule/redigo/redis"
	"github.com/julienschmidt/httprouter"
)

type GoCraftOptions struct {
	Namespace string
	Pool      *redis.Pool
}

type Context struct {
	namespace string
	pool      *redis.Pool
	client    *work.Client
}

// NewServer creates and returns a new server. The 'namespace' param is the redis namespace to use. The hostPort param is the address to bind on to expose the API.
func NewGoCraftApi(options GoCraftOptions) *Context {
	return &Context{
		namespace: options.Namespace,
		pool:      options.Pool,
		client:    work.NewClient(options.Namespace, options.Pool),
	}
}

func (c *Context) queues(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response, err := c.client.Queues()
	data := map[string]interface{}{
		"entries": response,
	}
	render(rw, data, err)
}

func (c *Context) workerPools(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	response, err := c.client.WorkerPoolHeartbeats()
	data := map[string]interface{}{
		"entries": response,
	}
	render(rw, data, err)
}

func (c *Context) busyWorkers(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	observations, err := c.client.WorkerObservations()
	if err != nil {
		renderError(rw, err)
		return
	}

	var busyObservations []*work.WorkerObservation
	for _, ob := range observations {
		if ob.IsBusy {
			busyObservations = append(busyObservations, ob)
		}
	}

	if len(busyObservations) > 0 {
		data := map[string]interface{}{
			"entries": busyObservations,
		}
		render(rw, data, err)
		return
	}

	render(rw, map[string]interface{}{
		"entries": []string{},
	}, err)
}

func (c *Context) retryJobs(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	page, err := parsePage(r)
	if err != nil {
		renderError(rw, err)
		return
	}

	jobs, count, err := c.client.RetryJobs(page)
	if err != nil {
		renderError(rw, err)
		return
	}

	response := struct {
		Count int64            `json:"count"`
		Jobs  []*work.RetryJob `json:"entries"`
	}{Count: count, Jobs: jobs}

	render(rw, response, err)
}

func (c *Context) scheduledJobs(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	page, err := parsePage(r)
	if err != nil {
		renderError(rw, err)
		return
	}

	jobs, count, err := c.client.ScheduledJobs(page)
	if err != nil {
		renderError(rw, err)
		return
	}

	response := struct {
		Count int64                `json:"count"`
		Jobs  []*work.ScheduledJob `json:"entries"`
	}{Count: count, Jobs: jobs}

	render(rw, response, err)
}

func (c *Context) deadJobs(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	page, err := parsePage(r)
	if err != nil {
		renderError(rw, err)
		return
	}

	jobs, count, err := c.client.DeadJobs(page)
	if err != nil {
		renderError(rw, err)
		return
	}

	response := struct {
		Count int64           `json:"count"`
		Jobs  []*work.DeadJob `json:"entries"`
	}{Count: count, Jobs: jobs}

	render(rw, response, err)
}

func (c *Context) deleteDeadJob(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	diedAt, err := strconv.ParseInt(ps.ByName("died_at"), 10, 64)
	if err != nil {
		renderError(rw, err)
		return
	}

	err = c.client.DeleteDeadJob(diedAt, ps.ByName("job_id"))

	render(rw, map[string]string{"status": "ok"}, err)
}

func (c *Context) retryDeadJob(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	diedAt, err := strconv.ParseInt(ps.ByName("died_at"), 10, 64)
	if err != nil {
		renderError(rw, err)
		return
	}

	err = c.client.RetryDeadJob(diedAt, ps.ByName("job_id"))

	render(rw, map[string]string{"status": "ok"}, err)
}

func (c *Context) deleteAllDeadJobs(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := c.client.DeleteAllDeadJobs()
	render(rw, map[string]string{"status": "ok"}, err)
}

func (c *Context) retryAllDeadJobs(rw http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := c.client.RetryAllDeadJobs()
	render(rw, map[string]string{"status": "ok"}, err)
}

func render(rw http.ResponseWriter, jsonable interface{}, err error) {
	if err != nil {
		renderError(rw, err)
		return
	}

	jsonData, err := json.MarshalIndent(jsonable, "", "\t")
	if err != nil {
		renderError(rw, err)
		return
	}
	rw.Write(jsonData)
}

func renderError(rw http.ResponseWriter, err error) {
	rw.WriteHeader(500)
	fmt.Fprintf(rw, `{"error": "%s"}`, err.Error())
}

func parsePage(r *http.Request) (uint, error) {
	err := r.ParseForm()
	if err != nil {
		return 0, err
	}

	pageStr := r.Form.Get("page")
	if pageStr == "" {
		pageStr = "1"
	}

	page, err := strconv.ParseUint(pageStr, 10, 0)
	return uint(page), err
}
