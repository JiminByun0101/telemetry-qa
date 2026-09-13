package fixtures

// ruleid: tqa-pii-rrn-literal
const sampleRRN = "999999-1234567"

func seedTestUser() {
	// ruleid: tqa-pii-rrn-literal
	notes := "test account, rrn 991231-2345678 for QA"
	_ = notes
}

// ok: tqa-pii-rrn-literal
const orderID = "202609-9800045"

func chargeCard() {
	// ruleid: tqa-pii-card-number-literal
	testCard := "4111-1111-1111-1111"
	_ = testCard
}

func trackShipment() {
	// ok: tqa-pii-card-number-literal
	trackingNumber := "TRK-58139204"
	_ = trackingNumber
}