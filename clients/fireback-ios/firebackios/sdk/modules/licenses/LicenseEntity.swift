import Foundation
class LicensePermissions : Codable, Identifiable {
    var capability: CapabilityEntity? = nil
    // var capabilityId: String? = nil
}
class LicenseEntity : Codable, Identifiable {
    var name: String? = nil
    var signedLicense: String? = nil
    var validityStartDate: Date? = nil
    var validityEndDate: Date? = nil
    var permissions: [LicensePermissions]? = nil
}
class LicenseEntityViewModel: ObservableObject {
    // improve the fields here
    @Published var name: String? = nil
    @Published var signedLicense: String? = nil
    func getDto() -> LicenseEntity {
        var dto = LicenseEntity()
        dto.name = self.name
        dto.signedLicense = self.signedLicense
        return dto
    }
}