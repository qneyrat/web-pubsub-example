package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app          = kingpin.New("chat", "CLI Chat client.")
	serverAPI    = app.Flag("api", "Server API address.").Default("127.0.0.1").String()
	serverWS     = app.Flag("ws", "Server WS address.").Default("127.0.0.1").String()
	conversation = app.Flag("conversation", "Conversation id.").Default("1").String()
	token        = app.Flag("token", "Auth token.").Default("").String()

	auth         = app.Command("auth", "sign in user.")
	authUsername = auth.Arg("username", "Username for user.").Required().String()
	authPassword = auth.Arg("password", "Password of user.").Required().String()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	case auth.FullCommand():
		token := authUser()
		go getMessages(token)
		scanBuffer(token)
	}
}

func authUser() string {
	data := url.Values{}
	data.Set("username", *authUsername)
	data.Add("password", *authPassword)

	client := &http.Client{}
	r, err := http.NewRequest("POST", "http://"+*serverAPI+"/login_check", strings.NewReader(data.Encode()))
	if err != nil {
		log.Println(err)
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	response, err := client.Do(r)
	if err != nil {
		log.Println(err)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	var authData map[string]interface{}
	if err := json.Unmarshal(contents, &authData); err != nil {
		panic(err)
	}

	return authData["token"].(string)
}

func getMessages(token string) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	c, _, err := websocket.DefaultDialer.Dial("ws://"+*serverWS+"/websocket?token="+token, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer c.Close()
		defer close(done)

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("ReadMessage:", err)
				return
			}

			log.Printf("New message: %s", message)
		}
	}()

	for {
		select {
		case <-interrupt:
			log.Println("interrupt")

			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}

			c.Close()
			os.Exit(1)
			return
		}
	}
}

func scanBuffer(token string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		postMessage(token, *conversation, scanner.Text())
	}

	if scanner.Err() != nil {
		log.Println(scanner.Err())
	}
}

func postMessage(token string, conversation string, message string) {
	client := &http.Client{}
	r, err := http.NewRequest(
		"POST",
		"http://"+*serverAPI+"/conversations/"+conversation+"/messages",
		strings.NewReader("{\n\t\"body\": \""+message+"\"\n}"),
	)
	if err != nil {
		log.Println(err)
	}
	r.Header.Add("authorization", "Bearer "+token)
	r.Header.Add("content-type", "application/json")
	r.Header.Add("cache-control", "no-cache")

	response, err := client.Do(r)
	if err != nil {
		log.Println(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Print(string(body))
}
