package database

import (
	"slices"
	"sync"
)

// Relationship maintains a unique set of (left, right) tuples to track the
// relation between entities of (one or two) specific types. This can be viewed
// as an in-memory implementation of foreign-key Relationship and JOIN
// operations in relational databases.
type Relationship struct {
	mutex  sync.RWMutex
	Tuples []RelationshipTuple
}

// RelationshipTuple represents a (left, right) tuple.
type RelationshipTuple struct {
	Left  string
	Right string
}

// NewRelationship creates a new Relationship instance.
func NewRelationship() *Relationship {
	return &Relationship{
		Tuples: []RelationshipTuple{},
	}
}

// GetLefts returns all "left" components of tuples that have a specific "right"
// component.
func (r *Relationship) GetLefts(right string) []string {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	result := []string{}
	for _, t := range r.Tuples {
		if t.Right == right {
			result = append(result, t.Left)
		}
	}
	return result
}

// GetRights returns all "right" components of tuples that have a specific "left"
// component.
func (r *Relationship) GetRights(left string) []string {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	result := []string{}
	for _, t := range r.Tuples {
		if t.Left == left {
			result = append(result, t.Right)
		}
	}
	return result
}

// Add adds a new tuple to the relationship. If the tuple already exists, the
// method returns false; otherwise returns true.
func (r *Relationship) Add(left string, right string) bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.add(left, right)
}

func (r *Relationship) add(left string, right string) bool {
	for _, t := range r.Tuples {
		if t.Left == left && t.Right == right {
			// Already exists
			return false
		}
	}
	r.Tuples = append(r.Tuples, RelationshipTuple{left, right})
	return true
}

// Remove removes a tuple from the relationship. If the tuple does not exist,
// the method returns false; otherwise returns true.
func (r *Relationship) Remove(left string, right string) bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.remove(left, right)
}

func (r *Relationship) remove(left string, right string) bool {
	old := r.Tuples
	r.Tuples = slices.DeleteFunc(r.Tuples, func(t RelationshipTuple) bool {
		return t.Left == left && t.Right == right
	})
	return len(old) != len(r.Tuples)
}

// RemoveLeft removes all tuples that have a specific "left" component. If no
// tuple was deleted, the method returns false; otherwise returns true.
func (r *Relationship) RemoveLeft(left string) bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	old := r.Tuples
	r.Tuples = slices.DeleteFunc(r.Tuples, func(t RelationshipTuple) bool {
		return t.Left == left
	})
	return len(old) != len(r.Tuples)
}

// RemoveRight removes all tuples that have a specific "right" component. If no
// tuple was deleted, the method returns false; otherwise returns true.
func (r *Relationship) RemoveRight(right string) bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	old := r.Tuples
	r.Tuples = slices.DeleteFunc(r.Tuples, func(t RelationshipTuple) bool {
		return t.Right == right
	})
	return len(old) != len(r.Tuples)
}

// PatchLeft adds and removes batches of "right" component values that for a
// specific "left" component. If nothing was added/removed, the method returns
// false; otherwise returns true.
func (r *Relationship) PatchLeft(left string, additions, removals []string) bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	result := false
	for _, id := range additions {
		result = r.add(left, id) || result
	}
	for _, id := range removals {
		result = r.remove(left, id) || result
	}
	return result
}

// PatchRight adds and removes batches of "left" component values that for a
// specific "right" component. If nothing was added/removed, the method returns
// false; otherwise returns true.
func (r *Relationship) PatchRight(right string, additions, removals []string) bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	result := false
	for _, id := range additions {
		result = r.add(id, right) || result
	}
	for _, id := range removals {
		result = r.remove(id, right) || result
	}
	return result
}

func mapStringSlice[T any](s []string, f func(string) T) []T {
	result := make([]T, 0, len(s))
	for _, v := range s {
		result = append(result, f(v))
	}
	return result
}
