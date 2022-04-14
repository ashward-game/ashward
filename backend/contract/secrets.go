package contract

type clientSecrets struct {
	Mnemonic            string
	DerivationPathOwner string
}

func NewClientSecret(mnemonic string,
	derivationPathOwner string) *clientSecrets {
	return &clientSecrets{
		Mnemonic:            mnemonic,
		DerivationPathOwner: derivationPathOwner,
	}
}
