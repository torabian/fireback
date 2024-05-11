import Foundation
class EmailProviderEntity : Codable, Identifiable {
    var type: String? = nil
    var apiKey: String? = nil
}
class EmailProviderEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var apiKey: String? = nil
  @Published var apiKeyErrorMessage: String? = nil
  func getDto() -> EmailProviderEntity {
      var dto = EmailProviderEntity()
    dto.apiKey = self.apiKey
      return dto
  }
}