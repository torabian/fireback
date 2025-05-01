import SwiftUI
struct ForgetPasswordForm: View {
    @State private var formModel = ForgetPasswordEntity()
    var body: some View {
        Form {
            Section {
                TextField("Email or Phone Number", text: $formModel.Status.toUnwrapped(defaultValue: ""))
                    .padding()
                    .autocapitalization(.none)
                    .keyboardType(.emailAddress)
                TextField("Email or Phone Number", text: $formModel.Otp.toUnwrapped(defaultValue: ""))
                    .padding()
                    .autocapitalization(.none)
                    .keyboardType(.emailAddress)
                TextField("Email or Phone Number", text: $formModel.RecoveryAbsoluteUrl.toUnwrapped(defaultValue: ""))
                    .padding()
                    .autocapitalization(.none)
                    .keyboardType(.emailAddress)
            }
        }
    }
}