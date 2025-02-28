package protocol

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/tokenized/smart-contract/pkg/txbuilder"
)

const (
	// OpReturn (OP_RETURN) is a script opcode is used to mark a transaction
	// output as invalid, and can be used to add data to a TX.
	OpReturn = byte(0x6a)

	// ProtocolID is the current protocol ID
	ProtocolID = "tokenized"

	// TestProtocolID is the current protocol ID for testing.
	TestProtocolID = "test.tokenized"

	// Version of the Tokenized protocol.
	Version = uint8(0)
)

// OpReturnMessage implements a base interface for all message types.
type OpReturnMessage interface {
	Type() string
	String() string
	serialize() ([]byte, error)
	write(b []byte) (int, error)
	Validate() error
}

// Deserialize returns a message, as an OpReturnMessage, from the OP_RETURN script.
func Deserialize(b []byte, isTest bool) (OpReturnMessage, error) {
	buf := bytes.NewBuffer(b)

	var opCode byte
	var err error

	// Parse OP_RETURN op code
	err = binary.Read(buf, DefaultEndian, &opCode)
	if err != nil {
		return nil, err
	}

	if opCode != OpReturn {
		return nil, fmt.Errorf("Not an op return output : %02x", opCode)
	}

	// Parse protocol ID
	var protocolID string
	if isTest {
		protocolID = TestProtocolID
	} else {
		protocolID = ProtocolID
	}

	protocolIDSize, err := txbuilder.ParsePushDataScript(buf)
	if err != nil {
		return nil, err
	}

	if int(protocolIDSize) != len(protocolID) {
		return nil, fmt.Errorf("Push not correct size for protocol ID : %d != %d", protocolIDSize, len(protocolID))
	}

	readProtocolID := make([]byte, int(protocolIDSize))
	_, err = buf.Read(readProtocolID)
	if err != nil {
		return nil, err
	}

	if !bytes.Equal(readProtocolID, []byte(protocolID)) {
		return nil, fmt.Errorf("Invalid protocol ID : %s", string(readProtocolID))
	}

	// Parse push op code for payload length + 3 for version and message type code
	var payloadSize uint64
	payloadSize, err = txbuilder.ParsePushDataScript(buf)
	if err != nil {
		return nil, err
	}

	if uint64(buf.Len()) < payloadSize {
		return nil, fmt.Errorf("Payload push op code is too large for message : %d", payloadSize)
	}

	// Parse version
	var version uint8
	err = binary.Read(buf, DefaultEndian, &version)
	if err != nil {
		return nil, err
	}

	if version != Version {
		return nil, fmt.Errorf("Unsupported version : %02x", version)
	}

	// Parse message type code
	code := make([]byte, 2)
	_, err = buf.Read(code)
	if err != nil {
		return nil, err
	}

	msg := TypeMapping(string(code))
	if msg == nil {
		return nil, fmt.Errorf("Unknown code : %s", code)
	}

	if _, err := msg.write(b[len(b)-buf.Len():]); err != nil {
		return nil, err
	}

	return msg, nil
}

// Serialize returns a complete op return script including the specified payload.
func Serialize(msg OpReturnMessage, isTest bool) ([]byte, error) {
	var buf bytes.Buffer
	var err error

	// Write OP_RETURN op code
	err = binary.Write(&buf, DefaultEndian, OpReturn)
	if err != nil {
		return nil, err
	}

	var protocolID string
	if isTest {
		protocolID = TestProtocolID
	} else {
		protocolID = ProtocolID
	}

	// Write protocol Id push op code
	_, err = buf.Write(txbuilder.PushDataScript(uint64(len(protocolID))))
	if err != nil {
		fmt.Printf("Failed to write push data : %s\n", err)
		return nil, err
	}

	// Write protocol Id
	_, err = buf.Write([]byte(protocolID))
	if err != nil {
		return nil, err
	}

	var payload []byte
	payload, err = msg.serialize()
	if err != nil {
		return nil, err
	}

	// Write push op code for payload length + 3 for version and message type code
	_, err = buf.Write(txbuilder.PushDataScript(uint64(len(payload)) + 3))
	if err != nil {
		return nil, err
	}

	// Write version
	err = binary.Write(&buf, DefaultEndian, Version)
	if err != nil {
		return nil, err
	}

	// Write message type code
	_, err = buf.Write([]byte(msg.Type()))
	if err != nil {
		return nil, err
	}

	// Write payload
	_, err = buf.Write(payload)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Code returns the identifying code from the OP_RETURN payload.
func Code(script []byte, isTest bool) (string, error) {
	buf := bytes.NewBuffer(script)

	var opCode byte
	var err error

	// Parse OP_RETURN op code
	err = binary.Read(buf, DefaultEndian, &opCode)
	if err != nil {
		return "", err
	}

	if opCode != OpReturn {
		return "", fmt.Errorf("Not an op return output : %02x", opCode)
	}

	// Parse protocol ID
	var protocolID string
	if isTest {
		protocolID = TestProtocolID
	} else {
		protocolID = ProtocolID
	}

	protocolIDSize, err := txbuilder.ParsePushDataScript(buf)
	if err != nil {
		return "", err
	}

	if int(protocolIDSize) != len(protocolID) {
		return "", fmt.Errorf("Push not correct size for protocol ID : %d != %d", protocolIDSize, len(protocolID))
	}

	readProtocolID := make([]byte, int(protocolIDSize))
	_, err = buf.Read(readProtocolID)
	if err != nil {
		return "", err
	}

	if !bytes.Equal(readProtocolID, []byte(protocolID)) {
		return "", fmt.Errorf("Invalid protocol ID : %s", string(readProtocolID))
	}

	// Parse push op code for payload length + 3 for version and message type code
	var payloadSize uint64
	payloadSize, err = txbuilder.ParsePushDataScript(buf)
	if err != nil {
		return "", err
	}

	if uint64(buf.Len()) < payloadSize {
		return "", fmt.Errorf("Payload push op code is too large for message : %d", payloadSize)
	}

	// Parse version
	var version uint8
	err = binary.Read(buf, DefaultEndian, &version)
	if err != nil {
		return "", err
	}

	if version != Version {
		return "", fmt.Errorf("Unsupported version : %02x", version)
	}

	// Parse message type code
	code := make([]byte, 2)
	_, err = buf.Read(code)
	if err != nil {
		return "", err
	}

	return string(code), nil
}

// ------------------------------------------------------------------------------------------------
// TxId represents a Bitcoin transaction ID. (Double SHA256 of tx data)
type TxId struct {
	data [32]byte
}

var zeroTxId TxId

func TxIdFromBytes(data []byte) *TxId {
	var result TxId
	copy(result.data[:], data)
	return &result
}

// Validate returns an error if the value is invalid
func (id *TxId) Validate() error {
	return nil
}

// IsZero returns true if the tx id is all zeros.
func (id *TxId) IsZero() bool {
	return bytes.Equal(id.data[:], zeroTxId.data[:])
}

// Equal returns true if the specified values are the same.
func (id *TxId) Equal(other TxId) bool {
	return bytes.Equal(id.data[:], other.data[:])
}

// Bytes returns the byte slice for the TxId.
func (id *TxId) Bytes() []byte {
	return id.data[:]
}

// String converts to a string
func (id *TxId) String() string {
	return fmt.Sprintf("%x", id.data[:])
}

// Serialize returns a byte slice with the TxId in it.
func (id *TxId) Serialize() ([]byte, error) {
	return id.data[:], nil
}

// Write reads a TxId from a bytes.Buffer
func (id *TxId) Write(buf *bytes.Buffer) error {
	return readLen(buf, id.data[:])
}

// MarshalJSON converts to json.
func (id *TxId) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%x\"", id.data)), nil
}

// UnmarshalJSON converts from json.
func (id *TxId) UnmarshalJSON(data []byte) error {
	if len(data) < 2 {
		return fmt.Errorf("Too short for TxId hex data : %d", len(data))
	}
	n, err := hex.Decode(id.data[:], data[1:len(data)-1])
	if err != nil {
		return err
	}
	if n != 32 {
		return fmt.Errorf("Invalid TxId size : %d", n)
	}
	return nil
}

// Set sets the value specified
func (id *TxId) Set(value []byte) error {
	if len(value) != len(id.data) {
		return errors.New("Invalid value length for txid")
	}
	copy(id.data[:], value)
	return nil
}

// ------------------------------------------------------------------------------------------------
// PublicKeyHash represents a Bitcoin Public Key Hash. Often used as an address to receive transactions.
type PublicKeyHash struct {
	data [20]byte
}

var zeroPKH PublicKeyHash

// Validate returns an error if the value is invalid
func (hash *PublicKeyHash) Validate() error {
	return nil
}

// IsZero returns true if the tx id is all zeros.
func (hash *PublicKeyHash) IsZero() bool {
	return bytes.Equal(hash.data[:], zeroPKH.data[:])
}

// Equal returns true if the specified values are the same.
func (hash *PublicKeyHash) Equal(other PublicKeyHash) bool {
	return bytes.Equal(hash.data[:], other.data[:])
}

// PublicKeyHashFromBytes returns a PublicKeyHash with the specified bytes.
func PublicKeyHashFromBytes(data []byte) *PublicKeyHash {
	var result PublicKeyHash
	copy(result.data[:], data)
	return &result
}

// Bytes returns a byte slice containing the PublicKeyHash.
func (hash *PublicKeyHash) Bytes() []byte {
	return hash.data[:]
}

// String converts to a string
func (hash *PublicKeyHash) String() string {
	return fmt.Sprintf("%x", hash.data[:])
}

// Serialize returns a byte slice with the PublicKeyHash in it.
func (hash *PublicKeyHash) Serialize() ([]byte, error) {
	return hash.data[:], nil
}

// Write reads a PublicKeyHash from a bytes.Buffer
func (hash *PublicKeyHash) Write(buf *bytes.Buffer) error {
	return readLen(buf, hash.data[:])
}

// MarshalJSON converts to json.
func (hash *PublicKeyHash) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%x\"", hash.data)), nil
}

// UnmarshalJSON converts from json.
func (hash *PublicKeyHash) UnmarshalJSON(data []byte) error {
	if len(data) < 2 {
		return fmt.Errorf("Too short for PublicKeyHash hex data : %d", len(data))
	}
	n, err := hex.Decode(hash.data[:], data[1:len(data)-1])
	if err != nil {
		return err
	}
	if n != 20 {
		return fmt.Errorf("Invalid PublicKeyHash size : %d", n)
	}
	return nil
}

// ------------------------------------------------------------------------------------------------
// AssetCode represents a unique identifier for a Tokenized asset.
type AssetCode struct {
	data [32]byte
}

// Validate returns an error if the value is invalid
func (code *AssetCode) Validate() error {
	return nil
}

// IsZero returns true if the AssetCode is all zeroes. (empty)
func (code *AssetCode) IsZero() bool {
	zero := make([]byte, 32, 32)
	return bytes.Equal(code.data[:], zero)
}

// Equal returns true if the specified asset code is the same value.
func (code *AssetCode) Equal(other AssetCode) bool {
	return bytes.Equal(code.data[:], other.data[:])
}

// AssetCodeFromContract generates a "unique" deterministic asset code from a contract public key
//   hash and an asset index.
func AssetCodeFromContract(contractPKH []byte, index uint64) *AssetCode {
	hash256 := sha256.New()
	hash256.Write(contractPKH)
	binary.Write(hash256, DefaultEndian, &index)
	hash := hash256.Sum(nil)

	var result AssetCode
	copy(result.data[:], hash)
	return &result
}

// AssetCodeFromBytes returns a AssetCode with the specified bytes.
func AssetCodeFromBytes(data []byte) *AssetCode {
	var result AssetCode
	copy(result.data[:], data)
	return &result
}

// Bytes returns a byte slice containing the AssetCode.
func (code *AssetCode) Bytes() []byte {
	return code.data[:]
}

// String converts to a string
func (code *AssetCode) String() string {
	return fmt.Sprintf("%x", code.data[:])
}

// Serialize returns a byte slice with the AssetCode in it.
func (code *AssetCode) Serialize() ([]byte, error) {
	return code.data[:], nil
}

// Write reads a AssetCode from a bytes.Buffer
func (code *AssetCode) Write(buf *bytes.Buffer) error {
	return readLen(buf, code.data[:])
}

// MarshalJSON converts to json.
func (code *AssetCode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%x\"", code.data)), nil
}

// UnmarshalJSON converts from json.
func (code *AssetCode) UnmarshalJSON(data []byte) error {
	if len(data) < 2 {
		return fmt.Errorf("Too short for AssetCode hex data : %d", len(data))
	}
	n, err := hex.Decode(code.data[:], data[1:len(data)-1])
	if err != nil {
		return err
	}
	if n != 32 {
		return fmt.Errorf("Invalid AssetCode size : %d", n)
	}
	return nil
}

// ------------------------------------------------------------------------------------------------
// ContractCode represents a unique identifier for a Tokenized static contract.
type ContractCode struct {
	data [32]byte
}

// Validate returns an error if the value is invalid
func (code *ContractCode) Validate() error {
	return nil
}

// IsZero returns true if the ContractCode is all zeroes. (empty)
func (code *ContractCode) IsZero() bool {
	zero := make([]byte, 32, 32)
	return bytes.Equal(code.data[:], zero)
}

// Equal returns true if the specified values are the same.
func (code *ContractCode) Equal(other ContractCode) bool {
	return bytes.Equal(code.data[:], other.data[:])
}

// AssetCodeFromBytes returns a ContractCode with the specified bytes.
func ContractCodeFromBytes(data []byte) *ContractCode {
	var result ContractCode
	copy(result.data[:], data)
	return &result
}

// Bytes returns a byte slice containing the ContractCode.
func (code *ContractCode) Bytes() []byte {
	return code.data[:]
}

// String converts to a string
func (code *ContractCode) String() string {
	return fmt.Sprintf("%x", code.data[:])
}

// Serialize returns a byte slice with the ContractCode in it.
func (code *ContractCode) Serialize() ([]byte, error) {
	return code.data[:], nil
}

// Write reads a ContractCode from a bytes.Buffer
func (code *ContractCode) Write(buf *bytes.Buffer) error {
	return readLen(buf, code.data[:])
}

// MarshalJSON converts to json.
func (code *ContractCode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%x\"", code.data)), nil
}

// UnmarshalJSON converts from json.
func (code *ContractCode) UnmarshalJSON(data []byte) error {
	if len(data) < 2 {
		return fmt.Errorf("Too short for ContractCode hex data : %d", len(data))
	}
	n, err := hex.Decode(code.data[:], data[1:len(data)-1])
	if err != nil {
		return err
	}
	if n != 32 {
		return fmt.Errorf("Invalid ContractCode size : %d", n)
	}
	return nil
}

// ------------------------------------------------------------------------------------------------
// Timestamp represents a time in the Tokenized protocol.
type Timestamp struct {
	nanoseconds uint64 // nanoseconds since Unix epoch
}

// NewTimestamp returns a new timestamp from nanoseconds.
func NewTimestamp(value uint64) Timestamp {
	return Timestamp{nanoseconds: value}
}

// CurrentTimestamp returns a Timestamp containing the current time.
func CurrentTimestamp() Timestamp {
	return Timestamp{nanoseconds: uint64(time.Now().UnixNano())}
}

// Validate returns an error if the value is invalid
func (time *Timestamp) Validate() error {
	return nil
}

// Equal returns true if the specified values are the same.
func (time *Timestamp) Equal(other Timestamp) bool {
	return time.nanoseconds == other.nanoseconds
}

// Nano returns the nanoseconds since the Unix epoch for the Timestamp.
func (time *Timestamp) Nano() uint64 {
	return time.nanoseconds
}

// Nano returns the seconds since the Unix epoch for the Timestamp.
func (time *Timestamp) Seconds() uint32 {
	return uint32(time.nanoseconds / 1000000000)
}

// String converts to a string
func (t *Timestamp) String() string {
	return time.Unix(int64(t.nanoseconds)/1000000000, 0).String()
}

// Serialize returns a byte slice with the Timestamp in it.
func (time *Timestamp) Serialize() ([]byte, error) {
	buf := new(bytes.Buffer)
	if err := write(buf, &time.nanoseconds); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Write reads a Timestamp from a bytes.Buffer
func (time *Timestamp) Write(buf *bytes.Buffer) error {
	if err := read(buf, &time.nanoseconds); err != nil {
		return err
	}
	return nil
}

// MarshalJSON converts to json.
func (time *Timestamp) MarshalJSON() ([]byte, error) {
	return []byte(strconv.FormatUint(time.nanoseconds, 10)), nil
}

// UnmarshalJSON converts from json.
func (time *Timestamp) UnmarshalJSON(data []byte) error {
	if len(data) == 2 && data[0] == '{' && data[1] == '}' {
		time.nanoseconds = 0
		return nil
	}

	var err error
	time.nanoseconds, err = strconv.ParseUint(string(data), 10, 64)
	return err
}
