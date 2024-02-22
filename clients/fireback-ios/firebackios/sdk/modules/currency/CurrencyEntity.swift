import Foundation
class CurrencyEntity : Codable, Identifiable {
    var symbol: String? = nil
    var name: String? = nil
    var symbolNative: String? = nil
    var decimalDigits: Int? = nil
    var rounding: Int? = nil
    var code: String? = nil
    var namePlural: String? = nil
}
class CurrencyEntityViewModel: ObservableObject {
    // improve the fields here
    @Published var symbol: String? = nil
    @Published var name: String? = nil
    @Published var symbolNative: String? = nil
    @Published var decimalDigits: Int? = nil
    @Published var rounding: Int? = nil
    @Published var code: String? = nil
    @Published var namePlural: String? = nil
    func getDto() -> CurrencyEntity {
        var dto = CurrencyEntity()
        dto.symbol = self.symbol
        dto.name = self.name
        dto.symbolNative = self.symbolNative
        dto.decimalDigits = self.decimalDigits
        dto.rounding = self.rounding
        dto.code = self.code
        dto.namePlural = self.namePlural
        return dto
    }
}