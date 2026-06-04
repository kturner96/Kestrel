package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
)

type ReedSwitch struct {
	Id int `json:"sensor_id"`
	DoorOpen bool `json:"door_open"`
	TriggeredAt time.Time `json:"triggered_at"`
}

func (h *Handler) HandlePost(w http.ResponseWriter, r *http.Request) {
	var payload ReedSwitch
	json.NewDecoder(r.Body).Decode(&payload)



	var nodeId int
	var sessionId int
	ctx := context.Background()

	tx, err := h.Pool.BeginTx(ctx, pgx.TxOptions{}) 
	if err != nil {
		http.Error(w, "starting transaction: ", http.StatusFailedDependency)
		return
	}
	

	defer tx.Rollback(ctx)

	err = h.Pool.QueryRow(ctx, "SELECT sensor_node_id FROM sensors WHERE sensor_id = $1", payload.Id).Scan(&nodeId)
	if err != nil {
		http.Error(w, "sensor not found ", http.StatusNotFound)
		return
	}

	err = tx.QueryRow(ctx, "INSERT INTO sessions (sensor_node_id, started_at) VALUES ($1, $2) RETURNING session_id", nodeId, payload.TriggeredAt).Scan(&sessionId)
	if err != nil {
		http.Error(w, "Failed to insert into sessions: ", http.StatusNotModified)
		return
	}


	_, err = tx.Exec(ctx, "INSERT INTO events (value_bool, recorded_at, session_id, sensor_id) VALUES ($1, $2, $3, $4)", payload.DoorOpen, payload.TriggeredAt, sessionId, payload.Id)
		if err != nil {
			http.Error(w, "Failed to insert into events: ", http.StatusNotModified)
			return
		}
	
	if err := tx.Commit(ctx); err != nil {
		http.Error(w, "Failed to commit transaction: ", http.StatusNotImplemented)
		return
	}

	// doorStatus := "closed."

	// if payload.DoorOpen {
	// 	doorStatus = "open."
	// }

	

	// fmt.Printf("---ReedSwitch---\n")
	// fmt.Printf("Sensor ID: %d\n Time: %s\n Door is %s\n ", payload.Id, payload.TriggeredAt,doorStatus)

}