def record_bad(span, request):
    # ruleid: tqa-card-raw-path-attribute-py
    span.set_attribute("http.target", request.path)

    # ruleid: tqa-card-raw-path-attribute-py
    span.set_attribute("http.uri", request.full_path)


def record_good(span, route_pattern):
    # ok: tqa-card-raw-path-attribute-py
    span.set_attribute("http.route", route_pattern)