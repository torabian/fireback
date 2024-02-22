import Foundation
struct LicenseFromActivationKeyDto : Codable {
    var activationKeyId: String? = nil
    var machineId: String? = nil
    func toJson() -> String? {
        let encoder = JSONEncoder()
        encoder.outputFormatting = .prettyPrinted
        do {
            let jsonData = try encoder.encode(self)
            let jsonString = String(data: jsonData, encoding: .utf8)
            return jsonString
        } catch {
            print("Failed to convert struct to JSON: \(error)")
            return nil
        }
    }
}
class LicenseFromActivationKeyDtoViewModel: ObservableObject {
    // improve the fields here
    @Published var activationKeyId: String? = nil
    @Published var machineId: String? = nil
    func getDto() -> LicenseFromActivationKeyDto {
        var dto = LicenseFromActivationKeyDto()
        dto.activationKeyId = self.activationKeyId
        dto.machineId = self.machineId
        return dto
    }
}