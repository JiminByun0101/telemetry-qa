import time
import uuid
from datetime import datetime


def record_bad(m):
    # ruleid: tqa-card-timestamp-label-value-py
    m.labels(str(datetime.now()))

    # ruleid: tqa-card-timestamp-label-value-py
    m.labels(str(uuid.uuid4()))

    # ruleid: tqa-card-timestamp-label-value-py
    m.labels(str(time.time()))

def record_good(m, route, status):
    # ok: tqa-card-timestamp-label-value-py
    m.labels(route, str(status))