package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	var codeTableFile string = "code.table"
	srcData, err := os.ReadFile(codeTableFile); if err != nil { fmt.Println("[ERROR]",err); return }

	var codeTable map[string]string
	err_2 := json.Unmarshal(srcData, &codeTable); if err_2 != nil { fmt.Println("[ERROR]",err_2); return }

	switch os.Args[1] {

	case "encrypt":
		encrypt(codeTable)

	case "decrypt":
		decrypt(codeTable)

	default :
		fmt.Printf("Unknow command \"%s\"\n",os.Args[1])
	}


}

func encrypt(codeTable map[string]string) {

	var filename string = os.Args[2]
	text, err := os.ReadFile(filename); if err != nil { fmt.Println("[ERROR]",err); return }

	var codeList []string
	for _, index := range(text) {

		code := codeTable[string(index)]

		if code == "" {
			fmt.Printf("Can't found char code of %q.\n",index)
			return
		}

		codeList = append(codeList, code)

	}

	data, err := json.Marshal(codeList); if err != nil { fmt.Println("[ERROR]",err); return }
	os.WriteFile(filename + ".enc", data, 0666)


}

func decrypt(codeTable map[string]string) {

	var filename string = os.Args[2]
	code, err := os.ReadFile(filename); if err != nil { fmt.Println("[ERROR]",err); return }

	var codeList []string
	err_2 := json.Unmarshal(code, &codeList); if err_2 != nil { fmt.Println("[ERROR]",err_2); return }

	decodeTable := make(map[string]string)
	for k, v := range(codeTable) {
		decodeTable[v] = k
	}
	
	var charList []string
	for _, index := range(codeList) {
		
		char := decodeTable[index]

		if char == "" {
			fmt.Printf("Can't found char code of \"%q\".\n",index)
			return
		}

		charList = append(charList,char)

	}

	text := ""

	for _, index := range(charList) {
		text += index
	}

	os.WriteFile(filename + ".txt", []byte(text),0666)

}