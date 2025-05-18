package gomlib

type OSType int

const (
	Android OSType = iota + 1
	IOS
	HarmonyOS
	HarmonyOSNext
	Windows
	MacOS
	Linux
	IPadOS OSType = 10
)

// deprecated
const (
	OSAndroid OSType = iota + 1
	OSIOS
	OSHarmony
	OSHarmonyNext
	OSWindows
	OSMacOS
	OSLinux
	OSIPadOS OSType = 10
)

func (t OSType) String() string {
	switch t {
	case Android:
		return "Android"
	case IOS:
		return "iOS"
	case IPadOS:
		return "iPadOS"
	case HarmonyOS:
		return "HarmonyOS"
	case HarmonyOSNext:
		return "HarmonyNext"
	case Windows:
		return "Windows"
	case MacOS:
		return "macOS"
	case Linux:
		return "Linux"
	default:
		return "Unknown"
	}
}

type DeviceEntry interface {
	Serial() string
	ID() int // for Android it is transport_id, for iOS it is DeviceID
	ConnType() string
	OSType() OSType
}

type ActionListener interface {
	OnAction(interface{})
}

type funcActionListener struct {
	onAction func(interface{})
}

func (l funcActionListener) OnAction(d interface{}) {
	if f := l.onAction; f != nil {
		f(d)
	}
}

func NewActionListener(action func(interface{})) ActionListener {
	return &funcActionListener{action}
}
