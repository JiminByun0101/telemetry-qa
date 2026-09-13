import logging

logger = logging.getLogger(__name__)


class Request:
    def __init__(self, user_id, path):
        self.user_id = user_id
        self.path = path


def log_bad(req):
    # ruleid: tqa-pii-object-dump-py
    logger.info("handling request", req.__dict__)


def log_bad2(req):
    # ruleid: tqa-pii-object-dump-py
    logger.info("request", vars(req))


def log_bad3(request):
    # ruleid: tqa-pii-object-dump-py
    logger.info("incoming headers", request.headers)


def log_bad4(request):
    # ruleid: tqa-pii-object-dump-py
    logger.info("incoming body", request.body)


def log_good(req):
    # ok: tqa-pii-object-dump-py
    logger.info("handling request", extra={"path": req.path})


def log_good2(status):
    # ok: tqa-pii-object-dump-py
    logger.info("responded with status %s", status)