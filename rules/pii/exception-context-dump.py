import logging

logger = logging.getLogger(__name__)


def process_payment_bad(card_number, amount):
    try:
        charge(card_number, amount)
    except Exception:
        # ruleid: tqa-pii-exception-context-dump-py
        logger.error("payment failed", locals())


def process_payment_bad2(user_token):
    try:
        validate(user_token)
    except Exception:
        # ruleid: tqa-pii-exception-context-dump-py
        logger.exception("validation failed", locals())


def process_payment_good(card_number, amount):
    try:
        charge(card_number, amount)
    except Exception as e:
        # ok: tqa-pii-exception-context-dump-py
        logger.error("payment failed", extra={"amount": amount, "error": str(e)})


def unrelated_function(data):
    # ok: tqa-pii-exception-context-dump-py
    logger.error("outside except block", locals())