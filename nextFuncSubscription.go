package pubsub

import (
	"github.com/bhbosman/goCommsDefinitions"
	"github.com/reactivex/rxgo/v2"
)

type NextFuncSubscription struct {
	isClosed bool
	nextFunc rxgo.NextFunc
}

func (self *NextFuncSubscription) Count() int {
	return 0
}

func (self *NextFuncSubscription) Flush() {
}

func (self *NextFuncSubscription) Close() {
	self.isClosed = true
}

func (self *NextFuncSubscription) Add(item interface{}) {
	if !self.isClosed {
		self.nextFunc(item)
	}
}

func (self *NextFuncSubscription) isIPubSubBag() goCommsDefinitions.IPubSubBag {
	return self
}

func NewNextFuncSubscription(nextFunc rxgo.NextFunc) *NextFuncSubscription {
	return &NextFuncSubscription{
		nextFunc: nextFunc,
	}
}
