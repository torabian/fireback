import Foundation
class LicensableProductEntity : Codable, Identifiable {
    var name: String? = nil
    var privateKey: String? = nil
    var publicKey: String? = nil
}
class LicensableProductEntityViewModel: ObservableObject {
    // improve the fields here
    @Published var name: String? = nil
    @Published var privateKey: String? = nil
    @Published var publicKey: String? = nil
    func getDto() -> LicensableProductEntity {
        var dto = LicensableProductEntity()
        dto.name = self.name
        dto.privateKey = self.privateKey
        dto.publicKey = self.publicKey
        return dto
    }
}