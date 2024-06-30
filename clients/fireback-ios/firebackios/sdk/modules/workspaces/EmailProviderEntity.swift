import Foundation
enum EmailProviderEntityType : Codable {
  case terminal
  case sendgrid
}
class EmailProviderEntity : Codable, Identifiable {
    var type: EmailProviderEntityType? = nil
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