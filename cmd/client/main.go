package main

import (
	"bufio"
	"context"
	"ditto/cookies"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	client := cookies.NewCookieTravelerClient(conn)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter a cookie flavor to send: ")
		flavor, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Oh no!... %s", err.Error())
			return
		}

		c := cookies.Cookie{
			Flavor: strings.TrimSpace(flavor),
		}
		cookie, err := client.SendCookie(context.Background(), &c)
		if err != nil {
			log.Println(err)
			log.Println("no cookie received...")
		} else {
			log.Printf("you received a %s cookie!", cookie.Flavor)
		}
	}
}
