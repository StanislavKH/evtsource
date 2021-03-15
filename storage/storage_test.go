package storage

import (
	"reflect"
	"testing"

	"github.com/StanislavKH/evtsource/account"
)

func TestStorage_GetAccount(t *testing.T) {
	a := account.New(321, "Wilma")
	m := make(map[account.ID]*account.Account)
	m[a.ID()] = a

	type fields struct {
		Accounts map[account.ID]*account.Account
	}
	type args struct {
		id    account.ID
		owner account.Owner
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *account.Account
		wantErr bool
	}{
		{
			name:    "get create account correct.",
			fields:  fields{Accounts: make(map[account.ID]*account.Account)},
			args:    args{id: 321, owner: "Wilma"},
			want:    a,
			wantErr: false,
		},
		{
			name:    "get create account incorrect.",
			fields:  fields{Accounts: m},
			args:    args{id: 321, owner: "Wilma"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				Accounts: tt.fields.Accounts,
			}
			got, err := s.GetAccount(tt.args.id, tt.args.owner)
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.GetAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Storage.GetAccount() = %v, want %v", got, tt.want)
			}
		})
	}
}
