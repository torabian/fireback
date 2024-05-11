import Foundation
class RegionalContentEntity : Codable, Identifiable {
    var content: html? = nil
    var region: String? = nil
    var title: String? = nil
    var languageId: String? = nil
    var keyGroup: enum? = nil
}
class RegionalContentEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var region: String? = nil
  @Published var regionErrorMessage: String? = nil
  @Published var title: String? = nil
  @Published var titleErrorMessage: String? = nil
  @Published var languageId: String? = nil
  @Published var languageIdErrorMessage: String? = nil
  func getDto() -> RegionalContentEntity {
      var dto = RegionalContentEntity()
    dto.region = self.region
    dto.title = self.title
    dto.languageId = self.languageId
      return dto
  }
}