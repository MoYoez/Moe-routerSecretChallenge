package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func timeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("tips", "Always Reverse")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	queryParams := r.URL.Query()
	timeUnixStr := queryParams.Get("timeunix")
	timezone := time.FixedZone("UTC+8", 8*60*60)
	currentTime := time.Now().In(timezone).Unix()
	if timeUnixStr == "" {
		fmt.Fprintf(w, "You Entered in a wrong way, maybe You should consider more ^^")
		return
	}
	if r.Method != http.MethodPost {
		fmt.Fprintf(w, "Data Always Reverse,isn't it?")
		return
	}
	ToInt64,err := strconv.ParseInt(timeUnixStr, 10, 64)
	if err != nil {
		fmt.Fprintf(w, "You Entered in a wrong way, maybe You should consider more ^^")
		return
	}
	oneHourAgo := currentTime - 3600
	oneHourLater := currentTime + 3600
	if ToInt64 >= oneHourAgo && ToInt64 <= oneHourLater {
		w.Write([]byte("You Are Right"))
		time.Sleep(time.Second *5)
		http.Redirect(w,r,"https://secret-mnfi.impart.icu/api/end?toolkit=work",http.StatusSeeOther)
	} else {
		w.Write([]byte("You Entered in a wrong way, please try more or use tips ^^ ("))
	}
}