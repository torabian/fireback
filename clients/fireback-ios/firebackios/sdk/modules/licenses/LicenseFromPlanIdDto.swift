import Foundation
struct LicenseFromPlanIdDto : Codable {
    var machineId: String? = nil
    var email: String? = nil
    var owner: String? = nil
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
class LicenseFromPlanIdDtoViewModel: ObservableObject {
    // improve the fields here
    @Published var machineId: String? = nil
    @Published var email: String? = nil
    @Published var owner: String? = nil
    func getDto() -> LicenseFromPlanIdDto {
        var dto = LicenseFromPlanIdDto()
        dto.machineId = self.machineId
        dto.email = self.email
        dto.owner = self.owner
        return dto
    }
}