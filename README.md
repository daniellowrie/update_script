# update_script

!!! FOR SECURITY TESTING PURPOSES ONLY !!! <br>
\******************************************************<br>
DO NOT USE THIS ON SYSTEMS THAT YOU DO NOT <br>
OWN UNLESS YOU HAVE EXPRESSED PERMISSION !

Fileless malware that bypasses Windows Defender using PowerShell and obfuscation. <br>
Just a simple Reverse Shell using a batch script to kick things off, after which everything is fileless.

YouTube video demonstration and explanation >>> https://youtu.be/BFVzmZXIbQk

[!] The `WinSecurityUpdate` file needs to be edited BY YOU to work. <br>
- You will need to decode the *base64* strings that request the `a1` and `r1` files. <br>
- Then you will need to modify the IP addresses in the `a1` and `r1` files to your own and re-encode with *base64*. <br>
- Replace the old `a1` and `r1` strings with your modified version and save.
 

Going to work on building a better mousetrap as this was just a PoC at making something 'malicious' that could bypass Defender.
<P></P>



https://user-images.githubusercontent.com/19762230/132001237-04ae01ae-102e-4070-9599-ad05128c3f73.mp4


