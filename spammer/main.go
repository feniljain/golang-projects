package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//for i := 0; i != -1; i++ {
	//	_, err := http.PostForm(
	//		"https://ccs.ieeecsvit.com/login",
	//		url.Values{
	//			"regno":    {"20BCE0562"},
	//			"password": {"Something@1234"},
	//		},
	//	)

	//	if err != nil {
	//		fmt.Println("An error occurred: ", err)
	//	}

	//	fmt.Println("Successful response: ", i)
	//}

	for i := 0; i < 1; i++ {
		regNo := "20BCE"
		if i < 10 {
			regNo += "000" + fmt.Sprintf("%02d", i)
			//fmt.Sprintf("%0002d", i)
		} else if i < 100 {
			//regNo += "00" + string(i)
			regNo += "00" + fmt.Sprintf("%02d", i)
		} else if i < 1000 {
			//regNo += "0"
			regNo += "0" + fmt.Sprintf("%02d", i)
		} else if i < 3000 {
			regNo += fmt.Sprintf("%02d", i)
		} else {
			break
		}

		//fmt.Println("Making request to: " + "20BCE" + fmt.Sprintf("%002d", i))
		url := "https://acmvit-recruitment-2020-bleeding-edge.azurewebsites.net/users"
		fmt.Println("Making request on : " + regNo)

		var jsonStr = []byte(`{ "links":",,,", "name": "Somone1 Shah", "password": "Something@1234", "phoneNumber": "9874561230", "registrationNumber": "20BCE0436",  "token": "03AGdBq25l5vhC1HCOcNJKg1MSb4p6CGyTOqzhRH7RGmJHFRKfWmA3GjSm362c1yWSOcK0cFZyiSyxbxp3NF3keM6xFbBATl9N1NXYk-_7raSogO5lIjLWcUMOwXrMBJIsEKqF_LBisgvETxNIhwZaKrSLnJlsSeRLlQV2aZd_MeBnEPMQ2fKSqoKMT0frb7_Sz0-bdhGcY0Ws4Iapguu1ILue7A6wUjkhCRcApfypqgKAwW5DJjjOnwiWQ9K6IxjoiBgFnDYdOR6NEO5YuCE8L09gjdz3ttmyVHkmz24yH020U6Al7yq5xOtry8y0evANgK2mCdWpxEo8ARBKwJOy4M-61OKHQC2XEPA1IGCw8NpgIiXlhPB_6FMgvC47h-OLO-0Ix2AqKXN8mbk49AvCoca-kzZvkFl95BGtWpZD_4FpOWgSwPq9EI31mD-uP5-kS8n-ondf_dXk1jgsU4eEzY1Diuvg-9glkvWdlnANlqMZ-DV2CdeKbrQ", "vitEmail": "someone1.shah2020@vitstudent.ac.in" }`)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("An error occurred: ", err)
		}
		defer resp.Body.Close()

		fmt.Println("Response Status", resp.Status)
		fmt.Println("Response Status Code", resp.StatusCode)
		fmt.Println("Response Headers", resp.Header)
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("response Body: ", string(body))

		//resp, err := http.PostForm(
		//	"https://acmvit-recruitment-2020-bleeding-edge.azurewebsites.net/users",
		//	url.Values{
		//		"hidden":               {""},
		//		"name":                 {string(i)},
		//		"regno":                {regNo},
		//		"email":                {string(i) + "@gmail.com"},
		//		"Password":             {"Something@1234"},
		//		"phone":                {"9874563210"},
		//		"gender":               {"male"},
		//		"g-recaptcha-response": {"03AGdBq25l5vhC1HCOcNJKg1MSb4p6CGyTOqzhRH7RGmJHFRKfWmA3GjSm362c1yWSOcK0cFZyiSyxbxp3NF3keM6xFbBATl9N1NXYk-_7raSogO5lIjLWcUMOwXrMBJIsEKqF_LBisgvETxNIhwZaKrSLnJlsSeRLlQV2aZd_MeBnEPMQ2fKSqoKMT0frb7_Sz0-bdhGcY0Ws4Iapguu1ILue7A6wUjkhCRcApfypqgKAwW5DJjjOnwiWQ9K6IxjoiBgFnDYdOR6NEO5YuCE8L09gjdz3ttmyVHkmz24yH020U6Al7yq5xOtry8y0evANgK2mCdWpxEo8ARBKwJOy4M-61OKHQC2XEPA1IGCw8NpgIiXlhPB_6FMgvC47h-OLO-0Ix2AqKXN8mbk49AvCoca-kzZvkFl95BGtWpZD_4FpOWgSwPq9EI31mD-uP5-kS8n-ondf_dXk1jgsU4eEzY1Diuvg-9glkvWdlnANlqMZ-DV2CdeKbrQ"},
		//	},
		//)
		//if err != nil {
		//	fmt.Println("An error occurred: ", err)
		//}

		//if resp.StatusCode != 200 {
		//	fmt.Println("Gone wrong response: ", i, "status code: ", resp.StatusCode)
		//} else {
		//	fmt.Println("Successful response: ", i, "status code: ", resp.StatusCode)
		//	_, err := ioutil.ReadAll(resp.Body)
		//	if err != nil {
		//		fmt.Println("Error decoding body")
		//	}
		//	//fmt.Println(string(byts))
		//}
	}
}
