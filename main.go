package main

import (
    "./src"
    "net/http"
    "encoding/json"
)

var blockchain *src.BlockChain

func main(){
    blockchain = src.NewBlockCHain()
    blockchain.SendData("Rose send 1 BTC to me")
    blockchain.SendData("I send 2 EOS to Rose")
    //blockchain.Print()

    http.HandleFunc("/get", blockchainGetHandler)
    http.HandleFunc("/write", blockchainWriteHandler)
    if http.ListenAndServe(":8081", nil) != nil {
        return
    }
}

func blockchainGetHandler(w http.ResponseWriter, r *http.Request){

    bytes,err := json.Marshal(blockchain)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Write(bytes)
    //io.WriteString(w, string(bytes))
}

func blockchainWriteHandler(w http.ResponseWriter, r *http.Request){
    data := r.URL.Query().Get("data")
    blockchain.SendData(data)
    blockchainGetHandler(w, r)
}