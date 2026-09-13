import logging

logger = logging.getLogger(__name__)


def login_bad(user, password):
    # ruleid: tqa-pii-sensitive-log-field-py
    logger.info("login attempt", extra={"user": user.id, "password": password})

    # ruleid: tqa-pii-sensitive-log-field-py
    logger.error("auth failed", access_token=user.token)


def login_good(user, password):
    # ok: tqa-pii-sensitive-log-field-py
    logger.info("login attempt", extra={"user_id": user.id, "password_hash": hash_pw(password)})

    # ok: tqa-pii-sensitive-log-field-py
    logger.error("auth failed", token_type="bearer", auth_method="oauth")

    # ok: tqa-pii-sensitive-log-field-py
    logger.info("payment", extra={"card_last4": card.last4})

    # ok: tqa-pii-sensitive-log-field-py
    logger.info("done", extra={"duration_ms": 12})