package main

import (
	"encoding/json"
	"fmt"
	"net"
	"wowpow/internal/conf"
	"wowpow/internal/proto"

	"github.com/catalinc/hashcash"
)

func main() {
	config := conf.NewClientConfig()
	servAddr := fmt.Sprintf("%v:%v", config.Host, config.Port)
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		panic(err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	for i := 0; i < config.Requests; i++ {
		rsp, err := sendRequest(conn, nil)
		if err != nil {
			panic(err)
		}

		showResponse(rsp)
		if rsp.Challenge != nil {
			h := hashcash.NewStd()

			stamp, err := h.Mint(rsp.Challenge.Text)
			fmt.Printf("solution: '%v'\n", stamp)

			if err != nil {
				panic(err)
			}

			rsp, err = sendRequest(conn, &stamp)
			if err != nil {
				panic(err)
			}

			showResponse(rsp)
		}
	}
}

func showResponse(rsp *proto.Response) {
	if rsp.Error != nil {
		fmt.Printf("error from server: '%v'\n", rsp.Error.Message)
	} else if rsp.Result != nil {
		fmt.Printf("result from server:\n%v\n\t - %v\n", rsp.Result.Text, rsp.Result.Author)
	} else if rsp.Challenge != nil {
		fmt.Printf("challenge from server: '%v'\n", rsp.Challenge.Text)
	} else {
		panic("wrong answer from server")
	}
}

func sendRequest(conn *net.TCPConn, stamp *string) (*proto.Response, error) {
	req := proto.Request{}

	if stamp != nil {
		req.Solution = &proto.Solution{
			Stamp: *stamp,
		}
	}

	bytes, err := json.Marshal(req)
	bytes = append(bytes, '\n')
	if err != nil {
		return nil, err
	}

	_, err = conn.Write(bytes)
	if err != nil {
		return nil, err
	}

	reply := make([]byte, 2048)

	n, err := conn.Read(reply)
	if err != nil {
		return nil, err
	}

	rsp := proto.Response{}
	err = json.Unmarshal(reply[:n], &rsp)
	if err != nil {
		return nil, err
	}

	return &rsp, nil
}
