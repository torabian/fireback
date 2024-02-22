import Foundation
class PriceTagVariations : Codable, Identifiable {
    var currency: CurrencyEntity? = nil
    // var currencyId: String? = nil
    var amount: Float64? = nil
}
class PriceTagEntity : Codable, Identifiable {
    var variations: [PriceTagVariations]? = nil
}
class PriceTagEntityViewModel: ObservableObject {
    // improve the fields here
    func getDto() -> PriceTagEntity {
        var dto = PriceTagEntity()
        return dto
    }
}