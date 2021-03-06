package retained

import (
	"github.com/danclive/mqtt/packets"
)

// IterateFn is the callback function used by Iterate()
// Return false means to stop the iteration.
type IterateFn func(message packets.Message) bool

// Store is the interface used by mqtt.server and external logic to handler the operations of retained messages.
// User can get the implementation from mqtt.Server interface.
// This interface provides the ability for extensions to interact with the retained message store.
// Notice:
// This methods will not trigger any mqtt hooks.
type Store interface {
	// GetRetainedMessage returns the message that equals the passed topic.
	GetRetainedMessage(topicName string) packets.Message
	// ClearAll clears all retained messages.
	ClearAll()
	// AddOrReplace adds or replaces a retained message.
	AddOrReplace(message packets.Message)
	// Remove removes a retained message.
	Remove(topicName string)
	// GetMatchedMessages returns the retained messages that match the passed topic filter.
	GetMatchedMessages(topicFilter string) []packets.Message
	// Iterate iterate all retained messages. The callback is called once for each message.
	// If callback return false, the iteration will be stopped.
	// Notice:
	// The results are not sorted in any way, no ordering of any kind is guaranteed.
	// This method will walk through all retained messages,
	// so this will be a expensive operation if there are a large number of retained messages.
	Iterate(fn IterateFn)
}
