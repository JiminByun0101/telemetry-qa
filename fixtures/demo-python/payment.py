import logging

from opentelemetry import trace

logger = logging.getLogger(__name__)
tracer = trace.get_tracer(__name__)


def charge_card(card_number, amount):
    try:
        _submit_charge(card_number, amount)
    except Exception:
        # NOTE: P003 - full local scope dumped, includes card_number
        logger.error("charge failed", locals())


def charge_card_good(card_number, amount):
    try:
        _submit_charge(card_number, amount)
    except Exception as e:
        # NOTE: clean - only non-sensitive context selected
        logger.error("charge failed", extra={"amount": amount, "error": str(e)})


def seed_test_account():
    # NOTE: P004 - obviously-invalid RRN literal used as seed data
    return {"rrn": "999999-1234567"}


def refund_bad():
    # NOTE: S004 - manual span created, never ended
    span = tracer.start_span("refund")
    _process_refund()


def refund_good():
    # NOTE: clean - context manager handles span lifecycle
    with tracer.start_as_current_span("refund"):
        _process_refund()