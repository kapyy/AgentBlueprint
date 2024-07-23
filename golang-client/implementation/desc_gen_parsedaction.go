package implementation

import (
	bpcontext "golang-client/bpcontext"
	proto "golang-client/message/proto"
	proto1 "google.golang.org/protobuf/proto"
)

type ParsedAction struct {
	embedded_emojilist *EmojiData
	*proto.ParsedAction
}

func (s *ParsedAction) Default() *proto.ParsedAction {
	return s.ParsedAction
}
func (s *ParsedAction) Set(parsedaction *proto.ParsedAction) {
	s.embedded_emojilist.Set(parsedaction.EmojiList)
	s.ParsedAction = parsedaction
}
func (s *ParsedAction) FullString() string {
	return s.EmojiListString()
}
func (s *ParsedAction) EmojiList() *EmojiData {
	return s.embedded_emojilist
}
func (s *ParsedAction) SetEmojiList(emojilist *proto.EmojiData) {
	s.embedded_emojilist.Set(emojilist)
	s.ParsedAction.EmojiList = emojilist
}
func (s *ParsedAction) EmojiListString() string {
	// TODO: implement me, this is where you write how you want you data to be recognized as natural language
	panic("implement me")
}
func (s *ParsedAction) Marshal() ([]byte, error) {
	return proto1.Marshal(s.ParsedAction)
}
func (s *ParsedAction) GetPropIndex(index uint64) (bpcontext.DataPropertyInterface, string) {
	switch index {
	case 0:
		return s, s.FullString()
	case 1:
		return s, s.EmojiListString()
	default:
		return s, ""
	}
}
