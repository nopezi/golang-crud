package lib

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ApiDefinition interface {
	Auth()
	ConsumeApi()
	Goresums()
}

type (
	Options struct {
		BaseUrl  string
		SSL      bool
		Payload  interface{}
		Request  *bytes.Buffer
		Method   string
		Header   []Header
		Auth     bool
		ClientID ClientID
	}
	ClientID struct {
		Clientid     string
		Clientsecret string
	}
	Header struct {
		Key   string
		Value string
	}
	Resp map[string]interface{}

	Resp2 struct {
		Success bool
		Message interface{}
	}
	Auth struct {
		Authorization string
		Type          string
	}
)

func AuthBearer(options Options, auth Auth) (response Resp, err error) {
	fmt.Println("Goresums | options => BaseUrl  string | SSL bool | Payload interface{} | Request  *bytes.Buffer | Method string | Header []Header | Auth bool | ClientID ClientID")
	fmt.Println("Goresums | options => ", options)
	resByte, _ := json.Marshal(options.Payload)

	options.Request = bytes.NewBuffer(resByte)
	fmt.Println("Goresums | request byte[] to NewBuffer =>", options.Request)

	transCfg := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: options.SSL}, // ignore expired SSL certificates
	}

	client := &http.Client{Transport: transCfg}
	req, err := http.NewRequest(
		options.Method,
		options.BaseUrl,
		options.Request)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Content-Type", "application/json")

	newAuth := Auth{
		Authorization: auth.Authorization,
	}

	req.Header.Add("Authorization", newAuth.Authorization)
	// if options.Auth {
	// }

	for _, header := range options.Header {
		req.Header.Add(header.Key, header.Value)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	// body := make(Response)
	// if json.NewDecoder(res.Body).Decode(&body); err != nil {
	// 	fmt.Println(err.Error())
	// }

	fmt.Println("Goresums | response res.Body =>", res.Body)
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	var responseObject Resp
	json.Unmarshal(bodyBytes, &responseObject)

	fmt.Println("Goresums | responseObject after json Unmarshal res.Body =>", responseObject)

	return responseObject, err
}

func BasicAuth(options Options, auth Auth) (response Resp, err error) {
	return nil, err
}

func OauthAuth(options Options, auth Auth) (response Resp, err error) {
	return nil, err
}

func Goresums(options Options, auth Auth) (response Resp, err error) {
	switch auth.Type {
	case "Bearer":
		AuthBearer(options, auth)
		return
	case "Basic":
		BasicAuth(options, auth)
		return
	case "Oauth2":
		OauthAuth(options, auth)
		return
	default:
		return nil, err
	}
}
