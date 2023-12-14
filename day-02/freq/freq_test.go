package main

import (
	"io"
	"reflect"
	"testing"
)

func Test_wordFrequency(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := wordFrequency(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("wordFrequency() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wordFrequency() = %v, want %v", got, tt.want)
			}
		})
	}
}
