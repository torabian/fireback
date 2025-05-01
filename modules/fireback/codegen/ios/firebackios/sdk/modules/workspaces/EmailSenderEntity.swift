import Foundation
class EmailSenderEntity : Codable, Identifiable {
    var fromName: String? = nil
    var fromEmailAddress: String? = nil
    var replyTo: String? = nil
    var nickName: String? = nil
}
class EmailSenderEntityViewModel: ObservableObject {
  // improve the fields here
  @Published var fromName: String? = nil
  @Published var fromNameErrorMessage: String? = nil
  @Published var fromEmailAddress: String? = nil
  @Published var fromEmailAddressErrorMessage: String? = nil
  @Published var replyTo: String? = nil
  @Published var replyToErrorMessage: String? = nil
  @Published var nickName: String? = nil
  @Published var nickNameErrorMessage: String? = nil
  func getDto() -> EmailSenderEntity {
      var dto = EmailSenderEntity()
    dto.fromName = self.fromName
    dto.fromEmailAddress = self.fromEmailAddress
    dto.replyTo = self.replyTo
    dto.nickName = self.nickName
      return dto
  }
}