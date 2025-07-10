package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// registerInEurekaServer: sent to discovery server, the api address
func registerInEurekaServer() {
	eurekaURL := os.Getenv("EUREKA_SERVER") + "/apps/USERS-API"

	payload := `
	<instance>
		<hostName>users-api-container</hostName>
		<instanceId>users-api</instanceId>
		<app>USERS-API</app>
		<ipAddr>users-api-container</ipAddr>
		<status>UP</status>
		<port enabled="true">8080</port>

		<homePageUrl>http://users-api-container:8080/</homePageUrl>
		<statusPageUrl>http://users-api-container:8080/actuator/health</statusPageUrl>
		<healthCheckUrl>http://users-api-container:8080/actuator/health</healthCheckUrl>

		<dataCenterInfo class="com.netflix.appinfo.InstanceInfo$DefaultDataCenterInfo">
			<name>MyOwn</name>
		</dataCenterInfo>
	</instance>`

	req, err := http.NewRequest("POST", eurekaURL, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/xml")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("Registry status: " + resp.Status)
}

// __sendHeartBeat: keep communication with the discovery server
func __sendHeartBeat() {
	url := os.Getenv("EUREKA_SERVER") + "/apps/USERS-API/users-api"
	req, _ := http.NewRequest("PUT", url, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Fail in heartBeat to eureka server")
		return
	}
	defer resp.Body.Close()
	fmt.Println("Heartbeat sended!", resp.Status)
}

func heartBeat() {
	ticker := time.NewTicker(30 * time.Second)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-ticker.C:
			__sendHeartBeat()
		case <-quit:
			ticker.Stop()
			fmt.Println("Exiting from heartbeat")
			os.Exit(0)
		}
	}
}
