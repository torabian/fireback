import Foundation
class PhoneConfirmationEntity : Codable, Identifiable {
    var user: UserEntity? = nil
    // var userId: String? = nil
    var status: String? = nil
    var phoneNumber: String? = nil
    var key: String? = nil
    var expiresAt: String? = nil
}
class PhoneConfirmationEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var status: String? = nil
  @Published var statusErrorMessage: String? = nil
  @Published var phoneNumber: String? = nil
  @Published var phoneNumberErrorMessage: String? = nil
  @Published var key: String? = nil
  @Published var keyErrorMessage: String? = nil
  @Published var expiresAt: String? = nil
  @Published var expiresAtErrorMessage: String? = nil
  func getDto() -> PhoneConfirmationEntity {
      var dto = PhoneConfirmationEntity()
    dto.status = self.status
    dto.phoneNumber = self.phoneNumber
    dto.key = self.key
    dto.expiresAt = self.expiresAt
      return dto
  }
}