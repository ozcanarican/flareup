package cloudflare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var apiKey string = ""

func Prepare() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No .env file")
	}
	apiKey = os.Getenv("API_KEY")
	fmt.Println(apiKey)

}

func apiURL(s string) string {
	return fmt.Sprintf("https://api.cloudflare.com/client/v4/%v", s)
}

func AddRecord(subdomain string, domain string, ip string, forceful bool) {
	zone := FindZoneByName(domain)
	records := GetRecordsByZone(zone)
	isExist := false
	var found RecordStruct
	for _, v := range records.Result {
		fmt.Println(v.Name)
		if v.Name == subdomain {
			found = v
			isExist = true
			break
		}
	}
	fmt.Println(isExist)
	if isExist {
		if forceful {

			deleteRecord(zone, found)
		} else {
			log.Fatalln("This record is already exists.\nUse -f flag to replace it")
		}
	}
	CreateRecord(zone, subdomain, ip)
}

func RemoveRecord(subdomain string, domain string) {
	zone := FindZoneByName(domain)
	records := GetRecordsByZone(zone)
	for _, v := range records.Result {
		if v.Name == subdomain {
			deleteRecord(zone, v)
			break
		}
	}
}

func deleteRecord(zone ZoneStruct, record RecordStruct) {
	url := apiURL(fmt.Sprintf(`zones/%v/dns_records/%v`, zone.ID, record.ID))
	client := &http.Client{}
	Prepare()
	req, err := http.NewRequest("DELETE", url, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", apiKey))
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(body))

}

func FindZoneByName(z string) ZoneStruct {
	Prepare()
	var found ZoneStruct

	client := &http.Client{}
	req, err := http.NewRequest("GET", apiURL(`zones`), nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", apiKey))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	var result ZoneListResult
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Cant convert to json")
	}
	for i, v := range result.Result {
		fmt.Printf("%v. %v\n", i, v.Name)
		if z == v.Name {
			found = v
		}
	}
	return found
}

func GetRecordsByZone(z ZoneStruct) RecordResponse {
	Prepare()
	url := apiURL(fmt.Sprintf("zones/%v/dns_records", z.ID))
	fmt.Println(url)

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", apiKey))
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(resp)
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	//fmt.Println(string(body))
	//fmt.Println("-------------------------")
	var result RecordResponse
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Cant convert to json")
	}
	return result
}

func CreateRecord(zone ZoneStruct, domain string, ip string) {
	Prepare()
	url := apiURL(fmt.Sprintf("zones/%v/dns_records", zone.ID))
	fmt.Println(url)

	fields := []byte(fmt.Sprintf(`{
	"type": "A",
	"name": "%v",
	"content": "%v",
	"proxied": false,
	"ttl": 1,
	"comment": "created by flareup"
	}`, domain, ip))

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(fields))
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %v", apiKey))
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	fmt.Println(string(body))

}
