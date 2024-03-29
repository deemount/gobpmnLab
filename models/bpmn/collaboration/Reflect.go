package collaboration

import "reflect"

// ReflectCollaborationGetMethodsToMap ...
func ReflectCollaborationGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Collaboration{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectCollaborationMethodsToMap ...
func ReflectCollaborationMethodsToMap() map[int]string {
	var ptr *Collaboration
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectCollaborationMethodsToSlice ...
func ReflectCollaborationMethodsToSlice() []string {
	var ptr *Collaboration
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}

// ReflectParticipantMethodsToMap ...
func ReflectParticipantGetMethodsToMap() map[int]string {
	m := make(map[int]string)
	t := reflect.TypeOf(Participant{})
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectParticipantMethodsToMap ...
func ReflectParticipantMethodsToMap() map[int]string {
	var ptr *Participant
	m := make(map[int]string)
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m[i] = t.Method(i).Name
	}
	return m
}

// ReflectParticipantMethodsToSlice ...
func ReflectParticipantMethodsToSlice() []string {
	var ptr *Participant
	m := []string{}
	t := reflect.TypeOf(ptr)
	for i := 0; i < t.NumMethod(); i++ {
		m = append(m, t.Method(i).Name)
	}
	return m
}
