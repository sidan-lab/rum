package data_test

import (
	"encoding/json"
	"testing"

	"github.com/sidan-lab/rum/common/data"
)

func TestPaymentPubKeyHash(t *testing.T) {
	correctPaymentPubKeyHash := `{"bytes":"hello"}`
	data := data.NewPaymentPubKeyHash("hello")
	dataJSON, _ := json.Marshal(data)
	if string(dataJSON) != correctPaymentPubKeyHash {
		t.Errorf("PaymentPubKeyHash incorrect")
	}
}

func TestPubKeyHash(t *testing.T) {
	correctPubKeyHash := `{"bytes":"hello"}`
	data := data.NewPubKeyHash("hello")
	dataJSON, _ := json.Marshal(data)
	if string(dataJSON) != correctPubKeyHash {
		t.Errorf("PubKeyHash incorrect")
	}
}

func TestMaybeStakingHash(t *testing.T) {
	correctMaybeStakingHash := `{"constructor":0,"fields":[{"constructor":0,"fields":[{"constructor":0,"fields":[{"bytes":"hello"}]}]}]}`
	stakeCredential := "hello"
	data := data.NewMaybeStakingHash(&stakeCredential, false)
	dataJSON, _ := json.Marshal(data)
	if string(dataJSON) != correctMaybeStakingHash {
		t.Errorf("MaybeStakingHash incorrect")
	}
}

func TestPubKeyAddress(t *testing.T) {
	correctPubKeyAddress := `{"constructor":0,"fields":[{"constructor":0,"fields":[{"bytes":"8f2ac4b2a57a90feb7717c7361c7043af6c3646e9db2b0e616482f73"}]},{"constructor":0,"fields":[{"constructor":0,"fields":[{"constructor":0,"fields":[{"bytes":"039506b8e57e150bb66f6134f3264d50c3b70ce44d052f4485cf388f"}]}]}]}]}`
	bytes := "8f2ac4b2a57a90feb7717c7361c7043af6c3646e9db2b0e616482f73"
	stakeCredential := "039506b8e57e150bb66f6134f3264d50c3b70ce44d052f4485cf388f"
	data := data.NewPubKeyAddress(bytes, &stakeCredential, false)
	dataJSON, _ := json.Marshal(data)
	if string(dataJSON) != correctPubKeyAddress {
		t.Errorf("PubKeyAddress incorrect")
	}
}

func TestScriptAddress(t *testing.T) {
	correctScriptAddress := `{"constructor":0,"fields":[{"constructor":1,"fields":[{"bytes":"hello"}]},{"constructor":1,"fields":[]}]}`
	data := data.NewScriptAddress("hello", nil, false)
	dataJSON, _ := json.Marshal(data)
	if string(dataJSON) != correctScriptAddress {
		t.Errorf("ScriptAddress incorrect")
	}
}

func TestScriptAddressScriptStakeKey(t *testing.T) {
	correctScriptAddress := `{"constructor":0,"fields":[{"constructor":1,"fields":[{"bytes":"8f2ac4b2a57a90feb7717c7361c7043af6c3646e9db2b0e616482f73"}]},{"constructor":0,"fields":[{"constructor":0,"fields":[{"constructor":1,"fields":[{"bytes":"039506b8e57e150bb66f6134f3264d50c3b70ce44d052f4485cf388f"}]}]}]}]}`
	stakeCredential := "039506b8e57e150bb66f6134f3264d50c3b70ce44d052f4485cf388f"
	data := data.NewScriptAddress("8f2ac4b2a57a90feb7717c7361c7043af6c3646e9db2b0e616482f73", &stakeCredential, true)
	dataJSON, _ := json.Marshal(data)
	if string(dataJSON) != correctScriptAddress {
		t.Errorf("ScriptAddress incorrect")
	}
}
