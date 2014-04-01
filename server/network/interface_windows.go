package network

// #cgo CFLAGS: -I../../include
// #cgo LDFLAGS: -lole32 -L../../lib -lwlanapi
// #define UNICODE
// #define  WLAN_NOTIFICATION_SOURCE_ALL 0X0000FFFF
// #include <windows.h>
// #include <wlanapi.h>
// #include <objbase.h>
// #include <wtypes.h>
// #include <stdio.h>
// #include <stdlib.h>
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
// void onNotification(PWLAN_NOTIFICATION_DATA data, PVOID context) {
//   printf("%x, %x, %s", data->NotificationSource, data->NotificationCode, data->InterfaceGuid);
// }
//
// char * setupAdhocNetwork(char * iface, char * ssid, unsigned int securityMode, unsigned int channel, char * password) {
//   HANDLE hClient = NULL;
//   DWORD dwMaxClient = 2;
//   DWORD dwCurVersion = 0;
//   PDWORD source = NULL;
//   PDWORD rc = NULL;
//   // WLAN_HOSTED_NETWORK_STATUS status = {0};
//   ULONG ssidLength = strlen(ssid);
//   int enabled = 1;
//   DOT11_SSID ssidStruct;
//   memset(&ssidStruct, 0, sizeof(ssidStruct));
//   ssidStruct.uSSIDLength = ssidLength;
//   memcpy(ssidStruct.ucSSID, ssid, ssidLength);
//
//   WLAN_HOSTED_NETWORK_CONNECTION_SETTINGS settings = { ssidStruct, 100 };
//   WlanOpenHandle(dwMaxClient, NULL, &dwCurVersion, &hClient);
//   WlanRegisterNotification(hClient, WLAN_NOTIFICATION_SOURCE_ALL, 1, &onNotification, NULL, NULL, source);
//   WlanHostedNetworkInitSettings(hClient, rc, NULL);
//   WlanHostedNetworkSetProperty(hClient, wlan_hosted_network_opcode_enable, sizeof(enabled), &enabled, rc, NULL);
//   WlanHostedNetworkSetProperty(hClient, wlan_hosted_network_opcode_connection_settings, sizeof(settings), &settings, rc, NULL);
//   WlanHostedNetworkSetSecondaryKey(hClient, strlen(password)+1, password, 1, 1, rc, NULL);
//   WlanHostedNetworkStartUsing(hClient, rc, NULL);
//   // WlanHostedNetworkQueryStatus(hClient, &status, NULL);
//
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
