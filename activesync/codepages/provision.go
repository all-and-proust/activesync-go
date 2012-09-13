package codepages

import (
	. "github.com/magicmonty/wbxml-go/wbxml"
)

const (
	NS_PROVISION                                           string = "Provision"
	TAG_PROVISION_PROVISION                                string = "Provision"
	TAG_PROVISION_POLICIES                                 string = "Policies"
	TAG_PROVISION_POLICY                                   string = "Policy"
	TAG_PROVISION_POLICYTYPE                               string = "PolicyType"
	TAG_PROVISION_POLICYKEY                                string = "PolicyKey"
	TAG_PROVISION_DATA                                     string = "Data"
	TAG_PROVISION_STATUS                                   string = "Status"
	TAG_PROVISION_REMOTEWIPE                               string = "RemoteWipe"
	TAG_PROVISION_EASPROVISIONDOC                          string = "EASProvisionDoc"
	TAG_PROVISION_DEVICEPASSWORDENABLED                    string = "DevicePasswordEnabled"
	TAG_PROVISION_ALPHANUMERICDEVICEPASSWORDREQUIRED       string = "AlphanumericDevicePasswordRequired"
	TAG_PROVISION_DEVICEENCRYPTIONENABLED                  string = "DeviceEncryptionEnabled"
	TAG_PROVISION_REQUIRESTORAGECARDENCRYPTION             string = "RequireStorageCardEncryption"
	TAG_PROVISION_PASSWORDRECOVERYENABLED                  string = "PasswordRecoveryEnabled"
	TAG_PROVISION_ATTACHMENTSENABLED                       string = "AttachmentsEnabled"
	TAG_PROVISION_MINDEVICEPASSWORDLENGTH                  string = "MinDevicePasswordLength"
	TAG_PROVISION_MAXINACTIVITYTIMEDEVICELOCK              string = "MaxInactivityTimeDeviceLock"
	TAG_PROVISION_MAXDEVICEPASSWORDFAILEDATTEMPTS          string = "MaxDevicePasswordFailedAttempts"
	TAG_PROVISION_MAXATTACHMENTSIZE                        string = "MaxAttachmentSize"
	TAG_PROVISION_ALLOWSIMPLEDEVICEPASSWORD                string = "AllowSimpleDevicePassword"
	TAG_PROVISION_DEVICEPASSWORDEXPIRATION                 string = "DevicePasswordExpiration"
	TAG_PROVISION_DEVICEPASSWORDHISTORY                    string = "DevicePasswordHistory"
	TAG_PROVISION_ALLOWSTORAGECARD                         string = "AllowStorageCard"
	TAG_PROVISION_ALLOWCAMERA                              string = "AllowCamera"
	TAG_PROVISION_REQUIREDEVICEENCRYPTION                  string = "RequireDeviceEncryption"
	TAG_PROVISION_ALLOWUNSIGNEDAPPLICATIONS                string = "AllowUnsignedApplications"
	TAG_PROVISION_ALLOWUNSIGNEDINSTALLATIONPACKAGES        string = "AllowUnsignedInstallationPackages"
	TAG_PROVISION_MINDEVICEPASSWORDCOMPLEXCHARACTERS       string = "MinDevicePasswordComplexCharacters"
	TAG_PROVISION_ALLOWWIFI                                string = "AllowWiFi"
	TAG_PROVISION_ALLOWTEXTMESSAGING                       string = "AllowTextMessaging"
	TAG_PROVISION_ALLOWPOPIMAPEMAIL                        string = "AllowPOPIMAPEmail"
	TAG_PROVISION_ALLOWBLUETOOTH                           string = "AllowBluetooth"
	TAG_PROVISION_ALLOWIRDA                                string = "AllowIrDA"
	TAG_PROVISION_REQUIREMANUALSYNCWHENROAMING             string = "RequireManualSyncWhenRoaming"
	TAG_PROVISION_ALLOWDESKTOPSYNC                         string = "AllowDesktopSync"
	TAG_PROVISION_MAXCALENDARAGEFILTER                     string = "MaxCalendarAgeFilter"
	TAG_PROVISION_ALLOWHTMLEMAIL                           string = "AllowHTMLEmail"
	TAG_PROVISION_MAXEMAILAGEFILTER                        string = "MaxEmailAgeFilter"
	TAG_PROVISION_MAXEMAILBODYTRUNCATIONSIZE               string = "MaxEmailBodyTruncationSize"
	TAG_PROVISION_MAXEMAILHTMLBODYTRUNCATIONSIZE           string = "MaxEmailHTMLBodyTruncationSize"
	TAG_PROVISION_REQUIRESIGNEDSMIMEMESSAGES               string = "RequireSignedSMIMEMessages"
	TAG_PROVISION_REQUIREENCRYPTEDSMIMEMESSAGES            string = "RequireEncryptedSMIMEMessages"
	TAG_PROVISION_REQUIRESIGNEDSMIMEALGORITHM              string = "RequireSignedSMIMEAlgorithm"
	TAG_PROVISION_REQUIREENCRYPTIONSMIMEALGORITHM          string = "RequireEncryptionSMIMEAlgorithm"
	TAG_PROVISION_ALLOWSMIMEENCRYPTIONALGORITHMNEGOTIATION string = "AllowSMIMEEncryptionAlgorithmNegotiation"
	TAG_PROVISION_ALLOWSMIMESOFTCERTS                      string = "AllowSMIMESoftCerts"
	TAG_PROVISION_ALLOWBROWSER                             string = "AllowBrowser"
	TAG_PROVISION_ALLOWCONSUMEREMAIL                       string = "AllowConsumerEmail"
	TAG_PROVISION_ALLOWREMOTEDESKTOP                       string = "AllowRemoteDesktop"
	TAG_PROVISION_ALLOWINTERNETSHARING                     string = "AllowInternetSharing"
	TAG_PROVISION_UNAPPROVEDINROMAPPLICATIONLIST           string = "UnapprovedInROMApplicationList"
	TAG_PROVISION_APPLICATIONNAME                          string = "ApplicationName"
	TAG_PROVISION_APPROVEDAPPLICATIONLIST                  string = "ApprovedApplicationList"
	TAG_PROVISION_HASH                                     string = "Hash"
)

const (
	CP_PROVISION                                          byte = 14
	ID_PROVISION_PROVISION                                byte = 0x05
	ID_PROVISION_POLICIES                                 byte = 0x06
	ID_PROVISION_POLICY                                   byte = 0x07
	ID_PROVISION_POLICYTYPE                               byte = 0x08
	ID_PROVISION_POLICYKEY                                byte = 0x09
	ID_PROVISION_DATA                                     byte = 0x0A
	ID_PROVISION_STATUS                                   byte = 0x0B
	ID_PROVISION_REMOTEWIPE                               byte = 0x0C
	ID_PROVISION_EASPROVISIONDOC                          byte = 0x0D
	ID_PROVISION_DEVICEPASSWORDENABLED                    byte = 0x0E
	ID_PROVISION_ALPHANUMERICDEVICEPASSWORDREQUIRED       byte = 0x0F
	ID_PROVISION_DEVICEENCRYPTIONENABLED                  byte = 0x10
	ID_PROVISION_REQUIRESTORAGECARDENCRYPTION             byte = 0x10
	ID_PROVISION_PASSWORDRECOVERYENABLED                  byte = 0x11
	ID_PROVISION_ATTACHMENTSENABLED                       byte = 0x13
	ID_PROVISION_MINDEVICEPASSWORDLENGTH                  byte = 0x14
	ID_PROVISION_MAXINACTIVITYTIMEDEVICELOCK              byte = 0x15
	ID_PROVISION_MAXDEVICEPASSWORDFAILEDATTEMPTS          byte = 0x16
	ID_PROVISION_MAXATTACHMENTSIZE                        byte = 0x17
	ID_PROVISION_ALLOWSIMPLEDEVICEPASSWORD                byte = 0x18
	ID_PROVISION_DEVICEPASSWORDEXPIRATION                 byte = 0x19
	ID_PROVISION_DEVICEPASSWORDHISTORY                    byte = 0x1A
	ID_PROVISION_ALLOWSTORAGECARD                         byte = 0x1B
	ID_PROVISION_ALLOWCAMERA                              byte = 0x1C
	ID_PROVISION_REQUIREDEVICEENCRYPTION                  byte = 0x1D
	ID_PROVISION_ALLOWUNSIGNEDAPPLICATIONS                byte = 0x1E
	ID_PROVISION_ALLOWUNSIGNEDINSTALLATIONPACKAGES        byte = 0x1F
	ID_PROVISION_MINDEVICEPASSWORDCOMPLEXCHARACTERS       byte = 0x20
	ID_PROVISION_ALLOWWIFI                                byte = 0x21
	ID_PROVISION_ALLOWTEXTMESSAGING                       byte = 0x22
	ID_PROVISION_ALLOWPOPIMAPEMAIL                        byte = 0x23
	ID_PROVISION_ALLOWBLUETOOTH                           byte = 0x24
	ID_PROVISION_ALLOWIRDA                                byte = 0x25
	ID_PROVISION_REQUIREMANUALSYNCWHENROAMING             byte = 0x26
	ID_PROVISION_ALLOWDESKTOPSYNC                         byte = 0x27
	ID_PROVISION_MAXCALENDARAGEFILTER                     byte = 0x28
	ID_PROVISION_ALLOWHTMLEMAIL                           byte = 0x29
	ID_PROVISION_MAXEMAILAGEFILTER                        byte = 0x2A
	ID_PROVISION_MAXEMAILBODYTRUNCATIONSIZE               byte = 0x2B
	ID_PROVISION_MAXEMAILHTMLBODYTRUNCATIONSIZE           byte = 0x2C
	ID_PROVISION_REQUIRESIGNEDSMIMEMESSAGES               byte = 0x2D
	ID_PROVISION_REQUIREENCRYPTEDSMIMEMESSAGES            byte = 0x2E
	ID_PROVISION_REQUIRESIGNEDSMIMEALGORITHM              byte = 0x2F
	ID_PROVISION_REQUIREENCRYPTIONSMIMEALGORITHM          byte = 0x30
	ID_PROVISION_ALLOWSMIMEENCRYPTIONALGORITHMNEGOTIATION byte = 0x31
	ID_PROVISION_ALLOWSMIMESOFTCERTS                      byte = 0x32
	ID_PROVISION_ALLOWBROWSER                             byte = 0x33
	ID_PROVISION_ALLOWCONSUMEREMAIL                       byte = 0x34
	ID_PROVISION_ALLOWREMOTEDESKTOP                       byte = 0x35
	ID_PROVISION_ALLOWINTERNETSHARING                     byte = 0x36
	ID_PROVISION_UNAPPROVEDINROMAPPLICATIONLIST           byte = 0x37
	ID_PROVISION_APPLICATIONNAME                          byte = 0x38
	ID_PROVISION_APPROVEDAPPLICATIONLIST                  byte = 0x39
	ID_PROVISION_HASH                                     byte = 0x3A
)

func Provision() CodePage {
	result := NewCodePage(NS_PROVISION, CP_PROVISION)

	result.AddTag(TAG_PROVISION_PROVISION, ID_PROVISION_PROVISION)
	result.AddTag(TAG_PROVISION_POLICIES, ID_PROVISION_POLICIES)
	result.AddTag(TAG_PROVISION_POLICY, ID_PROVISION_POLICY)
	result.AddTag(TAG_PROVISION_POLICYTYPE, ID_PROVISION_POLICYTYPE)
	result.AddTag(TAG_PROVISION_POLICYKEY, ID_PROVISION_POLICYKEY)
	result.AddTag(TAG_PROVISION_DATA, ID_PROVISION_DATA)
	result.AddTag(TAG_PROVISION_STATUS, ID_PROVISION_STATUS)
	result.AddTag(TAG_PROVISION_REMOTEWIPE, ID_PROVISION_REMOTEWIPE)
	result.AddTag(TAG_PROVISION_EASPROVISIONDOC, ID_PROVISION_EASPROVISIONDOC)
	result.AddTag(TAG_PROVISION_DEVICEPASSWORDENABLED, ID_PROVISION_DEVICEPASSWORDENABLED)
	result.AddTag(TAG_PROVISION_ALPHANUMERICDEVICEPASSWORDREQUIRED, ID_PROVISION_ALPHANUMERICDEVICEPASSWORDREQUIRED)
	result.AddTag(TAG_PROVISION_DEVICEENCRYPTIONENABLED, ID_PROVISION_DEVICEENCRYPTIONENABLED)
	result.AddTag(TAG_PROVISION_REQUIRESTORAGECARDENCRYPTION, ID_PROVISION_REQUIRESTORAGECARDENCRYPTION) // equivalent to DeviceEncryptionEnabled
	result.AddTag(TAG_PROVISION_PASSWORDRECOVERYENABLED, ID_PROVISION_PASSWORDRECOVERYENABLED)
	result.AddTag(TAG_PROVISION_ATTACHMENTSENABLED, ID_PROVISION_ATTACHMENTSENABLED)
	result.AddTag(TAG_PROVISION_MINDEVICEPASSWORDLENGTH, ID_PROVISION_MINDEVICEPASSWORDLENGTH)
	result.AddTag(TAG_PROVISION_MAXINACTIVITYTIMEDEVICELOCK, ID_PROVISION_MAXINACTIVITYTIMEDEVICELOCK)
	result.AddTag(TAG_PROVISION_MAXDEVICEPASSWORDFAILEDATTEMPTS, ID_PROVISION_MAXDEVICEPASSWORDFAILEDATTEMPTS)
	result.AddTag(TAG_PROVISION_MAXATTACHMENTSIZE, ID_PROVISION_MAXATTACHMENTSIZE)
	result.AddTag(TAG_PROVISION_ALLOWSIMPLEDEVICEPASSWORD, ID_PROVISION_ALLOWSIMPLEDEVICEPASSWORD)
	result.AddTag(TAG_PROVISION_DEVICEPASSWORDEXPIRATION, ID_PROVISION_DEVICEPASSWORDEXPIRATION)
	result.AddTag(TAG_PROVISION_DEVICEPASSWORDHISTORY, ID_PROVISION_DEVICEPASSWORDHISTORY)
	result.AddTag(TAG_PROVISION_ALLOWSTORAGECARD, ID_PROVISION_ALLOWSTORAGECARD)
	result.AddTag(TAG_PROVISION_ALLOWCAMERA, ID_PROVISION_ALLOWCAMERA)
	result.AddTag(TAG_PROVISION_REQUIREDEVICEENCRYPTION, ID_PROVISION_REQUIREDEVICEENCRYPTION)
	result.AddTag(TAG_PROVISION_ALLOWUNSIGNEDAPPLICATIONS, ID_PROVISION_ALLOWUNSIGNEDAPPLICATIONS)
	result.AddTag(TAG_PROVISION_ALLOWUNSIGNEDINSTALLATIONPACKAGES, ID_PROVISION_ALLOWUNSIGNEDINSTALLATIONPACKAGES)
	result.AddTag(TAG_PROVISION_MINDEVICEPASSWORDCOMPLEXCHARACTERS, ID_PROVISION_MINDEVICEPASSWORDCOMPLEXCHARACTERS)
	result.AddTag(TAG_PROVISION_ALLOWWIFI, ID_PROVISION_ALLOWWIFI)
	result.AddTag(TAG_PROVISION_ALLOWTEXTMESSAGING, ID_PROVISION_ALLOWTEXTMESSAGING)
	result.AddTag(TAG_PROVISION_ALLOWPOPIMAPEMAIL, ID_PROVISION_ALLOWPOPIMAPEMAIL)
	result.AddTag(TAG_PROVISION_ALLOWBLUETOOTH, ID_PROVISION_ALLOWBLUETOOTH)
	result.AddTag(TAG_PROVISION_ALLOWIRDA, ID_PROVISION_ALLOWIRDA)
	result.AddTag(TAG_PROVISION_REQUIREMANUALSYNCWHENROAMING, ID_PROVISION_REQUIREMANUALSYNCWHENROAMING)
	result.AddTag(TAG_PROVISION_ALLOWDESKTOPSYNC, ID_PROVISION_ALLOWDESKTOPSYNC)
	result.AddTag(TAG_PROVISION_MAXCALENDARAGEFILTER, ID_PROVISION_MAXCALENDARAGEFILTER)
	result.AddTag(TAG_PROVISION_ALLOWHTMLEMAIL, ID_PROVISION_ALLOWHTMLEMAIL)
	result.AddTag(TAG_PROVISION_MAXEMAILAGEFILTER, ID_PROVISION_MAXEMAILAGEFILTER)
	result.AddTag(TAG_PROVISION_MAXEMAILBODYTRUNCATIONSIZE, ID_PROVISION_MAXEMAILBODYTRUNCATIONSIZE)
	result.AddTag(TAG_PROVISION_MAXEMAILHTMLBODYTRUNCATIONSIZE, ID_PROVISION_MAXEMAILHTMLBODYTRUNCATIONSIZE)
	result.AddTag(TAG_PROVISION_REQUIRESIGNEDSMIMEMESSAGES, ID_PROVISION_REQUIRESIGNEDSMIMEMESSAGES)
	result.AddTag(TAG_PROVISION_REQUIREENCRYPTEDSMIMEMESSAGES, ID_PROVISION_REQUIREENCRYPTEDSMIMEMESSAGES)
	result.AddTag(TAG_PROVISION_REQUIRESIGNEDSMIMEALGORITHM, ID_PROVISION_REQUIRESIGNEDSMIMEALGORITHM)
	result.AddTag(TAG_PROVISION_REQUIREENCRYPTIONSMIMEALGORITHM, ID_PROVISION_REQUIREENCRYPTIONSMIMEALGORITHM)
	result.AddTag(TAG_PROVISION_ALLOWSMIMEENCRYPTIONALGORITHMNEGOTIATION, ID_PROVISION_ALLOWSMIMEENCRYPTIONALGORITHMNEGOTIATION)
	result.AddTag(TAG_PROVISION_ALLOWSMIMESOFTCERTS, ID_PROVISION_ALLOWSMIMESOFTCERTS)
	result.AddTag(TAG_PROVISION_ALLOWBROWSER, ID_PROVISION_ALLOWBROWSER)
	result.AddTag(TAG_PROVISION_ALLOWCONSUMEREMAIL, ID_PROVISION_ALLOWCONSUMEREMAIL)
	result.AddTag(TAG_PROVISION_ALLOWREMOTEDESKTOP, ID_PROVISION_ALLOWREMOTEDESKTOP)
	result.AddTag(TAG_PROVISION_ALLOWINTERNETSHARING, ID_PROVISION_ALLOWINTERNETSHARING)
	result.AddTag(TAG_PROVISION_UNAPPROVEDINROMAPPLICATIONLIST, ID_PROVISION_UNAPPROVEDINROMAPPLICATIONLIST)
	result.AddTag(TAG_PROVISION_APPLICATIONNAME, ID_PROVISION_APPLICATIONNAME)
	result.AddTag(TAG_PROVISION_APPROVEDAPPLICATIONLIST, ID_PROVISION_APPROVEDAPPLICATIONLIST)
	result.AddTag(TAG_PROVISION_HASH, ID_PROVISION_HASH)

	return result
}