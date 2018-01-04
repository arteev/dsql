package format

import (
	"errors"
	"strings"
	"sync"
)

//Errors
var (
	ErrEmpty = errors.New("Empty format name")
)

const (
	groot    = "root"
	sepGroup = ";"
	sepName  = ":"
)

//Format - groups of values for formatting
type Format struct {
	sync.RWMutex
	name   string
	groups map[string]*Group
	raw    string
}

//Group - group of values
type Group struct {
	sync.RWMutex
	name string
	vals map[string]string
}

//New parsing from string s and returns Format
func New(s string) (*Format, error) {
	if strings.TrimSpace(s) == "" {
		return nil, ErrEmpty
	}
	parts := strings.SplitN(s, ":", 2)
	sgroups := ""
	if len(parts) > 1 {
		sgroups = parts[1]
	}
	f := &Format{
		name:   parts[0],
		groups: make(map[string]*Group),
	}

	err := f.parse(sgroups)
	if err != nil {
		return nil, err
	}
	f.raw = sgroups

	return f, nil
}

func (f *Format) RawString() string {
	return f.raw
}

//Name returns the name of the format
func (f *Format) Name() string {
	return f.name
}

//Count returns the count of groups
func (f *Format) Count() int {
	f.RLock()
	defer f.RUnlock()
	return len(f.groups)
}

func (f *Format) parse(s string) error {
	f.groups[groot] = &Group{name: groot, vals: make(map[string]string)}

	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return nil
	}

	parts := strings.Split(s, sepGroup)
	for i, p := range parts {
		gn := strings.SplitN(p, sepName, 2)

		var gname string
		if len(gn) == 1 && i == 0 {
			//root
			//f.groups[groot].parse
			gname = groot
		} else {
			gname = strings.TrimSpace(gn[0])
		}

		if gname == "" {
			continue
		}

		item, ok := f.groups[gname]
		if !ok {
			item = &Group{
				name: gname,
				vals: make(map[string]string),
			}
		}
		gval := ""
		if len(gn) > 1 {

			gval = gn[1]

		} else {
			if gname == groot {
				gval = gn[0]
			}
		}

		item.parse(gval)
		f.groups[gname] = item
	}

	return nil
}

//Root Returns root group
func (f *Format) Root() *Group {
	f.RLock()
	defer f.RUnlock()
	return f.groups[groot]
}

//Count returns the count of values
func (g *Group) Count() int {
	g.RLock()
	defer g.RUnlock()
	return len(g.vals)
}

func (g *Group) parse(s string) {
	vals := strings.Split(s, ",")
	for _, v := range vals {
		if strings.TrimSpace(v) == "" {
			continue
		}
		kv := strings.SplitN(v, "=", 2)
		if len(kv) == 1 {
			g.vals[kv[0]] = ""
		} else {
			g.vals[kv[0]] = kv[1]
		}
	}
}

//Get returns a value by key
func (g *Group) Get(key string) (string, bool) {
	g.RLock()
	defer g.RUnlock()
	val, exist := g.vals[key]
	if !exist {
		return "", false
	}
	return val, true
}

//Name returns the name of the group
func (g *Group) Name() string {
	return g.name
}

//Groups returns group names
func (f *Format) Groups() []string {
	f.RLock()
	defer f.RUnlock()
	groups := make([]string, len(f.groups))
	i := 0
	for g := range f.groups {
		groups[i] = g
		i++
	}
	return groups
}

//Group returns a group by name
func (f *Format) Group(name string) (*Group, bool) {
	f.RLock()
	defer f.RUnlock()
	g, ok := f.groups[name]
	return g, ok
}
