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
//    TAGS
// =============================

func createTags(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	type payload struct {
		Name string `json:"name_tags"`
	}
	var p payload
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	err = application.SaveTags(p.Name)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusCreated, nil)
}

func getTag(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tagID, err := strconv.Atoi(ps.ByName("tag_id"))
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	tag, err := application.GetTag(tagID)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusOK, tag)
}

func getAllTag(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tags, err := application.GetAllTag()
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusOK, tags)
}

func removeTag(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	tagID, err := strconv.Atoi(ps.ByName("tag_id"))
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	err = application.RemoveTag(tagID)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusOK, nil)
}

func updateTag(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var p domain.Tags
	err := decoder.Decode(&p)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
	}

	tagID, err := strconv.Atoi(ps.ByName("tag_id"))
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	err = application.UpdateTag(p, tagID)
	if err != nil {
		Error(w, http.StatusNotFound, err, err.Error())
		return
	}

	JSON(w, http.StatusOK, nil)
}
