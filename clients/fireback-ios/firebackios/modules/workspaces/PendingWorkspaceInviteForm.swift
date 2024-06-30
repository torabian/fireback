import SwiftUI
struct PendingWorkspaceInviteForm: View {
    @State private var formModel = PendingWorkspaceInviteEntity()
    var body: some View {
        Form {
            Section {
                TextField("Email or Phone Number", text: $formModel.Value.toUnwrapped(defaultValue: ""))
                    .padding()
                    .autocapitalization(.none)
                    .keyboardType(.emailAddress)
                TextField("Email or Phone Number", text: $formModel.Type.toUnwrapped(defaultValue: ""))
                    .padding()
                    .autocapitalization(.none)
                    .keyboardType(.emailAddress)
                TextField("Email or Phone Number", text: $formModel.CoverLetter.toUnwrapped(defaultValue: ""))
                    .padding()
                    .autocapitalization(.none)
                    .keyboardType(.emailAddress)
                TextField("Email or Phone Number", text: $formModel.WorkspaceName.toUnwrapped(defaultValue: ""))
                    .padding()
                    .autocapitalization(.none)
                    .keyboardType(.emailAddress)
            }
        }
    }
}