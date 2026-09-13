from opentelemetry import trace

tracer = trace.get_tracer(__name__)


def bad_manual():
    # ruleid: tqa-silent-span-manual-not-ended-py
    span = tracer.start_span("bad_manual")
    do_work()
    return "done"


def good_context_manager():
    # ok: tqa-silent-span-manual-not-ended-py
    with tracer.start_as_current_span("good_cm") as span:
        do_work()
    return "done"


def good_manual_ended():
    # ok: tqa-silent-span-manual-not-ended-py
    span = tracer.start_span("good_manual")
    try:
        do_work()
    finally:
        span.end()
    return "done"


def good_with_wrapped():
    # ok: tqa-silent-span-manual-not-ended-py
    span = tracer.start_span("good_wrapped")
    with span:
        do_work()
    return "done"