package protocol

import (
	"bytes"
	"testing"
)

func TestAssetDefinition(t *testing.T) {
	// Create a randomized object
	initialMessage := AssetDefinition{}
	// AssetType (fixedchar)
	// fixedchar test not setup
	{
		{
			text := make([]byte, 0, 3)
			for i := uint64(0); i < 3; i++ {
				text = append(text, byte(65+i+0))
			}
			initialMessage.AssetType = string(text)
		}
	}

	// AssetAuthFlags (varbin)
	// varbin test not setup
	{
		initialMessage.AssetAuthFlags = make([]byte, 0, 8)
		for i := uint64(0); i < 8; i++ {
			initialMessage.AssetAuthFlags = append(initialMessage.AssetAuthFlags, byte(65+i+1))
		}
	}

	// TransfersPermitted (bool)
	// bool test not setup
	{
		initialMessage.TransfersPermitted = true
	}

	// TradeRestrictions (Polity[])
	// Polity[] test not setup
	{
		for i := 0; i < 5; i++ {
			var item [3]byte
			initialMessage.TradeRestrictions = append(initialMessage.TradeRestrictions, item)
		}
	}

	// EnforcementOrdersPermitted (bool)
	// bool test not setup
	{
		initialMessage.EnforcementOrdersPermitted = true
	}

	// VotingRights (bool)
	// bool test not setup
	{
		initialMessage.VotingRights = true
	}

	// VoteMultiplier (uint)
	// uint test not setup
	{
		// uint test not setup
	}

	// AdministrationProposal (bool)
	// bool test not setup
	{
		initialMessage.AdministrationProposal = true
	}

	// HolderProposal (bool)
	// bool test not setup
	{
		initialMessage.HolderProposal = true
	}

	// AssetModificationGovernance (uint)
	// uint test not setup
	{
		// uint test not setup
	}

	// TokenQty (uint)
	// uint test not setup
	{
		// uint test not setup
	}

	// AssetPayload (varbin)
	// varbin test not setup
	{
		initialMessage.AssetPayload = make([]byte, 0, 16)
		for i := uint64(0); i < 16; i++ {
			initialMessage.AssetPayload = append(initialMessage.AssetPayload, byte(65+i+11))
		}
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := AssetDefinition{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(initialEncoding, secondEncoding) {
		t.Errorf("Encoded value doesn't match.\ngot\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !cmp.Equal(&initialMessage, &decodedMessage) {
	// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// AssetType (fixedchar)
	if initialMessage.AssetType != decodedMessage.AssetType {
		t.Errorf("AssetType doesn't match : %s != %s", initialMessage.AssetType, decodedMessage.AssetType)
	}
	// AssetAuthFlags (varbin)
	if !bytes.Equal(initialMessage.AssetAuthFlags, decodedMessage.AssetAuthFlags) {
		t.Errorf("AssetAuthFlags doesn't match : %x != %x", initialMessage.AssetAuthFlags, decodedMessage.AssetAuthFlags)
	}
	// TransfersPermitted (bool)
	if initialMessage.TransfersPermitted != decodedMessage.TransfersPermitted {
		t.Errorf("TransfersPermitted doesn't match : %v != %v", initialMessage.TransfersPermitted, decodedMessage.TransfersPermitted)
	}
	// TradeRestrictions (Polity[])
	if len(initialMessage.TradeRestrictions) != len(decodedMessage.TradeRestrictions) {
		t.Errorf("TradeRestrictions lengths don't match : %d != %d", len(initialMessage.TradeRestrictions), len(decodedMessage.TradeRestrictions))
	}
	for i, value := range initialMessage.TradeRestrictions {
		if value != decodedMessage.TradeRestrictions[i] {
			t.Errorf("TradeRestrictions value %d doesn't match : %v != %v", i, value, decodedMessage.TradeRestrictions[i])
		}
	}
	// EnforcementOrdersPermitted (bool)
	if initialMessage.EnforcementOrdersPermitted != decodedMessage.EnforcementOrdersPermitted {
		t.Errorf("EnforcementOrdersPermitted doesn't match : %v != %v", initialMessage.EnforcementOrdersPermitted, decodedMessage.EnforcementOrdersPermitted)
	}
	// VotingRights (bool)
	if initialMessage.VotingRights != decodedMessage.VotingRights {
		t.Errorf("VotingRights doesn't match : %v != %v", initialMessage.VotingRights, decodedMessage.VotingRights)
	}
	// VoteMultiplier (uint)
	if initialMessage.VoteMultiplier != decodedMessage.VoteMultiplier {
		t.Errorf("VoteMultiplier doesn't match : %v != %v", initialMessage.VoteMultiplier, decodedMessage.VoteMultiplier)
	}
	// AdministrationProposal (bool)
	if initialMessage.AdministrationProposal != decodedMessage.AdministrationProposal {
		t.Errorf("AdministrationProposal doesn't match : %v != %v", initialMessage.AdministrationProposal, decodedMessage.AdministrationProposal)
	}
	// HolderProposal (bool)
	if initialMessage.HolderProposal != decodedMessage.HolderProposal {
		t.Errorf("HolderProposal doesn't match : %v != %v", initialMessage.HolderProposal, decodedMessage.HolderProposal)
	}
	// AssetModificationGovernance (uint)
	if initialMessage.AssetModificationGovernance != decodedMessage.AssetModificationGovernance {
		t.Errorf("AssetModificationGovernance doesn't match : %v != %v", initialMessage.AssetModificationGovernance, decodedMessage.AssetModificationGovernance)
	}
	// TokenQty (uint)
	if initialMessage.TokenQty != decodedMessage.TokenQty {
		t.Errorf("TokenQty doesn't match : %v != %v", initialMessage.TokenQty, decodedMessage.TokenQty)
	}
	// AssetPayload (varbin)
	if !bytes.Equal(initialMessage.AssetPayload, decodedMessage.AssetPayload) {
		t.Errorf("AssetPayload doesn't match : %x != %x", initialMessage.AssetPayload, decodedMessage.AssetPayload)
	}
}
func TestAssetCreation(t *testing.T) {
	// Create a randomized object
	initialMessage := AssetCreation{}
	// AssetType (fixedchar)
	// fixedchar test not setup
	{
		{
			text := make([]byte, 0, 3)
			for i := uint64(0); i < 3; i++ {
				text = append(text, byte(65+i+0))
			}
			initialMessage.AssetType = string(text)
		}
	}

	// AssetCode (AssetCode)
	// AssetCode test not setup
	{
		initialMessage.AssetCode = AssetCode{}
	}

	// AssetIndex (uint)
	// uint test not setup
	{
		// uint test not setup
	}

	// AssetAuthFlags (varbin)
	// varbin test not setup
	{
		initialMessage.AssetAuthFlags = make([]byte, 0, 8)
		for i := uint64(0); i < 8; i++ {
			initialMessage.AssetAuthFlags = append(initialMessage.AssetAuthFlags, byte(65+i+3))
		}
	}

	// TransfersPermitted (bool)
	// bool test not setup
	{
		initialMessage.TransfersPermitted = true
	}

	// TradeRestrictions (Polity[])
	// Polity[] test not setup
	{
		for i := 0; i < 5; i++ {
			var item [3]byte
			initialMessage.TradeRestrictions = append(initialMessage.TradeRestrictions, item)
		}
	}

	// EnforcementOrdersPermitted (bool)
	// bool test not setup
	{
		initialMessage.EnforcementOrdersPermitted = true
	}

	// VotingRights (bool)
	// bool test not setup
	{
		initialMessage.VotingRights = true
	}

	// VoteMultiplier (uint)
	// uint test not setup
	{
		// uint test not setup
	}

	// AdministrationProposal (bool)
	// bool test not setup
	{
		initialMessage.AdministrationProposal = true
	}

	// HolderProposal (bool)
	// bool test not setup
	{
		initialMessage.HolderProposal = true
	}

	// AssetModificationGovernance (uint)
	// uint test not setup
	{
		// uint test not setup
	}

	// TokenQty (uint)
	// uint test not setup
	{
		// uint test not setup
	}

	// AssetPayload (varbin)
	// varbin test not setup
	{
		initialMessage.AssetPayload = make([]byte, 0, 16)
		for i := uint64(0); i < 16; i++ {
			initialMessage.AssetPayload = append(initialMessage.AssetPayload, byte(65+i+13))
		}
	}

	// AssetRevision (uint)
	// uint test not setup
	{
		// uint test not setup
	}

	// Timestamp (Timestamp)
	// Timestamp test not setup
	{
		initialMessage.Timestamp = Timestamp{}
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := AssetCreation{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(initialEncoding, secondEncoding) {
		t.Errorf("Encoded value doesn't match.\ngot\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !cmp.Equal(&initialMessage, &decodedMessage) {
	// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// AssetType (fixedchar)
	if initialMessage.AssetType != decodedMessage.AssetType {
		t.Errorf("AssetType doesn't match : %s != %s", initialMessage.AssetType, decodedMessage.AssetType)
	}
	// AssetCode (AssetCode)
	// AssetCode test compare not setup
	// AssetIndex (uint)
	if initialMessage.AssetIndex != decodedMessage.AssetIndex {
		t.Errorf("AssetIndex doesn't match : %v != %v", initialMessage.AssetIndex, decodedMessage.AssetIndex)
	}
	// AssetAuthFlags (varbin)
	if !bytes.Equal(initialMessage.AssetAuthFlags, decodedMessage.AssetAuthFlags) {
		t.Errorf("AssetAuthFlags doesn't match : %x != %x", initialMessage.AssetAuthFlags, decodedMessage.AssetAuthFlags)
	}
	// TransfersPermitted (bool)
	if initialMessage.TransfersPermitted != decodedMessage.TransfersPermitted {
		t.Errorf("TransfersPermitted doesn't match : %v != %v", initialMessage.TransfersPermitted, decodedMessage.TransfersPermitted)
	}
	// TradeRestrictions (Polity[])
	if len(initialMessage.TradeRestrictions) != len(decodedMessage.TradeRestrictions) {
		t.Errorf("TradeRestrictions lengths don't match : %d != %d", len(initialMessage.TradeRestrictions), len(decodedMessage.TradeRestrictions))
	}
	for i, value := range initialMessage.TradeRestrictions {
		if value != decodedMessage.TradeRestrictions[i] {
			t.Errorf("TradeRestrictions value %d doesn't match : %v != %v", i, value, decodedMessage.TradeRestrictions[i])
		}
	}
	// EnforcementOrdersPermitted (bool)
	if initialMessage.EnforcementOrdersPermitted != decodedMessage.EnforcementOrdersPermitted {
		t.Errorf("EnforcementOrdersPermitted doesn't match : %v != %v", initialMessage.EnforcementOrdersPermitted, decodedMessage.EnforcementOrdersPermitted)
	}
	// VotingRights (bool)
	if initialMessage.VotingRights != decodedMessage.VotingRights {
		t.Errorf("VotingRights doesn't match : %v != %v", initialMessage.VotingRights, decodedMessage.VotingRights)
	}
	// VoteMultiplier (uint)
	if initialMessage.VoteMultiplier != decodedMessage.VoteMultiplier {
		t.Errorf("VoteMultiplier doesn't match : %v != %v", initialMessage.VoteMultiplier, decodedMessage.VoteMultiplier)
	}
	// AdministrationProposal (bool)
	if initialMessage.AdministrationProposal != decodedMessage.AdministrationProposal {
		t.Errorf("AdministrationProposal doesn't match : %v != %v", initialMessage.AdministrationProposal, decodedMessage.AdministrationProposal)
	}
	// HolderProposal (bool)
	if initialMessage.HolderProposal != decodedMessage.HolderProposal {
		t.Errorf("HolderProposal doesn't match : %v != %v", initialMessage.HolderProposal, decodedMessage.HolderProposal)
	}
	// AssetModificationGovernance (uint)
	if initialMessage.AssetModificationGovernance != decodedMessage.AssetModificationGovernance {
		t.Errorf("AssetModificationGovernance doesn't match : %v != %v", initialMessage.AssetModificationGovernance, decodedMessage.AssetModificationGovernance)
	}
	// TokenQty (uint)
	if initialMessage.TokenQty != decodedMessage.TokenQty {
		t.Errorf("TokenQty doesn't match : %v != %v", initialMessage.TokenQty, decodedMessage.TokenQty)
	}
	// AssetPayload (varbin)
	if !bytes.Equal(initialMessage.AssetPayload, decodedMessage.AssetPayload) {
		t.Errorf("AssetPayload doesn't match : %x != %x", initialMessage.AssetPayload, decodedMessage.AssetPayload)
	}
	// AssetRevision (uint)
	if initialMessage.AssetRevision != decodedMessage.AssetRevision {
		t.Errorf("AssetRevision doesn't match : %v != %v", initialMessage.AssetRevision, decodedMessage.AssetRevision)
	}
	// Timestamp (Timestamp)
	// Timestamp test compare not setup
}
func TestAssetModification(t *testing.T) {
	// Create a randomized object
	initialMessage := AssetModification{}
	// AssetType (fixedchar)
	// fixedchar test not setup
	{
		{
			text := make([]byte, 0, 3)
			for i := uint64(0); i < 3; i++ {
				text = append(text, byte(65+i+0))
			}
			initialMessage.AssetType = string(text)
		}
	}

	// AssetCode (AssetCode)
	// AssetCode test not setup
	{
		initialMessage.AssetCode = AssetCode{}
	}

	// AssetRevision (uint)
	// uint test not setup
	{
		// uint test not setup
	}

	// Amendments (Amendment[])
	// Amendment[] test not setup
	{
		for i := 0; i < 2; i++ {
			initialMessage.Amendments = append(initialMessage.Amendments, Amendment{})
		}
	}

	// RefTxID (TxId)
	// TxId test not setup
	{
		initialMessage.RefTxID = TxId{}
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := AssetModification{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(initialEncoding, secondEncoding) {
		t.Errorf("Encoded value doesn't match.\ngot\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !cmp.Equal(&initialMessage, &decodedMessage) {
	// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// AssetType (fixedchar)
	if initialMessage.AssetType != decodedMessage.AssetType {
		t.Errorf("AssetType doesn't match : %s != %s", initialMessage.AssetType, decodedMessage.AssetType)
	}
	// AssetCode (AssetCode)
	// AssetCode test compare not setup
	// AssetRevision (uint)
	if initialMessage.AssetRevision != decodedMessage.AssetRevision {
		t.Errorf("AssetRevision doesn't match : %v != %v", initialMessage.AssetRevision, decodedMessage.AssetRevision)
	}
	// Amendments (Amendment[])
	if len(initialMessage.Amendments) != len(decodedMessage.Amendments) {
		t.Errorf("Amendments lengths don't match : %d != %d", len(initialMessage.Amendments), len(decodedMessage.Amendments))
	}
	// RefTxID (TxId)
	// TxId test compare not setup
}
func TestContractOffer(t *testing.T) {
	// Create a randomized object
	initialMessage := ContractOffer{}
	// ContractName (varchar)
	// varchar test not setup
	{
		initialMessage.ContractName = "Text 0"
	}

	// BodyOfAgreementType (uint)
	// uint test not setup
	{
		// uint test not setup
	}

	// BodyOfAgreement (varbin)
	// varbin test not setup
	{
		initialMessage.BodyOfAgreement = make([]byte, 0, 32)
		for i := uint64(0); i < 32; i++ {
			initialMessage.BodyOfAgreement = append(initialMessage.BodyOfAgreement, byte(65+i+2))
		}
	}

	// ContractType (varchar)
	// varchar test not setup
	{
		initialMessage.ContractType = "Text 3"
	}

	// SupportingDocs (Document[])
	// Document[] test not setup
	{
		for i := 0; i < 2; i++ {
			initialMessage.SupportingDocs = append(initialMessage.SupportingDocs, Document{})
		}
	}

	// GoverningLaw (fixedchar)
	// fixedchar test not setup
	{
		{
			text := make([]byte, 0, 5)
			for i := uint64(0); i < 5; i++ {
				text = append(text, byte(65+i+5))
			}
			initialMessage.GoverningLaw = string(text)
		}
	}

	// Jurisdiction (fixedchar)
	// fixedchar test not setup
	{
		{
			text := make([]byte, 0, 5)
			for i := uint64(0); i < 5; i++ {
				text = append(text, byte(65+i+6))
			}
			initialMessage.Jurisdiction = string(text)
		}
	}

	// ContractExpiration (Timestamp)
	// Timestamp test not setup
	{
		initialMessage.ContractExpiration = Timestamp{}
	}

	// ContractURI (varchar)
	// varchar test not setup
	{
		initialMessage.ContractURI = "Text 8"
	}

	// Issuer (Entity)
	// Entity test not setup
	{
		initialMessage.Issuer = Entity{}
	}

	// IssuerLogoURL (varchar)
	// varchar test not setup
	{
		initialMessage.IssuerLogoURL = "Text 10"
	}

	// ContractOperatorIncluded (bool)
	// bool test not setup
	{
		initialMessage.ContractOperatorIncluded = true
	}

	// ContractOperator (Entity)
	if initialMessage.ContractOperatorIncluded {
		initialMessage.ContractOperator = Entity{}
	}

	// ContractAuthFlags (varbin)
	// varbin test not setup
	{
		initialMessage.ContractAuthFlags = make([]byte, 0, 16)
		for i := uint64(0); i < 16; i++ {
			initialMessage.ContractAuthFlags = append(initialMessage.ContractAuthFlags, byte(65+i+13))
		}
	}

	// ContractFee (uint)
	// uint test not setup
	{
		// uint test not setup
	}

	// VotingSystems (VotingSystem[])
	// VotingSystem[] test not setup
	{
		for i := 0; i < 2; i++ {
			initialMessage.VotingSystems = append(initialMessage.VotingSystems, VotingSystem{})
		}
	}

	// RestrictedQtyAssets (uint)
	// uint test not setup
	{
		// uint test not setup
	}

	// AdministrationProposal (bool)
	// bool test not setup
	{
		initialMessage.AdministrationProposal = true
	}

	// HolderProposal (bool)
	// bool test not setup
	{
		initialMessage.HolderProposal = true
	}

	// Oracles (Oracle[])
	// Oracle[] test not setup
	{
		for i := 0; i < 2; i++ {
			initialMessage.Oracles = append(initialMessage.Oracles, Oracle{})
		}
	}

	// MasterPKH (PublicKeyHash)
	// PublicKeyHash test not setup
	{
		initialMessage.MasterPKH = PublicKeyHash{}
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := ContractOffer{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(initialEncoding, secondEncoding) {
		t.Errorf("Encoded value doesn't match.\ngot\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !cmp.Equal(&initialMessage, &decodedMessage) {
	// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// ContractName (varchar)
	if initialMessage.ContractName != decodedMessage.ContractName {
		t.Errorf("ContractName doesn't match : %s != %s", initialMessage.ContractName, decodedMessage.ContractName)
	}
	// BodyOfAgreementType (uint)
	if initialMessage.BodyOfAgreementType != decodedMessage.BodyOfAgreementType {
		t.Errorf("BodyOfAgreementType doesn't match : %v != %v", initialMessage.BodyOfAgreementType, decodedMessage.BodyOfAgreementType)
	}
	// BodyOfAgreement (varbin)
	if !bytes.Equal(initialMessage.BodyOfAgreement, decodedMessage.BodyOfAgreement) {
		t.Errorf("BodyOfAgreement doesn't match : %x != %x", initialMessage.BodyOfAgreement, decodedMessage.BodyOfAgreement)
	}
	// ContractType (varchar)
	if initialMessage.ContractType != decodedMessage.ContractType {
		t.Errorf("ContractType doesn't match : %s != %s", initialMessage.ContractType, decodedMessage.ContractType)
	}
	// SupportingDocs (Document[])
	if len(initialMessage.SupportingDocs) != len(decodedMessage.SupportingDocs) {
		t.Errorf("SupportingDocs lengths don't match : %d != %d", len(initialMessage.SupportingDocs), len(decodedMessage.SupportingDocs))
	}
	// GoverningLaw (fixedchar)
	if initialMessage.GoverningLaw != decodedMessage.GoverningLaw {
		t.Errorf("GoverningLaw doesn't match : %s != %s", initialMessage.GoverningLaw, decodedMessage.GoverningLaw)
	}
	// Jurisdiction (fixedchar)
	if initialMessage.Jurisdiction != decodedMessage.Jurisdiction {
		t.Errorf("Jurisdiction doesn't match : %s != %s", initialMessage.Jurisdiction, decodedMessage.Jurisdiction)
	}
	// ContractExpiration (Timestamp)
	// Timestamp test compare not setup
	// ContractURI (varchar)
	if initialMessage.ContractURI != decodedMessage.ContractURI {
		t.Errorf("ContractURI doesn't match : %s != %s", initialMessage.ContractURI, decodedMessage.ContractURI)
	}
	// Issuer (Entity)
	// Entity test compare not setup
	// IssuerLogoURL (varchar)
	if initialMessage.IssuerLogoURL != decodedMessage.IssuerLogoURL {
		t.Errorf("IssuerLogoURL doesn't match : %s != %s", initialMessage.IssuerLogoURL, decodedMessage.IssuerLogoURL)
	}
	// ContractOperatorIncluded (bool)
	if initialMessage.ContractOperatorIncluded != decodedMessage.ContractOperatorIncluded {
		t.Errorf("ContractOperatorIncluded doesn't match : %v != %v", initialMessage.ContractOperatorIncluded, decodedMessage.ContractOperatorIncluded)
	}
	// ContractOperator (Entity)
	// Entity test compare not setup
	// ContractAuthFlags (varbin)
	if !bytes.Equal(initialMessage.ContractAuthFlags, decodedMessage.ContractAuthFlags) {
		t.Errorf("ContractAuthFlags doesn't match : %x != %x", initialMessage.ContractAuthFlags, decodedMessage.ContractAuthFlags)
	}
	// ContractFee (uint)
	if initialMessage.ContractFee != decodedMessage.ContractFee {
		t.Errorf("ContractFee doesn't match : %v != %v", initialMessage.ContractFee, decodedMessage.ContractFee)
	}
	// VotingSystems (VotingSystem[])
	if len(initialMessage.VotingSystems) != len(decodedMessage.VotingSystems) {
		t.Errorf("VotingSystems lengths don't match : %d != %d", len(initialMessage.VotingSystems), len(decodedMessage.VotingSystems))
	}
	// RestrictedQtyAssets (uint)
	if initialMessage.RestrictedQtyAssets != decodedMessage.RestrictedQtyAssets {
		t.Errorf("RestrictedQtyAssets doesn't match : %v != %v", initialMessage.RestrictedQtyAssets, decodedMessage.RestrictedQtyAssets)
	}
	// AdministrationProposal (bool)
	if initialMessage.AdministrationProposal != decodedMessage.AdministrationProposal {
		t.Errorf("AdministrationProposal doesn't match : %v != %v", initialMessage.AdministrationProposal, decodedMessage.AdministrationProposal)
	}
	// HolderProposal (bool)
	if initialMessage.HolderProposal != decodedMessage.HolderProposal {
		t.Errorf("HolderProposal doesn't match : %v != %v", initialMessage.HolderProposal, decodedMessage.HolderProposal)
	}
	// Oracles (Oracle[])
	if len(initialMessage.Oracles) != len(decodedMessage.Oracles) {
		t.Errorf("Oracles lengths don't match : %d != %d", len(initialMessage.Oracles), len(decodedMessage.Oracles))
	}
	// MasterPKH (PublicKeyHash)
	// PublicKeyHash test compare not setup
}
func TestContractFormation(t *testing.T) {
	// Create a randomized object
	initialMessage := ContractFormation{}
	// ContractName (varchar)
	// varchar test not setup
	{
		initialMessage.ContractName = "Text 0"
	}

	// BodyOfAgreementType (uint)
	// uint test not setup
	{
		// uint test not setup
	}

	// BodyOfAgreement (varbin)
	// varbin test not setup
	{
		initialMessage.BodyOfAgreement = make([]byte, 0, 32)
		for i := uint64(0); i < 32; i++ {
			initialMessage.BodyOfAgreement = append(initialMessage.BodyOfAgreement, byte(65+i+2))
		}
	}

	// ContractType (varchar)
	// varchar test not setup
	{
		initialMessage.ContractType = "Text 3"
	}

	// SupportingDocs (Document[])
	// Document[] test not setup
	{
		for i := 0; i < 2; i++ {
			initialMessage.SupportingDocs = append(initialMessage.SupportingDocs, Document{})
		}
	}

	// GoverningLaw (fixedchar)
	// fixedchar test not setup
	{
		{
			text := make([]byte, 0, 5)
			for i := uint64(0); i < 5; i++ {
				text = append(text, byte(65+i+5))
			}
			initialMessage.GoverningLaw = string(text)
		}
	}

	// Jurisdiction (fixedchar)
	// fixedchar test not setup
	{
		{
			text := make([]byte, 0, 5)
			for i := uint64(0); i < 5; i++ {
				text = append(text, byte(65+i+6))
			}
			initialMessage.Jurisdiction = string(text)
		}
	}

	// ContractExpiration (Timestamp)
	// Timestamp test not setup
	{
		initialMessage.ContractExpiration = Timestamp{}
	}

	// ContractURI (varchar)
	// varchar test not setup
	{
		initialMessage.ContractURI = "Text 8"
	}

	// Issuer (Entity)
	// Entity test not setup
	{
		initialMessage.Issuer = Entity{}
	}

	// IssuerLogoURL (varchar)
	// varchar test not setup
	{
		initialMessage.IssuerLogoURL = "Text 10"
	}

	// ContractOperatorIncluded (bool)
	// bool test not setup
	{
		initialMessage.ContractOperatorIncluded = true
	}

	// ContractOperator (Entity)
	if initialMessage.ContractOperatorIncluded {
		initialMessage.ContractOperator = Entity{}
	}

	// ContractAuthFlags (varbin)
	// varbin test not setup
	{
		initialMessage.ContractAuthFlags = make([]byte, 0, 16)
		for i := uint64(0); i < 16; i++ {
			initialMessage.ContractAuthFlags = append(initialMessage.ContractAuthFlags, byte(65+i+13))
		}
	}

	// ContractFee (uint)
	// uint test not setup
	{
		// uint test not setup
	}

	// VotingSystems (VotingSystem[])
	// VotingSystem[] test not setup
	{
		for i := 0; i < 2; i++ {
			initialMessage.VotingSystems = append(initialMessage.VotingSystems, VotingSystem{})
		}
	}

	// RestrictedQtyAssets (uint)
	// uint test not setup
	{
		// uint test not setup
	}

	// AdministrationProposal (bool)
	// bool test not setup
	{
		initialMessage.AdministrationProposal = true
	}

	// HolderProposal (bool)
	// bool test not setup
	{
		initialMessage.HolderProposal = true
	}

	// Oracles (Oracle[])
	// Oracle[] test not setup
	{
		for i := 0; i < 2; i++ {
			initialMessage.Oracles = append(initialMessage.Oracles, Oracle{})
		}
	}

	// MasterPKH (PublicKeyHash)
	// PublicKeyHash test not setup
	{
		initialMessage.MasterPKH = PublicKeyHash{}
	}

	// ContractRevision (uint)
	// uint test not setup
	{
		// uint test not setup
	}

	// Timestamp (Timestamp)
	// Timestamp test not setup
	{
		initialMessage.Timestamp = Timestamp{}
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := ContractFormation{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(initialEncoding, secondEncoding) {
		t.Errorf("Encoded value doesn't match.\ngot\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !cmp.Equal(&initialMessage, &decodedMessage) {
	// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// ContractName (varchar)
	if initialMessage.ContractName != decodedMessage.ContractName {
		t.Errorf("ContractName doesn't match : %s != %s", initialMessage.ContractName, decodedMessage.ContractName)
	}
	// BodyOfAgreementType (uint)
	if initialMessage.BodyOfAgreementType != decodedMessage.BodyOfAgreementType {
		t.Errorf("BodyOfAgreementType doesn't match : %v != %v", initialMessage.BodyOfAgreementType, decodedMessage.BodyOfAgreementType)
	}
	// BodyOfAgreement (varbin)
	if !bytes.Equal(initialMessage.BodyOfAgreement, decodedMessage.BodyOfAgreement) {
		t.Errorf("BodyOfAgreement doesn't match : %x != %x", initialMessage.BodyOfAgreement, decodedMessage.BodyOfAgreement)
	}
	// ContractType (varchar)
	if initialMessage.ContractType != decodedMessage.ContractType {
		t.Errorf("ContractType doesn't match : %s != %s", initialMessage.ContractType, decodedMessage.ContractType)
	}
	// SupportingDocs (Document[])
	if len(initialMessage.SupportingDocs) != len(decodedMessage.SupportingDocs) {
		t.Errorf("SupportingDocs lengths don't match : %d != %d", len(initialMessage.SupportingDocs), len(decodedMessage.SupportingDocs))
	}
	// GoverningLaw (fixedchar)
	if initialMessage.GoverningLaw != decodedMessage.GoverningLaw {
		t.Errorf("GoverningLaw doesn't match : %s != %s", initialMessage.GoverningLaw, decodedMessage.GoverningLaw)
	}
	// Jurisdiction (fixedchar)
	if initialMessage.Jurisdiction != decodedMessage.Jurisdiction {
		t.Errorf("Jurisdiction doesn't match : %s != %s", initialMessage.Jurisdiction, decodedMessage.Jurisdiction)
	}
	// ContractExpiration (Timestamp)
	// Timestamp test compare not setup
	// ContractURI (varchar)
	if initialMessage.ContractURI != decodedMessage.ContractURI {
		t.Errorf("ContractURI doesn't match : %s != %s", initialMessage.ContractURI, decodedMessage.ContractURI)
	}
	// Issuer (Entity)
	// Entity test compare not setup
	// IssuerLogoURL (varchar)
	if initialMessage.IssuerLogoURL != decodedMessage.IssuerLogoURL {
		t.Errorf("IssuerLogoURL doesn't match : %s != %s", initialMessage.IssuerLogoURL, decodedMessage.IssuerLogoURL)
	}
	// ContractOperatorIncluded (bool)
	if initialMessage.ContractOperatorIncluded != decodedMessage.ContractOperatorIncluded {
		t.Errorf("ContractOperatorIncluded doesn't match : %v != %v", initialMessage.ContractOperatorIncluded, decodedMessage.ContractOperatorIncluded)
	}
	// ContractOperator (Entity)
	// Entity test compare not setup
	// ContractAuthFlags (varbin)
	if !bytes.Equal(initialMessage.ContractAuthFlags, decodedMessage.ContractAuthFlags) {
		t.Errorf("ContractAuthFlags doesn't match : %x != %x", initialMessage.ContractAuthFlags, decodedMessage.ContractAuthFlags)
	}
	// ContractFee (uint)
	if initialMessage.ContractFee != decodedMessage.ContractFee {
		t.Errorf("ContractFee doesn't match : %v != %v", initialMessage.ContractFee, decodedMessage.ContractFee)
	}
	// VotingSystems (VotingSystem[])
	if len(initialMessage.VotingSystems) != len(decodedMessage.VotingSystems) {
		t.Errorf("VotingSystems lengths don't match : %d != %d", len(initialMessage.VotingSystems), len(decodedMessage.VotingSystems))
	}
	// RestrictedQtyAssets (uint)
	if initialMessage.RestrictedQtyAssets != decodedMessage.RestrictedQtyAssets {
		t.Errorf("RestrictedQtyAssets doesn't match : %v != %v", initialMessage.RestrictedQtyAssets, decodedMessage.RestrictedQtyAssets)
	}
	// AdministrationProposal (bool)
	if initialMessage.AdministrationProposal != decodedMessage.AdministrationProposal {
		t.Errorf("AdministrationProposal doesn't match : %v != %v", initialMessage.AdministrationProposal, decodedMessage.AdministrationProposal)
	}
	// HolderProposal (bool)
	if initialMessage.HolderProposal != decodedMessage.HolderProposal {
		t.Errorf("HolderProposal doesn't match : %v != %v", initialMessage.HolderProposal, decodedMessage.HolderProposal)
	}
	// Oracles (Oracle[])
	if len(initialMessage.Oracles) != len(decodedMessage.Oracles) {
		t.Errorf("Oracles lengths don't match : %d != %d", len(initialMessage.Oracles), len(decodedMessage.Oracles))
	}
	// MasterPKH (PublicKeyHash)
	// PublicKeyHash test compare not setup
	// ContractRevision (uint)
	if initialMessage.ContractRevision != decodedMessage.ContractRevision {
		t.Errorf("ContractRevision doesn't match : %v != %v", initialMessage.ContractRevision, decodedMessage.ContractRevision)
	}
	// Timestamp (Timestamp)
	// Timestamp test compare not setup
}
func TestContractAmendment(t *testing.T) {
	// Create a randomized object
	initialMessage := ContractAmendment{}
	// ChangeAdministrationAddress (bool)
	// bool test not setup
	{
		initialMessage.ChangeAdministrationAddress = true
	}

	// ChangeOperatorAddress (bool)
	// bool test not setup
	{
		initialMessage.ChangeOperatorAddress = true
	}

	// ContractRevision (uint)
	// uint test not setup
	{
		// uint test not setup
	}

	// Amendments (Amendment[])
	// Amendment[] test not setup
	{
		for i := 0; i < 2; i++ {
			initialMessage.Amendments = append(initialMessage.Amendments, Amendment{})
		}
	}

	// RefTxID (TxId)
	// TxId test not setup
	{
		initialMessage.RefTxID = TxId{}
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := ContractAmendment{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(initialEncoding, secondEncoding) {
		t.Errorf("Encoded value doesn't match.\ngot\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !cmp.Equal(&initialMessage, &decodedMessage) {
	// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// ChangeAdministrationAddress (bool)
	if initialMessage.ChangeAdministrationAddress != decodedMessage.ChangeAdministrationAddress {
		t.Errorf("ChangeAdministrationAddress doesn't match : %v != %v", initialMessage.ChangeAdministrationAddress, decodedMessage.ChangeAdministrationAddress)
	}
	// ChangeOperatorAddress (bool)
	if initialMessage.ChangeOperatorAddress != decodedMessage.ChangeOperatorAddress {
		t.Errorf("ChangeOperatorAddress doesn't match : %v != %v", initialMessage.ChangeOperatorAddress, decodedMessage.ChangeOperatorAddress)
	}
	// ContractRevision (uint)
	if initialMessage.ContractRevision != decodedMessage.ContractRevision {
		t.Errorf("ContractRevision doesn't match : %v != %v", initialMessage.ContractRevision, decodedMessage.ContractRevision)
	}
	// Amendments (Amendment[])
	if len(initialMessage.Amendments) != len(decodedMessage.Amendments) {
		t.Errorf("Amendments lengths don't match : %d != %d", len(initialMessage.Amendments), len(decodedMessage.Amendments))
	}
	// RefTxID (TxId)
	// TxId test compare not setup
}
func TestStaticContractFormation(t *testing.T) {
	// Create a randomized object
	initialMessage := StaticContractFormation{}
	// ContractName (varchar)
	// varchar test not setup
	{
		initialMessage.ContractName = "Text 0"
	}

	// ContractCode (ContractCode)
	// ContractCode test not setup
	{
		initialMessage.ContractCode = ContractCode{}
	}

	// BodyOfAgreementType (uint)
	// uint test not setup
	{
		// uint test not setup
	}

	// BodyOfAgreement (varbin)
	// varbin test not setup
	{
		initialMessage.BodyOfAgreement = make([]byte, 0, 32)
		for i := uint64(0); i < 32; i++ {
			initialMessage.BodyOfAgreement = append(initialMessage.BodyOfAgreement, byte(65+i+3))
		}
	}

	// ContractType (varchar)
	// varchar test not setup
	{
		initialMessage.ContractType = "Text 4"
	}

	// SupportingDocs (Document[])
	// Document[] test not setup
	{
		for i := 0; i < 2; i++ {
			initialMessage.SupportingDocs = append(initialMessage.SupportingDocs, Document{})
		}
	}

	// ContractRevision (uint)
	// uint test not setup
	{
		// uint test not setup
	}

	// GoverningLaw (fixedchar)
	// fixedchar test not setup
	{
		{
			text := make([]byte, 0, 5)
			for i := uint64(0); i < 5; i++ {
				text = append(text, byte(65+i+7))
			}
			initialMessage.GoverningLaw = string(text)
		}
	}

	// Jurisdiction (fixedchar)
	// fixedchar test not setup
	{
		{
			text := make([]byte, 0, 5)
			for i := uint64(0); i < 5; i++ {
				text = append(text, byte(65+i+8))
			}
			initialMessage.Jurisdiction = string(text)
		}
	}

	// EffectiveDate (Timestamp)
	// Timestamp test not setup
	{
		initialMessage.EffectiveDate = Timestamp{}
	}

	// ContractExpiration (Timestamp)
	// Timestamp test not setup
	{
		initialMessage.ContractExpiration = Timestamp{}
	}

	// ContractURI (varchar)
	// varchar test not setup
	{
		initialMessage.ContractURI = "Text 11"
	}

	// PrevRevTxID (TxId)
	// TxId test not setup
	{
		initialMessage.PrevRevTxID = TxId{}
	}

	// Entities (Entity[])
	// Entity[] test not setup
	{
		for i := 0; i < 2; i++ {
			initialMessage.Entities = append(initialMessage.Entities, Entity{})
		}
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := StaticContractFormation{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(initialEncoding, secondEncoding) {
		t.Errorf("Encoded value doesn't match.\ngot\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !cmp.Equal(&initialMessage, &decodedMessage) {
	// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// ContractName (varchar)
	if initialMessage.ContractName != decodedMessage.ContractName {
		t.Errorf("ContractName doesn't match : %s != %s", initialMessage.ContractName, decodedMessage.ContractName)
	}
	// ContractCode (ContractCode)
	// ContractCode test compare not setup
	// BodyOfAgreementType (uint)
	if initialMessage.BodyOfAgreementType != decodedMessage.BodyOfAgreementType {
		t.Errorf("BodyOfAgreementType doesn't match : %v != %v", initialMessage.BodyOfAgreementType, decodedMessage.BodyOfAgreementType)
	}
	// BodyOfAgreement (varbin)
	if !bytes.Equal(initialMessage.BodyOfAgreement, decodedMessage.BodyOfAgreement) {
		t.Errorf("BodyOfAgreement doesn't match : %x != %x", initialMessage.BodyOfAgreement, decodedMessage.BodyOfAgreement)
	}
	// ContractType (varchar)
	if initialMessage.ContractType != decodedMessage.ContractType {
		t.Errorf("ContractType doesn't match : %s != %s", initialMessage.ContractType, decodedMessage.ContractType)
	}
	// SupportingDocs (Document[])
	if len(initialMessage.SupportingDocs) != len(decodedMessage.SupportingDocs) {
		t.Errorf("SupportingDocs lengths don't match : %d != %d", len(initialMessage.SupportingDocs), len(decodedMessage.SupportingDocs))
	}
	// ContractRevision (uint)
	if initialMessage.ContractRevision != decodedMessage.ContractRevision {
		t.Errorf("ContractRevision doesn't match : %v != %v", initialMessage.ContractRevision, decodedMessage.ContractRevision)
	}
	// GoverningLaw (fixedchar)
	if initialMessage.GoverningLaw != decodedMessage.GoverningLaw {
		t.Errorf("GoverningLaw doesn't match : %s != %s", initialMessage.GoverningLaw, decodedMessage.GoverningLaw)
	}
	// Jurisdiction (fixedchar)
	if initialMessage.Jurisdiction != decodedMessage.Jurisdiction {
		t.Errorf("Jurisdiction doesn't match : %s != %s", initialMessage.Jurisdiction, decodedMessage.Jurisdiction)
	}
	// EffectiveDate (Timestamp)
	// Timestamp test compare not setup
	// ContractExpiration (Timestamp)
	// Timestamp test compare not setup
	// ContractURI (varchar)
	if initialMessage.ContractURI != decodedMessage.ContractURI {
		t.Errorf("ContractURI doesn't match : %s != %s", initialMessage.ContractURI, decodedMessage.ContractURI)
	}
	// PrevRevTxID (TxId)
	// TxId test compare not setup
	// Entities (Entity[])
	if len(initialMessage.Entities) != len(decodedMessage.Entities) {
		t.Errorf("Entities lengths don't match : %d != %d", len(initialMessage.Entities), len(decodedMessage.Entities))
	}
}
func TestContractAddressChange(t *testing.T) {
	// Create a randomized object
	initialMessage := ContractAddressChange{}
	// NewContractPKH (PublicKeyHash)
	// PublicKeyHash test not setup
	{
		initialMessage.NewContractPKH = PublicKeyHash{}
	}

	// Timestamp (Timestamp)
	// Timestamp test not setup
	{
		initialMessage.Timestamp = Timestamp{}
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := ContractAddressChange{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(initialEncoding, secondEncoding) {
		t.Errorf("Encoded value doesn't match.\ngot\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !cmp.Equal(&initialMessage, &decodedMessage) {
	// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// NewContractPKH (PublicKeyHash)
	// PublicKeyHash test compare not setup
	// Timestamp (Timestamp)
	// Timestamp test compare not setup
}
func TestOrder(t *testing.T) {
	// Create a randomized object
	initialMessage := Order{}
	// ComplianceAction (fixedchar)
	// fixedchar test not setup
	{
		// fixedchar test not setup
	}

	// AssetType (fixedchar)
	if initialMessage.ComplianceAction == 'F' || initialMessage.ComplianceAction == 'C' || initialMessage.ComplianceAction == 'R' {
		{
			text := make([]byte, 0, 3)
			for i := uint64(0); i < 3; i++ {
				text = append(text, byte(65+i+1))
			}
			initialMessage.AssetType = string(text)
		}
	}

	// AssetCode (AssetCode)
	if initialMessage.ComplianceAction == 'F' || initialMessage.ComplianceAction == 'C' || initialMessage.ComplianceAction == 'R' {
		initialMessage.AssetCode = AssetCode{}
	}

	// TargetAddresses (TargetAddress[])
	if initialMessage.ComplianceAction == 'F' || initialMessage.ComplianceAction == 'C' || initialMessage.ComplianceAction == 'R' {
		for i := 0; i < 2; i++ {
			initialMessage.TargetAddresses = append(initialMessage.TargetAddresses, TargetAddress{})
		}
	}

	// FreezeTxId (TxId)
	if initialMessage.ComplianceAction == 'T' {
		initialMessage.FreezeTxId = TxId{}
	}

	// FreezePeriod (Timestamp)
	if initialMessage.ComplianceAction == 'F' {
		initialMessage.FreezePeriod = Timestamp{}
	}

	// DepositAddress (PublicKeyHash)
	if initialMessage.ComplianceAction == 'C' {
		initialMessage.DepositAddress = PublicKeyHash{}
	}

	// AuthorityIncluded (bool)
	if initialMessage.ComplianceAction == 'F' || initialMessage.ComplianceAction == 'T' || initialMessage.ComplianceAction == 'C' {
		initialMessage.AuthorityIncluded = true
	}

	// AuthorityName (varchar)
	if initialMessage.AuthorityIncluded {
		initialMessage.AuthorityName = "Text 8"
	}

	// AuthorityPublicKey (varbin)
	if initialMessage.AuthorityIncluded {
		initialMessage.AuthorityPublicKey = make([]byte, 0, 8)
		for i := uint64(0); i < 8; i++ {
			initialMessage.AuthorityPublicKey = append(initialMessage.AuthorityPublicKey, byte(65+i+9))
		}
	}

	// SignatureAlgorithm (uint)
	if initialMessage.AuthorityIncluded {
		// uint test not setup
	}

	// OrderSignature (varbin)
	if initialMessage.AuthorityIncluded {
		initialMessage.OrderSignature = make([]byte, 0, 8)
		for i := uint64(0); i < 8; i++ {
			initialMessage.OrderSignature = append(initialMessage.OrderSignature, byte(65+i+11))
		}
	}

	// SupportingEvidenceHash (bin)
	// bin test not setup
	{
		// bin test not setup
	}

	// RefTxs (varbin)
	if initialMessage.ComplianceAction == 'R' {
		initialMessage.RefTxs = make([]byte, 0, 32)
		for i := uint64(0); i < 32; i++ {
			initialMessage.RefTxs = append(initialMessage.RefTxs, byte(65+i+13))
		}
	}

	// BitcoinDispersions (QuantityIndex[])
	if initialMessage.ComplianceAction == 'R' {
		for i := 0; i < 2; i++ {
			initialMessage.BitcoinDispersions = append(initialMessage.BitcoinDispersions, QuantityIndex{})
		}
	}

	// Message (varchar)
	// varchar test not setup
	{
		initialMessage.Message = "Text 15"
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Order{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(initialEncoding, secondEncoding) {
		t.Errorf("Encoded value doesn't match.\ngot\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !cmp.Equal(&initialMessage, &decodedMessage) {
	// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// ComplianceAction (fixedchar)
	if initialMessage.ComplianceAction != decodedMessage.ComplianceAction {
		t.Errorf("ComplianceAction doesn't match : %v != %v", initialMessage.ComplianceAction, decodedMessage.ComplianceAction)
	}
	// AssetType (fixedchar)
	if initialMessage.AssetType != decodedMessage.AssetType {
		t.Errorf("AssetType doesn't match : %s != %s", initialMessage.AssetType, decodedMessage.AssetType)
	}
	// AssetCode (AssetCode)
	// AssetCode test compare not setup
	// TargetAddresses (TargetAddress[])
	if len(initialMessage.TargetAddresses) != len(decodedMessage.TargetAddresses) {
		t.Errorf("TargetAddresses lengths don't match : %d != %d", len(initialMessage.TargetAddresses), len(decodedMessage.TargetAddresses))
	}
	// FreezeTxId (TxId)
	// TxId test compare not setup
	// FreezePeriod (Timestamp)
	// Timestamp test compare not setup
	// DepositAddress (PublicKeyHash)
	// PublicKeyHash test compare not setup
	// AuthorityIncluded (bool)
	if initialMessage.AuthorityIncluded != decodedMessage.AuthorityIncluded {
		t.Errorf("AuthorityIncluded doesn't match : %v != %v", initialMessage.AuthorityIncluded, decodedMessage.AuthorityIncluded)
	}
	// AuthorityName (varchar)
	if initialMessage.AuthorityName != decodedMessage.AuthorityName {
		t.Errorf("AuthorityName doesn't match : %s != %s", initialMessage.AuthorityName, decodedMessage.AuthorityName)
	}
	// AuthorityPublicKey (varbin)
	if !bytes.Equal(initialMessage.AuthorityPublicKey, decodedMessage.AuthorityPublicKey) {
		t.Errorf("AuthorityPublicKey doesn't match : %x != %x", initialMessage.AuthorityPublicKey, decodedMessage.AuthorityPublicKey)
	}
	// SignatureAlgorithm (uint)
	if initialMessage.SignatureAlgorithm != decodedMessage.SignatureAlgorithm {
		t.Errorf("SignatureAlgorithm doesn't match : %v != %v", initialMessage.SignatureAlgorithm, decodedMessage.SignatureAlgorithm)
	}
	// OrderSignature (varbin)
	if !bytes.Equal(initialMessage.OrderSignature, decodedMessage.OrderSignature) {
		t.Errorf("OrderSignature doesn't match : %x != %x", initialMessage.OrderSignature, decodedMessage.OrderSignature)
	}
	// SupportingEvidenceHash (bin)
	if initialMessage.SupportingEvidenceHash != decodedMessage.SupportingEvidenceHash {
		t.Errorf("SupportingEvidenceHash doesn't match : %v != %v", initialMessage.SupportingEvidenceHash, decodedMessage.SupportingEvidenceHash)
	}
	// RefTxs (varbin)
	if !bytes.Equal(initialMessage.RefTxs, decodedMessage.RefTxs) {
		t.Errorf("RefTxs doesn't match : %x != %x", initialMessage.RefTxs, decodedMessage.RefTxs)
	}
	// BitcoinDispersions (QuantityIndex[])
	if len(initialMessage.BitcoinDispersions) != len(decodedMessage.BitcoinDispersions) {
		t.Errorf("BitcoinDispersions lengths don't match : %d != %d", len(initialMessage.BitcoinDispersions), len(decodedMessage.BitcoinDispersions))
	}
	// Message (varchar)
	if initialMessage.Message != decodedMessage.Message {
		t.Errorf("Message doesn't match : %s != %s", initialMessage.Message, decodedMessage.Message)
	}
}
func TestFreeze(t *testing.T) {
	// Create a randomized object
	initialMessage := Freeze{}
	// AssetType (fixedchar)
	// fixedchar test not setup
	{
		{
			text := make([]byte, 0, 3)
			for i := uint64(0); i < 3; i++ {
				text = append(text, byte(65+i+0))
			}
			initialMessage.AssetType = string(text)
		}
	}

	// AssetCode (AssetCode)
	// AssetCode test not setup
	{
		initialMessage.AssetCode = AssetCode{}
	}

	// Quantities (QuantityIndex[])
	// QuantityIndex[] test not setup
	{
		for i := 0; i < 2; i++ {
			initialMessage.Quantities = append(initialMessage.Quantities, QuantityIndex{})
		}
	}

	// FreezePeriod (Timestamp)
	// Timestamp test not setup
	{
		initialMessage.FreezePeriod = Timestamp{}
	}

	// Timestamp (Timestamp)
	// Timestamp test not setup
	{
		initialMessage.Timestamp = Timestamp{}
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Freeze{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(initialEncoding, secondEncoding) {
		t.Errorf("Encoded value doesn't match.\ngot\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !cmp.Equal(&initialMessage, &decodedMessage) {
	// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// AssetType (fixedchar)
	if initialMessage.AssetType != decodedMessage.AssetType {
		t.Errorf("AssetType doesn't match : %s != %s", initialMessage.AssetType, decodedMessage.AssetType)
	}
	// AssetCode (AssetCode)
	// AssetCode test compare not setup
	// Quantities (QuantityIndex[])
	if len(initialMessage.Quantities) != len(decodedMessage.Quantities) {
		t.Errorf("Quantities lengths don't match : %d != %d", len(initialMessage.Quantities), len(decodedMessage.Quantities))
	}
	// FreezePeriod (Timestamp)
	// Timestamp test compare not setup
	// Timestamp (Timestamp)
	// Timestamp test compare not setup
}
func TestThaw(t *testing.T) {
	// Create a randomized object
	initialMessage := Thaw{}
	// FreezeTxId (TxId)
	// TxId test not setup
	{
		initialMessage.FreezeTxId = TxId{}
	}

	// Timestamp (Timestamp)
	// Timestamp test not setup
	{
		initialMessage.Timestamp = Timestamp{}
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Thaw{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(initialEncoding, secondEncoding) {
		t.Errorf("Encoded value doesn't match.\ngot\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !cmp.Equal(&initialMessage, &decodedMessage) {
	// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// FreezeTxId (TxId)
	// TxId test compare not setup
	// Timestamp (Timestamp)
	// Timestamp test compare not setup
}
func TestConfiscation(t *testing.T) {
	// Create a randomized object
	initialMessage := Confiscation{}
	// AssetType (fixedchar)
	// fixedchar test not setup
	{
		{
			text := make([]byte, 0, 3)
			for i := uint64(0); i < 3; i++ {
				text = append(text, byte(65+i+0))
			}
			initialMessage.AssetType = string(text)
		}
	}

	// AssetCode (AssetCode)
	// AssetCode test not setup
	{
		initialMessage.AssetCode = AssetCode{}
	}

	// Quantities (QuantityIndex[])
	// QuantityIndex[] test not setup
	{
		for i := 0; i < 2; i++ {
			initialMessage.Quantities = append(initialMessage.Quantities, QuantityIndex{})
		}
	}

	// DepositQty (uint)
	// uint test not setup
	{
		// uint test not setup
	}

	// Timestamp (Timestamp)
	// Timestamp test not setup
	{
		initialMessage.Timestamp = Timestamp{}
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Confiscation{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(initialEncoding, secondEncoding) {
		t.Errorf("Encoded value doesn't match.\ngot\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !cmp.Equal(&initialMessage, &decodedMessage) {
	// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// AssetType (fixedchar)
	if initialMessage.AssetType != decodedMessage.AssetType {
		t.Errorf("AssetType doesn't match : %s != %s", initialMessage.AssetType, decodedMessage.AssetType)
	}
	// AssetCode (AssetCode)
	// AssetCode test compare not setup
	// Quantities (QuantityIndex[])
	if len(initialMessage.Quantities) != len(decodedMessage.Quantities) {
		t.Errorf("Quantities lengths don't match : %d != %d", len(initialMessage.Quantities), len(decodedMessage.Quantities))
	}
	// DepositQty (uint)
	if initialMessage.DepositQty != decodedMessage.DepositQty {
		t.Errorf("DepositQty doesn't match : %v != %v", initialMessage.DepositQty, decodedMessage.DepositQty)
	}
	// Timestamp (Timestamp)
	// Timestamp test compare not setup
}
func TestReconciliation(t *testing.T) {
	// Create a randomized object
	initialMessage := Reconciliation{}
	// AssetType (fixedchar)
	// fixedchar test not setup
	{
		{
			text := make([]byte, 0, 3)
			for i := uint64(0); i < 3; i++ {
				text = append(text, byte(65+i+0))
			}
			initialMessage.AssetType = string(text)
		}
	}

	// AssetCode (AssetCode)
	// AssetCode test not setup
	{
		initialMessage.AssetCode = AssetCode{}
	}

	// Quantities (QuantityIndex[])
	// QuantityIndex[] test not setup
	{
		for i := 0; i < 2; i++ {
			initialMessage.Quantities = append(initialMessage.Quantities, QuantityIndex{})
		}
	}

	// Timestamp (Timestamp)
	// Timestamp test not setup
	{
		initialMessage.Timestamp = Timestamp{}
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Reconciliation{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(initialEncoding, secondEncoding) {
		t.Errorf("Encoded value doesn't match.\ngot\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !cmp.Equal(&initialMessage, &decodedMessage) {
	// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// AssetType (fixedchar)
	if initialMessage.AssetType != decodedMessage.AssetType {
		t.Errorf("AssetType doesn't match : %s != %s", initialMessage.AssetType, decodedMessage.AssetType)
	}
	// AssetCode (AssetCode)
	// AssetCode test compare not setup
	// Quantities (QuantityIndex[])
	if len(initialMessage.Quantities) != len(decodedMessage.Quantities) {
		t.Errorf("Quantities lengths don't match : %d != %d", len(initialMessage.Quantities), len(decodedMessage.Quantities))
	}
	// Timestamp (Timestamp)
	// Timestamp test compare not setup
}
func TestProposal(t *testing.T) {
	// Create a randomized object
	initialMessage := Proposal{}
	// Initiator (uint)
	// uint test not setup
	{
		// uint test not setup
	}

	// AssetSpecificVote (bool)
	// bool test not setup
	{
		initialMessage.AssetSpecificVote = true
	}

	// AssetType (fixedchar)
	if initialMessage.AssetSpecificVote {
		{
			text := make([]byte, 0, 3)
			for i := uint64(0); i < 3; i++ {
				text = append(text, byte(65+i+2))
			}
			initialMessage.AssetType = string(text)
		}
	}

	// AssetCode (AssetCode)
	if initialMessage.AssetSpecificVote {
		initialMessage.AssetCode = AssetCode{}
	}

	// VoteSystem (uint)
	// uint test not setup
	{
		// uint test not setup
	}

	// Specific (bool)
	// bool test not setup
	{
		initialMessage.Specific = true
	}

	// ProposedAmendments (Amendment[])
	if initialMessage.Specific {
		for i := 0; i < 2; i++ {
			initialMessage.ProposedAmendments = append(initialMessage.ProposedAmendments, Amendment{})
		}
	}

	// VoteOptions (varchar)
	// varchar test not setup
	{
		initialMessage.VoteOptions = "Text 7"
	}

	// VoteMax (uint)
	// uint test not setup
	{
		// uint test not setup
	}

	// ProposalDescription (varchar)
	// varchar test not setup
	{
		initialMessage.ProposalDescription = "Text 9"
	}

	// ProposalDocumentHash (bin)
	// bin test not setup
	{
		// bin test not setup
	}

	// VoteCutOffTimestamp (Timestamp)
	// Timestamp test not setup
	{
		initialMessage.VoteCutOffTimestamp = Timestamp{}
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Proposal{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(initialEncoding, secondEncoding) {
		t.Errorf("Encoded value doesn't match.\ngot\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !cmp.Equal(&initialMessage, &decodedMessage) {
	// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// Initiator (uint)
	if initialMessage.Initiator != decodedMessage.Initiator {
		t.Errorf("Initiator doesn't match : %v != %v", initialMessage.Initiator, decodedMessage.Initiator)
	}
	// AssetSpecificVote (bool)
	if initialMessage.AssetSpecificVote != decodedMessage.AssetSpecificVote {
		t.Errorf("AssetSpecificVote doesn't match : %v != %v", initialMessage.AssetSpecificVote, decodedMessage.AssetSpecificVote)
	}
	// AssetType (fixedchar)
	if initialMessage.AssetType != decodedMessage.AssetType {
		t.Errorf("AssetType doesn't match : %s != %s", initialMessage.AssetType, decodedMessage.AssetType)
	}
	// AssetCode (AssetCode)
	// AssetCode test compare not setup
	// VoteSystem (uint)
	if initialMessage.VoteSystem != decodedMessage.VoteSystem {
		t.Errorf("VoteSystem doesn't match : %v != %v", initialMessage.VoteSystem, decodedMessage.VoteSystem)
	}
	// Specific (bool)
	if initialMessage.Specific != decodedMessage.Specific {
		t.Errorf("Specific doesn't match : %v != %v", initialMessage.Specific, decodedMessage.Specific)
	}
	// ProposedAmendments (Amendment[])
	if len(initialMessage.ProposedAmendments) != len(decodedMessage.ProposedAmendments) {
		t.Errorf("ProposedAmendments lengths don't match : %d != %d", len(initialMessage.ProposedAmendments), len(decodedMessage.ProposedAmendments))
	}
	// VoteOptions (varchar)
	if initialMessage.VoteOptions != decodedMessage.VoteOptions {
		t.Errorf("VoteOptions doesn't match : %s != %s", initialMessage.VoteOptions, decodedMessage.VoteOptions)
	}
	// VoteMax (uint)
	if initialMessage.VoteMax != decodedMessage.VoteMax {
		t.Errorf("VoteMax doesn't match : %v != %v", initialMessage.VoteMax, decodedMessage.VoteMax)
	}
	// ProposalDescription (varchar)
	if initialMessage.ProposalDescription != decodedMessage.ProposalDescription {
		t.Errorf("ProposalDescription doesn't match : %s != %s", initialMessage.ProposalDescription, decodedMessage.ProposalDescription)
	}
	// ProposalDocumentHash (bin)
	if initialMessage.ProposalDocumentHash != decodedMessage.ProposalDocumentHash {
		t.Errorf("ProposalDocumentHash doesn't match : %v != %v", initialMessage.ProposalDocumentHash, decodedMessage.ProposalDocumentHash)
	}
	// VoteCutOffTimestamp (Timestamp)
	// Timestamp test compare not setup
}
func TestVote(t *testing.T) {
	// Create a randomized object
	initialMessage := Vote{}
	// Timestamp (Timestamp)
	// Timestamp test not setup
	{
		initialMessage.Timestamp = Timestamp{}
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Vote{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(initialEncoding, secondEncoding) {
		t.Errorf("Encoded value doesn't match.\ngot\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !cmp.Equal(&initialMessage, &decodedMessage) {
	// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// Timestamp (Timestamp)
	// Timestamp test compare not setup
}
func TestBallotCast(t *testing.T) {
	// Create a randomized object
	initialMessage := BallotCast{}
	// VoteTxId (TxId)
	// TxId test not setup
	{
		initialMessage.VoteTxId = TxId{}
	}

	// Vote (varchar)
	// varchar test not setup
	{
		initialMessage.Vote = "Text 1"
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := BallotCast{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(initialEncoding, secondEncoding) {
		t.Errorf("Encoded value doesn't match.\ngot\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !cmp.Equal(&initialMessage, &decodedMessage) {
	// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// VoteTxId (TxId)
	// TxId test compare not setup
	// Vote (varchar)
	if initialMessage.Vote != decodedMessage.Vote {
		t.Errorf("Vote doesn't match : %s != %s", initialMessage.Vote, decodedMessage.Vote)
	}
}
func TestBallotCounted(t *testing.T) {
	// Create a randomized object
	initialMessage := BallotCounted{}
	// VoteTxId (TxId)
	// TxId test not setup
	{
		initialMessage.VoteTxId = TxId{}
	}

	// Vote (varchar)
	// varchar test not setup
	{
		initialMessage.Vote = "Text 1"
	}

	// Quantity (uint)
	// uint test not setup
	{
		// uint test not setup
	}

	// Timestamp (Timestamp)
	// Timestamp test not setup
	{
		initialMessage.Timestamp = Timestamp{}
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := BallotCounted{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(initialEncoding, secondEncoding) {
		t.Errorf("Encoded value doesn't match.\ngot\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !cmp.Equal(&initialMessage, &decodedMessage) {
	// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// VoteTxId (TxId)
	// TxId test compare not setup
	// Vote (varchar)
	if initialMessage.Vote != decodedMessage.Vote {
		t.Errorf("Vote doesn't match : %s != %s", initialMessage.Vote, decodedMessage.Vote)
	}
	// Quantity (uint)
	if initialMessage.Quantity != decodedMessage.Quantity {
		t.Errorf("Quantity doesn't match : %v != %v", initialMessage.Quantity, decodedMessage.Quantity)
	}
	// Timestamp (Timestamp)
	// Timestamp test compare not setup
}
func TestResult(t *testing.T) {
	// Create a randomized object
	initialMessage := Result{}
	// AssetSpecificVote (bool)
	// bool test not setup
	{
		initialMessage.AssetSpecificVote = true
	}

	// AssetType (fixedchar)
	if initialMessage.AssetSpecificVote {
		{
			text := make([]byte, 0, 3)
			for i := uint64(0); i < 3; i++ {
				text = append(text, byte(65+i+1))
			}
			initialMessage.AssetType = string(text)
		}
	}

	// AssetCode (AssetCode)
	if initialMessage.AssetSpecificVote {
		initialMessage.AssetCode = AssetCode{}
	}

	// Specific (bool)
	// bool test not setup
	{
		initialMessage.Specific = true
	}

	// ProposedAmendments (Amendment[])
	if initialMessage.Specific {
		for i := 0; i < 2; i++ {
			initialMessage.ProposedAmendments = append(initialMessage.ProposedAmendments, Amendment{})
		}
	}

	// VoteTxId (TxId)
	// TxId test not setup
	{
		initialMessage.VoteTxId = TxId{}
	}

	// OptionTally (uint64[])
	// uint64[] test not setup
	{
		for i := 0; i < 5; i++ {
			var item uint64
			initialMessage.OptionTally = append(initialMessage.OptionTally, item)
		}
	}

	// Result (varchar)
	// varchar test not setup
	{
		initialMessage.Result = "Text 7"
	}

	// Timestamp (Timestamp)
	// Timestamp test not setup
	{
		initialMessage.Timestamp = Timestamp{}
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Result{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(initialEncoding, secondEncoding) {
		t.Errorf("Encoded value doesn't match.\ngot\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !cmp.Equal(&initialMessage, &decodedMessage) {
	// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// AssetSpecificVote (bool)
	if initialMessage.AssetSpecificVote != decodedMessage.AssetSpecificVote {
		t.Errorf("AssetSpecificVote doesn't match : %v != %v", initialMessage.AssetSpecificVote, decodedMessage.AssetSpecificVote)
	}
	// AssetType (fixedchar)
	if initialMessage.AssetType != decodedMessage.AssetType {
		t.Errorf("AssetType doesn't match : %s != %s", initialMessage.AssetType, decodedMessage.AssetType)
	}
	// AssetCode (AssetCode)
	// AssetCode test compare not setup
	// Specific (bool)
	if initialMessage.Specific != decodedMessage.Specific {
		t.Errorf("Specific doesn't match : %v != %v", initialMessage.Specific, decodedMessage.Specific)
	}
	// ProposedAmendments (Amendment[])
	if len(initialMessage.ProposedAmendments) != len(decodedMessage.ProposedAmendments) {
		t.Errorf("ProposedAmendments lengths don't match : %d != %d", len(initialMessage.ProposedAmendments), len(decodedMessage.ProposedAmendments))
	}
	// VoteTxId (TxId)
	// TxId test compare not setup
	// OptionTally (uint64[])
	if len(initialMessage.OptionTally) != len(decodedMessage.OptionTally) {
		t.Errorf("OptionTally lengths don't match : %d != %d", len(initialMessage.OptionTally), len(decodedMessage.OptionTally))
	}
	for i, value := range initialMessage.OptionTally {
		if value != decodedMessage.OptionTally[i] {
			t.Errorf("OptionTally value %d doesn't match : %v != %v", i, value, decodedMessage.OptionTally[i])
		}
	}
	// Result (varchar)
	if initialMessage.Result != decodedMessage.Result {
		t.Errorf("Result doesn't match : %s != %s", initialMessage.Result, decodedMessage.Result)
	}
	// Timestamp (Timestamp)
	// Timestamp test compare not setup
}
func TestMessage(t *testing.T) {
	// Create a randomized object
	initialMessage := Message{}
	// AddressIndexes (uint16[])
	// uint16[] test not setup
	{
		for i := 0; i < 5; i++ {
			var item uint16
			initialMessage.AddressIndexes = append(initialMessage.AddressIndexes, item)
		}
	}

	// MessageType (uint16)
	// uint16 test not setup
	{
		// uint16 test not setup
	}

	// MessagePayload (varbin)
	// varbin test not setup
	{
		initialMessage.MessagePayload = make([]byte, 0, 32)
		for i := uint64(0); i < 32; i++ {
			initialMessage.MessagePayload = append(initialMessage.MessagePayload, byte(65+i+2))
		}
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Message{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(initialEncoding, secondEncoding) {
		t.Errorf("Encoded value doesn't match.\ngot\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !cmp.Equal(&initialMessage, &decodedMessage) {
	// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// AddressIndexes (uint16[])
	if len(initialMessage.AddressIndexes) != len(decodedMessage.AddressIndexes) {
		t.Errorf("AddressIndexes lengths don't match : %d != %d", len(initialMessage.AddressIndexes), len(decodedMessage.AddressIndexes))
	}
	for i, value := range initialMessage.AddressIndexes {
		if value != decodedMessage.AddressIndexes[i] {
			t.Errorf("AddressIndexes value %d doesn't match : %v != %v", i, value, decodedMessage.AddressIndexes[i])
		}
	}
	// MessageType (uint16)
	if initialMessage.MessageType != decodedMessage.MessageType {
		t.Errorf("MessageType doesn't match : %v != %v", initialMessage.MessageType, decodedMessage.MessageType)
	}
	// MessagePayload (varbin)
	if !bytes.Equal(initialMessage.MessagePayload, decodedMessage.MessagePayload) {
		t.Errorf("MessagePayload doesn't match : %x != %x", initialMessage.MessagePayload, decodedMessage.MessagePayload)
	}
}
func TestRejection(t *testing.T) {
	// Create a randomized object
	initialMessage := Rejection{}
	// AddressIndexes (uint16[])
	// uint16[] test not setup
	{
		for i := 0; i < 5; i++ {
			var item uint16
			initialMessage.AddressIndexes = append(initialMessage.AddressIndexes, item)
		}
	}

	// RejectAddressIndex (uint)
	// uint test not setup
	{
		// uint test not setup
	}

	// RejectionCode (RejectionCode)
	// RejectionCode test not setup
	{
		// RejectionCode test not setup
	}

	// Message (varchar)
	// varchar test not setup
	{
		initialMessage.Message = "Text 3"
	}

	// Timestamp (Timestamp)
	// Timestamp test not setup
	{
		initialMessage.Timestamp = Timestamp{}
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Rejection{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(initialEncoding, secondEncoding) {
		t.Errorf("Encoded value doesn't match.\ngot\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !cmp.Equal(&initialMessage, &decodedMessage) {
	// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// AddressIndexes (uint16[])
	if len(initialMessage.AddressIndexes) != len(decodedMessage.AddressIndexes) {
		t.Errorf("AddressIndexes lengths don't match : %d != %d", len(initialMessage.AddressIndexes), len(decodedMessage.AddressIndexes))
	}
	for i, value := range initialMessage.AddressIndexes {
		if value != decodedMessage.AddressIndexes[i] {
			t.Errorf("AddressIndexes value %d doesn't match : %v != %v", i, value, decodedMessage.AddressIndexes[i])
		}
	}
	// RejectAddressIndex (uint)
	if initialMessage.RejectAddressIndex != decodedMessage.RejectAddressIndex {
		t.Errorf("RejectAddressIndex doesn't match : %v != %v", initialMessage.RejectAddressIndex, decodedMessage.RejectAddressIndex)
	}
	// RejectionCode (RejectionCode)
	if initialMessage.RejectionCode != decodedMessage.RejectionCode {
		t.Errorf("RejectionCode doesn't match : %v != %v", initialMessage.RejectionCode, decodedMessage.RejectionCode)
	}
	// Message (varchar)
	if initialMessage.Message != decodedMessage.Message {
		t.Errorf("Message doesn't match : %s != %s", initialMessage.Message, decodedMessage.Message)
	}
	// Timestamp (Timestamp)
	// Timestamp test compare not setup
}
func TestEstablishment(t *testing.T) {
	// Create a randomized object
	initialMessage := Establishment{}
	// Message (varchar)
	// varchar test not setup
	{
		initialMessage.Message = "Text 0"
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Establishment{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(initialEncoding, secondEncoding) {
		t.Errorf("Encoded value doesn't match.\ngot\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !cmp.Equal(&initialMessage, &decodedMessage) {
	// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// Message (varchar)
	if initialMessage.Message != decodedMessage.Message {
		t.Errorf("Message doesn't match : %s != %s", initialMessage.Message, decodedMessage.Message)
	}
}
func TestAddition(t *testing.T) {
	// Create a randomized object
	initialMessage := Addition{}
	// Message (varchar)
	// varchar test not setup
	{
		initialMessage.Message = "Text 0"
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Addition{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(initialEncoding, secondEncoding) {
		t.Errorf("Encoded value doesn't match.\ngot\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !cmp.Equal(&initialMessage, &decodedMessage) {
	// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// Message (varchar)
	if initialMessage.Message != decodedMessage.Message {
		t.Errorf("Message doesn't match : %s != %s", initialMessage.Message, decodedMessage.Message)
	}
}
func TestAlteration(t *testing.T) {
	// Create a randomized object
	initialMessage := Alteration{}
	// EntryTxID (TxId)
	// TxId test not setup
	{
		initialMessage.EntryTxID = TxId{}
	}

	// Message (varchar)
	// varchar test not setup
	{
		initialMessage.Message = "Text 1"
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Alteration{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(initialEncoding, secondEncoding) {
		t.Errorf("Encoded value doesn't match.\ngot\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !cmp.Equal(&initialMessage, &decodedMessage) {
	// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// EntryTxID (TxId)
	// TxId test compare not setup
	// Message (varchar)
	if initialMessage.Message != decodedMessage.Message {
		t.Errorf("Message doesn't match : %s != %s", initialMessage.Message, decodedMessage.Message)
	}
}
func TestRemoval(t *testing.T) {
	// Create a randomized object
	initialMessage := Removal{}
	// EntryTxID (TxId)
	// TxId test not setup
	{
		initialMessage.EntryTxID = TxId{}
	}

	// Message (varchar)
	// varchar test not setup
	{
		initialMessage.Message = "Text 1"
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Removal{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(initialEncoding, secondEncoding) {
		t.Errorf("Encoded value doesn't match.\ngot\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !cmp.Equal(&initialMessage, &decodedMessage) {
	// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// EntryTxID (TxId)
	// TxId test compare not setup
	// Message (varchar)
	if initialMessage.Message != decodedMessage.Message {
		t.Errorf("Message doesn't match : %s != %s", initialMessage.Message, decodedMessage.Message)
	}
}
func TestTransfer(t *testing.T) {
	// Create a randomized object
	initialMessage := Transfer{}
	// Assets (AssetTransfer[])
	// AssetTransfer[] test not setup
	{
		for i := 0; i < 2; i++ {
			initialMessage.Assets = append(initialMessage.Assets, AssetTransfer{})
		}
	}

	// OfferExpiry (Timestamp)
	// Timestamp test not setup
	{
		initialMessage.OfferExpiry = Timestamp{}
	}

	// ExchangeFee (uint)
	// uint test not setup
	{
		// uint test not setup
	}

	// ExchangeFeeAddress (PublicKeyHash)
	// PublicKeyHash test not setup
	{
		initialMessage.ExchangeFeeAddress = PublicKeyHash{}
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Transfer{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(initialEncoding, secondEncoding) {
		t.Errorf("Encoded value doesn't match.\ngot\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !cmp.Equal(&initialMessage, &decodedMessage) {
	// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// Assets (AssetTransfer[])
	if len(initialMessage.Assets) != len(decodedMessage.Assets) {
		t.Errorf("Assets lengths don't match : %d != %d", len(initialMessage.Assets), len(decodedMessage.Assets))
	}
	// OfferExpiry (Timestamp)
	// Timestamp test compare not setup
	// ExchangeFee (uint)
	if initialMessage.ExchangeFee != decodedMessage.ExchangeFee {
		t.Errorf("ExchangeFee doesn't match : %v != %v", initialMessage.ExchangeFee, decodedMessage.ExchangeFee)
	}
	// ExchangeFeeAddress (PublicKeyHash)
	// PublicKeyHash test compare not setup
}
func TestSettlement(t *testing.T) {
	// Create a randomized object
	initialMessage := Settlement{}
	// Assets (AssetSettlement[])
	// AssetSettlement[] test not setup
	{
		for i := 0; i < 2; i++ {
			initialMessage.Assets = append(initialMessage.Assets, AssetSettlement{})
		}
	}

	// Timestamp (Timestamp)
	// Timestamp test not setup
	{
		initialMessage.Timestamp = Timestamp{}
	}

	// Encode message
	initialEncoding, err := initialMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Initial encoding : %d bytes", len(initialEncoding))

	// Decode message
	decodedMessage := Settlement{}

	n, err := decodedMessage.write(initialEncoding)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Decoded : %d bytes", n)

	if n != len(initialEncoding) {
		t.Fatalf("got %v, want %v", n, len(initialEncoding))
	}

	// Serializing the message should give us the same bytes
	secondEncoding, err := decodedMessage.serialize()
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(initialEncoding, secondEncoding) {
		t.Errorf("Encoded value doesn't match.\ngot\n%+v\nwant\n%+v", initialEncoding, secondEncoding)
	}

	// if !cmp.Equal(&initialMessage, &decodedMessage) {
	// 	t.Errorf("Decoded value doesn't match.\ninitial : %+v\ndecoded : %+v", initialMessage, decodedMessage)
	// }

	// Compare re-serialized values
	// Assets (AssetSettlement[])
	if len(initialMessage.Assets) != len(decodedMessage.Assets) {
		t.Errorf("Assets lengths don't match : %d != %d", len(initialMessage.Assets), len(decodedMessage.Assets))
	}
	// Timestamp (Timestamp)
	// Timestamp test compare not setup
}
