package handler

import (
	"fmt"
	"net/http"
)

func Ends(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	key := queryParams.Get("toolkit")
	w.Header().Set("tips", "Tips can be used one more time.")
	if key == "work" {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, "~谢谢尝试过了前面的关卡\n因为是尝试所以没有考虑写太多，如果有兴趣的话，咱会试着再做一点好玩的\n不过，你还是来错了地方，猜一猜还剩下什么呢w")
		return
	}
	if key == "krow" {
		http.Redirect(w,r,"https://t.me/+fa6ycNIdJj0yM2M9",200)
		return
	}
	fmt.Fprintf(w,"?")
}
