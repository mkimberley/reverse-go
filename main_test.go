package main

import "testing"

func TestReverseFile(t *testing.T) {
	type args struct {
		filname string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "hello",
			args:    args{
				filname: "hello.txt",
			},
			want:    "Hello, World!",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReverseFile(tt.args.filname)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReverseFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReverseFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
