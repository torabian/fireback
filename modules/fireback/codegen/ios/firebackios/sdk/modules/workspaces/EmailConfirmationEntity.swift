import Foundation
class EmailConfirmationEntity : Codable, Identifiable {
    var user: UserEntity? = nil
    // var userId: String? = nil
    var status: String? = nil
    var email: String? = nil
    var key: String? = nil
    var expiresAt: String? = nil
}
class EmailConfirmationEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var status: String? = nil
  @Published var statusErrorMessage: String? = nil
  @Published var email: String? = nil
  @Published var emailErrorMessage: String? = nil
  @Published var key: String? = nil
  @Published var keyErrorMessage: String? = nil
  @Published var expiresAt: String? = nil
  @Published var expiresAtErrorMessage: String? = nil
  func getDto() -> EmailConfirmationEntity {
      var dto = EmailConfirmationEntity()
    dto.status = self.status
    dto.email = self.email
    dto.key = self.key
    dto.expiresAt = self.expiresAt
      return dto
  }
}