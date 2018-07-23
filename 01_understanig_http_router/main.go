package main

import (
	"net/http"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/altex/Matching-engine/order"
	"log"
	"encoding/json"
	"github.com/altex/Matching-engine/Matcher"
)

type Response struct {
	Order order.LimitOrder
	Matched interface{}

}

var i=0
func defaultHandler(w http.ResponseWriter,r *http.Request,_ httprouter.Params)  {
	fmt.Fprintln(w,"Hello",i)
	i++

}
func greeter(w http.ResponseWriter,r *http.Request, param httprouter.Params)  {
	 m:=order.LimitOrder{
	 }

	 res:=Response{
	 	Order:m,
	 }
	err:=json.NewDecoder(r.Body).Decode(&m)
	defer r.Body.Close()
	if err!=nil{
		log.Fatal(err," Error")
	}
	log.Println(m)
	ords:=Matcher.MatchLimitOrder(&m)
	res.Matched=ords
	log.Println("Matched orders: ",ords)
	err=json.NewEncoder(w).Encode(&res)
	if err!=nil{
		fmt.Println(m)
	}
	w.Header().Set("Content-Type","apllication/json")
	fmt.Fprintln(w,)

}
func main()  {
	router:=httprouter.New()
	router.GET("/",defaultHandler)
	router.POST("/index/:name",greeter)

	log.Fatal(http.ListenAndServe(":8080",router))
}

