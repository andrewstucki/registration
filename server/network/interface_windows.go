package network

// #cgo LDFLAGS: -lwlanapi -lole32
// #define UNICODE
// #include <windows.h>
// #include <wlanapi.h>
// #include <objbase.h>
// #include <wtypes.h>
// #include <stdio.h>
// #include <stdlib.h>
//
// char** enumerateInterfaces() {
//   HANDLE hClient = NULL;
//   DWORD dwMaxClient = 2;
//   DWORD dwCurVersion = 0;
//   WCHAR GuidString[39] = {0};
//   unsigned int i;
//   PWLAN_INTERFACE_INFO_LIST pIfList = NULL;
//   PWLAN_INTERFACE_INFO pIfInfo = NULL;
//   WlanOpenHandle(dwMaxClient, NULL, &dwCurVersion, &hClient);
//   WlanEnumInterfaces(hClient, NULL, &pIfList);
//   char **interfaces = (char **)malloc(sizeof(char*) * pIfList->dwNumberOfItems);
//   for (i = 0; i < (int) pIfList->dwNumberOfItems; i++) {
//     pIfInfo = (WLAN_INTERFACE_INFO *) &pIfList->InterfaceInfo[i];
//     int size = sizeof(GuidString)/sizeof(*GuidString);
//     StringFromGUID2(&pIfInfo->InterfaceGuid, (LPOLESTR) &GuidString, size);
//     interfaces[i] = (char *)malloc(size+1);
//     wcstombs_s(NULL, interfaces[i], size+1, GuidString, size+1);
//   }
//   return interfaces;
// }
//
// char * setupAdhocNetwork(char * iface, char * ssid, unsigned int securityMode, unsigned int channel, char * password) {
//   return "localhost";
// }
import "C"

import (
	"errors"
	"unsafe"
)

func getInterfaces() []string {
	return convertStringsToSlice(C.enumerateInterfaces())
}

func makeNetwork(iface string, ssid string, security uint, channel uint, pass string) (ip string, err error) {
	cSecurity := C.uint(security)
	cChannel := C.uint(channel)
	cIface := C.CString(iface)
	cSsid := C.CString(ssid)
	cPass := C.CString(pass)
	defer C.free(unsafe.Pointer(cIface))
	defer C.free(unsafe.Pointer(cSsid))
	defer C.free(unsafe.Pointer(cPass))

	cIP := C.setupAdhocNetwork(cIface, cSsid, cSecurity, cChannel, cPass)
	if cIP == nil {
		ip = ""
		err = errors.New("Could not make ad-hoc network")
	} else {
		defer C.free(unsafe.Pointer(cIP))
		ip = C.GoString(cIP)
	}
	return
}
