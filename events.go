package centrifuge

import "context"

// Mediator allows to proxy Centrifugo events to Go application code.
type Mediator struct {
	// Connect called every time client connects to node.
	Connect ConnectHandler
	// Disconnect called when client disconnected.
	Disconnect DisconnectHandler
	// Subscribe called when client subscribes on channel.
	Subscribe SubscribeHandler
	// Unsubscribe called when client unsubscribes from channel.
	Unsubscribe UnsubscribeHandler
	// Publish called when client publishes message into channel.
	Publish PublishHandler
	// Presence allows to register action to be executed on every
	// periodic client connection presence update.
	Presence PresenceHandler
	// Refresh called when it's time to refresh connection credentials.
	Refresh RefreshHandler
	// RPC allows to register custom logic on incoming RPC calls.
	RPC RPCHandler
	// Message called when client sent asynchronous message.
	Message MessageHandler
}

// EventContext added to all specific event contexts.
type EventContext struct {
	Client Client
}

// ConnectContext contains fields related to connect event.
type ConnectContext struct {
	EventContext
}

// ConnectReply contains fields determining the reaction on connect event.
type ConnectReply struct {
	Error      *Error
	Disconnect *Disconnect
}

// ConnectHandler ...
type ConnectHandler func(context.Context, ConnectContext) ConnectReply

// DisconnectContext contains fields related to disconnect event.
type DisconnectContext struct {
	EventContext
	Disconnect *Disconnect
}

// DisconnectReply contains fields determining the reaction on disconnect event.
type DisconnectReply struct{}

// DisconnectHandler ...
type DisconnectHandler func(context.Context, DisconnectContext) DisconnectReply

// SubscribeContext contains fields related to subscribe event.
type SubscribeContext struct {
	EventContext
	Channel string
}

// SubscribeReply contains fields determining the reaction on subscribe event.
type SubscribeReply struct {
	Error      *Error
	Disconnect *Disconnect
}

// SubscribeHandler ...
type SubscribeHandler func(context.Context, SubscribeContext) SubscribeReply

// UnsubscribeContext contains fields related to unsubscribe event.
type UnsubscribeContext struct {
	EventContext
	Channel string
}

// UnsubscribeReply contains fields determining the reaction on unsubscribe event.
type UnsubscribeReply struct {
}

// UnsubscribeHandler ...
type UnsubscribeHandler func(context.Context, UnsubscribeContext) UnsubscribeReply

// PublishContext contains fields related to publish event.
type PublishContext struct {
	EventContext
	Channel     string
	Publication *Publication
}

// PublishReply contains fields determining the reaction on publish event.
type PublishReply struct {
	Error      *Error
	Disconnect *Disconnect
}

// PublishHandler ...
type PublishHandler func(context.Context, PublishContext) PublishReply

// PresenceContext contains fields related to presence update event.
type PresenceContext struct {
	EventContext
	Channels []string
}

// PresenceReply contains fields determining the reaction on presence update event.
type PresenceReply struct {
	Disconnect *Disconnect
}

// PresenceHandler ...
type PresenceHandler func(context.Context, PresenceContext) PresenceReply

// RefreshContext contains fields related to refresh event.
type RefreshContext struct {
	EventContext
}

// RefreshReply contains fields determining the reaction on refresh event.
type RefreshReply struct {
	Exp  int64
	Info []byte
}

// RefreshHandler ...
type RefreshHandler func(context.Context, RefreshContext) RefreshReply

// RPCContext contains fields related to rpc request.
type RPCContext struct {
	EventContext
	Data Raw
}

// RPCReply contains fields determining the reaction on rpc request.
type RPCReply struct {
	Error      *Error
	Disconnect *Disconnect
	Data       Raw
}

// RPCHandler must handle incoming command from client.
type RPCHandler func(context.Context, RPCContext) RPCReply

// MessageContext contains fields related to message request.
type MessageContext struct {
	EventContext
	Data Raw
}

// MessageReply contains fields determining the reaction on message request.
type MessageReply struct {
	Disconnect *Disconnect
}

// MessageHandler must handle incoming async message from client.
type MessageHandler func(context.Context, MessageContext) MessageReply
