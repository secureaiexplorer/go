package main

import "testing"

func TestQueue_Publish(t *testing.T) {
	type fields struct {
		data        []string
		subscribers map[string]int
	}
	type args struct {
		message string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "TestCase1",
			fields: fields{
				data: []string{"message1", "message2", "message3"},
			},
		},
		{
			name: "TestCase2",
			fields: fields{
				data: []string{"m1", "m2"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &MyQueue{}
			for i := 0; i < len(tt.fields.data); i++ {
				q.Publish(tt.fields.data[i])
			}
			if len(q.Data) != len(tt.fields.data) {
				t.Error("Message Queue mismatch", len(q.Data), len(tt.fields.data))
				t.Fail()
			}
		})
	}
}

func TestQueue_Subscribe(t *testing.T) {
	type fields struct {
		data        []string
		subscribers map[string]int
	}
	type args struct {
		subsid string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "TestCase1",
			fields: fields{
				data:        []string{"message1", "message2", "message3"},
				subscribers: map[string]int{"subs1": 1, "subs2": 2},
			},
			args: args{"subs1"},
			want: "message2",
		},
		{
			name: "TestCase2",
			fields: fields{
				data:        []string{"m1", "m2"},
				subscribers: map[string]int{"subs1": 1, "subs2": 2},
			},
			args: args{"subs2"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := &MyQueue{
				Data:        tt.fields.data,
				Subscribers: tt.fields.subscribers,
			}
			if got := q.Subscribe(tt.args.subsid); got != tt.want {
				t.Errorf("Subscribe() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMyQueue_Publish(t *testing.T) {
	type fields struct {
		Data        []string
		Subscribers map[string]int
	}
	type args struct {
		message string
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
			q := &MyQueue{
				Data:        tt.fields.Data,
				Subscribers: tt.fields.Subscribers,
			}
			q.Publish(tt.args.message)
		})
	}
}
