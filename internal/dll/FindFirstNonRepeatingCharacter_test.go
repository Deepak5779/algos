package dll

import (
	"reflect"
	"testing"
)

func Test_newNode(t *testing.T) {
	type args struct {
		data byte
	}
	tests := []struct {
		name string
		args args
		want *DllNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newNode(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newDll(t *testing.T) {
	type args struct {
		data byte
	}
	tests := []struct {
		name string
		args args
		want *Dll
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newDll(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDll_insert(t *testing.T) {
	type fields struct {
		head *DllNode
	}
	type args struct {
		node *DllNode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dll := &Dll{
				head: tt.fields.head,
			}
			dll.insert(tt.args.node)
		})
	}
}

func TestDll_delete(t *testing.T) {
	type fields struct {
		head *DllNode
	}
	type args struct {
		node *DllNode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dll := &Dll{
				head: tt.fields.head,
			}
			dll.delete(tt.args.node)
		})
	}
}

func TestDll_toString(t *testing.T) {
	dll := newDll('d')
	dll.insertData('e')
	dll.insertData('e')
	dll.insertData('p')
	dll.insertData('a')
	dll.insertData('k')

	dll2 := newDll('a')
	dll2.insertData('b')
	n := newNode('c')
	dll2.insert(n)
	dll2.insertData('d')
	dll2.delete(n)

	type fields struct {
		dll *Dll
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"test1", fields{dll: dll}, "deepak"},
		{"test2", fields{dll: dll2}, "abd"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.dll.toString(); got != tt.want {
				t.Errorf("Dll.toString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstNonRepeatingCharecter(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{"test1", args{"cleartaxis"}, 'c'},
		{"test2", args{"cleartaxisc"}, 'l'},
		{"test3", args{"cleartaxiscl"}, 'e'},
		{"test4", args{"cleartaxiscle"}, 'r'},
		{"test5", args{"cleartaxisclea"}, 'r'},
		{"test3", args{"cleartaxisclear"}, 't'},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstNonRepeatingCharecter(tt.args.s); got != tt.want {
				t.Errorf("firstNonRepeatingCharecter() = %v, want %v", got, tt.want)
			}
		})
	}
}
