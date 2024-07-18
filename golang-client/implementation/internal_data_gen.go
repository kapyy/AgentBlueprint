package implementation

import protodata "golang-client/message/protoData"

type EmojiData struct {
	*protodata.EmojiData
}

func (s *EmojiData) Default() *protodata.EmojiData {
	return s.EmojiData
}
func (s *EmojiData) Set(emojidata *protodata.EmojiData) {
	s.EmojiData = emojidata
}
func (s *EmojiData) FullString() string {
	// TODO: implement me
	panic("implement me")
}
