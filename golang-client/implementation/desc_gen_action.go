package implementation

import (
	protodata "golang-client/message/protoData"
	"strconv"
)

type Action struct {
	*protodata.Action
}

func (s *Action) Default() *protodata.Action {
	return s.Action
}
func (s *Action) Set(action *protodata.Action) {
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
func (s *Action) GetPropIndex(index uint64) (interface{}, string) {
	switch index {
	case 0:
		return s.Default(), s.FullString()
	case 1:
		return &protodata.Action{ActionDescription: s.ActionDescription()}, s.ActionDescriptionString()
	case 2:
		return &protodata.Action{Duration: s.Duration()}, s.DurationString()
	case 4:
		return &protodata.Action{EndTime: s.EndTime()}, s.EndTimeString()
	case 3:
		return &protodata.Action{StartTime: s.StartTime()}, s.StartTimeString()
	default:
		return &protodata.Action{}, ""
	}
}

type ActionList struct {
	actionlist []*Action
}

func (sl *ActionList) Set(actionlist *protodata.ActionList) {
	sl.actionlist = make([]*Action, 0)
	for _, protoAction := range actionlist.ActionList {
		action := &Action{}
		action.Set(protoAction)
		sl.actionlist = append(sl.actionlist, action)
	}
}
func (sl *ActionList) GetPropIndex(index uint64) (interface{}, string) {
	protoList := make([]*protodata.Action, len(sl.actionlist))
	stringList := ""
	for i, s := range sl.actionlist {
		protoObj, stringObj := s.GetPropIndex(index)
		protoList[i] = protoObj.(*protodata.Action)
		stringList += strconv.Itoa(i) + ". " + stringObj + "\n"
	}
	return &protodata.ActionList{ActionList: protoList}, stringList
}
