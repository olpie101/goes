// TODO: move to eventstore package: eventstore.WithBus(store, bus)
package eventbus

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/modernice/goes/event"
)

// WithStoreOption is an option for WithStore.
type WithStoreOption func(*busWithStore)

type busWithStore struct {
	event.Bus

	s           event.Store
	storeFilter []func(event.Event) bool
}

// WithStore decorates an Event Bus with an Event Store. Events are inserted
// into the Store before they are published through the Bus. Events won't be
// published if the insertion into the Store fails.
func WithStore(b event.Bus, s event.Store, opts ...WithStoreOption) event.Bus {
	bs := &busWithStore{
		Bus: b,
		s:   s,
	}
	for _, opt := range opts {
		opt(bs)
	}
	return bs
}

// StoreFilter returns a WithStoreOption that adds the provided filter to the
// Bus. Use this Option to prevent Events from being inserted into the Event
// Store before they are published. Events that are published through the Bus
// are passed to every filter until one of them returns false. If a filter
// returns false for an Event, that Event won't be inserted into the event
// store and will only be published through the Bus.
func StoreFilter(filter ...func(event.Event) bool) WithStoreOption {
	return func(b *busWithStore) {
		b.storeFilter = append(b.storeFilter, filter...)
	}
}

// RequireAggregate returns a WithStoreOption that prevents Events that don't
// belong to an Aggregate to be inserted into the Event Store.
func RequireAggregate() WithStoreOption {
	return StoreFilter(func(evt event.Event) bool {
		return evt.AggregateName() != "" && evt.AggregateID() != uuid.Nil
	})
}

func (bus *busWithStore) Publish(ctx context.Context, events ...event.Event) error {
	if store := bus.storable(events...); len(store) > 0 {
		if err := bus.s.Insert(ctx, store...); err != nil {
			return fmt.Errorf("store: %w", err)
		}
	}

	if err := bus.Bus.Publish(ctx, events...); err != nil {
		return fmt.Errorf("publish: %w", err)
	}

	return nil
}

func (bus *busWithStore) storable(events ...event.Event) []event.Event {
	store := make([]event.Event, 0, len(events))
	for _, evt := range events {
		y := true
		for _, fn := range bus.storeFilter {
			if !fn(evt) {
				y = false
				break
			}
		}
		if y {
			store = append(store, evt)
		}
	}
	return store
}
