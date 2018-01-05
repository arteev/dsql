package format

import "testing"

func TestCommon(t *testing.T) {
	cases := []struct {
		data      string
		name      string
		mustError bool
		err       error
	}{
		{"", "", true, ErrEmpty},
		{"table", "table", false, nil},
		{"table:", "table", false, nil},
		{"table:test", "table", false, nil},
	}

	for _, c := range cases {
		f, err := New(c.data)
		if c.mustError {
			if err == nil {
				t.Fatal("Must be error")

			} else if err.Error() != c.err.Error() {
				t.Errorf("Must be error: %s,got %s", c.err, err)
			}
			continue
		}
		if !c.mustError && err != nil {
			t.Error(err)
			continue
		}

		if f.Name() != c.name {
			t.Errorf("Expected name %s,got %s", c.name, f.name)
		}
	}
}

func TestGroups(t *testing.T) {
	cases := []struct {
		data   string
		groups []string
	}{
		{"fmttest", []string{groot}},
		{"fmttest:val=123", []string{groot}},
		{"fmttest:;;;;;;", []string{groot}},
		{"fmttest:group1:val=223", []string{groot, "group1"}},
		{"fmttest:group1:;;;;", []string{groot, "group1"}},
		{"fmttest:group1:val=223;group2:val2=asd", []string{groot, "group1", "group2"}},
		{"fmttest:group1:val=223;group1:val2=333", []string{groot, "group1"}},
		{"fmttest:group1:;group2:;", []string{groot, "group1", "group2"}},
	}
	for _, c := range cases {
		f, err := New(c.data)
		if err != nil {
			t.Fatal(err)
		}
		if f.Count() != len(c.groups) {
			t.Errorf("Expected count groups %d, got %d", len(c.groups), f.Count())
		}
		for _, g := range c.groups {
			if _, exists := f.Group(g); !exists {
				t.Errorf("Expected exist group %q", g)
			}
		}

		gotGroups := f.Groups()
		if len(gotGroups) != len(c.groups) {
			t.Errorf("Expected Groups() %d, got %d", len(c.groups), len(gotGroups))
		}

		for _, g := range gotGroups {
			grGet, exists := f.Group(g)
			if !exists {
				t.Errorf("Expected exist group %q", g)
			}
			if grGet.Name() != g {
				t.Errorf("Expected group name %q,got %q", grGet.Name(), g)
			}

		}

	}
}

func TestRootValues(t *testing.T) {
	f, err := New("test:key=val")
	if err != nil {
		t.Fatal(err)
	}
	if got, _ := f.Group(groot); f.Root() != got {
		t.Errorf("Expected group:%q  %v,got %v", groot, f.Root(), got)
	}
}

func TestRaw(t *testing.T) {
	grs := "group1:;group2"
	f, err := New("test:" + grs)
	if err != nil {
		t.Fatal(err)
	}
	if got := f.RawString(); got != grs {
		t.Errorf("Expected raw %q,got %q", grs, got)
	}

	f, err = New("test")
	if err != nil {
		t.Fatal(err)
	}
	if got := f.RawString(); got != "" {
		t.Errorf("Expected raw %q,got %q", "", got)
	}
}

func TestExistValues(t *testing.T) {
	f, err := New("test:key=val;group1:key2=val2")
	if err != nil {
		t.Fatal(err)
	}
	if _, exist := f.Root().Get("key2"); exist {
		t.Errorf("Expected not exists key %q", "key2")
	}

	gr, _ := f.Group("group1")
	if _, exist := gr.Get("key33"); exist {
		t.Errorf("Expected not exists key %q", "key33")
	}
}
func TestGroupsValues(t *testing.T) {
	cases := []struct {
		data       string
		values     map[string]string
		checkGroup string
	}{
		{"test", map[string]string{}, groot},
		{"test:", map[string]string{}, groot},
		{"test:key", map[string]string{"key": ""}, groot},
		{"test:key=val", map[string]string{"key": "val"}, groot},
		{"test:key=val,key2=val2", map[string]string{"key": "val", "key2": "val2"}, groot},
		{"test:key=val,", map[string]string{"key": "val"}, groot},
		{"test:key=val,,,,,,", map[string]string{"key": "val"}, groot},
		{"test:key=val;group1:key1=val1", map[string]string{"key": "val"}, groot},
		{"test:group1:key1=val1", map[string]string{"key1": "val1"}, "group1"},
		{"test:k=v,a=b;group1:key1=val1", map[string]string{"key1": "val1"}, "group1"},
		{"test:k=v,a=b;group1:key1=val1;group2:x=y,z=k;group3:g=h", map[string]string{"x": "y", "z": "k"}, "group2"},
	}
	for _, c := range cases {
		f, err := New(c.data)
		if err != nil {
			t.Fatal(err)
		}

		group, exist := f.Group(c.checkGroup)
		if !exist {
			t.Fatalf("Expected group %s,but not exists", c.checkGroup)
		}

		if group.Count() != len(c.values) {
			t.Errorf("Expected %d,got %d", len(c.values), group.Count())
		}

		for k, v := range c.values {
			got, exist := group.Get(k)
			if !exist {
				t.Errorf("Expected exists the key %q", k)
			} else if got != v {
				t.Errorf("Expected value:%q,got %q", v, got)
			}
		}
	}

}
