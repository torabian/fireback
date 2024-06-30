import SwiftUI
struct WorkspaceTypeForm: View {
    @State private var formModel = WorkspaceTypeEntity()
    var body: some View {
        Form {
            Section {
                TextField("Email or Phone Number", text: $formModel.Title.toUnwrapped(defaultValue: ""))
                    .padding()
                    .autocapitalization(.none)
                    .keyboardType(.emailAddress)
                TextField("Email or Phone Number", text: $formModel.Description.toUnwrapped(defaultValue: ""))
                    .padding()
                    .autocapitalization(.none)
                    .keyboardType(.emailAddress)
                TextField("Email or Phone Number", text: $formModel.Slug.toUnwrapped(defaultValue: ""))
                    .padding()
                    .autocapitalization(.none)
                    .keyboardType(.emailAddress)
            }
        }
    }
}