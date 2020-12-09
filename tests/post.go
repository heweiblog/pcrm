package main

//post方法发送json数据
func ClientPost(url string, data interface{}, contentType string) string {
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}

func testPost(request models.Content) {
	url := "http://127.0.0.1:22222/"
	time.Sleep(time.Second * 5)
	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(request)

	fmt.Println(requestBody)

	req, err := http.NewRequest("POST", url, requestBody)
	req.Header.Set("Content-Type", "application/json")

	fmt.Println(req)
	client := &http.Client{}
	fmt.Println(client)
	resp, err := client.Do(req)
	fmt.Println(resp)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
}
