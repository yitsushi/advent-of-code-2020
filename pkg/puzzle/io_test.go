package puzzle

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

func Test_inputAsReader(t *testing.T) {
	type args struct {
		inputValue string
		inputFile  string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "from value",
			args: args{
				inputValue: "1234",
				inputFile:  "",
			},
			want:    "1234",
			wantErr: false,
		},
		{
			name: "from file",
			args: args{
				inputValue: "",
				inputFile:  "fixtures/input-file",
			},
			want:    "54321",
			wantErr: false,
		},
		{
			name: "no input",
			args: args{
				inputValue: "",
				inputFile:  "",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "prefer file over value",
			args: args{
				inputValue: "1234",
				inputFile:  "fixtures/input-file",
			},
			want:    "54321",
			wantErr: false,
		},
		{
			name: "file not found",
			args: args{
				inputValue: "",
				inputFile:  "fixtures/nothing",
			},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := inputAsReader(tt.args.inputValue, tt.args.inputFile)
			if (err != nil) != tt.wantErr {
				t.Errorf("inputAsReader() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if tt.wantErr {
				return
			}

			buf := new(strings.Builder)

			_, err = io.Copy(buf, got)
			if (err != nil) != tt.wantErr {
				t.Errorf("inputAsReader() error = %v, wantErr %v", err, tt.wantErr)

				return
			}

			if !reflect.DeepEqual(buf.String(), tt.want) {
				t.Errorf("inputAsReader() = %v, want %v", buf.String(), tt.want)
			}
		})
	}
}
