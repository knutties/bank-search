$version: "2.0"

namespace io.knutties.banksearch

/// Mixin holding the common bank-branch fields. Both the Branch lookup
/// response and a Search result item carry these.
@mixin
structure BranchFields {
    @required
    ifsc: String

    @required
    bank_code: String

    @required
    bank_name: String

    @required
    branch: String

    centre: String
    district: String
    state: String
    address: String
    city: String
    contact: String
    micr: String
    swift: String
    upi: Boolean
    neft: Boolean
    rtgs: Boolean
    imps: Boolean
}

/// Full branch record returned by GetBranch.
structure Branch with [BranchFields] {}

/// Look up a single branch by IFSC code.
@readonly
@http(method: "GET", uri: "/ifsc/{code}", code: 200)
operation GetBranch {
    input := {
        @httpLabel
        @required
        code: String
    }

    output: Branch

    errors: [BranchNotFound]
}

@error("client")
@httpError(404)
structure BranchNotFound {
    @required
    error: String
}
