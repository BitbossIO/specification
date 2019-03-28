package protocol

var (
	RejectionCodes = map[uint8][]byte{
		RejectionCodeInsufficientValue:          []byte("Fee Not Paid"),
		RejectionCodeMalformed:                  []byte("Malformed Request"),
		RejectionCodeContractAddress:            []byte("Contract Address"),
		RejectionCodeIssuerAddress:              []byte("Issuer Address"),
		RejectionCodeDuplicateAssetCode:         []byte("Duplicate Asset Code"),
		RejectionCodeFixedQuantity:              []byte("Fixed Quantity"),
		RejectionCodeContractExists:             []byte("Contract Exists"),
		RejectionCodeContractNotDynamic:         []byte("Contract Not Dynamic"),
		RejectionCodeContractQtyReduction:       []byte("Contract Qty Reduction"),
		RejectionCodeContractAuthFlags:          []byte("Contract Auth Flags"),
		RejectionCodeAssetAuthFlags:             []byte("Asset Auth Flags"),
		RejectionCodeContractExpiration:         []byte("Contract Expiration"),
		RejectionCodeContractUpdate:             []byte("Contract Update"),
		RejectionCodeVoteExists:                 []byte("Vote Exists"),
		RejectionCodeVoteNotFound:               []byte("Vote Not Found"),
		RejectionCodeVoteClosed:                 []byte("Vote Closed"),
		RejectionCodeBallotCounted:              []byte("Ballot Counted"),
		RejectionCodeAssetNotFound:              []byte("Asset Not Found"),
		RejectionCodeInsufficientAssets:         []byte("Insufficient Assets"),
		RejectionCodeTransferSelf:               []byte("Transfer Self"),
		RejectionCodeReceiverUnspecified:        []byte("Receiver Unspecified"),
		RejectionCodeUnknownAddress:             []byte("Unknown Address"),
		RejectionCodeFrozen:                     []byte("Frozen"),
		RejectionCodeContractRevision:           []byte("Contract Revision Incorrect"),
		RejectionCodeAssetRevision:              []byte("Asset Revision Incorrect"),
		RejectionCodeContractMissingNewIssuer:   []byte("Contract Issuer Change Missing"),
		RejectionCodeContractMissingNewOperator: []byte("Contract Operator Change Missing"),
		RejectionCodeInvalidProposal:            []byte("Invalid Proposal"),
	}
)
