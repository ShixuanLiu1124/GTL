package Vector

import (
	"reflect"
	"testing"
)

func TestNewUnsafeVector(t *testing.T) {
	type args struct {
		maxSize int
		values  []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *unsafeVector
		wantErr bool
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUnsafeVector(tt.args.maxSize, tt.args.values...)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUnsafeVector() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUnsafeVector() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewUnsafeVectorWithSlice(t *testing.T) {
	type args struct {
		maxSize int
		values  []interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    *unsafeVector
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUnsafeVectorWithSlice(tt.args.maxSize, tt.args.values)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUnsafeVectorWithSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUnsafeVectorWithSlice() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unsafeVector_At(t *testing.T) {
	type fields struct {
		s       []interface{}
		maxSize int
	}
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &unsafeVector{
				s:       tt.fields.s,
				maxSize: tt.fields.maxSize,
			}
			got, err := v.At(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("At() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("At() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unsafeVector_Clear(t *testing.T) {
	type fields struct {
		s       []interface{}
		maxSize int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &unsafeVector{
				s:       tt.fields.s,
				maxSize: tt.fields.maxSize,
			}
			v.Clear()
		})
	}
}

func Test_unsafeVector_CopyFromSlice(t *testing.T) {
	type fields struct {
		s       []interface{}
		maxSize int
	}
	type args struct {
		values []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &unsafeVector{
				s:       tt.fields.s,
				maxSize: tt.fields.maxSize,
			}
			if err := v.CatFromSlice(tt.args.values); (err != nil) != tt.wantErr {
				t.Errorf("CatFromSlice() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_unsafeVector_Empty(t *testing.T) {
	type fields struct {
		s       []interface{}
		maxSize int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &unsafeVector{
				s:       tt.fields.s,
				maxSize: tt.fields.maxSize,
			}
			if got := v.Empty(); got != tt.want {
				t.Errorf("Empty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unsafeVector_Fill(t *testing.T) {
	type fields struct {
		s       []interface{}
		maxSize int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &unsafeVector{
				s:       tt.fields.s,
				maxSize: tt.fields.maxSize,
			}
			if got := v.Fill(); got != tt.want {
				t.Errorf("Fill() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unsafeVector_Find(t *testing.T) {
	type fields struct {
		s       []interface{}
		maxSize int
	}
	type args struct {
		value interface{}
		less  func(interface{}, interface{}) bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &unsafeVector{
				s:       tt.fields.s,
				maxSize: tt.fields.maxSize,
			}
			if got := v.Find(tt.args.value, tt.args.less); got != tt.want {
				t.Errorf("Find() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unsafeVector_MarshalJSON(t *testing.T) {
	type fields struct {
		s       []interface{}
		maxSize int
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &unsafeVector{
				s:       tt.fields.s,
				maxSize: tt.fields.maxSize,
			}
			got, err := v.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unsafeVector_MaxSize(t *testing.T) {
	type fields struct {
		s       []interface{}
		maxSize int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &unsafeVector{
				s:       tt.fields.s,
				maxSize: tt.fields.maxSize,
			}
			if got := v.MaxSize(); got != tt.want {
				t.Errorf("MaxSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unsafeVector_PopBack(t *testing.T) {
	type fields struct {
		s       []interface{}
		maxSize int
	}
	tests := []struct {
		name    string
		fields  fields
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &unsafeVector{
				s:       tt.fields.s,
				maxSize: tt.fields.maxSize,
			}
			got, err := v.PopBack()
			if (err != nil) != tt.wantErr {
				t.Errorf("PopBack() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PopBack() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unsafeVector_PushBack(t *testing.T) {
	type fields struct {
		s       []interface{}
		maxSize int
	}
	type args struct {
		value interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &unsafeVector{
				s:       tt.fields.s,
				maxSize: tt.fields.maxSize,
			}
			if err := v.PushBack(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("PushBack() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_unsafeVector_Remove(t *testing.T) {
	type fields struct {
		s       []interface{}
		maxSize int
	}
	type args struct {
		start int
		end   int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &unsafeVector{
				s:       tt.fields.s,
				maxSize: tt.fields.maxSize,
			}
			if err := v.Remove(tt.args.start, tt.args.end); (err != nil) != tt.wantErr {
				t.Errorf("Remove() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_unsafeVector_SetMaxSize(t *testing.T) {
	type fields struct {
		s       []interface{}
		maxSize int
	}
	type args struct {
		maxSize int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &unsafeVector{
				s:       tt.fields.s,
				maxSize: tt.fields.maxSize,
			}
			if err := v.SetMaxSize(tt.args.maxSize); (err != nil) != tt.wantErr {
				t.Errorf("SetMaxSize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_unsafeVector_Size(t *testing.T) {
	type fields struct {
		s       []interface{}
		maxSize int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &unsafeVector{
				s:       tt.fields.s,
				maxSize: tt.fields.maxSize,
			}
			if got := v.Size(); got != tt.want {
				t.Errorf("Size() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unsafeVector_String(t *testing.T) {
	type fields struct {
		s       []interface{}
		maxSize int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &unsafeVector{
				s:       tt.fields.s,
				maxSize: tt.fields.maxSize,
			}
			if got := v.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unsafeVector_ToSlice(t *testing.T) {
	type fields struct {
		s       []interface{}
		maxSize int
	}
	tests := []struct {
		name   string
		fields fields
		want   []interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &unsafeVector{
				s:       tt.fields.s,
				maxSize: tt.fields.maxSize,
			}
			if got := v.ToSlice(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unsafeVector_UnmarshalJSON(t *testing.T) {
	type fields struct {
		s       []interface{}
		maxSize int
	}
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &unsafeVector{
				s:       tt.fields.s,
				maxSize: tt.fields.maxSize,
			}
			if err := v.UnmarshalJSON(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
