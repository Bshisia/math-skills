// calculations/calculations_test.go
package calculations

import (
	"os"
	"reflect"
	"testing"
)

// Helper function to create temporary files for testing
func createTempFile(t *testing.T, content string) *os.File {
	t.Helper()
	tmpfile, err := os.CreateTemp("", "example")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := tmpfile.Write([]byte(content)); err != nil {
		t.Fatal(err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}
	return tmpfile
}

func TestReadFile(t *testing.T) {
	tests := []struct {
		name    string
		content string
		want    []float64
		wantErr bool
	}{
		{
			name:    "Valid data",
			content: "1.1\n2.2\n3.3",
			want:    []float64{1.1, 2.2, 3.3},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a temporary file with the specified content
			tmpfile := createTempFile(t, tt.content)
			defer os.Remove(tmpfile.Name())

			// Call the ReadFile function with the temporary file's name
			got, err := ReadFile(tmpfile.Name())
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
