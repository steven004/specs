package deal

type StorageDealPaymentAction = int

const (
	ExpireStorageDeals StorageDealPaymentAction = 0
	CreditStorageDeals StorageDealPaymentAction = 1
)

type StorageDealSlashAction = int

const (
	SlashTerminatedFaults StorageDealSlashAction = 0
)
