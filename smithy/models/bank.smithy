$version: "2.0"

namespace io.knutties.banksearch

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
    member: Bank
}

structure Bank {
    @required
    bank_code: String

    @required
    bank_name: String
}
