$version: "2.0"

namespace io.knutties.banksearch

/// Lightweight liveness probe for load balancers.
@readonly
@http(method: "GET", uri: "/healthz", code: 200)
operation Healthz {
    output := {
        @required
        status: String
    }
}

/// Index version metadata and document count.
@readonly
@http(method: "GET", uri: "/status", code: 200)
operation Status {
    output := {
        @required
        status: String

        indexed_docs: Long
        release_tag: String
        rbi_update_date: String
        built_at: String
    }
}
