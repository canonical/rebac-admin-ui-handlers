package database

import (
	"slices"
	"sync"
)

type RelationshipTuple struct {
	Left  string
	Right string
}

type relationship struct {
	mutex  sync.RWMutex
	Tuples []RelationshipTuple
}

func NewRelationship() *relationship {
	return &relationship{
		Tuples: []RelationshipTuple{},
	}
}

func (r *relationship) GetAllLefts() []string {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	result := []string{}
	for _, t := range r.Tuples {
		result = append(result, t.Left)
	}
	return result
}

func (r *relationship) GetAllRights() []string {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	result := []string{}
	for _, t := range r.Tuples {
		result = append(result, t.Right)
	}
	return result
}

func (r *relationship) GetLefts(right string) []string {
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

func (r *relationship) GetRights(left string) []string {
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

func (r *relationship) Add(left string, right string) bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.add(left, right)
}

func (r *relationship) add(left string, right string) bool {
	for _, t := range r.Tuples {
		if t.Left == left && t.Right == right {
			// Already exists
			return false
		}
	}
	r.Tuples = append(r.Tuples, RelationshipTuple{left, right})
	return true
}

func (r *relationship) Remove(left string, right string) bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	return r.remove(left, right)
}

func (r *relationship) remove(left string, right string) bool {
	old := r.Tuples
	r.Tuples = slices.DeleteFunc(r.Tuples, func(t RelationshipTuple) bool {
		return t.Left == left && t.Right == right
	})
	return len(old) != len(r.Tuples)
}

func (r *relationship) RemoveLeft(left string) bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	old := r.Tuples
	r.Tuples = slices.DeleteFunc(r.Tuples, func(t RelationshipTuple) bool {
		return t.Left == left
	})
	return len(old) != len(r.Tuples)
}

func (r *relationship) RemoveRight(right string) bool {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	old := r.Tuples
	r.Tuples = slices.DeleteFunc(r.Tuples, func(t RelationshipTuple) bool {
		return t.Right == right
	})
	return len(old) != len(r.Tuples)
}

func (r *relationship) PatchLeft(left string, additions, removals []string) bool {
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

func (r *relationship) PatchRight(right string, additions, removals []string) bool {
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
