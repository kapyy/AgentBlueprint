package implementation

import bpcontext "golang-client/bpcontext"

func InitMgrComponent() {
	bpcontext.RegisterMgrComponent(1001, &ActionManager{})
	bpcontext.RegisterMgrComponent(4001, &ParsedActionManager{})
}
