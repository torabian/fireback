import Foundation
class TokenEntity : Codable, Identifiable {
    var user: UserEntity? = nil
    // var userId: String? = nil
    var validUntil: String? = nil
}
class TokenEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var validUntil: String? = nil
  @Published var validUntilErrorMessage: String? = nil
  func getDto() -> TokenEntity {
      var dto = TokenEntity()
    dto.validUntil = self.validUntil
      return dto
  }
}