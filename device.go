package gomlib

type OSType int

const (
	OSAndroid OSType = iota + 1
	OSIOS
)

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
