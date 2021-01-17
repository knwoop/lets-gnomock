package service

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/knwoop/lets-gnomock/src/testutils"
)

func TestUserService_Create(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		init    func(*UserService)
		arg     string
		want    string
		wantErr bool
	}{
		"create user": {
			init:    func(_ *UserService) {},
			arg:     "foo",
			want:    "foo",
			wantErr: false,
		},
		"depulicated username": {
			init: func(s *UserService) {
				ctx := context.Background()
				s.Create(ctx, "foo")
			},
			arg:     "foo",
			want:    "",
			wantErr: true,
		},
	}

	for name, tt := range tests {
		tt := tt
		ctx := context.Background()
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			dbClient, err := testutils.NewDBTestClient()
			if err != nil {
				t.Fatal(err)
			}
			defer dbClient.Close()

			s, err := NewUserService(dbClient.DB)
			if err != nil {
				t.Fatal(err)
			}
			tt.init(s)
			got, err := s.Create(ctx, tt.arg)
			if (err != nil) != tt.wantErr {
				t.Fatalf("failed to create a user: %s", err.Error())
			}

			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Fatalf("\n (-want, +got)\n%s", diff)
			}
		})
	}

}
