package transactions

import (
	"reflect"
	"testing"
	"time"

	"github.com/bvnk/bank/configuration"
	geo "github.com/paulmach/go.geo"
	"github.com/shopspring/decimal"
)

func TestLoadConfiguration(t *testing.T) {
	// Load app config
	_, err := configuration.LoadConfig()
	if err != nil {
		t.Errorf("loadDatabase does not pass. Configuration does not load, looking for %v, got %v", nil, err)
	}
}

func TestSavePainTransaction(t *testing.T) {
	config, _ := configuration.LoadConfig()
	SetConfig(&config)

	sender := AccountHolder{"accountNumSender", "bankNumSender"}
	receiver := AccountHolder{"accountNumReceiver", "bankNumReceiver"}
	p := geo.NewPoint(42.25, 120.2)
	trans := PAINTrans{1, 101, sender, receiver, decimal.NewFromFloat(0.), decimal.NewFromFloat(0.), *p, "Test desc", "approved", 123123}

	id, err := savePainTransaction(trans)
	if err != nil {
		t.Errorf("DoSavePainTransaction does not pass. Looking for %v, got %v", nil, err)
	}

	if reflect.TypeOf(id).Kind() != reflect.Int64 {
		t.Errorf("DoSavePainTransaction does not pass. Expected integer return. Looking for %v, got %v", "int64", reflect.TypeOf(id).Kind())
	}

	err = removePainTransaction(trans)
	if err != nil {
		t.Errorf("DoDeleteAccount does not pass. Looking for %v, got %v", nil, err)
	}
}

func BenchmarkSavePainTransaction(b *testing.B) {
	config, _ := configuration.LoadConfig()
	SetConfig(&config)

	for n := 0; n < b.N; n++ {
		sender := AccountHolder{"accountNumSender", "bankNumSender"}
		receiver := AccountHolder{"accountNumReceiver", "bankNumReceiver"}
		p := geo.NewPoint(42.25, 120.2)
		trans := PAINTrans{1, 101, sender, receiver, decimal.NewFromFloat(0.), decimal.NewFromFloat(0.), *p, "Test desc", "approved", 123123}

		_, _ = savePainTransaction(trans)
		_ = removePainTransaction(trans)
	}
}

func TestUpdateHoldingAccount(t *testing.T) {
	config, _ := configuration.LoadConfig()
	SetConfig(&config)

	ti := time.Now()
	sqlTime := int32(ti.Unix())

	err := updateBankHoldingAccount(decimal.NewFromFloat(0.), sqlTime)
	if err != nil {
		t.Errorf("DoUpdateHoldingAccount does not pass. Looking for %v, got %v", nil, err)
	}
}

func BenchmarkUpdateHoldingAccount(b *testing.B) {
	config, _ := configuration.LoadConfig()
	SetConfig(&config)

	for n := 0; n < b.N; n++ {
		ti := time.Now()
		sqlTime := int32(ti.Unix())
		_ = updateBankHoldingAccount(decimal.NewFromFloat(0.), sqlTime)
	}
}

// All of the below need active accounts to be run
// Check balance
// Deposit
// Credit
