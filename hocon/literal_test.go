package hocon

import (
	"reflect"
	"testing"
)

func TestHoconLiteral_GetArray(t *testing.T) {
	type fields struct {
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []*HoconValue
		wantErr bool
	}{
		{name: "empty", fields: fields{""}, want: nil, wantErr: true},
		{name: "text", fields: fields{"abc"}, want: nil, wantErr: true},
		{name: "array", fields: fields{"[a,b,c,]"}, want: nil, wantErr: true},
		{name: "int", fields: fields{"123"}, want: nil, wantErr: true},
		{name: "float", fields: fields{"123.456"}, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &HoconLiteral{
				value: tt.fields.value,
			}
			got, err := p.GetArray()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArray() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetArray() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHoconLiteral_GetString(t *testing.T) {
	type fields struct {
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{name: "empty", fields: fields{""}, want: "", wantErr: false},
		{name: "text", fields: fields{"abc"}, want: "abc", wantErr: false},
		{name: "array", fields: fields{"[a,b,c,]"}, want: "[a,b,c,]", wantErr: false},
		{name: "int", fields: fields{"123"}, want: "123", wantErr: false},
		{name: "float", fields: fields{"123.456"}, want: "123.456", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &HoconLiteral{
				value: tt.fields.value,
			}
			got, err := p.GetString()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetString() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHoconLiteral_IsArray(t *testing.T) {
	type fields struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "empty", fields: fields{""}, want: false},
		{name: "emptyArray", fields: fields{"[]"}, want: false},
		{name: "array", fields: fields{"[a,b,c,]"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &HoconLiteral{
				value: tt.fields.value,
			}
			if got := p.IsArray(); got != tt.want {
				t.Errorf("IsArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHoconLiteral_IsString(t *testing.T) {
	type fields struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{name: "empty", fields: fields{""}, want: true},
		{name: "text", fields: fields{"abc"}, want: true},
		{name: "array", fields: fields{"[a,b,c,]"}, want: true},
		{name: "int", fields: fields{"123"}, want: true},
		{name: "float", fields: fields{"123.456"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &HoconLiteral{
				value: tt.fields.value,
			}
			if got := p.IsString(); got != tt.want {
				t.Errorf("IsString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHoconLiteral_String(t *testing.T) {
	type fields struct {
		value string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{name: "empty", fields: fields{""}, want: ""},
		{name: "text", fields: fields{"abc"}, want: "abc"},
		{name: "array", fields: fields{"[a,b,c,]"}, want: "[a,b,c,]"},
		{name: "int", fields: fields{"123"}, want: "123"},
		{name: "float", fields: fields{"123.456"}, want: "123.456"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &HoconLiteral{
				value: tt.fields.value,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewHoconLiteral(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want *HoconLiteral
	}{
		{name: "simple", args: args{value: "string"}, want: NewHoconLiteral("string")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHoconLiteral(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHoconLiteral() = %v, want %v", got, tt.want)
			}
		})
	}
}