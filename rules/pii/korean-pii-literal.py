# ruleid: tqa-pii-rrn-literal
SAMPLE_RRN = "999999-1234567"


def seed_test_user():
    # ruleid: tqa-pii-rrn-literal
    notes = "test account, rrn 991231-2345678 for QA"
    return notes


# ok: tqa-pii-rrn-literal
ORDER_ID = "202609-9800045"


def charge_card():
    # ruleid: tqa-pii-card-number-literal
    test_card = "4111-1111-1111-1111"
    return test_card


def track_shipment():
    # ok: tqa-pii-card-number-literal
    tracking_number = "TRK-58139204"
    return tracking_number