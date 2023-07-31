package gomlib

type OSType int

const (
	OSAndroid OSType = iota + 0
	OSIOS
	OSReserved
	OSHarmony
	OSIPadOS
)

func (t OSType) String() string {
	switch t {
	case OSAndroid:
		return "Android"
	case OSIOS:
		return "iOS"
	case OSIPadOS:
		return "iPadOS"
	case OSHarmony:
		return "HarmonyOS"
	default:
		return "Unknown"
	}
}

type DeviceEntry interface {
	GetState() string
	GetSerial() string
	GetID() int // for Android it is transport_id, for iOS it is DeviceID
	GetConnType() string
	GetOSType() OSType
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
