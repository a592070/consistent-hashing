package consistent_hashing

import (
	"fmt"
	"testing"
)

func TestRing_AddNode(t *testing.T) {
	type fields struct {
		Nodes []string
	}
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name:   "add node",
			fields: fields{Nodes: []string{"1", "2", "3", "4"}},
			args:   args{id: "5"},
		},
		{
			name:   "add node",
			fields: fields{Nodes: []string{"1", "2", "3", "4"}},
			args:   args{id: "4"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRing()
			for _, node := range tt.fields.Nodes {
				r.AddNode(node)
			}
			for i, node := range r.Nodes {
				fmt.Println(i, node)
			}
			fmt.Println("====Added====")
			r.AddNode(tt.args.id)
			for i, node := range r.Nodes {
				fmt.Println(i, node)
			}
		})
	}
}

func TestRing_Get(t *testing.T) {
	type fields struct {
		Nodes []string
	}
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "case1",
			fields: fields{Nodes: []string{"1", "2", "3", "4"}},
			args:   args{id: "1"},
			want:   "1",
		},
		{
			name:   "case2",
			fields: fields{Nodes: []string{"1", "2", "3", "4"}},
			args:   args{id: "4"},
			want:   "4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRing()
			for _, node := range tt.fields.Nodes {
				r.AddNode(node)
			}
			for i, node := range r.Nodes {
				fmt.Println(i, node)
			}
			if got := r.Get(tt.args.id); got != tt.want {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRing_RemoveNode(t *testing.T) {
	type fields struct {
		Nodes []string
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:    "case1",
			fields:  fields{Nodes: []string{"1", "2", "3", "4"}},
			args:    args{id: "1"},
			wantErr: false,
		},
		{
			name:    "case2",
			fields:  fields{Nodes: []string{"1", "2", "3", "4"}},
			args:    args{id: "99"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRing()
			for _, node := range tt.fields.Nodes {
				r.AddNode(node)
			}
			for i, node := range r.Nodes {
				fmt.Println(i, node)
			}
			if err := r.RemoveNode(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("RemoveNode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRing_search(t *testing.T) {
	type fields struct {
		Nodes []string
	}
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "case1",
			fields: fields{Nodes: []string{"1", "2", "3", "4"}},
			args:   args{id: "2"},
			want:   0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRing()
			for _, node := range tt.fields.Nodes {
				r.AddNode(node)
			}
			for i, node := range r.Nodes {
				fmt.Println(i, node)
			}
			if got := r.search(tt.args.id); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
