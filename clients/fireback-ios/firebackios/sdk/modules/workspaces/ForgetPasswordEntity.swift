import Foundation
class ForgetPasswordEntity : Codable, Identifiable {
    var user: UserEntity? = nil
    // var userId: String? = nil
    var passport: PassportEntity? = nil
    // var passportId: String? = nil
    var status: String? = nil
    var validUntil: String? = nil
    var blockedUntil: String? = nil
    var secondsToUnblock: Int? = nil
    var otp: String? = nil
    var recoveryAbsoluteUrl: String? = nil
}
class ForgetPasswordEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var status: String? = nil
  @Published var statusErrorMessage: String? = nil
  @Published var secondsToUnblock: Int? = nil
  @Published var secondsToUnblockErrorMessage: Int? = nil
  @Published var otp: String? = nil
  @Published var otpErrorMessage: String? = nil
  @Published var recoveryAbsoluteUrl: String? = nil
  @Published var recoveryAbsoluteUrlErrorMessage: String? = nil
  func getDto() -> ForgetPasswordEntity {
      var dto = ForgetPasswordEntity()
    dto.status = self.status
    dto.secondsToUnblock = self.secondsToUnblock
    dto.otp = self.otp
    dto.recoveryAbsoluteUrl = self.recoveryAbsoluteUrl
      return dto
  }
}