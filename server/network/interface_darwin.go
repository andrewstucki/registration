package network

// #cgo CFLAGS: -x objective-c
// #cgo LDFLAGS: -framework Foundation -framework CoreWLAN
// #import <Foundation/Foundation.h>
// #import <CoreWLAN/CoreWLAN.h>
// #include <arpa/inet.h>
// #include <ifaddrs.h>
// #include <netdb.h>
// #include <stdlib.h>
//
// char** arrayToStrings(NSArray * array) {
//   int count = [array count];
//   char **strings = (char **) malloc(sizeof(char *) * (count + 1));
//
//   int i;
//   for(i = 0; i < count; i++) {
//       NSString *s = [array objectAtIndex:i];
//       const char *cstr = [s cStringUsingEncoding:NSUTF8StringEncoding];
//       int len = strlen(cstr);
//       char *cstr_copy = (char *) malloc(sizeof(char) * (len + 1));
//       strcpy(cstr_copy, cstr);
//       strings[i] = cstr_copy;
//   }
//   strings[i] = NULL;
//
//   return strings;
// }
//
// char** enumerateInterfaces() {
//   char** strings;
//   @autoreleasepool {
//     NSArray *array = [[CWInterface interfaceNames] allObjects];
//     strings = arrayToStrings(array);
//   }
//   return strings;
// }
//
// CWInterface * internalGetInterface(char * interface) {
//   return [CWInterface interfaceWithName: [NSString stringWithUTF8String:interface]];
// }
//
// void disconnectInterface(CWInterface *interface) {
//     [interface disassociate];
// }
//
// bool checkInterfaceForIBSS(CWInterface *interface) {
//     if ([interface.configuration requireAdministratorForAssociation]){
//         return false;
//     }
//     if ([interface.configuration requireAdministratorForIBSSMode]){
//         return false;
//     }
//     return true;
// }
//
// char * waitForIpaddress(CWInterface *interface) {
//     bool haveInterface = true;
//     printf("Waiting for self-assigned IP (this could take awhile).");
//     fflush(stdout);
//     while (haveInterface) {
//         sleep(1);
//         printf(".");
//         fflush(stdout);
//         haveInterface = false;
//         struct ifaddrs* interfaces = NULL;
//         struct ifaddrs* cur_addr = NULL;
//         int res = getifaddrs(&interfaces);
//         if (res == 0) {
//             cur_addr = interfaces;
//             while(cur_addr != NULL){
//                 if (strcmp(cur_addr->ifa_name,[interface.interfaceName cStringUsingEncoding:NSUTF8StringEncoding])==0) {
//                     haveInterface = true;
//                     if (cur_addr->ifa_addr->sa_family == AF_INET) {
//                         char *addrBuf = (char *) malloc(sizeof(char) * NI_MAXHOST);
//                         getnameinfo(cur_addr->ifa_addr, sizeof(struct sockaddr_in), addrBuf, NI_MAXHOST, NULL, 0, NI_NUMERICHOST);
//                         freeifaddrs(interfaces);
//                         printf("\n");
//                         return addrBuf;
//                     }
//                 }
//                 cur_addr = cur_addr->ifa_next;
//             }
//         }
//         freeifaddrs(interfaces);
//     }
//     return NULL;
// }
//
// char * setupAdhocNetwork(char * interface, char * ssid, unsigned int securityMode, unsigned int channel, char * password) {
//   @autoreleasepool {
//     CWInterface *iface = internalGetInterface(interface);
//     if (checkInterfaceForIBSS(iface)) {
//       NSError* error;
//       BOOL success = [iface startIBSSModeWithSSID:[[NSString stringWithUTF8String:ssid] dataUsingEncoding:NSUTF8StringEncoding] security:securityMode channel:channel password:[NSString stringWithUTF8String:password] error:&error];
//       if (success) {
//         return waitForIpaddress(iface);
//       }
//     }
//     return NULL;
//   }
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
