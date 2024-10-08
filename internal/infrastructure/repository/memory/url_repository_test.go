package memory

import (
	"reflect"
	"testing"
	"time"

	"github.com/golanglowell/quick-link/internal/domain"
)

func fakeURL() *domain.URL {
	return &domain.URL{
		ID:        "1",
		LongURL:   "https://fakerjs.dev/api/internet",
		ShortCode: "ABC123",
		CreatedAt: time.Date(2024, time.April, 1, 1, 1, 1, 1, time.Local),
		Clicks:    0,
	}
}

func TestURLRepository_Save(t *testing.T) {
	type args struct {
		url *domain.URL
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "save nil url",
			args:    args{},
			wantErr: true,
		},
		{
			name:    "save url",
			args:    args{fakeURL()},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewURLRepository()
			if err := u.Save(tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("URLRepository.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestURLRepository_FindByShortCode(t *testing.T) {
	type args struct {
		shortCode string
	}
	tests := []struct {
		name    string
		args    args
		want    *domain.URL
		wantErr bool
	}{
		{
			name:    "find empty short code",
			args:    args{},
			want:    nil,
			wantErr: true,
		},
		{
			name: "find non-existent short code",
			args: args{
				shortCode: "does not exist",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "find short code",
			args: args{
				shortCode: fakeURL().ShortCode,
			},
			want:    fakeURL(),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewURLRepository()

			if tt.name == "find short code" {
				u = func() *URLRepository {
					u.urls[fakeURL().ShortCode] = fakeURL()
					return u
				}()
			}

			got, err := u.FindByShortCode(tt.args.shortCode)
			if (err != nil) != tt.wantErr {
				t.Errorf("URLRepository.FindByShortCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("URLRepository.FindByShortCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
