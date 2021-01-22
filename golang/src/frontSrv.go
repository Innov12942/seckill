package main

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type frontHandler struct {
}
type req_t struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type req_token struct {
	Token    string `json:"token"`
	Getgoods string `json:"getgoods"`
}

type req_kill struct {
	Token  string `json:"token"`
	Goodid string `json:"goodid"`
}

type req_res struct {
	Token    string `json:"token"`
	Getreult string `json:"getresult"`
}

type response_t struct {
	Token string `json:"token"`
}

func (m *frontHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	// fmt.Println(buf.String())
	reqstr := buf.String()

	keystr := []string{"select", "drop", "create", "delete"}
	for _, kstr := range keystr {
		ridx := strings.Index(reqstr, kstr)
		if ridx >= 0 {
			w.WriteHeader(404)
			fmt.Fprintf(w, "Status:Failed\n")
			return
		}
	}

	if strings.Contains(reqstr, "username") && strings.Contains(reqstr, "password") {
		fmt.Println(reqstr)
		req := &req_t{}
		err := json.Unmarshal([]byte(reqstr), req)
		if err != nil {
			fmt.Println("Json Unmarshal error")
		}

		flagshort := len(req.Password) < 3
		match, _ := regexp.MatchString("(\\d+.*[a-zA-Z]+)|([a-zA-Z]+.*\\d+)", req.Password)

		if flagshort || !match {
			fmt.Println("password unregularized")
		}

		uinfo := sc_user{}
		res := db.Where("username=?", req.Username).First(&uinfo)

		if res.RowsAffected == 0 || uinfo.Password != req.Password {
			fmt.Printf("username:%s password:%s Incorrect\n", req.Username, req.Password)
			w.WriteHeader(404)
			fmt.Fprintf(w, "Status:Failed\n")
			return
		}

		curtime := time.Now().String()
		md5sum := md5.Sum([]byte(curtime + req.Username + req.Password))
		var md5hex string

		md5hex = fmt.Sprintf("%x", md5sum)
		InsertEntry(md5hex, uinfo.ID)

		resp := &response_t{md5hex}
		jsonresp, jerr := json.Marshal(resp)
		if jerr != nil {
			fmt.Println("Json Marshal error")
		}
		w.Write(jsonresp)
		return
	}

	fmt.Println(reqstr)
	if strings.Contains(reqstr, "getgoods") {
		reqtk := &req_token{}
		err := json.Unmarshal([]byte(reqstr), reqtk)
		if err != nil {
			fmt.Println("Json Unmarshal error")
		}

		//Wrong token or expired
		if FindEntry(reqtk.Token) == 0 {
			w.WriteHeader(404)
			fmt.Fprintf(w, "Status:Failed\n")
			return
		}

		finalstr := GetAll()
		finalstr += "A"

		fmt.Println(finalstr)
		fmt.Fprintf(w, "%s", finalstr)
	}

	if strings.Contains(reqstr, "goodid") {
		reqkl := &req_kill{}
		err := json.Unmarshal([]byte(reqstr), reqkl)
		if err != nil {
			fmt.Println("Json Unmarshal error")
		}

		//Wrong token or expired
		if FindEntry(reqkl.Token) == 0 {
			w.WriteHeader(404)
			fmt.Fprintf(w, "Status:Failed\n")
			return
		}

		uid := FindEntry(reqkl.Token)
		gidint, _ := strconv.Atoi(reqkl.Goodid)
		Killone(uid, gidint)
	}

	if strings.Contains(reqstr, "getresult") {
		reqres := &req_res{}
		err := json.Unmarshal([]byte(reqstr), reqres)
		if err != nil {
			fmt.Println("Json Unmarshal error")
		}

		//Wrong token or expired
		if FindEntry(reqres.Token) == 0 {
			w.WriteHeader(404)
			fmt.Fprintf(w, "Status:Failed\n")
			return
		}

		uid := FindEntry(reqres.Token)

		var gu []sc_good_user
		_ = db.Where("uid = ?", uid).Find(&gu)

		var finalstr string
		finalstr += ""
		for i, vgoods := range gu {
			vgj, _ := json.Marshal(vgoods)
			if i == 0 {
				finalstr += string(vgj)
			} else {
				finalstr += "&" + string(vgj)
			}
		}
		finalstr += "A"

		fmt.Printf("finalstr: %s \n", finalstr)
		fmt.Fprintf(w, "%s", finalstr)
	}

}

func Killone(uid int, gid int) {
	gu := sc_good_user{}
	gures := db.Where("uid = ? AND gid = ?", uid, gid).First(&gu)
	if gures.RowsAffected > 0 {
		return
	}

	ginfo := sc_good{}
	res := db.Where("id=?", gid).First(&ginfo)
	if res.RowsAffected > 0 && ginfo.Expire > time.Now().Unix() {
		fmt.Println("Gooods not expired!")
		return
	}
	if res.RowsAffected > 0 && ginfo.Remain > 0 {
		ginfo.Remain -= 1
		db.Save(&ginfo)

		gunew := sc_good_user{gid, uid}
		db.Create(&gunew)
	}
}

func FrontSrv() {
	InitRedis()
	err_ := http.ListenAndServe("127.0.0.1:12347", &frontHandler{})
	if err_ != nil {
		panic("ListenAndServe failed!")
	}
}
