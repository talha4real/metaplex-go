// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package metaplex

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// NOTE: Requires an AuctionManagerV1.
// Note: This requires that auction manager be in a Running state.
//
// If an auction is complete, you can redeem your bid for a specific item here. If you are the first to do this,
// The auction manager will switch from Running state to Disbursing state. If you are the last, this may change
// the auction manager state to Finished provided that no authorities remain to be delegated for Master Edition tokens.
//
// NOTE: Please note that it is totally possible to redeem a bid 2x - once for a prize you won and once at the RedeemParticipationBid point for an open edition
// that comes as a 'token of appreciation' for bidding. They are not mutually exclusive unless explicitly set to be that way.
type RedeemBid struct {

	// [0] = [WRITE] auctionManager
	// ··········· Auction manager
	//
	// [1] = [WRITE] safetyDepositTokenStorage
	// ··········· Safety deposit token storage account
	//
	// [2] = [WRITE] destinationAccount
	// ··········· Destination account.
	//
	// [3] = [WRITE] bidRedemptionKey
	// ··········· Bid redemption key -
	// ··········· Just a PDA with seed ['metaplex', auction_key, bidder_metadata_key] that we will allocate to mark that you redeemed your bid
	//
	// [4] = [WRITE] safetyDepositBox
	// ··········· Safety deposit box account
	//
	// [5] = [WRITE] vaultAccount
	// ··········· Vault account
	//
	// [6] = [WRITE] fractionMint
	// ··········· Fraction mint of the vault
	//
	// [7] = [] auction
	// ··········· Auction
	//
	// [8] = [] bidderMetadata
	// ··········· Your BidderMetadata account
	//
	// [9] = [SIGNER] bidder
	// ··········· [optional] Your Bidder account - Only needs to be signer if payer does not own
	//
	// [10] = [SIGNER] payer
	// ··········· Payer
	//
	// [11] = [] tokenProgram
	// ··········· Token program
	//
	// [12] = [] tokenVaultProgram
	// ··········· Token Vault program
	//
	// [13] = [] tokenMetadataProgram
	// ··········· Token metadata program
	//
	// [14] = [] store
	// ··········· Store
	//
	// [15] = [] system
	// ··········· System
	//
	// [16] = [] rentSysvar
	// ··········· Rent sysvar
	//
	// [17] = [] pdaBasedTransferAuthority
	// ··········· PDA-based Transfer authority to move the tokens from the store to the destination seed ['vault', program_id, vault key]
	// ··········· but please note that this is a PDA relative to the Token Vault program, with the 'vault' prefix
	//
	// [18] = [WRITE] masterEdition
	// ··········· Master edition (if Printing type of WinningConfig)
	//
	// [19] = [WRITE] reservationListPDA
	// ··········· Reservation list PDA ['metadata', program id, master edition key, 'reservation', auction manager key]
	// ··········· relative to token metadata program (if Printing type of WinningConfig)
	//
	// [20] = [] safetyDepositConfig
	// ··········· Safety deposit config pda of ['metaplex', program id, auction manager, safety deposit]
	// ··········· This account will only get used AND BE REQUIRED in the event this is an AuctionManagerV2
	//
	// [21] = [] auctionExtendedPDA
	// ··········· Auction extended (pda relative to auction of ['auction', program id, vault key, 'extended'])
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewRedeemBidInstructionBuilder creates a new `RedeemBid` instruction builder.
func NewRedeemBidInstructionBuilder() *RedeemBid {
	nd := &RedeemBid{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 22),
	}
	return nd
}

// SetAuctionManagerAccount sets the "auctionManager" account.
// Auction manager
func (inst *RedeemBid) SetAuctionManagerAccount(auctionManager ag_solanago.PublicKey) *RedeemBid {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(auctionManager).WRITE()
	return inst
}

// GetAuctionManagerAccount gets the "auctionManager" account.
// Auction manager
func (inst *RedeemBid) GetAuctionManagerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetSafetyDepositTokenStorageAccount sets the "safetyDepositTokenStorage" account.
// Safety deposit token storage account
func (inst *RedeemBid) SetSafetyDepositTokenStorageAccount(safetyDepositTokenStorage ag_solanago.PublicKey) *RedeemBid {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(safetyDepositTokenStorage).WRITE()
	return inst
}

// GetSafetyDepositTokenStorageAccount gets the "safetyDepositTokenStorage" account.
// Safety deposit token storage account
func (inst *RedeemBid) GetSafetyDepositTokenStorageAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetDestinationAccount sets the "destinationAccount" account.
// Destination account.
func (inst *RedeemBid) SetDestinationAccount(destinationAccount ag_solanago.PublicKey) *RedeemBid {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(destinationAccount).WRITE()
	return inst
}

// GetDestinationAccount gets the "destinationAccount" account.
// Destination account.
func (inst *RedeemBid) GetDestinationAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetBidRedemptionKeyAccount sets the "bidRedemptionKey" account.
// Bid redemption key -
// Just a PDA with seed ['metaplex', auction_key, bidder_metadata_key] that we will allocate to mark that you redeemed your bid
func (inst *RedeemBid) SetBidRedemptionKeyAccount(bidRedemptionKey ag_solanago.PublicKey) *RedeemBid {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(bidRedemptionKey).WRITE()
	return inst
}

// GetBidRedemptionKeyAccount gets the "bidRedemptionKey" account.
// Bid redemption key -
// Just a PDA with seed ['metaplex', auction_key, bidder_metadata_key] that we will allocate to mark that you redeemed your bid
func (inst *RedeemBid) GetBidRedemptionKeyAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetSafetyDepositBoxAccount sets the "safetyDepositBox" account.
// Safety deposit box account
func (inst *RedeemBid) SetSafetyDepositBoxAccount(safetyDepositBox ag_solanago.PublicKey) *RedeemBid {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(safetyDepositBox).WRITE()
	return inst
}

// GetSafetyDepositBoxAccount gets the "safetyDepositBox" account.
// Safety deposit box account
func (inst *RedeemBid) GetSafetyDepositBoxAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetVaultAccount sets the "vaultAccount" account.
// Vault account
func (inst *RedeemBid) SetVaultAccount(vaultAccount ag_solanago.PublicKey) *RedeemBid {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(vaultAccount).WRITE()
	return inst
}

// GetVaultAccount gets the "vaultAccount" account.
// Vault account
func (inst *RedeemBid) GetVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetFractionMintAccount sets the "fractionMint" account.
// Fraction mint of the vault
func (inst *RedeemBid) SetFractionMintAccount(fractionMint ag_solanago.PublicKey) *RedeemBid {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(fractionMint).WRITE()
	return inst
}

// GetFractionMintAccount gets the "fractionMint" account.
// Fraction mint of the vault
func (inst *RedeemBid) GetFractionMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

// SetAuctionAccount sets the "auction" account.
// Auction
func (inst *RedeemBid) SetAuctionAccount(auction ag_solanago.PublicKey) *RedeemBid {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(auction)
	return inst
}

// GetAuctionAccount gets the "auction" account.
// Auction
func (inst *RedeemBid) GetAuctionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[7]
}

// SetBidderMetadataAccount sets the "bidderMetadata" account.
// Your BidderMetadata account
func (inst *RedeemBid) SetBidderMetadataAccount(bidderMetadata ag_solanago.PublicKey) *RedeemBid {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(bidderMetadata)
	return inst
}

// GetBidderMetadataAccount gets the "bidderMetadata" account.
// Your BidderMetadata account
func (inst *RedeemBid) GetBidderMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[8]
}

// SetBidderAccount sets the "bidder" account.
// [optional] Your Bidder account - Only needs to be signer if payer does not own
func (inst *RedeemBid) SetBidderAccount(bidder ag_solanago.PublicKey) *RedeemBid {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(bidder).SIGNER()
	return inst
}

// GetBidderAccount gets the "bidder" account.
// [optional] Your Bidder account - Only needs to be signer if payer does not own
func (inst *RedeemBid) GetBidderAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[9]
}

// SetPayerAccount sets the "payer" account.
// Payer
func (inst *RedeemBid) SetPayerAccount(payer ag_solanago.PublicKey) *RedeemBid {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// Payer
func (inst *RedeemBid) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[10]
}

// SetTokenProgramAccount sets the "tokenProgram" account.
// Token program
func (inst *RedeemBid) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *RedeemBid {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
// Token program
func (inst *RedeemBid) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[11]
}

// SetTokenVaultProgramAccount sets the "tokenVaultProgram" account.
// Token Vault program
func (inst *RedeemBid) SetTokenVaultProgramAccount(tokenVaultProgram ag_solanago.PublicKey) *RedeemBid {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(tokenVaultProgram)
	return inst
}

// GetTokenVaultProgramAccount gets the "tokenVaultProgram" account.
// Token Vault program
func (inst *RedeemBid) GetTokenVaultProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[12]
}

// SetTokenMetadataProgramAccount sets the "tokenMetadataProgram" account.
// Token metadata program
func (inst *RedeemBid) SetTokenMetadataProgramAccount(tokenMetadataProgram ag_solanago.PublicKey) *RedeemBid {
	inst.AccountMetaSlice[13] = ag_solanago.Meta(tokenMetadataProgram)
	return inst
}

// GetTokenMetadataProgramAccount gets the "tokenMetadataProgram" account.
// Token metadata program
func (inst *RedeemBid) GetTokenMetadataProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[13]
}

// SetStoreAccount sets the "store" account.
// Store
func (inst *RedeemBid) SetStoreAccount(store ag_solanago.PublicKey) *RedeemBid {
	inst.AccountMetaSlice[14] = ag_solanago.Meta(store)
	return inst
}

// GetStoreAccount gets the "store" account.
// Store
func (inst *RedeemBid) GetStoreAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[14]
}

// SetSystemAccount sets the "system" account.
// System
func (inst *RedeemBid) SetSystemAccount(system ag_solanago.PublicKey) *RedeemBid {
	inst.AccountMetaSlice[15] = ag_solanago.Meta(system)
	return inst
}

// GetSystemAccount gets the "system" account.
// System
func (inst *RedeemBid) GetSystemAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[15]
}

// SetRentSysvarAccount sets the "rentSysvar" account.
// Rent sysvar
func (inst *RedeemBid) SetRentSysvarAccount(rentSysvar ag_solanago.PublicKey) *RedeemBid {
	inst.AccountMetaSlice[16] = ag_solanago.Meta(rentSysvar)
	return inst
}

// GetRentSysvarAccount gets the "rentSysvar" account.
// Rent sysvar
func (inst *RedeemBid) GetRentSysvarAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[16]
}

// SetPdaBasedTransferAuthorityAccount sets the "pdaBasedTransferAuthority" account.
// PDA-based Transfer authority to move the tokens from the store to the destination seed ['vault', program_id, vault key]
// but please note that this is a PDA relative to the Token Vault program, with the 'vault' prefix
func (inst *RedeemBid) SetPdaBasedTransferAuthorityAccount(pdaBasedTransferAuthority ag_solanago.PublicKey) *RedeemBid {
	inst.AccountMetaSlice[17] = ag_solanago.Meta(pdaBasedTransferAuthority)
	return inst
}

// GetPdaBasedTransferAuthorityAccount gets the "pdaBasedTransferAuthority" account.
// PDA-based Transfer authority to move the tokens from the store to the destination seed ['vault', program_id, vault key]
// but please note that this is a PDA relative to the Token Vault program, with the 'vault' prefix
func (inst *RedeemBid) GetPdaBasedTransferAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[17]
}

// SetMasterEditionAccount sets the "masterEdition" account.
// Master edition (if Printing type of WinningConfig)
func (inst *RedeemBid) SetMasterEditionAccount(masterEdition ag_solanago.PublicKey) *RedeemBid {
	inst.AccountMetaSlice[18] = ag_solanago.Meta(masterEdition).WRITE()
	return inst
}

// GetMasterEditionAccount gets the "masterEdition" account.
// Master edition (if Printing type of WinningConfig)
func (inst *RedeemBid) GetMasterEditionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[18]
}

// SetReservationListPDAAccount sets the "reservationListPDA" account.
// Reservation list PDA ['metadata', program id, master edition key, 'reservation', auction manager key]
// relative to token metadata program (if Printing type of WinningConfig)
func (inst *RedeemBid) SetReservationListPDAAccount(reservationListPDA ag_solanago.PublicKey) *RedeemBid {
	inst.AccountMetaSlice[19] = ag_solanago.Meta(reservationListPDA).WRITE()
	return inst
}

// GetReservationListPDAAccount gets the "reservationListPDA" account.
// Reservation list PDA ['metadata', program id, master edition key, 'reservation', auction manager key]
// relative to token metadata program (if Printing type of WinningConfig)
func (inst *RedeemBid) GetReservationListPDAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[19]
}

// SetSafetyDepositConfigAccount sets the "safetyDepositConfig" account.
// Safety deposit config pda of ['metaplex', program id, auction manager, safety deposit]
// This account will only get used AND BE REQUIRED in the event this is an AuctionManagerV2
func (inst *RedeemBid) SetSafetyDepositConfigAccount(safetyDepositConfig ag_solanago.PublicKey) *RedeemBid {
	inst.AccountMetaSlice[20] = ag_solanago.Meta(safetyDepositConfig)
	return inst
}

// GetSafetyDepositConfigAccount gets the "safetyDepositConfig" account.
// Safety deposit config pda of ['metaplex', program id, auction manager, safety deposit]
// This account will only get used AND BE REQUIRED in the event this is an AuctionManagerV2
func (inst *RedeemBid) GetSafetyDepositConfigAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[20]
}

// SetAuctionExtendedPDAAccount sets the "auctionExtendedPDA" account.
// Auction extended (pda relative to auction of ['auction', program id, vault key, 'extended'])
func (inst *RedeemBid) SetAuctionExtendedPDAAccount(auctionExtendedPDA ag_solanago.PublicKey) *RedeemBid {
	inst.AccountMetaSlice[21] = ag_solanago.Meta(auctionExtendedPDA)
	return inst
}

// GetAuctionExtendedPDAAccount gets the "auctionExtendedPDA" account.
// Auction extended (pda relative to auction of ['auction', program id, vault key, 'extended'])
func (inst *RedeemBid) GetAuctionExtendedPDAAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[21]
}

func (inst RedeemBid) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_RedeemBid),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst RedeemBid) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *RedeemBid) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.AuctionManager is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.SafetyDepositTokenStorage is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.DestinationAccount is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.BidRedemptionKey is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.SafetyDepositBox is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.VaultAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.FractionMint is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Auction is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.BidderMetadata is not set")
		}

		// [9] = Bidder is optional

		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.TokenVaultProgram is not set")
		}
		if inst.AccountMetaSlice[13] == nil {
			return errors.New("accounts.TokenMetadataProgram is not set")
		}
		if inst.AccountMetaSlice[14] == nil {
			return errors.New("accounts.Store is not set")
		}
		if inst.AccountMetaSlice[15] == nil {
			return errors.New("accounts.System is not set")
		}
		if inst.AccountMetaSlice[16] == nil {
			return errors.New("accounts.RentSysvar is not set")
		}
		if inst.AccountMetaSlice[17] == nil {
			return errors.New("accounts.PdaBasedTransferAuthority is not set")
		}

		// [18] = MasterEdition is optional

		// [19] = ReservationListPDA is optional

		if inst.AccountMetaSlice[20] == nil {
			return errors.New("accounts.SafetyDepositConfig is not set")
		}
		if inst.AccountMetaSlice[21] == nil {
			return errors.New("accounts.AuctionExtendedPDA is not set")
		}
	}
	return nil
}

func (inst *RedeemBid) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("RedeemBid")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=22]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("           auctionManager", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("safetyDepositTokenStorage", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("              destination", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("         bidRedemptionKey", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("         safetyDepositBox", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("                    vault", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("             fractionMint", inst.AccountMetaSlice[6]))
						accountsBranch.Child(ag_format.Meta("                  auction", inst.AccountMetaSlice[7]))
						accountsBranch.Child(ag_format.Meta("           bidderMetadata", inst.AccountMetaSlice[8]))
						accountsBranch.Child(ag_format.Meta("                   bidder", inst.AccountMetaSlice[9]))
						accountsBranch.Child(ag_format.Meta("                    payer", inst.AccountMetaSlice[10]))
						accountsBranch.Child(ag_format.Meta("             tokenProgram", inst.AccountMetaSlice[11]))
						accountsBranch.Child(ag_format.Meta("        tokenVaultProgram", inst.AccountMetaSlice[12]))
						accountsBranch.Child(ag_format.Meta("     tokenMetadataProgram", inst.AccountMetaSlice[13]))
						accountsBranch.Child(ag_format.Meta("                    store", inst.AccountMetaSlice[14]))
						accountsBranch.Child(ag_format.Meta("                   system", inst.AccountMetaSlice[15]))
						accountsBranch.Child(ag_format.Meta("               rentSysvar", inst.AccountMetaSlice[16]))
						accountsBranch.Child(ag_format.Meta("pdaBasedTransferAuthority", inst.AccountMetaSlice[17]))
						accountsBranch.Child(ag_format.Meta("            masterEdition", inst.AccountMetaSlice[18]))
						accountsBranch.Child(ag_format.Meta("       reservationListPDA", inst.AccountMetaSlice[19]))
						accountsBranch.Child(ag_format.Meta("      safetyDepositConfig", inst.AccountMetaSlice[20]))
						accountsBranch.Child(ag_format.Meta("       auctionExtendedPDA", inst.AccountMetaSlice[21]))
					})
				})
		})
}

func (obj RedeemBid) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *RedeemBid) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewRedeemBidInstruction declares a new RedeemBid instruction with the provided parameters and accounts.
func NewRedeemBidInstruction(
	// Accounts:
	auctionManager ag_solanago.PublicKey,
	safetyDepositTokenStorage ag_solanago.PublicKey,
	destinationAccount ag_solanago.PublicKey,
	bidRedemptionKey ag_solanago.PublicKey,
	safetyDepositBox ag_solanago.PublicKey,
	vaultAccount ag_solanago.PublicKey,
	fractionMint ag_solanago.PublicKey,
	auction ag_solanago.PublicKey,
	bidderMetadata ag_solanago.PublicKey,
	bidder ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	tokenVaultProgram ag_solanago.PublicKey,
	tokenMetadataProgram ag_solanago.PublicKey,
	store ag_solanago.PublicKey,
	system ag_solanago.PublicKey,
	rentSysvar ag_solanago.PublicKey,
	pdaBasedTransferAuthority ag_solanago.PublicKey,
	masterEdition ag_solanago.PublicKey,
	reservationListPDA ag_solanago.PublicKey,
	safetyDepositConfig ag_solanago.PublicKey,
	auctionExtendedPDA ag_solanago.PublicKey) *RedeemBid {
	return NewRedeemBidInstructionBuilder().
		SetAuctionManagerAccount(auctionManager).
		SetSafetyDepositTokenStorageAccount(safetyDepositTokenStorage).
		SetDestinationAccount(destinationAccount).
		SetBidRedemptionKeyAccount(bidRedemptionKey).
		SetSafetyDepositBoxAccount(safetyDepositBox).
		SetVaultAccount(vaultAccount).
		SetFractionMintAccount(fractionMint).
		SetAuctionAccount(auction).
		SetBidderMetadataAccount(bidderMetadata).
		SetBidderAccount(bidder).
		SetPayerAccount(payer).
		SetTokenProgramAccount(tokenProgram).
		SetTokenVaultProgramAccount(tokenVaultProgram).
		SetTokenMetadataProgramAccount(tokenMetadataProgram).
		SetStoreAccount(store).
		SetSystemAccount(system).
		SetRentSysvarAccount(rentSysvar).
		SetPdaBasedTransferAuthorityAccount(pdaBasedTransferAuthority).
		SetMasterEditionAccount(masterEdition).
		SetReservationListPDAAccount(reservationListPDA).
		SetSafetyDepositConfigAccount(safetyDepositConfig).
		SetAuctionExtendedPDAAccount(auctionExtendedPDA)
}
