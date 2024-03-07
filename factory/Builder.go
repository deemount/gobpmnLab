package factory

import (
	"crypto/rand"
	"fmt"
	"hash/fnv"
	"log"
	"reflect"
	"strings"

	"github.com/deemount/gobpmnLab/models/bpmn/core"
	"github.com/deemount/gobpmnLab/models/bpmn/events/elements"
	"github.com/deemount/gobpmnLab/models/bpmn/pool"
	"github.com/deemount/gobpmnLab/utils"
)

type Def core.DefinitionsRepository

// Builder ...
type Builder struct {
	Reflect
	Settings
	Map       map[string][]interface{}
	Suffix    string
	HashTable []string
}

// Hash ...
func (h *Builder) Hash() string {
	if h.isZero() {
		r, _ := h.hash()
		h.Suffix = r.Suffix
	}
	return h.Suffix
}

// Inject all anonymous fields with a hash value by fields with type Builder
// It also setup reflected fields with boolean type and checks out the configuration by wording
// usage:
// e.g.
// var p CProcess
// p = h.inject(CProcess{}).(CProcess)
func (h *Builder) Inject(p interface{}) interface{} { return h.inject(p) }

// Create receives the definitions repository by the app in p argument
// and calls the main elements to set the maps, including process parameters
// n of process.
func (h *Builder) Build(p interface{}) { h.build(p) }

/*
 * @private
 */

// hash generates a hash value by using the crypto/rand package
// and the hash/fnv package to generate a 32-bit FNV-1a hash.
// It returns a Builder with a hash value and an error.
// If the error is not nil, it means that the hash value could not be generated.
// The hash value is a string of 32 hexadecimal characters.
// The hash value is used to generate a unique suffix for each process.
// The suffix is used to generate a unique ID for each element of a process.
// These elements are:
// - sub process
// - message flow
// - sequence flow
// - data object
// - data store
// - group
// - participant
// - lane
// - text annotation
// - association
// - process
// - start event
// - end event
// - intermediate catch event
// - intermediate throw event
// - boundary event
// - task
// - service task
// - user task
// - manual task
// - script task
// - business rule task
// - receive task
// - send task
// - call activity.

func (h Builder) hash() (Builder, error) {

	n := 8
	b := make([]byte, n)
	c := fnv.New32a()

	if _, err := rand.Read(b); err != nil {
		return Builder{}, err
	}
	s := fmt.Sprintf("%x", b)

	if _, err := c.Write([]byte(s)); err != nil {
		return Builder{}, err
	}
	defer c.Reset()

	r := Builder{Suffix: fmt.Sprintf("%x", string(c.Sum(nil)))}

	return r, nil
}

func (h *Builder) mapped(m map[string][]interface{}) *map[string][]interface{} {
	h.Map = m
	return &h.Map
}

// inject itself reflects a given struct and inject
// signed fields with hash values.
// There are two conditions to assign fields of a strcut:
// a) The struct has anonymous fields
// b) The struct has no anymous fields
// It also counts the element in their specification to know
// how much elements of each package needs to be mapped later then.
func (h *Builder) inject(p interface{}) interface{} {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("factory.builder: inject recovered", r)
		}
	}()

	log.Println("factory/Builder.go (inject): > NewReflect of interface{} >>p<<. Start injecting fields")

	ref := NewReflect(p)
	ref.Interface().New().Maps().Reflection()

	// TODO: ref.Analyze fields can be used here
	log.Printf("ref.Analyze.Anonym %+v", ref.Analyze.Anonym)
	log.Printf("ref.Analyze.Builder %+v", ref.Analyze.Builder)
	log.Printf("ref.Analyze.Config %+v", ref.Analyze.Config)
	log.Printf("ref.Analyze.Words %+v", ref.Analyze.Words)

	// anonymous field are reflected
	if ref.hasAnonymous() {

		length := len(ref.Anonym)
		log.Printf("factory/Builder.go (inject): > length of anonymous fields is %d", length)

		// create processMap, anonymMap and hashMap
		processMap := make(map[string]map[int][]interface{})
		anonymMap := make(map[int][]interface{}, length)
		//anonymMap := make(map[int][]AnonymMap, length)
		hashMap := make(map[string][]interface{}, length)

		// walk through the map with names of anonymous fields
		// e.g. starts with customer support, customer, ...
		log.Println("factory/Builder.go (inject): BEGIN OF FOR Anonymous Fields")
		for index, field := range ref.Anonym {

			log.Printf("factory/Builder.go (inject): >> start range through anonymous fields %d of %d", index+1, length)

			// get the reflected value of field
			n := ref.Temporary.FieldByName(field)

			// create the field map
			fieldMap := make(map[int][]interface{}, n.NumField())
			log.Printf("factory/Builder.go (inject): >> create fieldMap %#v with length %d", fieldMap, n.NumField())
			// create the hash slice
			hashSlice := make([]interface{}, n.NumField())
			log.Printf("factory/Builder.go (inject): >> create hashSlice %#v with length %d", hashSlice, n.NumField())

			// append to anonymMap the name of anonymous field
			anonymMap[index] = append(anonymMap[index], n.Type().Name())
			log.Printf("factory/Builder.go (inject): >> append to anonymMap[%d] the name %s", index, n.Type().Name())

			// walk through the values of fields assigned to the interface {}
			log.Printf("factory/Builder.go (inject): BEGIN OF FOR Index at %d", index)
			for i := 0; i < n.NumField(); i++ {

				// get the name of field
				name := n.Type().Field(i).Name
				// append to fieldMap the name of field
				fieldMap[i] = append(fieldMap[i], name)
				log.Printf("factory/Builder.go (inject): --> append to fieldMap[%d] the name %s", i, name)

				if strings.Contains(name, field) {
					log.Printf("factory/Builder.go (inject): --> detected field %s contains anonymous field %s", name, field)
				}

				// set by kind of reflected value above
				switch n.Field(i).Kind() {

				// kind is a bool
				case reflect.Bool:

					// only the first field, which IsExecutable is set to true,
					// means, only one process in a collaboration can be executed at runtime
					// this can be changed in the future, if the engine fits for more execution
					// options
					if strings.Contains(name, boolIsExecutable) && i == 0 {
						n.Field(0).SetBool(true)
						hashSlice[i] = bool(true)
						log.Printf("factory/Builder.go (inject): inject first bool field %s once", name)
					} else {
						hashSlice[i] = bool(false)
						log.Printf("factory/Builder.go (inject): second bool field %s stays false", name)
					}

					log.Printf("factory/Builder.go (inject): >> field injection %d of %d done", i+1, n.NumField())

				// kind is a struct
				case reflect.Struct:

					// TODO: The counting list could become bigger, when different elements needs to be counted
					h.countPool(field, name)    // counts processes, participants and their shapes
					h.countMessage(field, name) // counts message flows and their edges
					h.countStartEvent(name)     // counts startevents
					h.countTask(name)           // counts tasks
					h.countEndEvent(name)       // counts endevent

					// Injecting by each index of the given process structs
					// e.g. starts with customer support, customer,
					strHash := fmt.Sprintf("%s", n.Field(i).FieldByName("Suffix"))
					if strHash == "" {
						hash, _ := h.hash() // generate hash value
						hashSlice[i] = hash.Suffix
						n.Field(i).Set(reflect.ValueOf(hash)) // inject the field
						log.Printf("factory/Builder.go (inject): inject first at process index [%d] => %v on current fieldname %s (%v) with hash value ", index, ref.Analyze.Words[index], n.Type().Field(i).Name, n.Field(i).FieldByName("Suffix"))
					} else {
						log.Printf("factory/Builder.go (inject): current fieldname %s (%v) has got hash value before", n.Type().Field(i).Name, n.Field(i).FieldByName("Suffix"))
					}

					// TODO: Check only previous element; maybe a solution trying to erate the hash slice

					if i > 0 {

						if n.Field(i-1).Kind() != reflect.Bool {
							log.Printf("factory/Builder.go (inject): previous fieldname was %s (%v)", n.Type().Field(i-1).Name, n.Field(i-1).FieldByName("Suffix"))
						}
					} else {
						log.Printf("factory/Builder.go (inject): current fieldname %s has no previous field", n.Type().Field(i).Name)
					}

					// TODO: Check next element and set hash value; maybe a solution trying to erate the hash slice
					if i+1 < n.NumField() {
						nexthash, _ := h.hash() // generate hash value
						hashSlice[i+1] = nexthash.Suffix
						n.Field(i + 1).Set(reflect.ValueOf(nexthash)) // inject the field
						log.Printf("factory/Builder.go (inject): inject next fieldname at %v: %s (%v)", n.Type().Name(), n.Type().Field(i+1).Name, n.Field(i+1).FieldByName("Suffix"))
					} else {
						log.Printf("factory/Builder.go (inject): current fieldname %s has no fieldname to the next", n.Type().Field(i).Name)
					}

					log.Printf("factory/Builder.go (inject): >> field injection %d of %d done", i+1, n.NumField())

				}

			}
			log.Printf("factory/Builder.go (inject): END OF FOR Index at %d", index)
			log.Printf("factory/Builder.go (inject): >> show current hashSlice %#v", hashSlice)
			// merge the hashSlice with the hashMap
			mergeStringSliceToMap(hashMap, n.Type().Name(), hashSlice)
			log.Printf("factory/Builder.go (inject): >> show current hashMap %#v", hashMap)
			// merge the fieldMap with the processMap
			processMap[n.Type().Name()] = MergeMaps(anonymMap, fieldMap)
			log.Printf("factory/Builder.go (inject): >> show current processMap %#v", processMap)
		}
		log.Println("factory/Builder.go (inject): END OF FOR 1")
		log.Printf("factory/Builder.go (inject): > merged hashSlice with hashMap %+v", hashMap)
		log.Printf("factory/Builder.go (inject): > merged anonymMap with fieldMap to processMap %+v", processMap)
	} else {
		log.Println("factory/Builder.go (inject): > no anonymous fields detected")
	}

	// anonymous field aren't reflected
	if ref.hasNotAnonymous() {

		// walk through the map with names of builder fields
		for _, builderField := range ref.Builder {

			// get the reflected value of field
			n := ref.Temporary.FieldByName(builderField)

			if strings.Contains(builderField, fieldProcess) {
				h.NumProc++
			}

			if strings.Contains(builderField, fieldStartEvent) && utils.After(builderField, "From") == fieldStartEvent {
				h.NumStartEvent++
				h.NumShape++
			}

			if strings.Contains(builderField, fieldTask) && utils.After(builderField, "From") == fieldTask {
				h.NumTask++
				h.NumShape++
			}

			if strings.Contains(builderField, fieldEndEvent) {
				h.NumEndEvent++
				h.NumShape++
			}

			hash, _ := h.hash()          // generate hash value
			n.Set(reflect.ValueOf(hash)) // inject the field

			log.Printf("factory/Builder.go (inject): inject struct field %v", builderField)

		}

		// walk through the map with names of boolean fields
		for _, configField := range ref.Config {

			// get the reflected value of field
			n2 := ref.Temporary.FieldByName(configField)

			// only the first field, which IsExecutable is set to true
			n2.SetBool(true)

			log.Printf("factory/Builder.go (inject): inject config field %v", configField)

		}

	} else {
		log.Println("factory/Builder.go (inject): > no builder fields detected")
	}

	p = ref.Set()
	ref.countWords()

	log.Printf("factory/Builder.go (inject): eof inject; detect %d of reflected public methods", reflect.TypeOf(p).NumMethod())

	return p

}

// build contains the reflected process definition (p interface{})
// and can call SetMainElements from the core package or
// calls it by the reflected method name.
// This method hides some of the setters by building the BPMN
// with reflection
func (h *Builder) build(p interface{}) {

	// el is the interface {}
	el := reflect.ValueOf(&p).Elem()

	// Allocate a temporary variable with type of the struct.
	// el.Elem() is the value contained in the interface
	definitions := reflect.New(el.Elem().Type()).Elem() // *core.Definitions
	definitions.Set(el.Elem())                          // reflected process definitions el will be assigned to the core definitions

	// set collaboration, process and diagram
	collaboration := definitions.MethodByName(MethodSetCollaboration)
	collaboration.Call([]reflect.Value{})

	process := definitions.MethodByName(MethodSetProcess)
	process.Call([]reflect.Value{reflect.ValueOf(h.NumProc)}) // h.numProc represents number of processes

	diagram := definitions.MethodByName(MethodSetDiagram)
	diagram.Call([]reflect.Value{reflect.ValueOf(1)})

}

/*** Semantic, Analyze and Statistic ***/

// inPool ...
func (h *Builder) countPool(field, builderField string) {
	if h.isPool(field) {
		if strings.Contains(builderField, fieldProcess) {
			h.NumProc++
		}
		if strings.Contains(builderField, fieldID) {
			h.NumPart++
			h.NumShape++
		}
	}
}

// inMessage ...
func (h *Builder) countMessage(field, builderField string) {
	if h.isMessage(field) {
		if strings.Contains(builderField, fieldMessage) {
			h.NumMsg++
			h.NumEdge++
		}
	}
}

// isStartEvent ...
func (h *Builder) countStartEvent(builderField string) {
	if strings.Contains(builderField, fieldStartEvent) && utils.After(builderField, "From") == "" {
		h.NumStartEvent++
		h.NumShape++
	}
}

// isTask ...
func (h *Builder) countTask(builderField string) {
	if strings.Contains(builderField, fieldTask) && utils.After(builderField, "From") == "" {
		h.NumTask++
		h.NumShape++
	}
}

// isEndEvent ...
func (h *Builder) countEndEvent(builderField string) {
	if strings.Contains(builderField, fieldEndEvent) {
		h.NumEndEvent++
		h.NumShape++
	}
}

// isZero ...
func (h *Builder) isZero() bool { return h.Suffix == "" }

// isPool ...
func (h *Builder) isPool(field string) bool { return pool.IsPool(field) }

// isMessage ...
func (h *Builder) isMessage(field string) bool { return elements.IsMessage(field) }

// isUnknown ...
func (h *Builder) isUnknown(field string) bool { return true }

func MergeMaps[M ~map[K]V, K comparable, V any](src ...M) M {
	merged := make(M)
	for _, m := range src {
		for k, v := range m {
			merged[k] = v
		}
	}
	return merged
}

func mergeStringSliceToMap(m map[string][]interface{}, k string, v []interface{}) {
	if m[k] == nil {
		m[k] = make([]interface{}, len(v))
		for i, s := range v {
			m[k][i] = interface{}(s)
		}
	} else {
		m[k] = append(m[k], v...)
	}
}
