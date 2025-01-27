import SwiftUI
struct BackupTableMetaForm: View {
    @State private var formModel = BackupTableMetaEntity()
    var body: some View {
        Form {
            Section {
                TextField("Email or Phone Number", text: $formModel.TableNameInDb.toUnwrapped(defaultValue: ""))
                    .padding()
                    .autocapitalization(.none)
                    .keyboardType(.emailAddress)
            }
        }
    }
}