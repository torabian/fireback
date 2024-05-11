import Foundation
class GsmProviderEntity : Codable, Identifiable {
    var apiKey: String? = nil
    var mainSenderNumber: String? = nil
//    var type: enum? = nil
    var invokeUrl: String? = nil
    var invokeBody: String? = nil
}
class GsmProviderEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var apiKey: String? = nil
  @Published var apiKeyErrorMessage: String? = nil
  @Published var mainSenderNumber: String? = nil
  @Published var mainSenderNumberErrorMessage: String? = nil
  @Published var invokeUrl: String? = nil
  @Published var invokeUrlErrorMessage: String? = nil
  @Published var invokeBody: String? = nil
  @Published var invokeBodyErrorMessage: String? = nil
  func getDto() -> GsmProviderEntity {
      var dto = GsmProviderEntity()
    dto.apiKey = self.apiKey
    dto.mainSenderNumber = self.mainSenderNumber
    dto.invokeUrl = self.invokeUrl
    dto.invokeBody = self.invokeBody
      return dto
  }
}
