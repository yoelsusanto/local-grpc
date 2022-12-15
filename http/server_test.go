package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"testing"
	"time"

	"github.com/yoelsusanto/local-grpc/utils"
)

var resultSimple *SimpleResponse

func BenchmarkSimpleHTTP(b *testing.B) {
	b.StopTimer()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	port := 50001 + rand.Int31n(100)

	utils.StartSubProcess(ctx, "main2", "http", fmt.Sprintf("%d", port))

	time.Sleep(2 * time.Second)

	serverAddr := fmt.Sprintf("http://127.0.0.1:%d", port)

	b.StartTimer()

	var simpleResponse *SimpleResponse

	for n := 0; n < b.N; n++ {
		sr := SimpleRequest{
			Text: "hello_world",
		}

		jsonBody, err := json.Marshal(sr)
		if err != nil {
			log.Fatal("Failed to marshal JSON", err)
		}
		bodyReader := bytes.NewReader(jsonBody)

		resp, err := http.Post(fmt.Sprintf("%s/simple", serverAddr), "application/json", bodyReader)
		if err != nil {
			log.Fatal("Failed to send request", err)
		}
		defer resp.Body.Close()

		err = json.NewDecoder(resp.Body).Decode(&simpleResponse)
		if err != nil {
			log.Fatal("Failed to unmarshal JSON", err)
		}

	}

	b.StopTimer()

	resultSimple = simpleResponse
	fmt.Println(simpleResponse)
}
