import Foundation
class PreferenceEntity : Codable, Identifiable {
    var timezone: String? = nil
}
class PreferenceEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var timezone: String? = nil
  @Published var timezoneErrorMessage: String? = nil
  func getDto() -> PreferenceEntity {
      var dto = PreferenceEntity()
    dto.timezone = self.timezone
      return dto
  }
}