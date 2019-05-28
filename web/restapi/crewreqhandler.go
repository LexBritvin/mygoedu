package restapi

import (
	"encoding/json"
	"fmt"
	"mygoedu/dblayer"
	"net/http"
	"strconv"
)

type MyCrewReqHandler struct {
	dbConn dblayer.DBLayer
}

func NewCrewReqHandler() *MyCrewReqHandler {
	return new(MyCrewReqHandler)
}

func (hcwreq *MyCrewReqHandler) connect(o, conn string) error {
	db, err := dblayer.ConnectDatabase(o, conn)
	if err != nil {
		return err
	}
	hcwreq.dbConn = db
	return nil
}

func (hcwreq *MyCrewReqHandler) handleCrewRequests(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ids := r.RequestURI[len("/api/crew/"):]
		id, err := strconv.Atoi(ids)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "id %s provided is not of valid number. \n", ids)
			return
		}
		cm, err := hcwreq.dbConn.FindMember(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error %s occured when search for id %d \n ", err.Error(), id)
			return
		}
		json.NewEncoder(w).Encode(&cm)

	case "POST":
		cm := new(dblayer.CrewMember)
		err := json.NewDecoder(r.Body).Decode(cm)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Error %s occured", err)
			return
		}

		err = hcwreq.dbConn.AddMember(cm)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error %s occured while adding a crew member to the DB", err)
			return
		}
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Successfully inserted id %d \n", cm.ID)

	}
}
