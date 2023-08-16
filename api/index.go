package handler

import (
	"fmt"
	"net/http"

	wt "github.com/MoYoez/text-watermark"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("tips", "Some Water Dropped On the Text, Can you help me to clean it?")
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	secretPath1 := "Achieved The First Goal. | Router To next: time | Requests Args: ?timeunix=?"
	encText := wt.AddWaterMarkToText("To identify the world", secretPath1)
	fmt.Fprintf(w, encText+"\n"+"You should consider more what I want to give.")
}
