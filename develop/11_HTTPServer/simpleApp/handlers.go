package simpleapp

import (
"encoding/json"
"net/http"

)

func (s *simpleApp) getEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	json.NewEncoder(w).Encode(s.events)
}

func (s *simpleApp) createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// var ev Event
	// _ = json.NewDecoder(r.Body).Decode(&ev)

	// if err := serializeToJSON("./one.json", ev); err != nil {
	// 	log.Fatal(err)
	// }


}

// func deserialize(filename string, intr interface{}) (interface{}, error) {
// 	file, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if err := json.Unmarshal(file, &intr); err != nil {
// 		return nil, err
// 	}

// 	return intr, nil
// }

// func serializeToJSON(file string, ev Event) error {
// 	_, err := deserialize(file, &evs)
// 	if err != nil {
// 		return err
// 	}

// 	evs = append(evs, ev)
// 	rowData, err := json.Marshal(evs)
// 	if err != nil {
// 		return err
// 	}

// 	if err := ioutil.WriteFile("./one.json", rowData, 0); err != nil {
// 		return err
// 	}

// 	return nil
// }
