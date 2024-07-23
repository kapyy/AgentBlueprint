package implementation

import (
	bpcontext "golang-client/bpcontext"
	proto "golang-client/message/proto"
	proto1 "google.golang.org/protobuf/proto"
	"strconv"
)

type Action struct {
	*proto.Action
}

func (s *Action) Default() *proto.Action {
	return s.Action
}
func (s *Action) Set(action *proto.Action) {
	s.Action = action
}
func (s *Action) FullString() string {
	return s.ActionDescriptionString() + s.DurationString() + s.EndTimeString() + s.StartTimeString()
}
func (s *Action) ActionDescription() string {
	return s.Action.ActionDescription
}
func (s *Action) SetActionDescription(actiondescription string) {
	s.Action.ActionDescription = actiondescription
}
func (s *Action) ActionDescriptionString() string {
	// TODO: implement me, this is where you write how you want you data to be recognized as natural language
	panic("implement me")
}
func (s *Action) Duration() int32 {
	return s.Action.Duration
}
func (s *Action) SetDuration(duration int32) {
	s.Action.Duration = duration
}
func (s *Action) DurationString() string {
	// TODO: implement me, this is where you write how you want you data to be recognized as natural language
	panic("implement me")
}
func (s *Action) EndTime() uint64 {
	return s.Action.EndTime
}
func (s *Action) SetEndTime(endtime uint64) {
	s.Action.EndTime = endtime
}
func (s *Action) EndTimeString() string {
	// TODO: implement me, this is where you write how you want you data to be recognized as natural language
	panic("implement me")
}
func (s *Action) StartTime() uint64 {
	return s.Action.StartTime
}
func (s *Action) SetStartTime(starttime uint64) {
	s.Action.StartTime = starttime
}
func (s *Action) StartTimeString() string {
	// TODO: implement me, this is where you write how you want you data to be recognized as natural language
	panic("implement me")
}
func (s *Action) Marshal() ([]byte, error) {
	return proto1.Marshal(s.Action)
}
func (s *Action) GetPropIndex(index uint64) (bpcontext.DataPropertyInterface, string) {
	switch index {
	case 0:
		return s, s.FullString()
	case 1:
		return s, s.ActionDescriptionString()
	case 2:
		return s, s.DurationString()
	case 4:
		return s, s.EndTimeString()
	case 3:
		return s, s.StartTimeString()
	default:
		return s, ""
	}
}

type ActionList struct {
	actionlist []*Action
}

func (sl *ActionList) Set(actionlist *proto.ActionList) {
	sl.actionlist = make([]*Action, 0)
	for _, protoAction := range actionlist.ActionList {
		action := &Action{}
		action.Set(protoAction)
		sl.actionlist = append(sl.actionlist, action)
	}
}
func (sl *ActionList) Marshal() ([]byte, error) {
	actionList := &proto.ActionList{}
	for _, action := range sl.actionlist {
		actionList.ActionList = append(actionList.ActionList, action.Action)
	}
	return proto1.Marshal(actionList)
}
func (sl *ActionList) GetPropIndex(index uint64) (bpcontext.DataPropertyInterface, string) {
	stringList := ""
	for i, s := range sl.actionlist {
		_, stringObj := s.GetPropIndex(index)
		stringList += strconv.Itoa(i) + ". " + stringObj + "\n"
	}
	return sl, stringList
}
