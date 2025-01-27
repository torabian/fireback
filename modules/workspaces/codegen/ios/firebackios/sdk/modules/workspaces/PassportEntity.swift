import Foundation
class PassportEntity : Codable, Identifiable {
    var type: String? = nil
    var user: UserEntity? = nil
    // var userId: String? = nil
    var value: String? = nil
    var password: String? = nil
    var confirmed: Bool? = nil
    var accessToken: String? = nil
}
class PassportEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var type: String? = nil
  @Published var typeErrorMessage: String? = nil
  @Published var value: String? = nil
  @Published var valueErrorMessage: String? = nil
  @Published var password: String? = nil
  @Published var passwordErrorMessage: String? = nil
  @Published var confirmed: Bool? = nil
  @Published var confirmedErrorMessage: Bool? = nil
  @Published var accessToken: String? = nil
  @Published var accessTokenErrorMessage: String? = nil
  func getDto() -> PassportEntity {
      var dto = PassportEntity()
    dto.type = self.type
    dto.value = self.value
    dto.password = self.password
    dto.confirmed = self.confirmed
    dto.accessToken = self.accessToken
      return dto
  }
}