package main

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type RoomDb struct {
	rooms []bool
	m     sync.Mutex
}

var roomDb = RoomDb{rooms: make([]bool, 45), m: sync.Mutex{}}
var wg = sync.WaitGroup{}

func main() {

	http.HandleFunc("/bookRoom", bookRoom)
	http.HandleFunc("/emptyRoom", emptyRoom)
	http.ListenAndServe(":8080", nil)
}

func emptyRoom(w http.ResponseWriter, r *http.Request) {
	defer roomDb.m.Unlock()
	if r.Method != "POST" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	r.ParseForm()
	roomNo, err := strconv.ParseInt(r.PostForm.Get("roomNo"), 10, 8)
	if err != nil {
		http.Error(w, "Invalid Room No", http.StatusBadRequest)
		return
	}
	if roomNo > 45 || roomNo < 0 {
		http.Error(w, "Invalid Room No", http.StatusBadRequest)
		return
	}
	roomDb.m.Lock()
	if roomDb.rooms[roomNo] == false {
		http.Error(w, "Room Already Empty", http.StatusBadRequest)
		return
	}
	roomDb.rooms[roomNo] = false
	fmt.Println(roomNo)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Room Status Updated Successfully")
}

func bookRoom(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	r.ParseForm()

	roomNo, err := strconv.ParseInt(r.PostForm.Get("roomNo"), 10, 8)
	if err != nil {
		http.Error(w, "Invalid Room No", http.StatusBadRequest)
		return
	}
	if roomNo > 45 || roomNo < 0 {
		http.Error(w, "Invalid Room No", http.StatusBadRequest)
		return
	}

	wg.Add(1)
	go processRoom(roomNo, &w, &wg)
	wg.Wait()
}

func processRoom(roomNo int64, w *http.ResponseWriter, wg *sync.WaitGroup) {
	defer wg.Done()
	defer roomDb.m.Unlock()
	roomDb.m.Lock()
	time.Sleep(5 * time.Second)
	if roomDb.rooms[roomNo] == true {
		http.Error(*w, "Room Already Occupied", http.StatusBadRequest)
		return
	}
	roomDb.rooms[roomNo] = true
	(*w).WriteHeader(http.StatusOK)
	fmt.Fprint(*w, "Room Status Updated")
	fmt.Println(roomNo)

}
