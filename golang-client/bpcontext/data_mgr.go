package bpcontext

type Listener func(...any)

type InternalEventInterface interface {
	AddListener(eventType int, listener Listener)
	Invoke(eventType int, args ...any)
}

type DataEventManager struct {
	list map[int][]Listener
}

func (b *DataEventManager) AddListener(eventType int, listener Listener) {
	b.list[eventType] = append(b.list[eventType], listener)
}

func (b *DataEventManager) Invoke(eventType int, args ...any) {
	for _, listener := range b.list[eventType] {
		listener(args)
	}
}

type DistributedEvent struct {
	eventManager InternalEventInterface
}

func (s *DistributedEvent) ConnectToEventManager(eventManager InternalEventInterface) {
	s.eventManager = eventManager
}
func (s *DistributedEvent) AddListener(eventType int, listener Listener) {
	s.eventManager.AddListener(eventType, listener)
}

func (s *DistributedEvent) Invoke(eventType int, args ...any) {
	s.eventManager.Invoke(eventType, args) // TODO: undefined recursive call?
}

type ContextEventManager struct {
	list map[int][]Listener
}

func (c *ContextEventManager) AddListener(eventType int, listener Listener) {
	c.list[eventType] = append(c.list[eventType], listener)
}

func (c *ContextEventManager) Invoke(eventType int, args ...any) {
	for _, listener := range c.list[eventType] {
		listener(args)
	}
}
