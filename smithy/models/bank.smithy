$version: "2.0"

namespace io.knutties.banksearch

/// A bank registered with the index. Identified by its 4-char IFSC bank
/// code. Today the only bound operation is List; future Get/Put/Delete
/// will hang off this resource without changing existing URLs.
//
// `properties` is intentionally omitted: Smithy requires every declared
// property to be referenced by a create or instance (read/update/delete)
// operation, and we have none yet. When GetBank lands, `bank_name` should
// move from BankSummary into a `properties` block here, and BankSummary
// can be re-projected from the resource via `for Bank` + `$bank_name`.
resource Bank {
    identifiers: {
        bank_code: String
    }

    list: ListBanks
}

/// List the distinct banks present in the index, sorted by bank_code.
@readonly
@http(method: "GET", uri: "/list", code: 200)
operation ListBanks {
    output := {
        @required
        total: Integer

        @required
        banks: BankList
    }
}

list BankList {
    member: BankSummary
}

/// Wire shape for a bank in the listing.
structure BankSummary {
    @required
    bank_code: String

    @required
    bank_name: String
}
