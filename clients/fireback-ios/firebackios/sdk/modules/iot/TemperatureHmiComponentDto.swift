import Foundation
struct TemperatureHmiComponentDto : Codable {
    var viewMode: String? = nil
    var units: String? = nil
    var maximumTemperature: Float64? = nil
    var minimumTemperature: Float64? = nil
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
class TemperatureHmiComponentDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var viewMode: String? = nil
  @Published var viewModeErrorMessage: String? = nil
  @Published var units: String? = nil
  @Published var unitsErrorMessage: String? = nil
  func getDto() -> TemperatureHmiComponentDto {
      var dto = TemperatureHmiComponentDto()
    dto.viewMode = self.viewMode
    dto.units = self.units
      return dto
  }
}
