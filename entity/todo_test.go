package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTodoID(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    TodoID
		wantErr bool
	}{
		{
			name: "return TodoID",
			args: args{
				value: "test_id",
			},
			want:    "test_id",
			wantErr: false,
		},
		{
			name: "return error when value is invalid",
			args: args{
				value: "",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTodoID(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTodoID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got, "ParseTodoID() = %v, want %v", got, tt.want)
		})
	}
}

func TestTodoID_Validate(t *testing.T) {
	tests := []struct {
		name    string
		i       TodoID
		wantErr bool
	}{
		{
			name:    "return nil when value is valid",
			i:       "test_id",
			wantErr: false,
		},
		{
			name:    "return error when value is empty",
			i:       "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.i.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("TodoID.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTodoID_String(t *testing.T) {
	tests := []struct {
		name string
		i    TodoID
		want string
	}{
		{
			name: "returned string value",
			i:    "test_id",
			want: "test_id",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.i.String()
			assert.Equal(t, tt.want, got, "TodoID.String() = %v, want %v", got, tt.want)
		})
	}
}

func TestTodoID_isEmpty(t *testing.T) {
	tests := []struct {
		name string
		i    TodoID
		want bool
	}{
		{
			name: "return true when value is empty",
			i:    "",
			want: true,
		},
		{
			name: "return false when value is not empty",
			i:    "test_id",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.i.isEmpty()
			assert.Equal(t, tt.want, got, "TodoID.isEmpty() = %v, want %v", got, tt.want)
		})
	}
}

func TestNewTodo(t *testing.T) {
	type args struct {
		id   TodoID
		text string
		done bool
	}
	tests := []struct {
		name    string
		args    args
		want    *Todo
		wantErr bool
	}{
		{
			name: "return Todo",
			args: args{
				id:   "test_id",
				text: "my todo",
				done: true,
			},
			want: &Todo{
				id:   "test_id",
				text: "my todo",
				done: true,
			},
			wantErr: false,
		},
		{
			name: "return error when todo is invalid",
			args: args{
				id:   "",
				text: "my todo",
				done: false,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTodo(tt.args.id, tt.args.text, tt.args.done)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTodo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			assert.Equal(t, tt.want, got, "NewTodo() = %v, want %v", got, tt.want)
		})
	}
}

func TestTodo_Validate(t *testing.T) {
	type fields struct {
		id   TodoID
		text string
		done bool
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "return nil when todo is valid",
			fields: fields{
				id:   "test_id",
				text: "my todo",
				done: true,
			},
			wantErr: false,
		},
		{
			name: "return error when id is invalid",
			fields: fields{
				id:   "",
				text: "my todo",
				done: true,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Todo{
				id:   tt.fields.id,
				text: tt.fields.text,
				done: tt.fields.done,
			}
			if err := tr.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Todo.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTodo_ID(t *testing.T) {
	type fields struct {
		id   TodoID
		text string
		done bool
	}
	tests := []struct {
		name   string
		fields *fields
		want   TodoID
	}{
		{
			name: "return todo ID",
			fields: &fields{
				id:   "test_id",
				text: "my todo",
				done: false,
			},
			want: "test_id",
		},
		{
			name:   "return empty when todo is nil",
			fields: nil,
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tr *Todo
			if tt.fields != nil {
				tr = &Todo{
					id:   tt.fields.id,
					text: tt.fields.text,
					done: tt.fields.done,
				}
			}
			got := tr.ID()
			assert.Equal(t, tt.want, got, "Todo.ID() = %v, want %v", got, tt.want)
		})
	}
}

func TestTodo_Text(t *testing.T) {
	type fields struct {
		id   TodoID
		text string
		done bool
	}
	tests := []struct {
		name   string
		fields *fields
		want   string
	}{
		{
			name: "return todo text",
			fields: &fields{
				id:   "test_id",
				text: "my todo",
				done: false,
			},
			want: "my todo",
		},
		{
			name:   "return empty when todo is nil",
			fields: nil,
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tr *Todo
			if tt.fields != nil {
				tr = &Todo{
					id:   tt.fields.id,
					text: tt.fields.text,
					done: tt.fields.done,
				}
			}
			got := tr.Text()
			assert.Equal(t, tt.want, got, "Todo.Text() = %v, want %v", got, tt.want)
		})
	}
}

func TestTodo_Done(t *testing.T) {
	type fields struct {
		id   TodoID
		text string
		done bool
	}
	tests := []struct {
		name   string
		fields *fields
		want   bool
	}{
		{
			name: "return todo text",
			fields: &fields{
				id:   "test_id",
				text: "my todo",
				done: true,
			},
			want: true,
		},
		{
			name:   "return empty when todo is nil",
			fields: nil,
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tr *Todo
			if tt.fields != nil {
				tr = &Todo{
					id:   tt.fields.id,
					text: tt.fields.text,
					done: tt.fields.done,
				}
			}
			got := tr.Done()
			assert.Equal(t, tt.want, got, "Todo.Done() = %v, want %v", got, tt.want)
		})
	}
}
