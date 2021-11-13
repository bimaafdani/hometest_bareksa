package interfaces

import (
	"HoteTestBareksa/application"
	"HoteTestBareksa/domain"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// =============================
//    NEWS
// =============================

func getNews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	param := ps.ByName("param")

	// if param is numeric than search by news_id, otherwise
	// if alphabetic then search by topic.Slug
	newsID, err := strconv.Atoi(param)
	if err != nil {
		// param is alphabetic
		news, err2 := application.GetNewsByTopic(param)
		if err2 != nil {
			Error(w, http.StatusNotFound, err2, err2.Error())
			return
		}

		JSON(w, http.StatusOK, news)
		return
	}

	// param is numeric
	news, err := application.GetNews(newsID)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusOK, news)
}

func getAllNews(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	queryValues := r.URL.Query()
	status := queryValues.Get("status")

	// if status parameter exist draft|deleted|publish
	if status == "draft" || status == "deleted" || status == "publish" {
		news, err := application.GetAllNewsByFilter(status)
		if err != nil {
			Error(w, http.StatusNotFound, err, err.Error())
			return
		}

		JSON(w, http.StatusOK, news)
		return
	}

	limit := queryValues.Get("limit")
	page := queryValues.Get("page")

	// if custom pagination exist news?limit=15&page=2
	if limit != "" && page != "" {
		limit, _ := strconv.Atoi(limit)
		page, _ := strconv.Atoi(page)

		if limit != 0 && page != 0 {
			news, err := application.GetAllNews(limit, page)
			if err != nil {
				Error(w, http.StatusNotFound, err, err.Error())
				return
			}

			JSON(w, http.StatusOK, news)
			return
		}
	}

	news, err := application.GetAllNews(15, 1) // 15, 1 default pagination
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusOK, news)
}

func createNews(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var p domain.News
	if err := decoder.Decode(&p); err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	err := application.AddNews(p)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusCreated, nil)
}

func removeNews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	newsID, err := strconv.Atoi(ps.ByName("news_id"))
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	err = application.RemoveNews(newsID)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusOK, nil)
}

func updateNews(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var p domain.News
	err := decoder.Decode(&p)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	newsID, err := strconv.Atoi(ps.ByName("news_id"))
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	err = application.UpdateNews(p, newsID)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusOK, nil)
}
