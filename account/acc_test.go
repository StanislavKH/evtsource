package account

import (
	"reflect"
	"testing"
)

func TestNewFromEvents(t *testing.T) {
	a := New(321, "Bob")
	a.Deposit(100)

	c := NewFromEvents(a.EventsList())

	type args struct {
		events []Event
	}
	tests := []struct {
		name string
		args args
		want *Account
	}{
		{
			name: "create account from events",
			args: args{events: a.EventsList()},
			want: c,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFromEvents(tt.args.events); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFromEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	c := New(321, "Bob")

	type args struct {
		id    ID
		owner Owner
	}
	tests := []struct {
		name string
		args args
		want *Account
	}{
		{
			name: "create new account.",
			args: args{id: 321, owner: "Bob"},
			want: c,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.id, tt.args.owner); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_ID(t *testing.T) {
	type fields struct {
		id      ID
		owner   Owner
		balance Amount
		updates []Event
		version int
	}
	tests := []struct {
		name   string
		fields fields
		want   ID
	}{
		{
			name: "get ID",
			fields: fields{id: 321},
			want: 321,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Account{
				id:      tt.fields.id,
				owner:   tt.fields.owner,
				balance: tt.fields.balance,
				updates: tt.fields.updates,
				version: tt.fields.version,
			}
			if got := a.ID(); got != tt.want {
				t.Errorf("Account.ID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_Owner(t *testing.T) {
	type fields struct {
		id      ID
		owner   Owner
		balance Amount
		updates []Event
		version int
	}
	tests := []struct {
		name   string
		fields fields
		want   Owner
	}{
		{
			name: "get owner",
			fields: fields{owner: "Bob"},
			want: "Bob",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Account{
				id:      tt.fields.id,
				owner:   tt.fields.owner,
				balance: tt.fields.balance,
				updates: tt.fields.updates,
				version: tt.fields.version,
			}
			if got := a.Owner(); got != tt.want {
				t.Errorf("Account.Owner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_Balance(t *testing.T) {
	type fields struct {
		id      ID
		owner   Owner
		balance Amount
		updates []Event
		version int
	}
	tests := []struct {
		name   string
		fields fields
		want   Amount
	}{
		{
			name: "get Balance",
			fields: fields{balance: 155},
			want: 155,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Account{
				id:      tt.fields.id,
				owner:   tt.fields.owner,
				balance: tt.fields.balance,
				updates: tt.fields.updates,
				version: tt.fields.version,
			}
			if got := a.Balance(); got != tt.want {
				t.Errorf("Account.Balance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_UpdateOwner(t *testing.T) {
	type fields struct {
		id      ID
		owner   Owner
		balance Amount
		updates []Event
		version int
	}
	type args struct {
		owner Owner
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Owner
		wantErr bool
	}{
		{
			name: "update owner",
			fields: fields{id: 321, owner: "Bob"},
			args: args{owner: "Wilma"},
			want: "Wilma",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				id:      tt.fields.id,
				owner:   tt.fields.owner,
				balance: tt.fields.balance,
				updates: tt.fields.updates,
				version: tt.fields.version,
			}
			if err := a.UpdateOwner(tt.args.owner); (err != nil) != tt.wantErr {
				t.Errorf("Account.UpdateOwner() error = %v, wantErr %v", err, tt.wantErr)
			}
			if got := a.Owner(); got != tt.want {
				t.Errorf("Account.Owner() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccount_Deposit(t *testing.T) {
	type fields struct {
		id      ID
		owner   Owner
		balance Amount
		updates []Event
		version int
	}
	type args struct {
		amount Amount
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "add deposit correct",
			fields: fields{id: 321, owner: "Bob"},
			args: args{amount: 173},
			wantErr: false,
		},
		{
			name: "add deposit incorrect",
			fields: fields{id: 321, owner: "Bob"},
			args: args{amount: -25},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				id:      tt.fields.id,
				owner:   tt.fields.owner,
				balance: tt.fields.balance,
				updates: tt.fields.updates,
				version: tt.fields.version,
			}
			if err := a.Deposit(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("Account.Deposit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAccount_Withdrawal(t *testing.T) {
	type fields struct {
		id      ID
		owner   Owner
		balance Amount
		updates []Event
		version int
	}
	type args struct {
		amount Amount
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "withdrawal correct",
			fields: fields{id: 321, owner: "Bob", balance: 50},
			args: args{amount: 50},
			wantErr: false,
		},
		{
			name: "withdrawal incorrect",
			fields: fields{id: 321, owner: "Bob", balance: 50},
			args: args{amount: 500},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Account{
				id:      tt.fields.id,
				owner:   tt.fields.owner,
				balance: tt.fields.balance,
				updates: tt.fields.updates,
				version: tt.fields.version,
			}
			if err := a.Withdrawal(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("Account.Withdrawal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

