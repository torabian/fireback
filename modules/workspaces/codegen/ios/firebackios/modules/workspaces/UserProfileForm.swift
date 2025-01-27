import SwiftUI
struct UserProfileForm: View {
    @State private var formModel = UserProfileEntity()
    var body: some View {
        Form {
            Section {
                TextField("Email or Phone Number", text: $formModel.FirstName.toUnwrapped(defaultValue: ""))
                    .padding()
                    .autocapitalization(.none)
                    .keyboardType(.emailAddress)
                TextField("Email or Phone Number", text: $formModel.LastName.toUnwrapped(defaultValue: ""))
                    .padding()
                    .autocapitalization(.none)
                    .keyboardType(.emailAddress)
            }
        }
    }
}