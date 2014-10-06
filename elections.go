package main

import (
  "os"
  "fmt"
  "time"
  "strconv"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

func percent(x interface{}, y interface{}) float64 {
  fx, _ := strconv.ParseFloat(x.(string), 32)
  fy, _ := strconv.ParseFloat(y.(string), 32)
  return fy / fx * 100.0
}

func work() {
  resp, err := http.Get("http://divulga.tse.jus.br/2014/divulgacao/oficial/143/dadosdivweb/br/br-0001-e001431-w.js")
  if err != nil || resp.StatusCode != 200 {
    fmt.Println("Something went wrong")
    os.Exit(1)
  }
  defer resp.Body.Close()
  data := make(map[string]interface{})
  body, _ := ioutil.ReadAll(resp.Body)
  err = json.Unmarshal(body, &data)
  fmt.Println("\n\nÚltima apuração em:", data["ht"] )
  fmt.Printf("Votos apurados: %.5f%% \n", percent(data["e"], data["ea"]))

  candidates, _ := data["cand"].([]interface{})

  for i:= 0; i < 3; i++ {
    candidate := candidates[i].(map[string]interface{})
    number, name, percentage := candidate["n"].(string), candidate["nm"].(string), percent(data["vv"], candidate["v"])

    fmt.Printf("\n[%v] %s - %.2f%%", number, name, percentage)
  }
}

func main() {
  for {
    work()
    time.Sleep(60 * time.Second)
  }
}
