package implementation

import proto "golang-client/message/proto"

type EmojiData struct {
	*proto.EmojiData
}

func (s *EmojiData) Default() *proto.EmojiData {
	return s.EmojiData
}
func (s *EmojiData) Set(emojidata *proto.EmojiData) {
	s.EmojiData = emojidata
}
func (s *EmojiData) FullString() string {
	// TODO: implement me
	panic("implement me")
}
