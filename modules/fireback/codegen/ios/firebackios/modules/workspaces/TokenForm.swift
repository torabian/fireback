import SwiftUI
struct TokenForm: View {
    @State private var formModel = TokenEntity()
    var body: some View {
        Form {
            Section {
                TextField("Email or Phone Number", text: $formModel.ValidUntil.toUnwrapped(defaultValue: ""))
                    .padding()
                    .autocapitalization(.none)
                    .keyboardType(.emailAddress)
            }
        }
    }
}