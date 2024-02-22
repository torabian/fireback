import Foundation
struct ReactiveSearchResultDto : Codable {
    var uniqueId: String? = nil
    var phrase: String? = nil
    var icon: String? = nil
    var description: String? = nil
    var group: String? = nil
    var uiLocation: String? = nil
    var actionFn: String? = nil
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
class ReactiveSearchResultDtoViewModel: ObservableObject {
  // improve the fields here
  @Published var uniqueId: String? = nil
  @Published var uniqueIdErrorMessage: String? = nil
  @Published var phrase: String? = nil
  @Published var phraseErrorMessage: String? = nil
  @Published var icon: String? = nil
  @Published var iconErrorMessage: String? = nil
  @Published var description: String? = nil
  @Published var descriptionErrorMessage: String? = nil
  @Published var group: String? = nil
  @Published var groupErrorMessage: String? = nil
  @Published var uiLocation: String? = nil
  @Published var uiLocationErrorMessage: String? = nil
  @Published var actionFn: String? = nil
  @Published var actionFnErrorMessage: String? = nil
  func getDto() -> ReactiveSearchResultDto {
      var dto = ReactiveSearchResultDto()
    dto.uniqueId = self.uniqueId
    dto.phrase = self.phrase
    dto.icon = self.icon
    dto.description = self.description
    dto.group = self.group
    dto.uiLocation = self.uiLocation
    dto.actionFn = self.actionFn
      return dto
  }
}
