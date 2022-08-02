package pubsub

//type ChannelSubscription struct {
//	Data chan interface{}
//}
//
//func (self *ChannelSubscription) Count() int {
//	return len(self.Data)
//}
//
//func (self *ChannelSubscription) Flush() {
//	for range self.Data {
//	}
//}
//
//func (self *ChannelSubscription) Close() {
//	close(self.Data)
//}
//
//func (self *ChannelSubscription) Add(item interface{}) {
//	self.Data <- item
//}
//
//func (self *ChannelSubscription) isIPubSubBag() goCommsDefinitions.IPubSubBag {
//	return self
//}
//
//func NewChannelSubscription(capacity int) *ChannelSubscription {
//	return &ChannelSubscription{
//		Data: make(chan interface{}, capacity),
//	}
//}
