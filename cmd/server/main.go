package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net"
	"strings"
	"wowpow/internal/conf"
	"wowpow/internal/proto"
	"wowpow/internal/res"

	"github.com/catalinc/hashcash"
	"github.com/google/uuid"
)

func main() {
	res.LoadQuotes()
	config := conf.NewServerConfig()
	ln, err := net.Listen("tcp", fmt.Sprintf(":%v", config.Port))
	if err != nil {
		panic(err)
	}

	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func sendResponse(conn net.Conn, rsp *proto.Response) {
	bytes, err := json.Marshal(rsp)
	bytes = append(bytes, '\n')
	if err != nil {
		panic(err)
	}

	conn.Write(bytes)
}

func sendError(conn net.Conn, err string) {
	rsp := proto.Response{
		Error: &proto.Error{
			Message: err,
		},
	}

	sendResponse(conn, &rsp)
}

func sendQuote(conn net.Conn, quote *proto.Quote) {
	rsp := proto.Response{
		Result: quote,
	}

	sendResponse(conn, &rsp)
}

func sendChallenge(conn net.Conn, text string) {
	rsp := proto.Response{
		Challenge: &proto.Challenge{
			Text: text,
		},
	}

	sendResponse(conn, &rsp)
}

func handleConnection(conn net.Conn) {
	var idStr string
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		bytes, err := reader.ReadBytes('\n')
		if err != nil {
			sendError(conn, fmt.Sprintf("read error: '%v'", err.Error()))
			return
		}

		var req proto.Request
		err = json.Unmarshal(bytes, &req)
		if err != nil {
			sendError(conn, fmt.Sprintf("parse error: '%v'", err.Error()))
			return
		}

		if req.Solution == nil {
			if len(idStr) > 0 {
				sendError(conn, "proto error: no challenge response")
				return
			}

			id, err := uuid.NewUUID()
			if err != nil {
				panic(err)
			}

			idStr = id.String()
			sendChallenge(conn, idStr)
		} else {
			if len(idStr) == 0 {
				sendError(conn, "proto error: no challenge requested")
				return
			}

			if !strings.Contains(req.Solution.Stamp, idStr) {
				sendError(conn, "challenge error: fraud stamp")
				return
			}

			h := hashcash.NewStd()
			if !h.Check(req.Solution.Stamp) {
				sendError(conn, "challenge error: invalid stamp")
				return
			}

			idStr = ""
			sendQuote(conn, res.GetRandomQuote())
		}
	}
}
